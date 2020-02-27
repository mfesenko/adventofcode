.PHONY: mocks install-mockgen

PROJECT_ROOT  = github.com/mfesenko/adventofcode
GOMOCK = $(go get github.com/golang/mock/gomock github.com/golang/mock/mockgen)

install-mockgen:
	GO111MODULE=on go get github.com/golang/mock/mockgen@latest

mocks: install-mockgen 	
	$(call local_mockgen,2019/async,Executable)

test: mocks
	go test -cover ./...

define run_mockgen $(GOMOCK)
  	mkdir -p $(2) && rm -rf $(2)/*
  	mockgen -destination=$(2)/mocks.go $(1) $(subst $(semicolon),$(comma),$(3))
endef

define local_mockgen
	$(call run_mockgen,$(PROJECT_ROOT)/$(1),.mocks/$(1),$(2))
endef
