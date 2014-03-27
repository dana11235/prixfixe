// Core cache class
package prixfixe

import (
	"sort"
  "strings"
)

type Cache struct {
	root *CacheItem
}

func NewCache() *Cache {
	c := new(Cache)
	c.root = NewCacheItem("")
	return c
}

func (c Cache) Put(iKey string, tokens map[string]string) *CacheItem {
  key := strings.ToLower(iKey)
	currItem := c.root
	for index := 0; index <= len(key); index++ {
		val, present := currItem.descendants[key[:index]]
		if present {
			currItem = val
		} else {
			currItem.descendants[key[:index]] = NewCacheItem(key[:index])
			currItem = currItem.descendants[key[:index]]
		}
	}
	currItem.Tokens = tokens
	return currItem
}

func (c Cache) Get(iKey string) *CacheItem {
  key := strings.ToLower(iKey)
	currItem := c.root
	for index := 0; index <= len(key); index++ {
		_, present := currItem.descendants[key[:index]]
		if !present {
			return nil
		} else {
			currItem = currItem.descendants[key[:index]]
		}
	}

	// If there are no tokens, it is nil
	if len(currItem.Tokens) > 0 {
		return currItem
	} else {
		return nil
	}
}

func (c Cache) Delete(key string) {
	c.Put(key, nil)
}

func (ci CacheItem) retrieveDescendants() CacheItems {
	var matches []*CacheItem
	if len(ci.Tokens) > 0 {
		matches = append(matches, &ci)
	}
	for k := range ci.descendants {
		currItem := ci.descendants[k]
		newMatches := currItem.retrieveDescendants()
		matches = append(matches, newMatches...)
	}
	return matches
}

func (c Cache) PrefixSearch(iKey string) []*CacheItem {
  key := strings.ToLower(iKey)
	if len(key) == 0 {
		return nil
	}
	currItem := c.root
	for index := 0; index <= len(key); index++ {
		_, present := currItem.descendants[key[:index]]
		if !present {
			return nil
		} else {
			currItem = currItem.descendants[key[:index]]
		}
	}
	return currItem.retrieveDescendants()
}

func (c Cache) SortedPrefixSearch(key string) []*CacheItem {
	descendants := c.PrefixSearch(key)
	sort.Sort(ByKey{descendants})
	return descendants
}
