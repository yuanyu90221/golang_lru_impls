package lru

type NodeV2 struct {
	K, V       int
	Prev, Next *NodeV2
}

type LRUCacheV2 struct {
	head, tail *NodeV2
	keys       map[int]*NodeV2
	capacity   int
}

func ConstructorV2(capacity int) LRUCacheV2 {
	return LRUCacheV2{keys: make(map[int]*NodeV2), capacity: capacity}
}

func (cache *LRUCacheV2) Get(key int) int {
	if node, ok := cache.keys[key]; ok {
		cache.Remove(node)
		cache.AddToFirst(node)
		return node.V
	}
	return -1
}

func (cache *LRUCacheV2) Put(key int, value int) {
	if node, ok := cache.keys[key]; ok {
		node.V = value
		cache.Remove(node)
		cache.AddToFirst(node)
		return
	} else {
		node = &NodeV2{K: key, V: value}
		cache.keys[key] = node
		cache.AddToFirst(node)
	}
	if len(cache.keys) > cache.capacity {
		delete(cache.keys, cache.tail.K)
		cache.Remove(cache.tail)
	}
}

func (cache *LRUCacheV2) AddToFirst(node *NodeV2) {
	node.Prev = nil
	node.Next = cache.head
	if cache.head != nil {
		cache.head.Prev = node
	}
	cache.head = node
	if cache.tail == nil {
		cache.tail = node
		cache.tail.Next = nil
	}
}

func (cache *LRUCacheV2) Remove(node *NodeV2) {
	if node == cache.head {
		cache.head = node.Next
		node.Next = nil
		return
	}
	if node == cache.tail {
		cache.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
