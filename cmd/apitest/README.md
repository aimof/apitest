# Using docker

Requirements: docker-compose 18.06.0+

```sh
# before building docker you must build apitest with env CGI_ENABLED=0
# Set GOOS and GOARCH if you want.
CGO_ENABLED=0 go build -o apitest
docker build .
```