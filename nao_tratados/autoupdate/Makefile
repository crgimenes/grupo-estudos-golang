all: linux windows zip 

linux:
	GOOS=linux GOARCH=amd64 go build -o launcher cmd/launcher/main.go
	GOOS=linux GOARCH=amd64 go build -o helper cmd/helper/main.go
	
windows:
	GOOS=windows GOARCH=amd64 go build -o launcher.exe cmd/launcher/main.go
	GOOS=windows GOARCH=amd64 go build -o helper.exe cmd/helper/main.go

zip:
	zip helper_windows.zip helper.exe
	zip helper_linux.zip helper

clean:
	rm -f launcher launcher.exe helper helper.exe helper_windows.zip helper_linux.zip

