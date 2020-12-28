
docker:
	docker build --network=host --tag tf03kmon:latest .

run:
	docker run --name tf03kmon_agent -d --restart=always --privileged tf03kmon:latest

pibin:
	env GOOS=linux GOARCH=arm GOARM=7 go build

clean:
	docker stop tf03kmon_agent
	docker rm tf03kmon_agent

.PHONY: docker clean run