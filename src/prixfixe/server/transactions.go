// Functions for dealing with writing/replaying transactions
package server

import (
  . "prixfixe"
  "bufio"
  "os"
  "log"
  "encoding/json"
  "io"
)

type JsonRecord struct {
  Key string
  Tokens map[string]string
}

func processTransaction(line []byte) (JsonRecord, error) {
  var record JsonRecord
  err := json.Unmarshal(line, &record)
  if err != nil {
    log.Fatal(err)
  }
  return record, err
}

func processTransactions(r io.Reader, write bool) error {
  scanner := bufio.NewScanner(r)
  var records []JsonRecord
  for scanner.Scan() {
    line := []byte(scanner.Text())
    record, err := processTransaction(line)
    if err != nil {
      return err
    } else {
      records = append(records, record)
    }
  }
  for index := range records {
    record := records[index]
    cacheItem := staticCache.Put(record.Key, record.Tokens)
    if write {
      writeTransaction(cacheItem)
    }
  }
  return nil
}

func replayTransactionLog() {
    if _, err := os.Stat(*fileName); !os.IsNotExist(err) {
      file, err := os.Open(*fileName)
      if err != nil {
        log.Fatal(err)
      } else {
        processTransactions(file, false)
      }
    }
}

func openTransactionLog() {
  if _, err := os.Stat(*fileName); os.IsNotExist(err) {
    transactionLog, _ = os.Create(*fileName)
  } else {
    file, err := os.OpenFile(*fileName, os.O_RDWR|os.O_APPEND, 0660);
    if err != nil {
      log.Fatal(err)
    } else {
      transactionLog = file
    }
  }
}

func writeTransaction(c *CacheItem) {
  if transactionLog != nil {
    encodedValue, err := json.Marshal(c)
    if err == nil {
      n, err := transactionLog.Write(encodedValue)
      if err == nil {
        log.Println("Wrote", n, "bytes")
        transactionLog.WriteString("\n")
      } else {
        log.Println(err.Error())
      }
    } else {
      log.Println(err.Error())
    }
  }
}
