run:
	go run main.go

build:
	go build .

optmized:
	 CGO_ENABLED=0 go build -o hs -ldflags="-s -w" main.go

