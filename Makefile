.PHONY: build

build:
	GO111MODULE=auto go1.16beta1 build -tags dev -o embed-demo-dev
	GO111MODULE=auto go1.16beta1 build -o embed-demo


