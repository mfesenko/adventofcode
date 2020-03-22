.PHONY: mocks install-mockgen

PROJECT_ROOT = github.com/mfesenko/adventofcode

install-mockgen:
	GO111MODULE=on go get -v -u github.com/golang/mock/mockgen

mocks: install-mockgen
	$(call local_mockgen,2019/arcade,Computer)
	$(call local_mockgen,2019/async,Executable)
	$(call local_mockgen,2019/drawing,Computer)
	$(call local_mockgen,2019/navigation,Graph)
	$(call local_mockgen,2019/repair,Computer)

test: mocks
	go test -cover ./...

define run_mockgen
  	mkdir -p $(2) && rm -rf $(2)/*
  	${HOME}/go/bin/mockgen -destination=$(2)/mocks.go $(1) $(subst $(semicolon),$(comma),$(3))
endef

define local_mockgen
	$(call run_mockgen,$(PROJECT_ROOT)/$(1),.mocks/$(1),$(2))
endef
