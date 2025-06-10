M := .git/.makes
$(shell [ -d $M ] || git clone -q https://github.com/makeplus/makes $M)
include $M/init.mk
include $M/go.mk

override PATH := $(ROOT):$(PATH)
export PATH

PROGRAM := yaml-node-dump

TEST-FILES := $(wildcard test/*.yaml)

ifneq (,$(file))
TEST-FILES := $(file)
endif


# Print Makefile targets summary
default::
	@printf '%s\n' $(TEST-FILES)

test-go: $(GO)
	go test ./...

test: $(TEST-FILES)

# $(MAKE) -s $(TEST-FILES) | less -FRX

.PHONY: $(TEST-FILES)
$(TEST-FILES):: $(PROGRAM)
	@printf "$@ "
	@$< < $@ > $(LOCAL-TMP)/got
	@if diff -q $(@:test/%.yaml=test/%.out) $(LOCAL-TMP)/got; then \
	  echo "PASS"; \
	else \
	  echo "FAIL"; \
	fi
	@diff $(@:test/%.yaml=test/%.out) $(LOCAL-TMP)/got


# @( \
#   echo "==== Input — $@"; \
#   echo; \
#   cat $@; \
#   echo; \
#   echo '==== Output:'; \
#   echo; \
#   $< < $@; \
# )

build: $(PROGRAM)

clean:
	$(RM) $(PROGRAM)

$(PROGRAM): yaml.go $(GO)
	go mod tidy
	go fmt
	go build
