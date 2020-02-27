.PHONY: mocks install-mockgen

PROJECT_ROOT  = github.com/mfesenko/adventofcode

install-mockgen:
	GO111MODULE=on go get -v -u github.com/golang/mock/mockgen
	GO111MODULE=on go install -v github.com/golang/mock/mockgen

mocks: install-mockgen
	$(call local_mockgen,2019/async,Executable)

test: mocks
	go test -cover ./...

define run_mockgen
  	mkdir -p $(2) && rm -rf $(2)/*
  	mockgen -destination=$(2)/mocks.go $(1) $(subst $(semicolon),$(comma),$(3))
endef

define local_mockgen
	$(call run_mockgen,$(PROJECT_ROOT)/$(1),.mocks/$(1),$(2))
endef
