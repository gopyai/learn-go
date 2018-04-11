Don't build go binary as follow. It won't work, need to compile without dynamic library.
	go build -o main .

Do as follow instead. It uses static library.
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


This go example requires certificate to run. Copy from golang image as follow:
Run the image:
	docker run --rm -v "$PWD/certs":/tmp -w /tmp -it golang 
and then execute:
	cp /etc/ssl/certs/ca-certificates.crt .
It will copy certificate into certs directory.


Build docker image and run it.
	docker build -t arief/smallgo .
	docker run --rm arief/smallgo