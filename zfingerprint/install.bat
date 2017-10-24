set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1
set buildmode=c-shared 
go build -a -installsuffix cgo -o main .
pause
