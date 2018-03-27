package main

import "sync"

type RoundRobinBalancer struct {
	serversNum    int
	nextServerIdx int
	stats         []int
	mutex         sync.Mutex
}

func (balancer *RoundRobinBalancer) Init(serversCount int) {
	if serversCount < 1 {
		panic("Can't start balancer with < 1 servers")
	}

	balancer.serversNum = serversCount
	balancer.stats = make([]int, serversCount, serversCount)
	balancer.nextServerIdx = 0
}

func (balancer *RoundRobinBalancer) GiveStat() []int {
	return balancer.stats
}

func (balancer *RoundRobinBalancer) GiveNode() (sendToServer int) {
	balancer.mutex.Lock()

	sendToServer = balancer.nextServerIdx
	balancer.stats[sendToServer] += 1

	if sendToServer == balancer.serversNum-1 {
		balancer.nextServerIdx = 0
	} else {
		balancer.nextServerIdx++
	}

	balancer.mutex.Unlock()
	return
}
