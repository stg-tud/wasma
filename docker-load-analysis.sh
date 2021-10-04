name=$1

mkdir -p $HOME/wasma/bin/$name

containerId=$(docker ps -aqf "name=wasma")
docker cp $HOME/wasma/analyses/$name/$name.go $containerId:/usr/wasma/cmd/$name/$name.go