$Env:GOOS = "linux"; $Env:GOARCH = "amd64"
go build -o dist/v2zf_linux_amd64

$Env:GOOS = "linux"; $Env:GOARCH = "arm64"
go build -o dist/v2zf_linux_arm64

$Env:GOOS = "windows"; $Env:GOARCH = "amd64"
go build -o dist/v2zf_windows_amd64.exe

$Env:GOOS = "darwin"; $Env:GOARCH = "amd64"
go build -o dist/v2zf_darwin_amd64