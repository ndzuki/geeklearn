module cache

go 1.17

replace lru => ../lru

require (
	consistenthash v0.0.0-00010101000000-000000000000
	lru v0.0.0-00010101000000-000000000000
	singleflight v0.0.0-00010101000000-000000000000
)

replace consistenthash => ../consistenthash

replace singleflight => ../singleflight
