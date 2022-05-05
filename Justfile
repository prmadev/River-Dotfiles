build:
	go build -o init .
restart:
	./init
push:
	git add .
	comet
	git push
