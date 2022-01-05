package main

import (
	"fmt"
	"log"
	"net/http"

	geecache "github.com/NDzuki/geeklearn/geecache/cache"
)

var db = map[string]string{
	"Tom":  "530",
	"Jack": "385",
	"Sam":  "324",
}

func main() {
	geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := geecache.NewHTTPPool(addr)
	log.Println("cache is running at ", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
