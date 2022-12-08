default: test

test: 
	cd calculator && go test -v ./...

build:
	cd cmd/cal && go build -o cal