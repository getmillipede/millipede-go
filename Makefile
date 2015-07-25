dist:
	docker build -t millipede-go-builder .
	@docker rm millipede-go-builder 2>/dev/null || true
	mkdir -p dist
	docker run --name=millipede-go-builder millipede-go-builder tar -cf - /etc/ssl > dist/ssl.tar
	docker cp millipede-go-builder:/go/bin tmp
	docker rm millipede-go-builder
	touch tmp/bin/*
	mv tmp/bin/* dist
	rm -rf tmp
