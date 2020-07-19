#!/bin/bash
repository=$1
ref=$2
commit_sha=$3

CONTAINER_NAME=memezis
IMAGE_NAME=docker.pkg.github.com/cherya/memezis/memezis:latest

docker pull $IMAGE_NAME
docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME
docker run --mount type=bind,source=$(pwd)/production.env,target=/app/production.env --name $CONTAINER_NAME -d --net=host -d $IMAGE_NAME

if [ ! "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
  msg="Update failed for *${repository}* \nref: *${ref}* \ncommit: (${commit_sha})[https://github.com/cherya/memezis/commit/${commit_sha}]"
  bash scripts/notify.sh "$msg"
else
  msg="Successfully updated *${repository}* \nref: *${ref}* \ncommit: [${commit_sha}](https://github.com/cherya/memezis/commit/${commit_sha})"
  bash scripts/notify.sh "$msg"
fi