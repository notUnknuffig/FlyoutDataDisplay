GOOS=windows GOARCH=amd64 go build -o ./out/MFD.exe -ldflags "-H=windowsgui" ./cmd/client
GOOS=windows GOARCH=amd64 go build -o "./out/MFD (Debug).exe" -ldflags "-H=windowsgui" ./cmd/client