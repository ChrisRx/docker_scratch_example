deps:
	@go get -u github.com/labstack/echo

build:
	@go build -o tiny-server server.go

build-cgo:
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tiny-server server.go

docker:
	@docker build -t tiny-server .
	@docker run -p 8080:8080 tiny-server
