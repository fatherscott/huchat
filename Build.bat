SET "GOOS=linux"
SET "GOARCH=amd64"
go build -o HuChat main.go

SET "GOOS=windows"
SET "GOARCH=amd64"
go build -o HuChat.exe main.go

SET "GOOS=darwin"
SET "GOARCH=amd64"
go build -o HuChat.EXEC main.go

pause