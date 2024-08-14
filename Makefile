.DEFAULT_GOAL := build

fmt:
		go fmt ./...

vet: fmt
		go vet ./...

clean:
		go clean

test: vet
		go test -v -cover ./...

build: test
		go build -o bin

run: build
		./bin

build-image: test
		docker build -t url_shortener .

deploy:
		kubectl apply -f k8s/deployment.yaml -n development

build-deploy: build-image deploy

delete:
		kubectl delete -f k8s/deployment.yaml -n development
		