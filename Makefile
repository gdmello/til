test:
	ginkgo *

build: test
	godep go build -o til
