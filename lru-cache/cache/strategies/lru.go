package strategies

import (
	"errors"
	"lru-cache/dell"
)

// LRUCache implements cache.Cache when used from package cache.
type LRUCache struct {
	capacity int
	dell     *dell.Dell
	hashMap  map[string]*dell.Node
}

// NewLRUCache builds an LRU-backed cache with the given capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		dell:     dell.New(),
		hashMap:  make(map[string]*dell.Node),
	}
}

func (c *LRUCache) Get(key string) (int, error) {
	node, exists := c.hashMap[key]
	if exists {
		c.dell.MoveNodeToLast(node)
		return node.GetValue(), nil
	}
	return 0, errors.New("key not found")
}

func (c *LRUCache) updateValue(key string, value int) error {
	node, _ := c.hashMap[key]
	node.UpdateValue(value)
	c.dell.MoveNodeToLast(node)
	return nil
}

func (c *LRUCache) Put(key string, value int) error {
	node, exists := c.hashMap[key]
	if exists {
		return c.updateValue(key, value)
	}

	if c.dell.GetLength() < c.capacity {
		node, err := c.dell.Add(key, value)
		if err != nil {
			return err
		}
		c.hashMap[key] = node
		return nil
	}

	firstNode, _ := c.dell.GetFirstNode()
	delete(c.hashMap, firstNode.GetKey())
	c.dell.Delete(firstNode)
	node, err := c.dell.Add(key, value)
	if err != nil {
		return err
	}
	c.hashMap[key] = node
	return nil
}

func (c *LRUCache) Print() {
	c.dell.Print()
}
