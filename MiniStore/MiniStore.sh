#!/bin/bash

docker pull ethereum/solc:0.6.0
docker run --rm -v $(pwd):/root ethereum/solc:0.6.0 --abi /root/MiniStore.sol --bin /root/MiniStore.sol -o /root --overwrite
abigen --bin=./MiniStore.bin --abi=./MiniStore.abi --pkg=MiniStore --out=MiniStore.go
