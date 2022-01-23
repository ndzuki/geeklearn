# How to use redis-benchmark

Redis 自带了一个叫 redis-benchmark 的工具来模拟 N 个客户端同时发出 M 个请求。
```shell
Usage: redis-benchmark [-h <host>] [-p <port>] [-c <clients>] [-n <requests]> [-k <boolean>]

 -h <hostname>      Server hostname (default 127.0.0.1)
 -p <port>          Server port (default 6379)
 -s <socket>        Server socket (overrides host and port)
 -a <password>      Password for Redis Auth
 -c <clients>       Number of parallel connections (default 50)
 -n <requests>      Total number of requests (default 100000)
 -d <size>          Data size of SET/GET value in bytes (default 2)
 -dbnum <db>        SELECT the specified db number (default 0)
 -k <boolean>       1=keep alive 0=reconnect (default 1)
 -r <keyspacelen>   Use random keys for SET/GET/INCR, random values for SADD
  Using this option the benchmark will expand the string __rand_int__
  inside an argument with a 12 digits number in the specified range
  from 0 to keyspacelen-1. The substitution changes every time a command
  is executed. Default tests use this to hit random keys in the
  specified range.
 -P <numreq>        Pipeline <numreq> requests. Default 1 (no pipeline).
 -q                 Quiet. Just show query/sec values
 --csv              Output in CSV format
 -l                 Loop. Run the tests forever
 -t <tests>         Only run the comma separated list of tests. The test
                    names are the same as the ones produced as output.
 -I                 Idle mode. Just open N idle connections and wait.
```

## 一般这样启动测试
```shell
redis-benchmark -q -n 100000
```

## 只运行一些测试用例的子集
你不必每次都运行 redis-benchmark 默认的所有测试。
使用 -t 参数可以选择你需要运行的测试用例，比如下面的范例:
```shell
redis-benchmark -t set,lpush -n 100000 -q
```

你不必每次都运行 redis-benchmark 默认的所有测试。
使用 -t 参数可以选择你需要运行的测试用例，比如下面的范例:
```shell
redis-benchmark -n 100000 -q script load "redis.call('set','foo','bar')"
```

## 选择测试键的范围大小
默认情况下面，基准测试使用单一的 key。
在一个基于内存的数据库里， 单一 key 测试和真实情况下面不会有巨大变化。
当然，使用一个大的 key 范围空间， 可以模拟现实情况下面的缓存不命中情况。

这时候我们可以使用 -r 命令。
比如，假设我们想设置 10 万随机 key 连续 SET 100 万次，我们可以使用下列的命令:
```bash
#cleanall
redis-cli flushall

#start benchmark test
redis-benchmark -t set -r 100000 -n 1000000

#checkout
redis-cli dbsize
```
## 使用 pipelining
默认情况下，每个客户端都是在一个请求完成之后才发送下一个请求 （benchmark 会模拟 50 个客户端除非使用 -c 指定特别的数量）， 这意味着服务器几乎是按顺序读取每个客户端的命令。Also RTT is payed as well.

真实世界会更复杂，Redis 支持 /topics/pipelining，使得可以一次性执行多条命令成为可能。 Redis pipelining 可以提高服务器的 TPS。 下面这个案例是在 Macbook air 11” 上使用 pipelining 组织 16 条命令的测试范例：
```shell
redis-benchmark -n 1000000 -t set,get -P 16 -q
```
**记得在多条命令需要处理时候使用** pipelining。

详细：http://www.redis.cn/topics/benchmarks.html

## 作业内容：

1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

docker 启动redis ，并写一个小脚本进行检验：
```yaml
version: '3'
services:
  redis:
    image: redis:6.0
    container_name: redis
    tty: true
    volumes:
      - /home/nd/go/github.com/NDzuki/geeklearn/8th_week/redis/benchmark_test.sh:/tmp/benchmark_test.sh
    # command: /bin/sh /tmp/benchmark_test.sh

```
```bash
#!/bin/env bash

set -eo

array_num=(10 20 50 100 200 1000 5000)

for i in ${array_num[@]}; do
    redis-benchmark -t set,get -d $i
    sleep 1
done
```
进入redis容器并执行脚本：
```shell
/bin/bash /tmp/benchmark_test.sh
```
2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。