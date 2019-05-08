build-windows:
	go build -ldflags -H=windowsgui -o build/randpaper.exe ./...

create-icon:
	go get github.com/cratonica/2goarray
	2goarray Icon main < icon.ico > icon.go
