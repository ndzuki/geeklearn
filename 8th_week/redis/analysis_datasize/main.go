package main

import (
	"context"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func InitRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	return client, client.Ping(context.TODO()).Err()
}

func main() {
	dataSize := []int{
		10000,
		50000,
		100000,
		200000,
		300000,
		500000,
		1000000,
	}

	ctx := context.Background()
	wg := &sync.WaitGroup{}
	rdb, err := InitRedisClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, s := range dataSize {
		rdb.FlushDB(ctx)

		before, err := rdb.Info(ctx, "memory").Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		for i := 0; i < s; i++ {
			wg.Add(1)
			go func(data int) {
				rdb.Set(ctx, uuid.NewString(), 0, 0)
				wg.Done()
			}(i)
		}
		wg.Wait()

		after, err := rdb.Info(ctx, "memory").Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		ioutil.WriteFile(
			"./analysis_results_"+strconv.Itoa(s),
			[]byte("--->Before<---"+"\n"+before+"\n\n"+"--->After<---"+"\n"+after),
			0644,
		)

		rdb.FlushDB(ctx)
		time.Sleep(time.Second * 3)
	}

	defer rdb.Close()

}
