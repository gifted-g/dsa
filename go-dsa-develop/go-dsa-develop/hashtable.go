package main

import (
	"hash/fnv"
)

type KeyValue struct {
	Key   string
	Value string
}

type HashTable struct {
	buckets [][]KeyValue
	length  int
	cap     int
}

func NewHashTable() *HashTable {
	const defaultBucketSize = 5

	return &HashTable{
		buckets: make([][]KeyValue, defaultBucketSize),
		length:  0,
		cap:     defaultBucketSize,
	}
}

func GenerateHash(key string, bucketSize int) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % bucketSize
}

func (ht *HashTable) Put(key, value string) {
	hash := GenerateHash(key, ht.cap)

	// if key already exist
	for i, kv := range ht.buckets[hash] {
		if kv.Key == key {
			ht.buckets[hash][i].Value = value
			return
		}
	}

	kv := KeyValue{
		Key:   key,
		Value: value,
	}

	ht.buckets[hash] = append(ht.buckets[hash], kv)

	ht.length++
	if ht.loadFactor() > 0.5 {
		ht.resize()
	}
}

func (ht *HashTable) loadFactor() float32 {
	return float32(ht.length) / float32(len(ht.buckets))
}

func (ht *HashTable) resize() {
	newCap := ht.cap * 2
	newItems := make([][]KeyValue, newCap)

	for _, buckets := range ht.buckets {
		for _, k := range buckets {
			hash := GenerateHash(k.Key, newCap)
			newItems[hash] = append(newItems[hash], KeyValue{k.Key, k.Value})
		}
	}

	ht.buckets = newItems
	ht.cap = newCap
}

func (ht *HashTable) Get(key string) (string, bool) {
	hash := GenerateHash(key, ht.cap)
	for _, kv := range ht.buckets[hash] {
		if kv.Key == key {
			return kv.Value, true
		}
	}

	return "", false
}
