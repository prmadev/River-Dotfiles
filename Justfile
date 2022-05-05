build:
	go build -o init .
restart:
	./init
push:
	git add .
	git cz
	git push
