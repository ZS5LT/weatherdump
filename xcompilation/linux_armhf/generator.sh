#!/bin/bash

PACKAGE_NAME=weatherdump-cli-linux-armhf

mkdir -p ./dist ./dist/$PACKAGE_NAME ./dist/export
CXX=arm-linux-gnueabihf-g++ CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -o dist/$PACKAGE_NAME/weatherdump ./main.go
cd ./dist && tar --xform s:'./':: -czvf $PACKAGE_NAME.tar.gz ./$PACKAGE_NAME
cd - && mv ./dist/$PACKAGE_NAME.tar.gz ./dist/export