mongo:
	podman run --name learn-mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=secret -d mongo:latest

createdb:
	podman 

.PHONY: create-db