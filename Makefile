# Project-specific variables
BINARIES ?=	millipede-fuse millipede-go millipede-http
CONVEY_PORT ?=	9042

# Common variables
rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))
SOURCES :=      $(call rwildcard,./cmd/ ./pkg/,*.go) glide.lock
GOENV ?=	GO15VENDOREXPERIMENT=1
GO ?=		$(GOENV) go
USER ?=		$(shell whoami)

all:	build


.PHONY: build
build:	$(BINARIES)


$(BINARIES):	$(SOURCES)
	$(GO) build -i -o $@ ./cmd/$@


.PHONY: test
test:
	#$(GO) get -t .
	$(GO) test -i .
	$(GO) test -v .


.PHONY: clean
clean:
	rm -f $(BINARIES)


.PHONY: re
re:	clean all


.PHONY: convey
convey:
	$(GO) get github.com/smartystreets/goconvey
	goconvey -cover -port=$(CONVEY_PORT) -workDir="$(realpath .)" -depth=1


.PHONY:	cover
cover:	profile.out


profile.out:	$(SOURCES)
	rm -f $@
	$(GO) test -covermode=count -coverpkg=. -coverprofile=$@ .


.PHONY: docker-build
docker-build:
	$(GO) get github.com/laher/goxc
	rm -rf contrib/docker/linux_386
	for binary in $(BINARIES); do                                             \
	  goxc -bc="linux,386" -d . -pv contrib/docker -n $$binary xc;            \
	  mv contrib/docker/linux_386/$$binary contrib/docker/entrypoint;         \
	  docker build -t $(USER)/$$binary contrib/docker;                        \
	  docker run -it --rm $(USER)/$$binary || true;                           \
	  docker inspect --type=image --format="{{ .Id }}" moul/$$binary || true; \
	  echo "Now you can run 'docker push $(USER)/$$binary'";                  \
	done
