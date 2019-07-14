# Mower exo with golang

* author : thomas rayez
* date : 14/07/2019

## pré-requis 

* golang 1.12
* docker 

## Init

In the root directorie run 

```
export GOPATH="$(pwd)"
```

## Run 

``` 
go run src/exo/main.go src/ressource/in.dat 
```

## Run Test 

```
go test exo/orientation exo/mower exo/parser
```

## Build 

```
# Build
go build -o bin/mower src/exo/main.go
# RUN
bin/mower src/ressource/in.dat
```

## Package

```
docker build -t mower:latest .

docker run -ti -v src/ressource:/input mower:latest /usr/bin/mower /input/in.dat 
```


