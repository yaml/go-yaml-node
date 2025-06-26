# Using the "Makes" Makefile setup - https://github.com/makeplus/makes
M := .cache/makes
$(shell [ -d $M ] || git clone -q https://github.com/makeplus/makes $M)
include $M/init.mk
include $M/clean.mk
GO-CMDS-SKIP := install
include $M/go.mk

override PATH := $(ROOT):$(PATH)
export PATH

PROGRAM := go-yaml-node

TEST-FILES := $(wildcard test/*.yaml)

ifneq (,$(file))
TEST-FILES := $(file)
endif

MAKES-CLEAN := $(PROGRAM)


default::

test: $(TEST-FILES)

test-update: $(PROGRAM)
	@for f in $(TEST-FILES); do \
		printf "Updating %s... " "$$f"; \
		./$(PROGRAM) < "$$f" > "$${f%.yaml}.want"; \
		echo "done"; \
	done

.PHONY: $(TEST-FILES)
$(TEST-FILES):: $(PROGRAM)
	@printf "$@ "
	@$< < $@ > $(LOCAL-TMP)/got
	@if diff -q $(@:test/%.yaml=test/%.want) $(LOCAL-TMP)/got; then \
	  echo "PASS"; \
	else \
	  echo "FAIL"; \
	fi
	@diff -u $(@:test/%.yaml=test/%.want) $(LOCAL-TMP)/got || true

build: $(PROGRAM)

install: $(PROGRAM)
ifndef PREFIX
	$(error PREFIX is not set)
else
	install -m 0755 $(PROGRAM) $(PREFIX)/bin/$(PROGRAM)
endif

$(PROGRAM): yaml.go $(GO)
	go mod tidy
	go fmt
	go build
