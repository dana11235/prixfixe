package server

import (
  "fmt"
  "net/http"
  "encoding/json"
)

func loadHandlers() {
    http.HandleFunc("/put", putHandler)
    http.HandleFunc("/putall", putAllHandler)
    http.HandleFunc("/get", getHandler)
    http.HandleFunc("/search", searchHandler)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
    key := r.FormValue("key")
    jsonValues := r.FormValue("value")
    encodedValues := []byte(jsonValues)
    if len(key) > 0 && len(jsonValues) > 0 {
      var values map[string]string
      err := json.Unmarshal(encodedValues, &values)
      if err == nil {
        insertedItem := staticCache.Put(key, values)
        writeTransaction(insertedItem)
        fmt.Fprintf(w, "OK")
      } else {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
      }
    } else {
      http.Error(w, "Invalid Input: Must Specify Key and Value", http.StatusServiceUnavailable)
    }
}

func putAllHandler(w http.ResponseWriter, r *http.Request) {
    /*
    values := r.FormValue("values")
    if len(values) > 0 {
      encodedValues := []byte(values)
      var values map[string]string
      err := json.Unmarshal(encodedValues, &values)
      if err == nil {
        insertedItem := staticCache.Put(key, values)
        writeTransaction(insertedItem)
        fmt.Fprintf(w, "OK")
      } else {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
      }
    } else {
      http.Error(w, "Invalid Input: Must Specify Values", http.StatusServiceUnavailable)
    }
    */
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
          output := padJsonp(r.FormValue("jsonp"), string(encodedValue))
          fmt.Fprintf(w, output)
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
          output := padJsonp(r.FormValue("jsonp"), string(encodedValue))
          fmt.Fprintf(w, output)
        } else {
          http.Error(w, err.Error(), http.StatusInternalServerError)
        }
      }
    } else {
      http.Error(w, "Invalid Input: Must Specify Key", http.StatusServiceUnavailable)
    }
}

// Helper functions used by handlers
func padJsonp(jsonp string, jsonresp string) string {
  if len(jsonp) == 0 {
    return jsonresp
  } else {
    return fmt.Sprintf("%s(%s);", jsonp, jsonresp)
  }
}
