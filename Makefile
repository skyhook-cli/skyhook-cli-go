all: macos_bin linux_bin windows_bin

macos_bin: src/main.go
	cd src && env GOOS=darwin GOARCH=amd64 go build -o ../bin/macos/skyhook

linux_bin: src/main.go
	cd src && env GOOS=linux GOARCH=amd64 go build -o ../bin/linux/skyhook

windows_bin: src/main.go
	cd src && env GOOS=windows GOARCH=amd64 go build -o ../bin/windows/skyhook.exe

format:
	gofmt -w .

clean:
	rm -rf bin
	rm -rf Tetherfile **/Tetherfile
