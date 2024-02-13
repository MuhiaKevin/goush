run:
	go run cmd/web/*
fmt:
	go fmt cmd/web/*; go fmt internal/models/*

post:
	curl -v -X POST http://localhost:4000/link/create -H 'Content-Type: application/json' -d '{"original_url" : "https://github.com/rustdesk/rustdesk/releases"}'

dev: 
	air --build.cmd "go build -o bin/goush cmd/web/*" --build.bin "bin/goush"
