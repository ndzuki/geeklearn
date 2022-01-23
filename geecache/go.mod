module geecache

go 1.17

replace cache => ./cache

replace lru => ./lru

require cache v0.0.0-00010101000000-000000000000

require (
	consistenthash v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	lru v0.0.0-00010101000000-000000000000 // indirect
)

replace consistenthash => ./consistenthash
