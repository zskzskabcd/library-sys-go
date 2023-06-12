go build -o out/lib-sys-go -trimpath -ldflags="-s -w" ../main.go 

# windows
GOOS=windows GOARCH=amd64 go build -o out/lib-sys-go-windows.exe -trimpath -ldflags="-s -w" ../main.go
