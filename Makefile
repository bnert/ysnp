build-release:
	mkdir -p ./dist
	go build -o ./dist/ysnp *.go
