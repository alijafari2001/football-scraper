example1: example2 example3
	go run cmd/main.go
example2:
	go run cmd/main.go -s=19/20 -l=bu -t
example3:
	go build cmd/main.go
	./main -l=se -s=21/22 -f=11