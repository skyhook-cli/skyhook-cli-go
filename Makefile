bins: macos_bin linux_bin windows_bin

macos_bin: main.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/macos/skyhook

linux_bin: main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/skyhook

windows_bin: main.go
	env GOOS=windows GOARCH=amd64 go build -o bin/windows/skyhook.exe

zips: macos_zip linux_zip windows_zip

macos_zip:
	cd bin/macos && zip -q ./skyhook-cli-go-macos-v$${VERSION_NUMBER:-local}-x64.zip skyhook

linux_zip:
	cd bin/linux && zip -q ./skyhook-cli-go-linux-v$${VERSION_NUMBER:-local}-x64.zip skyhook

windows_zip:
	cd bin/windows && zip -q ./skyhook-cli-go-windows-v$${VERSION_NUMBER:-local}-x64.zip skyhook.exe

format:
	gofmt -w .

clean:
	rm -rf bin
	rm -rf Tetherfile **/Tetherfile
