install:
	go-bindata -pkg audio -o audio/audio.go audio/
	go install .

build:
	go-bindata -pkg audio -o audio/audio.go audio/
	go build .