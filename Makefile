install:
	@echo "> Installing whoami..."
	gh extension install .

remove:
	@echo "< Removing whoami..."
	gh extension remove whoami

reload: remove install