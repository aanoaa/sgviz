DB_FILE	:= $(HOME)/.local/share/sgviz/sqlite3.db
DB_DIR	:= $(dir $(DB_FILE))
SQLITE	:= sqlite3
APP	:= bin/sgviz

all: build

.PHONY: build $(APP)
build: $(APP)

$(APP):
	go build -o $@ main.go

.PHONY: db
db: $(DB_FILE)

$(DB_FILE): $(DB_DIR)
	$(SQLITE) $@ < configs/init.sql
	$(SQLITE) $@ < configs/test-data.sql

$(DB_DIR):
	mkdir -p $@

.PHONY: clean
clean:
	rm -f $(DB_FILE)

.PHONY: schema
schema:
	sqlboiler sqlite3

.PHONY: import
import:
	$(SQLITE) $(DB_FILE) < data/sql/data.sql

.PHONY: export
export:
	$(SQLITE) $(DB_FILE) .schema > data/sql/schema.sql
	$(SQLITE) $(DB_FILE) .dump > data/sql/dump.sql
	grep -vx -f data/sql/schema.sql data/sql/dump.sql > data/sql/data.sql
