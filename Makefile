composeup:
	podman-compose up -d

composedown:
	podman-compose down

server:
	go run main.go

.PHONY: composeup composedown server