# Minio -> Golang Example

## Run Locally
You will need to create a new env file via `cp .env.sample .env`.

This project is fully docker-ized. Run with `docker-compose up -d`.

To just rebuild the golang binary api container, run `docker-compose up -d --build api`

### Interaction

To interact w/ this, send a post request to `http://localhost/` with json data shaped like:
```json
{
  "file_name": "test.txt", 
  "content": "these are the files contents"
}
```

You should then find a `test.txt` file in the `./var/minio/private` folder, containing the text
`these are the files contents`.

### Minio portal

You can access the minio portal from your web browser by navigating to `http://localhost:9000`.

You will be presented w/ a login portal where you can enter the access key id and secret from your `.env` file.
Once logged in, you can interact w/ the contents of the buckets.
