# AnyWind
Run libraries on your machine using anywind commands
# Quick start
## GO
You can try anywind using **go**, in order to install it you can run:
```sh
cd /cmd # Go to anywind to build binaries
go install anywind # Install anywind using go
anywind # this will run the anywind command
```
>Make sure go is installed on your host machine.
## Docker
You can try to run this project using docker, you will first have to build the image:
```sh
docker build -t anywind -f ./docker/dockerfile ./cmd
```
then run:
```sh
docker run anywind get --all # this will get you all the libs in libs.json
```
