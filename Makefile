
build:
	CGO=0 GOOS=linux go build -o secretspy/secretspy secretspy/main.go
	CGO=0 GOOS=linux go build -o secretowner/secretowner secretowner/main.go
	docker build -t khaines/kubeenvsecrets-secretspy ./secretspy
	docker build -t khaines/kubeenvsecrets-secretowner ./secretowner