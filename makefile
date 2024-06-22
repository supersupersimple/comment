build-fe:
	npx tailwindcss -i ./app/assets/input_comment.css -o ./app/assets/dist/css/comment.css --minify
	npx tailwindcss -i ./app/assets/input_admin.css -o ./app/assets/dist/css/admin.css --minify
	npx uglifyjs ./app/assets/sssc.js -o ./app/assets/dist/js/sssc.js
	npx uglifyjs ./app/assets/comment.js -o ./app/assets/dist/js/comment.js

build-be:
	go generate ./ent
	templ fmt .
	templ generate --path app/templates/

build:
	make build-fe
	make build-be

build-dev:
	npx tailwindcss -i ./app/assets/input_comment.css -o ./app/assets/dist/css/comment.css
	npx tailwindcss -i ./app/assets/input_admin.css -o ./app/assets/dist/css/admin.css
	cp ./app/assets/sssc.js ./app/assets/dist/js/sssc.js
	cp ./app/assets/comment.js ./app/assets/dist/js/comment.js
	make build-be

run:
	go run main.go serve --https=false

build-run:
	make build-dev
	make run

compile:
	go mod tidy
	make build
	go build -o comment

compile-air:
	make build-dev
	go build -o ./tmp/main .

docker-go-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o comment

docker-node-build:
	make build-fe
