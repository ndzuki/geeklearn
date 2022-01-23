#!/bin/env bash

set -eo

array_num=(10 20 50 100 200 1000 5000)

for i in ${array_num[@]}; do
    redis-benchmark -t set,get -d $i
    sleep 1
done