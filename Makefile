all:
	go get github.com/chzyer/readline
	go get golang.org/x/tools/cmd/goyacc
	goyacc -o hoc.go -p Hoc hoc.y
	go build 
	cp hoc ${GOPATH}/bin/hoc

