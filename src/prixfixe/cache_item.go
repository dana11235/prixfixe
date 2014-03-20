// Cache item. Contains some methods for sorting
package prixfixe

type CacheItem struct {
  Key string
  Tokens map[string]string
  descendants map[string]*CacheItem
}

func NewCacheItem(key string) *CacheItem {
  ci := new(CacheItem)
  ci.Key = key
  ci.descendants = make(map[string]*CacheItem)
  return ci
}

/* Functions used for sorting cache items */

type CacheItems []*CacheItem

// Len is part of sort.Interface.
func (s CacheItems) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s CacheItems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Key Comparator
type ByKey struct{ CacheItems }
func (s ByKey) Less(i, j int) bool {
  return s.CacheItems[i].Key < s.CacheItems[j].Key
}
