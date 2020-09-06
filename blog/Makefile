
setup: run

bundle:
	go-bindata -o=assets/bindata.go --pkg=assets templates/... markdown/...

run:
	export HTTP_PORT=8099 && \
	go run .

build:
	export GOOS=linux && \ 
	go build -o ./blog .