




deploy-rabbitmq:
	@echo "deploying a simple RabbitMQ with docker-compose"


build-sender: 
	go mod vendor
	docker build -f ./build/sender/Dockerfile-sender -t sender-amqp:latest .
	rm -fr vendor 

