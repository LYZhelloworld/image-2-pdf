GO_ENV = GO111MODULE=on
MAIN = cmd/image2pdf/main.go

.PHONY: all build build_windows build_linux clean
all: build

build: build_windows build_linux

build_windows: $(MAIN)
	$(GO_ENV) GOOS=windows GOARCH=amd64 go build -o bin/image2pdf.exe $(MAIN)

build_linux: $(MAIN)
	$(GO_ENV) GOOS=linux GOARCH=amd64 go build -o bin/image2pdf $(MAIN)

clean:
	rm -rf bin/
