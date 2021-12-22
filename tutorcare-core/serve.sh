#!/bin/sh
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
ng serve &
gin --port 4201 --path . --build ./src/main/ --i --all &

wait