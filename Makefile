
docker:
	docker build --network=host --tag tf03mqtt:latest .

run:
	docker run --name tf03mqtt_agent -d --restart=always --privileged tf03kmon:latest

pibin:
	env GOOS=linux GOARCH=arm GOARM=7 go build

clean:
	docker stop tf03mqtt_agent
	docker rm tf03mqtt_agent

.PHONY: docker clean run