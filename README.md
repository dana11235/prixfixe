PrixFixe
========

Simple in-memory prefix search server. Uses a Trie as the underlying
implementation. It is designed to be used for autocomplete, since ActiveRecord
is too slow to do this quickly and efficiently.

Right now, it just allows you to set/get keys, and prefix search. Hopefully more later.

## Running The Server
You can use optional parameters -port and -file
    go run src/main.go -file={test.json} -port={8080}

## Running The Tests 
    go test prixefixe
