#!/bin/sh
APP="git-hours"
# Build for each OS.
GOOS=linux GOARCH=amd64 go build -o ./bin/linux/${APP} git-hours.go timefunc.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows/${APP}.exe git-hours.go timefunc.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/${APP} git-hours.go timefunc.go

# Compress for Github Release upload
cd ./bin/linux/ && tar -zcvf ../${APP}_linux_x86-64.tgz . && cd -
cd ./bin/windows/ && tar -zcvf ../${APP}_windows_x86-64.tgz . && cd -
cd ./bin/darwin/ && tar -zcvf ../${APP}_darwin_x86-64.tgz . && cd -

# Cleanup
rm -rf ./bin/linux
rm -rf ./bin/windows
rm -rf ./bin/darwin
