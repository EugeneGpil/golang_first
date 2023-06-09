dev:
	docker compose up --build --remove-orphans --detach

exec:
	docker compose exec --user=app golang bash

exec-root:
	docker compose exec --user=root golang bash

stop:
	docker compose stop
