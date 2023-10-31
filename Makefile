build:
	GOOS=linux go build .

container: build
	docker build . -t playtechnique/play-controller:pre-release

push: container
	docker push playtechnique/play-controller:pre-release
