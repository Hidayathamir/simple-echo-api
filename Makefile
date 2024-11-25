run:
	go run .

build-docker:
	docker build -t hidayathamir/simple-echo-api .

push-docker:
	docker push hidayathamir/simple-echo-api

run-docker:
	docker run -p 8080:8080 hidayathamir/simple-echo-api

