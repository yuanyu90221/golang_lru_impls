package lru

import "container/list"

type Node struct {
	K, V int
}

type LRUCache struct {
	capacity int
	list     *list.List
	cacheMap map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cacheMap: make(map[int]*list.Element, capacity),
	}
}

func (lru *LRUCache) Get(key int) int {
	if ele, ok := lru.cacheMap[key]; ok {
		lru.list.MoveToFront(ele)
		return ele.Value.(*Node).V
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if ele, ok := lru.cacheMap[key]; ok {
		lru.list.MoveToFront(ele)
		ele.Value.(*Node).V = value
		return
	}
	if lru.list.Len() == lru.capacity {
		last := lru.list.Back()
		node := last.Value.(*Node)
		delete(lru.cacheMap, node.K)
		lru.list.Remove(last)
	}
	ele := lru.list.PushFront(&Node{K: key, V: value})
	lru.cacheMap[key] = ele
}
