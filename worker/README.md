## First, you should run command
docker compose up -d 

## After that you need to check my-mongodb container DNS with CLI

docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' my-mongodb

## Enviroment variables (.env)
MONGO_URI=mongodb://admin:password@192.168.237.2(you should change with your my-mongodb container DNS):27017
MONGO_DB=golangcimri
ACCESS_KEY=AKIA5OUQ2DJ2KKE2JV7P
SECRET_KEY=phuUfwcRSjSRNgflhDvTzfEl9JyLKJz5KBvZnK6a
REGION=eu-central-1
BUCKET_NAME=cimri-casestudy

## Due to the problems with the docker compiler, you must enter the input data statically from the objectKey in the main.go section. 

objectKeys := []string{"products-1.jsonl", "products-2.jsonl", "products-3.jsonl", "products-4.jsonl"}

## When you stand up the container with the same files, you will see that the same data is not written to the db. you can also check this with the following command

docker-compose logs my-go-worker