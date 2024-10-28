GO=go
PKG=github.com/alvii147/gloop
COV=coverage.out

TEST_OPTS=-skip Example.* -coverprofile $(COV)
ifdef TESTCASE
	TEST_OPTS=-run $(TESTCASE)
endif

ifeq ($(VERBOSE), 1)
	TEST_OPTS:=$(TEST_OPTS) -v
endif

.PHONY: test
test:
	$(GO) test $(TEST_OPTS) $(PKG)
	@if [ -z "$(TESTCASE)" ]; then\
		$(GO) tool cover -func $(COV);\
	fi
	@rm -f $(COV)
