GOARCH=arm64 GOOS=linux go build -o go/getmodel go/getmodel.go
adb push go/getmodel /data/local/tmp 