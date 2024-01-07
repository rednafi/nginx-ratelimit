lint:
	prettier --write .

test:
	cd app && go test -v .
