# build bundle
GOARCH=arm64 GOOS=linux go build -o remote-lock src/remote-lock.go
# push to device
adb push remote-lock /data/local/tmp