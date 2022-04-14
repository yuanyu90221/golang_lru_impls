package lru

import "fmt"

func RunLRU(actions []string, value [][]int) []string {
	length := len(actions)
	result := make([]string, length)
	result[0] = "null"
	lru := Constructor(value[0][0])
	for idx := 1; idx < length; idx++ {
		action := actions[idx]
		switch action {
		case "put":
			lru.Put(value[idx][0], value[idx][1])
			result[idx] = "null"
		case "get":
			result[idx] = fmt.Sprintf("%d", lru.Get(value[idx][0]))
		}
	}
	return result
}

func RunLRUV2(actions []string, value [][]int) []string {
	length := len(actions)
	result := make([]string, length)
	result[0] = "null"
	lru := ConstructorV2(value[0][0])
	for idx := 1; idx < length; idx++ {
		action := actions[idx]
		switch action {
		case "put":
			lru.Put(value[idx][0], value[idx][1])
			result[idx] = "null"
		case "get":
			result[idx] = fmt.Sprintf("%d", lru.Get(value[idx][0]))
		}
	}
	return result
}
