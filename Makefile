install:
	cd client && yarn install

build:
	cd client && yarn build
	ENV=prod go build -buildvcs=false -o ./bin/aircup ./main.go

dev:
	cd client && yarn dev & air && fg
