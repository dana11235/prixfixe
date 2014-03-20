package prixfixe

import "testing"

func TestPutNormalElement(t *testing.T) {
  c := NewCache()
  c.Put("Dana", map[string]string{"first_name": "Dana", "last_name": "Levine"})
  el := c.Get("Dana")
  if el == nil {
    t.Error("Element Not Inserted")
    t.Fail()
  }
  if len(el.Tokens) == 0 {
    t.Error("Tokens Not Inserted")
    t.Fail()
  }
  el = c.Get("Dan")
  if el != nil {
    t.Error("Element Present When Should Be Nil")
    t.Fail()
  }
}

func TestGetNullElement(t *testing.T) {
  c := NewCache()
  el := c.Get("Dana")
  if el != nil {
    t.Error("Element Present When Should Be Nil")
    t.Fail()
  }
}

func TestOverWriteElement(t *testing.T) {
  c := NewCache()
  c.Put("Dana", map[string]string{"first_name": "Dana", "last_name": "Levine"})
  c.Put("Dana", map[string]string{"first_name": "Dana", "last_name": "Jones"})
  el := c.Get("Dana")
  if el == nil {
    t.Error("Element Not Inserted")
    t.Fail()
  }
  last_name := el.Tokens["last_name"]
  if last_name == "Levine" || last_name != "Jones" {
    t.Error("Last Name Not Replaced Correctly")
    t.Fail()
  }
}

func TestDeleteElement(t *testing.T) {
  c := NewCache()
  c.Put("Dana", map[string]string{"first_name": "Dana", "last_name": "Levine"})
  c.Delete("Dana")
  el := c.Get("Dana")
  if el != nil {
    t.Error("Element Not Deleted")
    t.Fail()
  }
}

func insertSampleItems(c *Cache) {
  c.Put("Dana", map[string]string{"first_name": "Dana", "last_name": "Levine"})
  c.Put("Dando", map[string]string{"first_name": "Dando", "last_name": "Levindo"})
  c.Put("Danalope", map[string]string{
    "first_name": "Danalope", "last_name": "Levindolop"})
  c.Put("Alissa", map[string]string{"first_name": "Alissa", "last_name": "Wong"})
}

func TestPrefixSearch(t *testing.T) {
  c := NewCache()
  insertSampleItems(c)
  results := c.PrefixSearch("Dan")
  if len(results) != 3 {
    t.Error("Wrong Number of Results")
    t.Fail()
  }
  names := map[string]bool{"Dana":true, "Dando":true, "Danalope":true}
  for _, e := range results {
    _, prs := names[e.Key]
    if !prs {
      t.Error("Unexpected result present: " + e.Key)
    }
  }
}

func TestFullPrefixSearch(t *testing.T) {
  c := NewCache()
  insertSampleItems(c)
  results := c.PrefixSearch("Alissa")
  if len(results) != 1 {
    t.Error("Wrong Number of Results")
    t.Fail()
  }
}

func TestEmptyPrefixSearch(t *testing.T) {
  c := NewCache()
  insertSampleItems(c)
  results := c.PrefixSearch("")
  if len(results) != 0 {
    t.Error("Wrong Number of Results")
    t.Fail()
  }
}
