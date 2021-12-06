echo "Build for Linux"
GOOS=linux go build -o bin/backupmygit_linux main.go

echo "Build for MacOS"
GOOS=darwin go build -o bin/backupmygit_mac main.go

echo "Build for Windows"
GOOS=windows go build -o bin/backupmygit_win main.go