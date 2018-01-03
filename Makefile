# Variables
BUILD_DIR 		:= dist
GITHASH 			:= $(shell git rev-parse HEAD)
LINT_PATHS		:= ./cmd/... achievements.go
FORMAT_PATHS 	:= ./cmd achievements.go
DEADLINE 			:= 60s

# Compilation variables
CC 						:= go build
DFLAGS 				:= -race
CFLAGS 				:= -X github.com/FlorentinDUBOIS/achievements/cmd.githash=$(GITHASH)
CROSS					:= GOOS=linux GOARCH=amd64

# Makefile variables
VPATH 				:= $(BUILD_DIR)

# Function definitions
rwildcard			:= $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

.SECONDEXPANSION:
.PHONY: all
all: init format lint release

.PHONY: init
init:
	go get -u github.com/golang/dep
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install --update
	dep ensure

.PHONY: format
format:
	gofmt -w -s $(FORMAT_PATHS)

.PHONY: lint
lint:
	gometalinter --vendor --deadline=$(DEADLINE) --disable-all --enable=gofmt $(LINT_PATHS)
	gometalinter --vendor --deadline=$(DEADLINE) --disable-all --enable=vet $(LINT_PATHS)
	gometalinter --vendor --deadline=$(DEADLINE) --disable-all --enable=golint $(LINT_PATHS)
	gometalinter --vendor --deadline=$(DEADLINE) --disable-all --enable=ineffassign $(LINT_PATHS)
	gometalinter --vendor --deadline=$(DEADLINE) --disable-all --enable=misspell $(LINT_PATHS)
	gometalinter --vendor --deadline=$(DEADLINE) --disable-all --enable=staticcheck $(LINT_PATHS)
	gometalinter --vendor --deadline=$(DEADLINE) --disable-all --enable=gas $(LINT_PATHS)

.PHONY: dev
dev: format lint build

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: build
build: achievements.go $$(call rwildcard, ./, *.go)
	$(CC) $(DFLAGS) -ldflags "$(CFLAGS)" -o $(BUILD_DIR)/achievements.build achievements.go

.PHONY: release
release: achievements.go $$(call rwildcard, ./, *.go)
	$(CC) -ldflags "-s -w $(CFLAGS)" -o $(BUILD_DIR)/achievements.release achievements.go

.PHONY: dist
dist: achievements.go $$(call rwildcard, ./, *.go)
	$(CROSS) $(CC) -ldflags "-s -w $(CFLAGS)" -o $(BUILD_DIR)/achievements.dist achievements.go
