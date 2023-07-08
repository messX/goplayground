package lru

import (
	log "github.com/sirupsen/logrus"
)

type Lru struct {
	Capacity int
	items    []LruItems
}

type LruItems struct {
	Key   string
	Value string
}

func NewLru(cap int) *Lru {
	return &Lru{
		Capacity: cap,
		items:    make([]LruItems, 0),
	}
}

func (lru *Lru) Get(key string) string {
	log.Infof("Looking for %s", key)
	var value string
	if len(lru.items) == 0 {
		log.Info("No data in the cache")
	}
	for index, data := range lru.items {
		if data.Key == key {
			lru.items = moveTop(lru.items, index)
			value = data.Value
			break
		}
	}
	if value == "" {
		log.Warnf("Key not found %s", key)
	}
	return value
}

func (lru *Lru) Put(key string, val string) {
	log.Infof("Putting data for: %s", key)
	if len(lru.items) == lru.Capacity {
		lru.items = append(lru.items[1:], LruItems{
			Key:   key,
			Value: val,
		})
	} else {
		for index, data := range lru.items {
			if data.Key == key {
				data.Value = val
				lru.items = moveTop(lru.items, index)
				return
			}
		}
	}
	newData := []LruItems{{
		Key:   key,
		Value: val,
	}}
	if len(lru.items) < lru.Capacity {
		lru.items = append(lru.items, newData...)
	} else {
		lru.items = append(lru.items[:lru.Capacity-1], newData...)
	}

}

func moveTop(lruItems []LruItems, index int) []LruItems {
	if len(lruItems) == 1 {
		return lruItems
	}

	if index == 0 {
		lruItems = append(lruItems[1:], lruItems[index])
	} else {
		data := lruItems[index]
		lruItems = append(append(lruItems[:index], lruItems[index+1:]...),
			data)
	}
	return lruItems
}

func (lru *Lru) Print() {
	log.Info(lru.items)
}

func Test() {
	lru := NewLru(5)
	lru.Put("test1", "1")
	lru.Put("test2", "2")
	log.Info("it 0")
	lru.Print()
	lru.Get("test1")
	log.Info("it 1")
	lru.Print()
	lru.Put("test3", "3")
	lru.Put("test4", "4")
	lru.Put("test5", "5")
	lru.Put("test6", "6")
	log.Info("it 2")
	lru.Print()
	lru.Get("test3")
	log.Info("it 3")
	lru.Print()
	lru.Get("test2")
	log.Info("it 4")
	lru.Print()
}
