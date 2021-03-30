#Docker env
dev-build:
	docker build -f docker/dev/Dockerfile -t test.com/mutant-checker:dev .

dev-run:
	docker run --rm -it \
	-p 4000:4001 \
	-v ${PWD}/cmd:/app/cmd \
	-v ${PWD}/infrastructure:/app/infrastructure \
	-v ${PWD}/domain:/app/domain \
	--env-file ./.env \
	test.com/mutant-checker:dev

dev-start: dev-build dev-run

# if you check the image for production works!
prod-build:
	docker build -f docker/prod/Dockerfile -t mutant-checker:prod .

prod-run:
	docker run --rm -it \
	-p 4000:4002 \
	--env-file ./.env \
	mutant-checker:prod

tests:
	go test -v ./... -coverprofile coverage/cover.out && go tool cover -html=coverage/cover.out -o coverage/cover.html
clean:
	- rm -rf coverage
	- rm -rf tmp
	- rm -rf vendor
	- rm gin-bin
	- rm main
	- rm courier