VERSION=0.0.1

clean:
	@rm -rf ./build

build: clean
	@goxc \
	  -bc="darwin,amd64" \
	  -pv=$(VERSION) \
	  -d=build \
	  -build-ldflags "-X main.VERSION=$(VERSION)"

version:
	@echo $(VERSION)
