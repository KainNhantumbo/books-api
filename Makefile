build:
	go build -o server main.go

run:
	go run main.go

watch:
	reflex -s -r '\.go$$' make run
