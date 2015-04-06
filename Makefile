all:
	go build
run:
	./infor-you-mation -alsologtostderr=true
test:
	go test ./...
