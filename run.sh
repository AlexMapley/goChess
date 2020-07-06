#!/bin/sh

docker build . --tag goChess

docker run -ti goChess