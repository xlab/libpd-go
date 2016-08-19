all:
	cgogen core.yml

clean:
	rm -f core/cgo_helpers.go core/cgo_helpers.c core/cgo_helpers.h core/doc.go core/types.go core/const.go
	rm -f core/core.go

test:
	cd core && go build
	cd core && go test
	go test
