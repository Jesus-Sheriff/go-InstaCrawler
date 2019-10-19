# Go parameters
#Plantilla de: https://sohlich.github.io/post/go_makefile/
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u -v
BINARY_NAME=go-InstaCrawler
BINARY_UNIX=$(BINARY_NAME)_unix
# INSTAGRAM_USERNAME=_nombre_usuario_
# INSTAGRAM_PASSWORD=_contraseña_usuario_
GOPATH=$(HOME)/go
#PATH=$(PATH):/usr/local/go/bin
EXPORT=export

all: test run
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: deps
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run: deps
	$(GORUN) goinsta.v2/examples/show-latest-image/main.go
deps:
	$(GOGET) github.com/ahmdrz/goinsta
variables:
	$(EXPORT) $(GOPATH)
	$(EXPORT) $(PATH):/usr/local/go/bin 
	echo "GOPATH: " $(GOPATH) 
	echo "PATH: " $(PATH)
user:
	$(EXPORT) $(INSTAGRAM_USERNAME)
	$(EXPORT) $(INSTAGRAM_PASSWORD)
	echo "Usuario actual: " $(INSTAGRAM_USERNAME)

