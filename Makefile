
IMAGE := beeman/go-static-server

docker-build:
	docker build -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)

docker-run:
	docker run -it --rm --name go-static-server -p8080:8080 $(IMAGE)