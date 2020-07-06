#!/bin/sh

docker build . --tag gochess

docker run -ti gochess