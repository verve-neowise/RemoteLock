# build bundle
GOARCH=arm64 GOOS=$1 go build -o remote-lock src/*

echo "build for OS $1" 
