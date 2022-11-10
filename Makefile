run:
	cd web/app && yarn install && yarn build
	GIN_MODE=release go run main.go
web-dev:
	cd web/app && yarn dev
