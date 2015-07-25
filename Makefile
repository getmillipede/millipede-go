NAME = millipede-go


all: $(NAME)


$(NAME): $(shell find . -name "*.go")
	go build -o $@ .


clean:
	rm -f $(NAME)


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
