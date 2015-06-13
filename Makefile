debug:
	DEBUG=infor-you-mation:server node bin/www
run:
	node bin/www
startmongo:
	mongod -f /usr/local/etc/mongod.conf
