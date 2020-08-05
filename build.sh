#!/bin/bash

set -e
set -x

module="go_project_name"

# 编译时使用vendor中的库文件，而不是~/gopath/pkg/mod/中的
export GOFLAGS=-mod=vendor
commitid=$(git rev-parse HEAD)
version=$(git rev-parse --abbrev-ref HEAD | grep -v HEAD || git describe --exact-match HEAD || git rev-parse HEAD)
time=$(date | tr ' ' '-')
flags="-X main.CommitID=${commitid} -X main.Version=${version} -X main.BuildTime=${time}"
go build -ldflags="$flags" -o ${module}

rm -rf output/
mkdir -p output/bin
mkdir -p output/conf
mkdir -p output/log

cp run.sh output/
cp ${module} output/bin/
cp -r conf/* output/conf/
