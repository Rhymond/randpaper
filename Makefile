build-windows:
	go build -ldflags -H=windowsgui -o build/randpaper.exe main.go

.PHONY: build-windows