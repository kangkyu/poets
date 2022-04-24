# Make sure you use tabs instead of spaces in the commands after the tasks.
test:
	curl localhost:8080/submit --data-raw '{"message":"Here we go"}' -H 'Content-Type: application/json'

.PHONY: test
