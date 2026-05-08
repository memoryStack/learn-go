package cache

import (
	"errors"
	"lru-cache/dell"
)

type Cache struct {
	Size int
	dell *dell.Dell
	hashMap map[string]*dell.Node
}

func New(size int) *Cache {
	return &Cache {
		dell: dell.New(size),
		Size: size,
		hashMap: make(map[string]*dell.Node),
	}
}

func (c *Cache) Get(key string) (int, error) {
	node, exists := c.hashMap[key]
	if (exists) {
		c.dell.MoveNodeToLast(node)
		return node.Value, nil
	}
	return 0, errors.New("key not found")
}

// the key will always exist here
func (c *Cache) updateValue(key string, value int) error {
	node, _ := c.hashMap[key]
	node.Value = value
	c.dell.MoveNodeToLast(node)
	return nil
}

func (c *Cache) Put(key string, value int) error {
	node, exists := c.hashMap[key]
	if (exists) {
		return c.updateValue(key, value)
	}

	if (c.dell.Length < c.Size) {
		node, err := c.dell.Add(key, value)
		if (err != nil) {
			return err
		}
		c.hashMap[key] = node
		return nil
	}

	firstNode, _ := c.dell.GetFirstNode()
	delete(c.hashMap, firstNode.Key)
	node, err := c.dell.Add(key, value) // this logic doesn't belong in the dell package
	if (err != nil) {
		return err
	}
	c.hashMap[key] = node
	return nil
}

func (c *Cache) Print() {
	c.dell.Print()
}
