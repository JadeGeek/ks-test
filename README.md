# Duckpass Project (SUNTECH) 

## Upload Service


```
✓ usage: make [target]

build-no-cache                 - Build the dpass upload docker image based on scratch with no cache
build                          - Build the dpass upload docker image based on scratch
help                           - Show help message
ls                             - List 'dpass-upload' docker images
push-to-azure                  - Push docker image to azurecr.io container registry
push-to-gcp                    - Push docker image to gcr.io container registry
run                            - Run the dpass upload docker image based on scratch
```

## Quickstart 

```
make build && make run
```

## Local Dev
```
skaffold dev
```

## Tests
```
curl -X POST http://192.168.16.96/upload -F "file=@test.png"  -H "Content-Type: multipart/form-data"
```
