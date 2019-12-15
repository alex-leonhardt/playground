# Playground

## prom/pushgateway

```
docker run --rm -p 9091:9091 prom/pushgateway
```

then

http://localhost:9091

## Run

```
go test -v ./... -test.run TestMain/MyTest -count=1
```
