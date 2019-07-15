# Mower exo with golang - Multithread version

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
# Create docker image
docker build -t mower:latest .

# Run 
docker run -ti -v "$(pwd)/src/ressource:/input" mower:latest /usr/bin/mower /input.dat 

# Test Error
docker run -ti -v "$(pwd)/src/ressource:/input" mower:latest /usr/bin/mower /input/conflict.dat
docker run -ti -v "$(pwd)/src/ressource:/input" mower:latest /usr/bin/mower /input/oofr.dat

```


