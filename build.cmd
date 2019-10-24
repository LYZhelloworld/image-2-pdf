@echo off
set GO111MODULE=on
go build -o bin/image2pdf.exe cmd/image2pdf/main.go
