NAME = terraform-provider-convox
VERSION := v0.1.3
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
XC_OSARCH := darwin/amd64 linux/amd64
BUILD_DIR ?= build
DIST_DIR ?= dist

bindir ?= ~/.terraform.d/plugins

all: deps build

build: XC_OSARCH=$(GOOS)/$(GOARCH)
build: build-all

dev: bindir=./terraform
dev: install

install: $(BUILD_DIR)/$(GOOS)_$(GOARCH)/$(NAME)
	install -m 755 $< $(DESTDIR)$(bindir)/$(NAME)

deps:
	go get -u github.com/mitchellh/gox
	go get -u github.com/github/hub
	glide install

#$(addsuffix /$(NAME), $(addprefix $(BUILD_DIR)/, $(subst /,_,$(XC_OSARCH)))): $(BUILD_DIR)/%/$(NAME):
	#$*
	#$@

build-all:
	-rm -rf dist
	gox \
		-ldflags "-X main.version=${VERSION}" \
		-osarch="${XC_OSARCH}" \
		-output="${BUILD_DIR}/{{.OS}}_{{.Arch}}/${NAME}"

clean:
	-rm -rf $(BUILD_DIR)
	-rm -rf $(DIST_DIR)

ARTIFACTS := $(shell find $(BUILD_DIR) -mindepth 1 -maxdepth 1 -type d)
dist:
	mkdir -p ${DIST_DIR}
	$(foreach dir, $(ARTIFACTS), \
		tar -zcf "${DIST_DIR}/${NAME}_${VERSION}_$(shell basename ${dir}).tgz" -C $(dir) $(NAME);)

release:
	git log --no-merges \
		--format='%C(auto,green)* %s%C(auto,reset)%n%w(0,2,2)%+b' \
		--reverse "$(git describe --abbrev=0 --tags)..HEAD" \
		| hub release create --draft -m ${VERSION} -F - "${VERSION}" $(foreach file,$(wildcard ${DIST_DIR}/*), -a $(file))

.PHONY: all build clean
