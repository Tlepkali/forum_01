silent:

.PHONY: migrate

migrate: silent
	sqlite3 forum.sqlite3 < migrations/01_add_user_table.sql
	sqlite3 forum.sqlite3 < migrations/02_add_post_table.sql
	sqlite3 forum.sqlite3 < migrations/03_add_category_table.sql
	sqlite3 forum.sqlite3 < migrations/04_add_session_table.sql
	sqlite3 forum.sqlite3 < migrations/05_add_comment_table.sql
	sqlite3 forum.sqlite3 < migrations/06_add_post_vote_table.sql
	sqlite3 forum.sqlite3 < migrations/07_insert_categories.sql
	sqlite3 forum.sqlite3 < migrations/08_add_post_categories.sql
	sqlite3 forum.sqlite3 < migrations/09_add_comment_vote_table.sql
drop: silent
	sqlite3 forum.sqlite3 < migrations/00_drop_all_tables.sql

.PHONY: docker-build

docker-build: silent
	docker build -t forum .

.PHONY: docker-run

docker-run: silent
	docker run -p 8080:8080 --rm --name forum forum:latest

.PHONY: initdb

initdb: silent
	docker exec -it forum sh initdb.sh 


#############################################################################################
# Quality Control																			#	
#############################################################################################

.PHONY: audit

audit: silent
	golangci-lint run
	go mod tidy
	go verify
	go fmt ./...
	go vet ./...
	statciheck ./...

.PHONY: vendor

vendor: silent
	go mod vendor