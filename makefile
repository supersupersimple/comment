build:
	go generate ./ent
	templ fmt .
	templ generate --path app/templates/
	npx tailwindcss -i ./app/assets/input_comment.css -o ./app/assets/dist/css/comment.css --minify
	npx tailwindcss -i ./app/assets/input_admin.css -o ./app/assets/dist/css/admin.css --minify
	uglifyjs ./app/assets/sssc.js -o ./app/assets/dist/js/sssc.js

run:
	make build
	go run main.go serve

compile:
	go mod tidy
	make build
	go build -o comment

compile-air:
	make build
	go build -o ./tmp/main .