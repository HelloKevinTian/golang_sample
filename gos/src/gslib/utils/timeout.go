package utils

// import (
// 	"time"
// )

// func WARP_TIMEOUT(fun func(), seconds int, err_msg ...interface{}) {
// 	select {
// 	case fun():
// 	case <-time.After(seconds * time.Second):
// 		WARN(err_msg)
// 	}
// }

// func GEN_TIMEOUT(fun func(), err_msg ...interface{}) {
// 	select {
// 	case fun():
// 	case <-time.After(GEN_SERVER_TIMEOUT * time.Second):
// 		WARN(err_msg)
// 	}
// }
