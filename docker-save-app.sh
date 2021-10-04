name=$1

containerId=$(docker ps -aqf "name=wasma")
docker cp $containerId:/usr/wasma/bin/$name $HOME/wasma/bin/$name