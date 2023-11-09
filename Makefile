build:
	GOOS='windows' GOARCH='amd64' go build .
	go build .

clean:
	rm -rf ./influxmonitor
	rm -rf ./influxmonitor.exe

tidy:
	go mod tidy

.PHONY: clean build tidy