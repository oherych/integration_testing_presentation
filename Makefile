run:
	@ DATABASE_HOST=localhost \
	DATABASE_USER=user \
	DATABASE_NAME=dev \
	DATABASE_PASSWORD=password \
	STORAGE_ENDPOINT=localhost:9000 \
	STORAGE_ACCESS_KEY=BH0K3LEZSZX2KFM53LLS \
	STORAGE_SECRET_KEY=Pzfn+HrTbw+oPO8Tz5NFnj/1RbWSjH1qQ+cqCJE6 \
	STORAGE_LOCATION=eu-central-1 \
	STORAGE_PAYLOAD_BUCKET=test \
	PORT=:8080 \
	go run main.go