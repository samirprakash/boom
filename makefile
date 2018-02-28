VERSION=0.0.1
PATH_BUILD=build/
FILE_COMMAND=heft
FILE_ARCH=darwin_amd64

clean:
	@rm -rf ./build

build: clean
	@goxc \
	  -bc="darwin,amd64" \
	  -pv=$(VERSION) \
	  -d=$(PATH_BUILD) \
	  -build-ldflags "-X main.VERSION=$(VERSION)"

version:
	@echo $(VERSION)