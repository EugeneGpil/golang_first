dev:
	cd container &&
		docker compose up --build --remove-orphans --detach &&
		docker compose exec --user=app golang bash

stop:
	cd container && docker compose stop
