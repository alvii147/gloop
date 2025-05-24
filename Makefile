GO=go
PKG=github.com/alvii147/gloop
COV=coverage.out

TEST_OPTS=-skip Example.* -coverprofile $(COV)
ifdef TESTCASE
	TEST_OPTS=-run $(TESTCASE)
endif

TEST_OPTS:=$(TEST_OPTS) -v -count=1

.PHONY: test
test:
	$(GO) test $(TEST_OPTS) $(PKG)

.PHONY: race
race:
	$(GO) test -race $(TEST_OPTS) $(PKG)

.PHONY: cover
cover:
	$(GO) tool cover -func $(COV);
