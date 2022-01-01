#!/bin/sh
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
ng serve --proxy-config proxy.config.json