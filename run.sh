
#!/bin/bash
container_name=template-echo-golang
local=
docker stop $container_name
docker rm $container_name
docker rmi $(docker images $container_name -q)
docker build -t $container_name:$1 . || exit 1
docker run --name $container_name \
    -d \
    --restart always \
    -v $local:/go/src/template-echo-golang/logs \
    -it \
    -p 8030:1323 \
    $container_name:$1 || exit 1