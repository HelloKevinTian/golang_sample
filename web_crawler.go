package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var websites []string

func main() {
	//runtime.GOMAXPROCS(4)
	// 各大电商网站首页数据量大小检测
	websites = []string{
		"http://www.51buy.com/", "http://www.360buy.com/", "http://www.tmall.com/", "http://www.taobao.com/",
		"http://china.alibaba.com/", "http://www.paipai.com/", "http://shop.qq.com/", "http://www.lightinthebox.com/",
		"http://www.amazon.cn/", "http://www.newegg.com.cn/", "http://www.vancl.com/", "http://www.yihaodian.com/",
		"http://www.dangdang.com/", "http://www.m18.com/", "http://www.suning.com/", "http://www.hstyle.com/",
		"http://shop.vipshop.com/home.php"}
	// 并发5个运行
	pnum := 10 // 默认设置10个并发测试
	parallelRequest(pnum, websites)
}

func parallelRequest(pnum int, websites []string) { // 并行抓取
	total := len(websites)
	if pnum <= 0 { // 在设定为0时，全部并发
		pnum = total
	}
	if pnum > total {
		pnum = total
	}
	startTime := time.Now().UnixNano()
	fetchData := make(map[string]string, total) // 反馈抓取后的数据结果，可以写入文件查看
	execChans := make(chan bool, pnum)          // 控制并发数量的通道，第二个参数指定通道可以容纳的数量，会阻塞执行
	doneChans := make(chan bool, 1)             // 用来传递完成信号，完成信号只需要设定容纳一位即可，完成后再次读取新的任务
	for i := 0; i < total; i++ {
		go request(i, websites[i], execChans, doneChans, fetchData) // 以协程方式运行
	}

	for i := 0; i < total; i++ {
		r := <-doneChans // 完成一个，同时获取下一个任务
		<-execChans      // 紧接着读取下一个任务，类是于beanstalkd的任务分发机制
		if !r {          // 获取失败时，打印该网址失败。
			log.Printf("第 %s 项获取失败", i)
		}
	}
	close(doneChans)                                            // 关闭完成信号
	close(execChans)                                            // 关闭执行信号
	processed := float32(time.Now().UnixNano()-startTime) / 1e9 // 统计总耗时
	log.Printf("全部完成，并发数量：%d, 共耗时：%.3fs", pnum, processed)
	log.Printf("data: %q", fetchData)
}

func request(i int, url string, execChans chan bool, doneChans chan bool, fetchData map[string]string) {
	execChans <- true // 放在函数的开始处，用来阻塞执行，如果通道里的数量超过设定数量，在没有读取完成前，不会运行
	log.Printf("start=>NO: %02d, url: %s", i, url)
	isOk := false
	startTime := time.Now().UnixNano()
	resp, _ := http.Get(url)
	defer (func() {
		resp.Body.Close()
		doneChans <- isOk
		processed := float32(time.Now().UnixNano()-startTime) / 1e9
		log.Printf("end  =>NO: %02d, url: %s, status: %t, time: %.3fs", i, url, isOk, processed)
	})()
	body, err := ioutil.ReadAll(resp.Body)
	len := len(body)
	fetchData[url] = fmt.Sprintf("len: %d", len)
	if err == nil {
		isOk = true
	}
}

//output
// 2015/08/28 10:37:31 start=>NO: 16, url: http://shop.vipshop.com/home.php
// 2015/08/28 10:37:31 start=>NO: 00, url: http://www.51buy.com/
// 2015/08/28 10:37:31 start=>NO: 01, url: http://www.360buy.com/
// 2015/08/28 10:37:31 start=>NO: 02, url: http://www.tmall.com/
// 2015/08/28 10:37:31 start=>NO: 03, url: http://www.taobao.com/
// 2015/08/28 10:37:31 start=>NO: 04, url: http://china.alibaba.com/
// 2015/08/28 10:37:31 start=>NO: 05, url: http://www.paipai.com/
// 2015/08/28 10:37:31 start=>NO: 06, url: http://shop.qq.com/
// 2015/08/28 10:37:31 start=>NO: 07, url: http://www.lightinthebox.com/
// 2015/08/28 10:37:31 start=>NO: 08, url: http://www.amazon.cn/
// 2015/08/28 10:37:31 end  =>NO: 05, url: http://www.paipai.com/, status: true, time: 0.026s
// 2015/08/28 10:37:31 start=>NO: 09, url: http://www.newegg.com.cn/
// 2015/08/28 10:37:31 end  =>NO: 03, url: http://www.taobao.com/, status: true, time: 0.100s
// 2015/08/28 10:37:31 start=>NO: 10, url: http://www.vancl.com/
// 2015/08/28 10:37:31 end  =>NO: 10, url: http://www.vancl.com/, status: true, time: 0.055s
// 2015/08/28 10:37:31 start=>NO: 11, url: http://www.yihaodian.com/
// 2015/08/28 10:37:31 end  =>NO: 02, url: http://www.tmall.com/, status: true, time: 0.168s
// 2015/08/28 10:37:31 start=>NO: 12, url: http://www.dangdang.com/
// 2015/08/28 10:37:31 end  =>NO: 01, url: http://www.360buy.com/, status: true, time: 0.204s
// 2015/08/28 10:37:31 start=>NO: 13, url: http://www.m18.com/
// 2015/08/28 10:37:31 end  =>NO: 00, url: http://www.51buy.com/, status: true, time: 0.207s
// 2015/08/28 10:37:31 start=>NO: 14, url: http://www.suning.com/
// 2015/08/28 10:37:32 end  =>NO: 12, url: http://www.dangdang.com/, status: true, time: 0.055s
// 2015/08/28 10:37:32 start=>NO: 15, url: http://www.hstyle.com/
// 2015/08/28 10:37:32 end  =>NO: 04, url: http://china.alibaba.com/, status: true, time: 0.269s
// 2015/08/28 10:37:32 end  =>NO: 14, url: http://www.suning.com/, status: true, time: 0.116s
// 2015/08/28 10:37:32 end  =>NO: 16, url: http://shop.vipshop.com/home.php, status: true, time: 0.330s
// 2015/08/28 10:37:32 end  =>NO: 13, url: http://www.m18.com/, status: true, time: 0.445s
// 2015/08/28 10:37:32 end  =>NO: 11, url: http://www.yihaodian.com/, status: true, time: 0.733s
// 2015/08/28 10:37:33 end  =>NO: 08, url: http://www.amazon.cn/, status: true, time: 1.356s
// 2015/08/28 10:37:33 end  =>NO: 06, url: http://shop.qq.com/, status: true, time: 1.469s
// 2015/08/28 10:37:34 end  =>NO: 07, url: http://www.lightinthebox.com/, status: true, time: 2.417s
// 2015/08/28 10:37:34 end  =>NO: 09, url: http://www.newegg.com.cn/, status: true, time: 2.710s
// 2015/08/28 10:37:39 end  =>NO: 15, url: http://www.hstyle.com/, status: true, time: 7.859s
// 2015/08/28 10:37:39 全部完成，并发数量：10, 共耗时：8.083s
// 2015/08/28 10:37:39 data: map["http://www.dangdang.com/":"len: 66620" "http://www.yihaodian.com/":"len: 164295"
// "http://www.newegg.com.cn/":"len: 288761" "http://www.hstyle.com/":"len: 264440" "http://www.paipai.com/":"len: 103970"
// "http://www.vancl.com/":"len: 49211" "http://www.360buy.com/":"len: 162286" "http://www.51buy.com/":"len: 151092"
// "http://www.suning.com/":"len: 254846" "http://www.m18.com/":"len: 246157" "http://shop.qq.com/":"len: 162286"
// "http://www.taobao.com/":"len: 152567" "http://www.tmall.com/":"len: 53902" "http://china.alibaba.com/":"len: 40772"
// "http://shop.vipshop.com/home.php":"len: 115869" "http://www.amazon.cn/":"len: 571912" "http://www.lightinthebox.com/":"len: 259075"]
