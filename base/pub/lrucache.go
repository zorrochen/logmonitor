package pub

import (
	"container/list"
	"errors"
)

type CacheNode struct {
	Key, Value interface{}
}

func (cnode *CacheNode) NewCacheNode(k, v interface{}) *CacheNode {
	return &CacheNode{k, v}
}

type LRUCache struct {
	Capacity int
	dlist    *list.List
	cacheMap map[interface{}]*list.Element
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Capacity: cap,
		dlist:    list.New(),
		cacheMap: make(map[interface{}]*list.Element)}
}

func (lru *LRUCache) Size() int {
	return lru.dlist.Len()
}

func (lru *LRUCache) Set(k, v interface{}) error {

	if lru.dlist == nil {
		return errors.New("LRUCache init failed.")
	}

	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pElement)
		pElement.Value.(*CacheNode).Value = v
		return nil
	}

	newElement := lru.dlist.PushFront(&CacheNode{k, v})
	lru.cacheMap[k] = newElement

	if lru.dlist.Len() > lru.Capacity {
		//移掉最后一个
		lastElement := lru.dlist.Back()
		if lastElement == nil {
			return nil
		}
		cacheNode := lastElement.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode.Key)
		lru.dlist.Remove(lastElement)
	}
	return nil
}

// func (lru *LRUCache) Get(k interface{}) (v interface{}, ret bool, err error) {
// 	if lru.cacheMap == nil {
// 		return v, false, errors.New("LRUCache init failed.")
// 	}
//
// 	if pElement, ok := lru.cacheMap[k]; ok {
// 		lru.dlist.MoveToFront(pElement)
// 		return pElement.Value.(*CacheNode).Value, true, nil
// 	}
// 	return v, false, nil
// }

func (lru *LRUCache) GetAllWithoutUpdate() (retlist []interface{}, err error) {
	if lru.cacheMap == nil {
		return retlist, errors.New("LRUCache init failed.")
	}

	for e := lru.dlist.Front(); e != nil; e = e.Next() {
		retlist = append(retlist, e.Value.(*CacheNode).Value)
	}

	return retlist, nil
}

func (lru *LRUCache) Remove(k interface{}) bool {

	if lru.cacheMap == nil {
		return false
	}

	if pElement, ok := lru.cacheMap[k]; ok {
		cacheNode := pElement.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode.Key)
		lru.dlist.Remove(pElement)
		return true
	}
	return false
}
