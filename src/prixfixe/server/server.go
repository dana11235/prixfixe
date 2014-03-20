// An extremely simple server that reads/writes key/value pairs and allows
// prefix searching
package server

import (
  . "prixfixe"
  "fmt"
  "encoding/json"
  "flag"
  "log"
  "os"
  "net/http"
  "bufio"
)

// The static instance that this server uses
var staticCache *Cache = NewCache()

// The transaction file used to write the transaction log
var transactionLog *os.File

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
    // Replay the log
    replayJsonLog()
    // Open the log for writing
    openJsonLog()
  }
}

type JsonRecord struct {
  Key string
  Tokens map[string]string
}

func replayJsonLog() {
    file, err := os.Open(*fileName)
    //ioutil.ReadFile(*fileName)
    if err != nil {
	    log.Fatal(err)
    } else {
      scanner := bufio.NewScanner(file)
      var record JsonRecord
      for scanner.Scan() {
        line := []byte(scanner.Text())
        err := json.Unmarshal(line, &record)
        if err == nil {
            staticCache.Put(record.Key, record.Tokens)
        } else {
          log.Fatal(err)
        }
      }
      /* var jsonRecords []JsonRecord
      err := json.Unmarshal(file, &jsonRecords)
      if err == nil {
        for _, record := range jsonRecords {
          staticCache.Put(record.Key, record.Tokens)
        }
      } else {
        log.Fatal(err)
      }*/
    }
}

func openJsonLog() {
  file, err := os.OpenFile(*fileName, os.O_RDWR|os.O_APPEND, 0660);
  if err != nil {
    log.Fatal(err)
  } else {
    transactionLog = file
  }
}

func writeTransaction(c *CacheItem) {
  if transactionLog != nil {
    encodedValue, err := json.Marshal(c)
    if err == nil {
      n, err := transactionLog.Write(encodedValue)
      if err == nil {
        log.Println("Wrote", n, "bytes")
      } else {
        log.Println(err.Error())
      }
    } else {
      log.Println(err.Error())
    }
  }
}
