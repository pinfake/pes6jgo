docker run --rm -e CGO_ENABLED=0 -v "%~dp0:/go/src/github.com/pinfake/pes6go" golang /bin/bash -c "go get github.com/pinfake/pes6go/... && go build -ldflags '-linkmode external -extldflags -static' -o /go/src/github.com/pinfake/pes6go/bin/pes6go /go/src/github.com/pinfake/pes6go/main/pes6go.go"