package cache

import (
	"errors"
	"sync"
)

// Cache struct
type Cache struct {
	mutex sync.RWMutex
	pages map[string]interface{}
}

// Page interface
type Page interface {
	GetSlug() string
}

// New returns a new cache
func New() *Cache {
	return &Cache{
		pages: make(map[string]interface{}),
	}
}

// StorePage adds a page to the cache
func (c *Cache) StorePage(slug string, page interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.pages[slug] = page
}

// FindPage returns a page if found or an error
func (c *Cache) FindPage(slug string) (interface{}, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.pages[slug]; ok == false {
		return nil, errors.New("Page not in cache")
	}

	return c.pages[slug], nil
}

// All returns all pages in cache
func (c *Cache) All() map[string]interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.pages
}
