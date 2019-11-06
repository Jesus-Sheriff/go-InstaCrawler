# Go parameters
#Plantilla de: https://sohlich.github.io/post/go_makefile/
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -v
BINARY_NAME=go-InstaCrawler

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
run-pmgo: deps
	pmgo start goinsta.v2/examples/show-latest-image/ app
run-travis: deps
	pmgo start github.com/Jesus-Sheriff/go-InstaCrawler/goinsta.v2/examples/show-latest-image/ app
runcircle: deps
	pwd
	$(GORUN) goinsta.v2/examples/show-latest-image/main.go
	pmgo start ~/project/go-InstaCrawler/goinsta.v2/examples/show-latest-image/ app 
	pmgo start go-InstaCrawler/goinsta.v2/examples/show-latest-image/ app 
	pmgo stop app
	pmgo delete app
	cd ~/project/go-InstaCrawler/
	pmgo start goinsta.v2/examples/show-latest-image/ app
	pmgo stop app
	pmgo delete app
	cd ~
	pmgo start ~/project/go-InstaCrawler/goinsta.v2/examples/show-latest-image/ app
	pmgo stop app
	pmgo delete app
	# _/home/circleci/project/goinsta.v2/examples/show-latest-image
stop:
	pmgo stop app
	pmgo delete app
deps:
	$(GOGET) github.com/ahmdrz/goinsta
	$(GOGET) github.com/gorilla/mux
	$(GOGET) github.com/joho/godotenv
	$(GOGET) github.com/struCoder/pmgo
variables:
	. ./environment.sh

