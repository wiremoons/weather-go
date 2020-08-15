#
#	Makefile for Go Language code
#
# set the default build target name:
.DEFAULT_GOAL := defTarget

# define binary output file name and path:
SRC=*.go 
OUTNAME=bin/weather-go
# Go compiler settings
CC=go build
# normal build flags
CFLAGS=-i
# release build flags
RELCFLAGS=-i -ldflags="-s -w"
# small build flags (similar to release)
SMLCFLAGS=-i -gcflags=all=-dwarf=false -ldflags="-s -w"
# build and run flags
RFLAGS=run
#
# To build for Linux 32bit ARM
ARM32=GOOS=linux GOARCH=arm
# To build for Linux 64bit ARM aarch64
ARM64=GOOS=linux GOARCH=arm64
# To build for Linux 32bit
LIN32=GOOS=linux GOARCH=386
# To build for Linux 64bit
LIN64=GOOS=linux GOARCH=amd64
# To build Windows 32 bit version:
WIN32=GOOS=windows GOARCH=386
# To build Windows 64 bit version:
WIN64=GOOS=windows GOARCH=amd64
# To build Mac OS X 32 bit version:
MAC32=GOOS=darwin GOARCH=386
# To build Mac OS X 64 bit version:
MAC64=GOOS=darwin GOARCH=amd64
# To build FreeBSD 64 bit version:
FREE64=GOOS=freebsd GOARCH=amd64

defTarget: $(SRC)
	$(CC) -o $(OUTNAME) $(CFLAGS) $(SRC)

release: $(SRC)
	$(CC) -o $(OUTNAME) $(RELCFLAGS) $(SRC)
	upx $(OUTNAME)

shared:  $(SRC)
	$(CC) -o $(OUTNAME)-shared $(RELCFLAGS) -linkshared $(SRC)

nodebug: $(SRC)
	$(CC) -o $(OUTNAME)-nodebug $(SMLCFLAGS) $(SRC)

arm32: $(SRC)
	$(ARM32) $(CC) -o $(OUTNAME)-arm32 $(RELCFLAGS) $(SRC)
	upx $(OUTNAME)-arm32
	
arm64: $(SRC)
	$(ARM64) $(CC) -o $(OUTNAME)-aarch64 $(RELCFLAGS) $(SRC)
	upx $(OUTNAME)-aarch64

lin32: $(SRC)
	$(LIN32) $(CC) -o $(OUTNAME)-x386 $(RELCFLAGS) $(SRC)
	upx $(OUTNAME)

lin64: $(SRC)
	$(LIN64) $(CC) -o $(OUTNAME)-amd64 $(RELCFLAGS) $(SRC)
	upx $(OUTNAME)-amd64

win32: $(SRC)
	$(WIN32) $(CC) -o $(OUTNAME)-x386.exe $(RELCFLAGS) $(SRC)
	upx $(OUTNAME)-x386.exe

win64: $(SRC)
	$(WIN64) $(CC) -o $(OUTNAME)-x64.exe $(RELCFLAGS) $(SRC)
	upx $(OUTNAME)-x64.exe

mac32: $(SRC)
	$(MAC32) $(CC) -o $(OUTNAME)-mac386 $(RELCFLAGS) $(SRC)

mac64: $(SRC)
	$(MAC64) $(CC) -o $(OUTNAME)-macx64 $(RELCFLAGS) $(SRC)
	upx $(OUTNAME)-macx64

free64: $(SRC)
	$(FREE64) $(CC) -o $(OUTNAME)-freebsd64 $(RELCFLAGS) $(SRC)

run: $(SRC)
	$(CC) $(RFLAGS) $(SRC)

clean:
	rm $(OUTNAME) $(OUTNAME)-arm32 $(OUTNAME)-aarch64 $(OUTNAME)-x64.exe $(OUTNAME)-x386.exe $(OUTNAME)-amd64 $(OUTNAME)-x386 $(OUTNAME)-macx64 $(OUTNAME)-mac386 $(OUTNAME)-freebsd64

all: arm32 arm64 lin64 win64 mac64 free64
