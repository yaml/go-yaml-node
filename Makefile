# Using the "Makes" Makefile setup - https://github.com/makeplus/makes
M := .git/.makes
$(shell [ -d $M ] || git clone -q https://github.com/makeplus/makes $M)
include $M/init.mk
include $M/go.mk

override PATH := $(ROOT):$(PATH)
export PATH

PROGRAM := go-yaml-node

TEST-FILES := $(wildcard test/*.yaml)

ifneq (,$(file))
TEST-FILES := $(file)
endif


default::

test: $(TEST-FILES)

.PHONY: $(TEST-FILES)
$(TEST-FILES):: $(PROGRAM)
	@printf "$@ "
	@$< < $@ > $(LOCAL-TMP)/got
	@if diff -q $(@:test/%.yaml=test/%.out) $(LOCAL-TMP)/got; then \
	  echo "PASS"; \
	else \
	  echo "FAIL"; \
	fi
	@diff -u $(@:test/%.yaml=test/%.out) $(LOCAL-TMP)/got || true

build: $(PROGRAM)

install: $(PROGRAM)
ifndef PREFIX
	$(error PREFIX is not set)
else
	install -m 0755 $(PROGRAM) $(PREFIX)/bin/$(PROGRAM)
endif

tidy: $(GO)
	go mod tidy

fmt: $(GO)
	go fmt

clean:
	$(RM) $(PROGRAM)

$(PROGRAM): yaml.go $(GO)
	go mod tidy
	go fmt
	go build
