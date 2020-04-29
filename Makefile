_installmockgen:
	GO111MODULE=off go get github.com/golang/mock/mockgen
genmocks: _installmockgen
	mockgen -source=storage/storage.go -package mock -destination mock/storage.go

test: genmocks
	go test ./...
