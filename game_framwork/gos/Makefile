all:
	./tools/gen_protocol
	./tools/gen_routes
	bundle exec rake generate_tables
	go install server

start:
	./bin/server

console:
	go install server
	./bin/server

install:
	bundle install
	bundle exec rake db:create
	bundle exec rake db:migrate
