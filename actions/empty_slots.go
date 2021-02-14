package actions

import (
	"container/heap"
	"errors"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	topEl := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return topEl
}

type EmptySlots struct {
	h *IntHeap
}

func NewEmptySlots(numSlots int) *EmptySlots {
	slots := make([]int, numSlots)
	for i := 0; i < numSlots; i++ {
		slots[i] = i
	}
	h := IntHeap(slots)
	heap.Init(&h)
	return &EmptySlots{&h}
}

func (eSlots *EmptySlots) GetEmpty() (int, error) {
	if eSlots.h.Len() == 0 {
		return -1, errors.New("Parking lot Full")
	}
	slotIdx := heap.Pop(eSlots.h).(int)
	return slotIdx, nil
}

func (eSlots *EmptySlots) SetEmpty(slot int) {
	heap.Push(eSlots.h, slot)
}
