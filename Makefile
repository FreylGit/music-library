LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN="$(LOCAL_BIN)" go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN="$(LOCAL_BIN)" go install github.com/gojuno/minimock/v3/cmd/minimock@v3.3.6


LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=5435 dbname=song user=song-user password=song-password sslmode=disable"

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v
