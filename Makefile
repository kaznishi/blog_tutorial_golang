init :
	docker-compose build --no-cache
	docker-compose run --rm golang dep ensure

dep :
	docker-compose run --rm golang dep ensure

start :
	docker-compose up