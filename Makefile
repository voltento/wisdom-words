# Makefile

# Variables
SERVER_IMAGE_NAME = wisdom-server
CLIENT_IMAGE_NAME = wisdom-client
SERVER_CONTAINER_NAME = wisdom-server-container
CLIENT_CONTAINER_NAME = wisdom-client-container

# Build server Docker image
build-server:
	docker build -t $(SERVER_IMAGE_NAME) -f ./ops/server/Dockerfile .

# Build client Docker image
build-client:
	docker build -t $(CLIENT_IMAGE_NAME) -f ./ops/client/Dockerfile .

# Run server container
run-server: stop-server build-server
	docker run --name $(SERVER_CONTAINER_NAME) -p 8080:8080 $(SERVER_IMAGE_NAME) --difficulty=5

# Run client container
run-client: stop-client build-client
	docker run --name $(CLIENT_CONTAINER_NAME) --link $(SERVER_CONTAINER_NAME):server $(CLIENT_IMAGE_NAME) --address=$(SERVER_CONTAINER_NAME):8080

# Stop and remove server container
stop-server:
	docker stop $(SERVER_CONTAINER_NAME) || echo ''
	docker rm $(SERVER_CONTAINER_NAME) || echo ''

# Stop and remove client container
stop-client:
	docker stop $(CLIENT_CONTAINER_NAME) || echo ''
	docker rm $(CLIENT_CONTAINER_NAME)  || echo ''

# Remove server and client images
clean-images:
	docker rmi $(SERVER_IMAGE_NAME)
	docker rmi $(CLIENT_IMAGE_NAME)

# Run all
all: build-server build-client run-server run-client

# Stop all
stop-all: stop-server stop-client

# Clean everything
clean: stop-all clean-images
