all:
	go build -o infor-you-mation
run:
	./infor-you-mation -alsologtostderr=true
debug:
	./infor-you-mation -alsologtostderr=true -v=5
test:
	go test ./...
