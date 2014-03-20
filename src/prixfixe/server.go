// An extremely simple server that reads/writes key/value pairs and allows
// prefix searching
package prixfixe

import (
  "fmt"
  "net/http"
  "encoding/json"
  "flag"
  "log"
  "io/ioutil"
)

// The static instance that this server uses
var staticCache *Cache = NewCache()

func RunServer() {
    parseFlags()
    loadHandlers()
    log.Println("Listening on Port", *port)
    http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}


var fileName = flag.String("file", "", "CSV File to load data from")
var port = flag.Int("port", 8080, "Port to bind the server to")
func parseFlags() {
  flag.Parse()
  if len(*fileName) > 0 {
    loadJsonFile(fileName)
  }
}

type JsonRecord struct {
  Key string
  Tokens map[string]string
}

func loadJsonFile(fileName* string) {
    file, err := ioutil.ReadFile(*fileName)
    if err != nil {
	    log.Fatal(err)
    } else {
      var jsonRecords []JsonRecord
      err := json.Unmarshal(file, &jsonRecords)
      if err == nil {
        for _, record := range jsonRecords {
          staticCache.Put(record.Key, record.Tokens)
        }
      } else {
        log.Fatal(err)
      }
    }
}

func loadHandlers() {
    http.HandleFunc("/put", putHandler)
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
        staticCache.Put(key, values)
        fmt.Fprintf(w, "OK")
      } else {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
      }
    } else {
      http.Error(w, "Invalid Input: Must Specify Key and Value", http.StatusServiceUnavailable)
    }
}

func padJsonp(jsonp string, jsonresp string) string {
  if len(jsonp) == 0 {
    return jsonresp
  } else {
    return fmt.Sprintf("%s(%s);", jsonp, jsonresp)
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
