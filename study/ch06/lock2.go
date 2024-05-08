package main

import "sync"

/*
	我们有两组协程，其中一组负责协，另一个负责读，web大多是读多写少，
	虽然有多个协程，但是仔细分析我们会发现，读协程之间应该并发，读和写之间应该串行，读和读之间不应该并行
*/

func main() {

	var rwlock sync.RWMutex

}
