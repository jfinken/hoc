all:
	goyacc -o hoc.go -p Hoc hoc.y
	go build 
	cp hoc ${GOPATH}/bin/hoc

