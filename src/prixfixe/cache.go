// Core cache class
package prixfixe

import (
  "sort"
)

type Cache struct {
  root *CacheItem
}

func NewCache() *Cache {
  c := new(Cache)
  c.root = NewCacheItem("")
  return c
}

func (c Cache) Put(key string, tokens map[string]string) {
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
}

func (c Cache) Get(key string) *CacheItem {
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
  for k := range ci.descendants {
    currItem := ci.descendants[k]
    if len(currItem.Tokens) > 0 {
      matches = append(matches, currItem)
    }
    if len(currItem.descendants) > 0 {
      newMatches := currItem.retrieveDescendants()
      matches = append(matches, newMatches...)
    }
  }
  return matches
}

func (c Cache) PrefixSearch(key string) []*CacheItem {
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
