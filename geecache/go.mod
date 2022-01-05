module geecache

go 1.17

replace cache => ./cache

replace lru => ./lru

require cache v0.0.0-00010101000000-000000000000

require lru v0.0.0-00010101000000-000000000000 // indirect
