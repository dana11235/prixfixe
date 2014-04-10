// An extremely simple server that reads/writes key/value pairs and allows
// prefix searching
package server

import (
  . "prixfixe"
  "fmt"
  "flag"
  "net/http"
  "log"
  "os"
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
var authKey = flag.String("auth_key", "", "Key to authenticate write requests")
func parseFlags() {
  flag.Parse()
  if len(*fileName) > 0 {
    // Replay the log
    replayTransactionLog()
    // Open the log for writing
    openTransactionLog()
  }
}
