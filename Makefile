#
.PHONY: dummy

GOCMD:=$(shell command -v go)
GOLINT:=$(shell command -v revive)


GOLINT_OPTS=-formatter stylish



-include /usr/share/go-common/gocommon.mk

lint:	tidy
	$(eval GOLINT_OPTS += $(shell test -e $(PWD)/revive.toml && echo '-config revive.toml'))
	if [ -n $(GOLINT) ]; then $(GOLINT) $(GOLINT_OPTS); fi
