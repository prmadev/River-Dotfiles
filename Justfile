build:
	go get -u
	go mod tidy
	go build -o init .
restart:
	./init
push:
	git add .
	git cz	
	git push
