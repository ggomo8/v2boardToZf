SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o dist/v2zf_linux_amd64


SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o dist/v2zf_linux_arm64

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o dist/v2zf_windows_amd64.exe

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o dist/v2zf_darwin_amd64