/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 * 6月 29 2020
 * 哨兵+双向指针+map
 */
package leetcode

type CacheRecord struct {
	pre *CacheRecord
	next *CacheRecord
	key int
	value int
}

type LRUCache struct {
	head *CacheRecord
	tail *CacheRecord
	m map[int]*CacheRecord
	capacity int
}


func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		&CacheRecord{},
		&CacheRecord{},
		make(map[int]*CacheRecord),
		capacity,
	}

	lruCache.head.next = lruCache.tail
	lruCache.tail.pre = lruCache.head

	return lruCache
}


func (this *LRUCache) Get(key int) int {
	record, exist := this.m[key]
	if exist {
		this.remove(this.m[key])
		this.setHead(this.m[key])
		return record.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	cacheRecord := &CacheRecord{nil,nil, key, value}
	_, exist := this.m[key]
	// not exist
	if len(this.m) < this.capacity && !exist {
		this.setHead(cacheRecord)
		this.m[key] = cacheRecord
	} else if exist {
		this.remove(this.m[key])
		this.m[key] = cacheRecord
		this.setHead(this.m[key])
	} else {
		// length equals capacity && not exists
		deleteKey := this.tail.pre.key
		delete(this.m, deleteKey)
		this.remove(this.tail.pre)

		// add new cache record
		this.setHead(cacheRecord)
		this.m[key] = cacheRecord
	}

}

func (this *LRUCache) remove(record *CacheRecord) {
	if len(this.m) == 0 {
		return
	} else {
		record.pre.next = record.next
		record.next.pre = record.pre
	}
}

func (this *LRUCache) setHead(record *CacheRecord) {
	record.next = this.head.next
	record.pre = this.head
	this.head.next.pre = record
	this.head.next = record
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

