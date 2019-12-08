DIST := dist
GO ?= go

# Executable
ifeq ($(OS), Windows_NT)
	EXECUTABLE := fritzbox-graphite_exe
else
	EXECUTABLE := gitea
endif

# Version
ifneq ($(DRONE_TAG),)
	VERSION ?= $(subst v,,$(DRONE_TAG))
	GITEA_VERSION := $(VERSION)
else
	ifneq ($(DRONE_BRANCH),)
		VERSION ?= $(subst release/v,,$(DRONE_BRANCH))
	else
		VERSION ?= master
	endif
	GITEA_VERSION := $(shell git describe --tags --always | sed 's/-/+/' | sed 's/^v//')
endif

# Make Tasks
release: release-dirs release-windows release-linux release-darwin release-copy release-compress release-check

release-dirs:
	mkdir -p $(DIST)/binaries $(DIST)/release

release-linux:
	@hash xgo > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/karalabe/xgo; \
	fi
	xgo -dest $(DIST)/binaries -tags 'netgo $(TAGS)' -ldflags '-linkmode external -extldflags "-static" $(LDFLAGS)' -targets 'linux/amd64,linux/386,linux/arm' -out fritzbox-graphite_$(VERSION) .

release-windows:
	@hash xgo > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/karalabe/xgo; \
	fi
	xgo -dest $(DIST)/binaries -tags 'netgo $(TAGS)' -ldflags '-linkmode external -extldflags "-static" $(LDFLAGS)' -targets 'windows/*' -out fritzbox-graphite_$(VERSION) .

release-darwin:
	@hash xgo > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/karalabe/xgo; \
	fi
	xgo -dest $(DIST)/binaries -tags 'netgo $(TAGS)' -ldflags '-linkmode external -extldflags "-static" $(LDFLAGS)' -targets 'darwin-10.10/*' -out fritzbox-graphite_$(VERSION) .

release-copy:
	$(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(notdir $(file));)

release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

release-compress:
	@hash gxz > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/ulikunitz/xz/cmd/gxz; \
	fi
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),gxz -k -9 $(notdir $(file));)
