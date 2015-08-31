package consistent_hash

import (
	"fmt"
	"testing"
)

func TestConsitentHashing(t *testing.T) {
	ch := new(ConsistentHashing)
	ch.Init()
	t.Log("testing")
	t.Log(1234, fmt.Sprint(ch.GetNode(1234)))
	t.Log(12345, fmt.Sprint(ch.GetNode(12345)))
	t.Log(22345, fmt.Sprint(ch.GetNode(22345)))
	t.Log(32345, fmt.Sprint(ch.GetNode(32345)))
	t.Log("adding nodes a, b, c")
	ch.AddNode("a", 10000)
	ch.AddNode("a", 10000)
	ch.AddNode("b", 20000)
	ch.AddNode("c", 30000)

	t.Log("testing")
	t.Log(1234, fmt.Sprint(ch.GetNode(1234)))
	t.Log(12345, fmt.Sprint(ch.GetNode(12345)))
	t.Log(22345, fmt.Sprint(ch.GetNode(22345)))
	t.Log(32345, fmt.Sprint(ch.GetNode(32345)))

	t.Log("remove node", 20000)
	ch.RemoveNode(20000)
	t.Log("testing")
	t.Log(1234, fmt.Sprint(ch.GetNode(1234)))
	t.Log(12345, fmt.Sprint(ch.GetNode(12345)))
	t.Log(22345, fmt.Sprint(ch.GetNode(22345)))
	t.Log(32345, fmt.Sprint(ch.GetNode(32345)))
	t.Log("remove all node")
	ch.RemoveNode(10000)
	ch.RemoveNode(30000)
	t.Log("testing")
	t.Log(1234, fmt.Sprint(ch.GetNode(1234)))
	t.Log(12345, fmt.Sprint(ch.GetNode(12345)))
	t.Log(22345, fmt.Sprint(ch.GetNode(22345)))
	t.Log(32345, fmt.Sprint(ch.GetNode(32345)))
}
