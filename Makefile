NAME = millipede-go
SRC = cmd/millipede-go
PACKAGES = .


GOCMD ?=        go
GOBUILD ?=      $(GOCMD) build
GOCLEAN ?=      $(GOCMD) clean
GOINSTALL ?=    $(GOCMD) install
GOTEST ?=       $(GOCMD) test
GOFMT ?=        gofmt -w


BUILD_LIST = $(foreach int, $(SRC), $(int)_build)
CLEAN_LIST = $(foreach int, $(SRC) $(PACKAGES), $(int)_clean)
INSTALL_LIST = $(foreach int, $(SRC), $(int)_install)
IREF_LIST = $(foreach int, $(SRC) $(PACKAGES), $(int)_iref)
TEST_LIST = $(foreach int, $(SRC) $(PACKAGES), $(int)_test)
BENCH_LIST = $(foreach int, $(SRC) $(PACKAGES), $(int)_bench)
TRAVIS_LIST = $(foreach int, $(SRC) $(PACKAGES), $(int)_travis)
COVER_LIST = $(foreach int, $(PACKAGES), $(int)_cover)
FMT_LIST = $(foreach int, $(SRC) $(PACKAGES), $(int)_fmt)


.PHONY: $(CLEAN_LIST) $(TEST_LIST) $(FMT_LIST) $(INSTALL_LIST) $(BUILD_LIST) $(IREF_LIST) $(BENCH_LIST) $(TRAVIS_LIST) $(COVER_LIST)


all: build


.PHONY: build
build: $(BUILD_LIST)
clean: $(CLEAN_LIST)
install: $(INSTALL_LIST)
test: $(TEST_LIST)
bench: $(BENCH_LIST)
travis: $(TRAVIS_LIST)
cover:
	rm -f profile.out
	$(MAKE) $(COVER_LIST)
	echo "mode: set" | cat - profile.out > profile.out.tmp && mv profile.out.tmp profile.out
iref: $(IREF_LIST)
fmt: $(FMT_LIST)


$(BUILD_LIST): %_build: %_fmt %_iref
	$(GOBUILD) -o $(NAME) ./$*
$(CLEAN_LIST): %_clean:
	$(GOCLEAN) ./$*
$(INSTALL_LIST): %_install:
	$(GOINSTALL) ./$*
$(IREF_LIST): %_iref:
	$(GOTEST) -i ./$*
$(TEST_LIST): %_test:
	$(GOTEST) ./$*
$(BENCH_LIST): %_bench:
	$(GOTEST) -bench . ./$*
$(TRAVIS_LIST): %_travis:
	$(GOTEST) -v ./$*
$(COVER_LIST): %_cover:
	$(GOTEST) -coverprofile=file-profile.out ./$*
	if [ -f file-profile.out ]; then cat file-profile.out | grep -v "mode: set" >> profile.out || true; rm -f file-profile.out; fi
$(FMT_LIST): %_fmt:
	$(GOFMT) ./$*


.PHONY: re
re: clean all


.PHONY: dist
dist:
	docker build -t $(NAME)-builder .
	@docker rm $(NAME)-builder 2>/dev/null || true
	mkdir -p dist
	docker run --name=$(NAME)-builder $(NAME)-builder tar -cf - /etc/ssl > dist/ssl.tar
	docker cp $(NAME)-builder:/go/bin tmp
	docker rm $(NAME)-builder
	touch tmp/bin/*
	mv tmp/bin/* dist
	rm -rf tmp
