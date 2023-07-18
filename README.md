# sequence

A package to generate sequential numbers


## Install

``` shell
git clone git@github.com:kzrl/sequence.git
cd sequence
```

## Run sequenced
```
# optionally, set a port with
# EXPORT SEQUENCED_PORT=":8081" 
# (it needs the colon in the string for simplicity, a real world tool would have friendlier error handling)
go run cmd/sequenced/main.go
```

## Tests

``` shell
go test -race ./...
```

## Usage

``` shell
curl localhost:8081 && echo
```


