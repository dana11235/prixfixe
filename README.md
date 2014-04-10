PrixFixe
========

Simple in-memory prefix search server. Uses a Trie as the underlying
implementation. It is designed to be used for autocomplete, since ActiveRecord
is too slow to do this quickly and efficiently.

Right now, it allows you to set/get keys, and prefix search.

## Running The Server
You can use optional parameters -port and -file
    go run src/main.go -file={test.json} -port={8080}

The server also takes an optional "auth_key" flag, which will require all write requests to have an auth header with this value

## Making Requests
Here's an example of a put (without authentication)
    curl "http://localhost:8080/put?key=dana+jonez&value=\{\"email\":\"DanaBWinston@armyspy.com\",\"first_name\":\"Dana\",\"id\":\"61753\",\"last_name\":\"Jonez\"\}"

Here's an example of a put with auth
    curl -H "XPrixfixeAuth: asdf" "http://localhost:8080/put?key=dana+jonez&value=\{\"email\":\"DanaBWinston@armyspy.com\",\"first_name\":\"Dana\",\"id\":\"61753\",\"last_name\":\"Jonez\"\}"

## Running The Tests
    go test prixefixe
