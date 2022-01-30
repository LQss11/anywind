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
then to add new lib run:

```sh
docker run -it --rm --mount type=bind,source="$(pwd)"/cmd/libs.json,target=/libs.json anywind add --name lib3 --version 5.0 
```
to get lib you can run:
```sh
# this will get you all the libs in libs.json
docker run -it --rm --mount type=bind,source="$(pwd)"/cmd/libs.json,target=/libs.json anywind get --all
```

>Windows user update `"$(pwd)"` with `%cd%`