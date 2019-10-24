GO_BUILD := go build
GO_ENV = GO111MODULE=on
BINDIR := $(PWD)/bin
CMDDIR := $(PWD)/cmd
MAIN = main.go
BIN = image2pdf

.PHONY: all
all: build

.PHONY: build
build: $(CMDDIR)/$(BIN)/$(MAIN)
	$(GO_ENV) $(GO_BUILD) -o $(BINDIR)/$(BIN) $(CMDDIR)/$(BIN)/$(MAIN)

.PHONY: clean
clean:
	rm -rf $(BINDIR)
