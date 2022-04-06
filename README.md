# golang_lru_impls

LRU(Least Recently Used) 是一種演算法把最近未使用的 Cache 值給移除掉


## 觀察

首先是 cache 有一個最大容量 K 

並且在超過這個容量 K 時 ， 有值要插入需要把最少存取的值移除

再在把這個新的值放入

## 初步想法是

首先是 

1 必須儲存 cache 容量值 cap

2 所有的 cache 值是 key value 並且要被有順序性保存，舉例來說，最新被存儲放最前面，最晚被存除放最後面

3 所有 cache 值由於需要快速存儲可以透過 map 來以 key value 方式做存放


簡單的想法是

透過 container 中 list 來紀錄所有值，最新存取放最前面

case 1: 放入新值時

每次存放新的值先檢查 cache 是否已經有該 key 值， 如果有更新 value 值 並且把該 key value 移到 list 最前面

如果不具有該 key值， 則先檢查 list 是否已達到 cap

如果達到 cap 則移除 list 最後元素以及在 map 之中的 key 

把 key value 放到 list 最前面並且放入 map 之中

case 2: 存取值某個 key 時

如果 key 不存在於 map 中，則回傳 -1

如果 key 存在則把 元素 key value 放到最前面， 並且回傳 value

## 實作

```golang
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

```