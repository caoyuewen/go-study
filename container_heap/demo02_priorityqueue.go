/**
 * Description: 
 * User: 1067
 * Date: 2018-09-12
 * Time: 17:31
 */

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))

	i := 0

	for value, priority := range items {

		pq[i] = &Item{
			value,
			priority,
			i,
		}
		i++
	}

	heap.Init(&pq)

	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	pq.Update(item, item.value, 5)

	for len(pq) > 0 {
		item := pq.Pop().(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}

}

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)

}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0:n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
