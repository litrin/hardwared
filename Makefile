INSTALL_PATH=/usr/bin
APPLICATION=hardwared
BUILDTAGS=./main.go

all:
	go build -tags "$(BUILDTAGS)" -o $(APPLICATION)

install:
	cp $(APPLICATION) $(INSTALL_PATH)

clean:
	rm $(APPLICATION)