name=$1

mkdir -p $HOME/wasma/analyses/$name

containerId=$(docker ps -aqf "name=wasma")
docker cp $containerId:/usr/wasma/cmd/$name/$name.go $HOME/wasma/analyses/$name/$name.go