OUT=sfs

run:
	go run main.go

build: clean
	GOOS=linux GOARCH=amd64 go build -o bin/$(OUT)-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/$(OUT)-darwin-amd64

clean:
	find bin ! -name '.gitkeep' -type f -exec rm {} +
