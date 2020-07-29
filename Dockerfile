FROM golang:1.14

WORKDIR /go/src/fibonacci
COPY . .

RUN cd .. \
	cd .. \
	cd bin \
	go install -v ./main.go 



CMD main
