GO=go
PKG=github.com/alvii147/gloop
COV=coverage.out

TEST_OPTS=-skip Example.* -coverprofile $(COV)
ifdef TESTCASE
	TEST_OPTS=-run $(TESTCASE)
endif

TEST_OPTS:=$(TEST_OPTS) -v

.PHONY: test
test:
	$(GO) test $(TEST_OPTS) $(PKG)
	@if [ -z "$(TESTCASE)" ]; then\
		$(GO) tool cover -func $(COV);\
	fi
	@rm -f $(COV)
