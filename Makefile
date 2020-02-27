.PHONY: mocks install-mockgen

PROJECT_ROOT  = github.com/mfesenko/adventofcode

install-mockgen:
	GO111MODULE=on go get -v -u github.com/golang/mock/mockgen

mocks: install-mockgen
	$(call local_mockgen,2019/async,Executable)

test: mocks
	go test -cover ./...

define run_mockgen
  	mkdir -p $(2) && rm -rf $(2)/*
	echo "GOPATH = ${GOPATH}"
	echo "GOBIN = ${GOBIN}"
	echo "HOME = ${HOME}"
	echo "GOROOT = ${GOROOT}"
	echo "GOTOOLDIR = ${GOTOOLDIR}"
	ls -l "${GOROOT}/bin"
	ls -l "${HOME}/go/bin"
	ls -l "${GOBIN}"
	"${GOBIN}/mockgen" -version
  	"${GOBIN}/mockgen" -destination=$(2)/mocks.go $(1) $(subst $(semicolon),$(comma),$(3))
endef

define local_mockgen
	$(call run_mockgen,$(PROJECT_ROOT)/$(1),.mocks/$(1),$(2))
endef
