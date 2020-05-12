# Makefile configured from and for OSX. It may work on other OS however it is not actively supported to do so.
# Requires semver https://github.com/markchalloner/git-semver to set the semantic tag version, and to retrieve it.

QT_DIR =/usr/local/opt/qt

.PHONY: build desktop all windows clean

GIT_COMMIT := $(shell git rev-list -1 HEAD)
DATE			?= $(shell date +%FT%T%z)
TAGGING			?= $(shell git describe --tags)#potentially out of datedel
SEMVER			?= $(shell git semver get)# uses https://github.com/markchalloner/git-semver --> git semver get
HASH			?= $(shell git rev-parse HEAD)
DBNAME ?= pickleit.db

# PRODUCTION ?= TRUE
LDFLAGS			+= -X "main.buildTagging=${TAGGING}"
LDFLAGS			+= -X "main.buildVersion=${SEMVER}"
LDFLAGS			+= -X "main.buildHash=${HASH}"
LDFLAGS			+= -X "main.buildDate=${DATE}"
LDFLAGS			+= -X "main.dbName=${DBNAME}"
LDFLAGS			+= -X "main.persistLogs=TRUE"
# LDFLAGS			+= -s -w
CGO_LDFLAGS_ALLOW			="-Wl,-rpath,@executable_path/Frameworks"

build:
	export GO111MODULE=off; GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.gitCommit=$(GIT_COMMIT)" .
desktop:
	export GO111MODULE=off; QT_DIR=$(QT_DIR) CGO_LDFLAGS_ALLOW=$(CGO_LDFLAGS_ALLOW) qtdeploy -ldflags="$(LDFLAGS)" test desktop
#osxtest: qtmoc -slow && go run -tags "qml" .

local:
	export GO111MODULE=off; QT_DIR=$(QT_DIR) CGO_LDFLAGS_ALLOW=$(CGO_LDFLAGS_ALLOW) qtdeploy -tags "qml" -ldflags="$(LDFLAGS) -X "main.production=FALSE" -X "main.flavour=local"" test desktop
production:
	export GO111MODULE=off; GO111MODULE=off QT_DIR=$(QT_DIR) CGO_LDFLAGS_ALLOW=$(CGO_LDFLAGS_ALLOW) qtdeploy -tags "qml" -ldflags="$(LDFLAGS) -X "main.production=TRUE" -X "main.flavour=production"" build desktop
windows:
	GO111MODULE=off QT_DIR=$(QT_DIR) CGO_LDFLAGS_ALLOW=$(CGO_LDFLAGS_ALLOW) qtdeploy -tags "qml" -ldflags="$(LDFLAGS) -X "main.production=TRUE"" -docker build windows_64_shared
debug:
	export GO111MODULE=off; QT_DIR=$(QT_DIR) CGO_LDFLAGS_ALLOW=$(CGO_LDFLAGS_ALLOW) qtdeploy -debug -tags "qml" -ldflags="$(LDFLAGS) -X "main.production=FALSE"" test desktop
fast:
	export GO111MODULE=off; QT_DIR=$(QT_DIR) CGO_LDFLAGS_ALLOW=$(CGO_LDFLAGS_ALLOW) qtdeploy -fast -tags "qml" -ldflags="$(LDFLAGS) -X "main.production=FALSE"" test desktop
virtualqml:
	LDFLAGS			+= -X "main.virtual=TRUE"
	export GO111MODULE=off; QT_DIR=$(QT_DIR) CGO_LDFLAGS_ALLOW=$(CGO_LDFLAGS_ALLOW) qtdeploy -tags "qml" -ldflags="$(LDFLAGS) -X "main.production=FALSE"" test desktop
mockgui:
	export GO111MODULE=off; QT_DIR=$(QT_DIR) CGO_LDFLAGS_ALLOW=$(CGO_LDFLAGS_ALLOW) qtdeploy -tags "qml mock" -ldflags="$(LDFLAGS) -X "main.production=FALSE"" test desktop
