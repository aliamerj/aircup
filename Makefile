APP_NAME=aircup

.PHONY: web build run clean dev

web:
	cd web && npm install && npm run build

build: web
	go build -o bin/$(APP_NAME) .

run: build
	./bin/$(APP_NAME) serve

dev:
	concurrently \
		"cd web && npm run dev" \
		"go run . serve"

clean:
	rm -rf bin
	rm -rf web/dist
