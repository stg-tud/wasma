name=$1

mkdir -p $HOME/wasma/bin/$name

containerId=$(docker ps -aqf "name=wasma")
docker cp $HOME/wasma/bin/$name.go $containerId:/usr/wasma/bin/$name.go