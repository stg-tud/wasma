mkdir -p $HOME/wasma
mkdir -p $HOME/wasma/data
mkdir -p $HOME/wasma/bin
mkdir -p $HOME/wasma/analyses
docker build -t wasma .
docker create -it --name wasma -v $HOME/wasma/data:/usr/wasma/data wasma