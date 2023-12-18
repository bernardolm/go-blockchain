.PHONY:

run:
	reset ; goimports -w ./.. ; go mod tidy ; go run main.go

debug:
	dlv debug --listen=:2345 --headless=true --api-version=2
