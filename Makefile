all:
	go build
run: all
	./infor-you-mation -alsologtostderr=true
