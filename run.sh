#!/bin/bash
container_name=
dockerfile=
local=
port=

docker stop $container_name
docker rm $container_name
docker rmi $(docker images $container_name -q)
docker build -f $dockerfile -t $container_name:1.0.0 . || exit 1
docker run --name $container_name \
    -d \
    --restart always \
    -v $local:/go/src/template-echo-golang/logs \
    -it \
    -p $port:1323 \
    $container_name:1.0.0 || exit 1