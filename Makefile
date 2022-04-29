



build-sender: 
	go mod vendor
	docker build -t sender-amqp:latest .
	rm -fr vendor 