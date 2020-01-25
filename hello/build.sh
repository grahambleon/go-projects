#! /bin/bash

rm -rf bin
mkdir bin
cd src
go build hello.go
mv hello ../bin/hello
cd ..
echo "bin/hello" > run.sh
chmod 777 run.sh
