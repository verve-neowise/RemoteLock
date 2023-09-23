# build bundle
GOARCH=arm64 GOOS=linux go build -o remote-lock src/*
# push to device
adb push remote-lock /data/local/tmp