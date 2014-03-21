// Functions for dealing with writing/replaying transactions
package server

import (
  . "prixfixe"
  "bufio"
  "os"
  "log"
  "encoding/json"
)

type JsonRecord struct {
  Key string
  Tokens map[string]string
}

func replayTransactionLog() {
    file, err := os.Open(*fileName)
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
    }
}

func openTransactionLog() {
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
