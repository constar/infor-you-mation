development:
	./node_modules/grunt-cli/bin/grunt && PORT=3001 node ./bin/www
production:
	PORT=3000 node ./bin/www
