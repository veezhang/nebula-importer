package pos

import (
	"container/heap"
	"fmt"
	"sync"
)

func Push(p int) {
	gPos.Lock()
	heap.Push(&gPos, p)
	if p > gPos.max {
		gPos.max = p
	}
	gPos.Unlock()
}

func Pop() int {
	gPos.Lock()
	p := heap.Pop(&gPos)
	gPos.Unlock()
	return p.(int)
}

func Remove(p int) {
	gPos.Lock()
	defer gPos.Unlock()
	for i := 0; i < gPos.Len(); i++ {
		if gPos.posSlice[i] == p {
			heap.Remove(&gPos, i)
			return
		}
	}
	panic(fmt.Sprintf("pos %d not exists", p))
}

func Top() int {
	gPos.Lock()
	p := gPos.max
	if gPos.Len() > 0 {
		p = gPos.posSlice[0]
	}
	gPos.Unlock()
	return p
}

var gPos = struct {
	sync.Mutex
	posSlice
	max int
}{}

type posSlice []int

func (p posSlice) Len() int           { return len(p) }
func (p posSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p posSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *posSlice) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *posSlice) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}
