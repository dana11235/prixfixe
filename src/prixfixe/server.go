// An extremely simple server that reads/writes key/value pairs and allows
// prefix searching
package prixfixe

import (
  "fmt"
  "net/http"
  "encoding/json"
)

// The static instance that this server uses
var staticCache *Cache = NewCache()

func RunServer() {
    http.HandleFunc("/put", putHandler)
    http.HandleFunc("/get", getHandler)
    http.HandleFunc("/search", searchHandler)
    http.ListenAndServe(":8080", nil)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
    key := r.FormValue("key")
    jsonValues := r.FormValue("value")
    encodedValues := []byte(jsonValues)
    if len(key) > 0 && len(jsonValues) > 0 {
      var values map[string]string
      err := json.Unmarshal(encodedValues, &values)
      if err == nil {
        staticCache.Put(key, values)
        fmt.Fprintf(w, "OK")
      } else {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
      }
    } else {
      http.Error(w, "Invalid Input: Must Specify Key and Value", http.StatusServiceUnavailable)
    }
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    key := r.FormValue("key")
    if len(key) > 0 {
      value := staticCache.Get(key)
      if value == nil {
        http.Error(w, "Value Not Found", http.StatusNotFound)
      } else {
        encodedValue, err := json.Marshal(value)
        if err == nil {
          fmt.Fprintf(w, string(encodedValue))
        } else {
          http.Error(w, err.Error(), http.StatusInternalServerError)
        }
      }
    } else {
      http.Error(w, "Invalid Input: Must Specify Key", http.StatusServiceUnavailable)
    }
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    key := r.FormValue("key")
    if len(key) > 0 {
      value := staticCache.SortedPrefixSearch(key)
      if value == nil {
        http.Error(w, "Value Not Found", http.StatusNotFound)
      } else {
        encodedValue, err := json.Marshal(value)
        if err == nil {
          fmt.Fprintf(w, string(encodedValue))
        } else {
          http.Error(w, err.Error(), http.StatusInternalServerError)
        }
      }
    } else {
      http.Error(w, "Invalid Input: Must Specify Key", http.StatusServiceUnavailable)
    }
}
