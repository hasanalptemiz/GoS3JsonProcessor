## First, you should run command
docker compose up -d 

## After that you need to check my-mongodb container DNS with CLI

docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' my-mongodb

## Enviroment variables (.env)
API_PORT=8080
API_TOKEN = QJl6eEuAJAfunviF
MONGO_URI=mongodb://admin:password@192.168.228.2:27017
MONGO_DB= golangcimri

## After that, all you have to do is go to the post office and test the get record enpoint. I got a very fast get endpoint using compound indexing. I shared the postman link with this address review@cimri.com. Api can be tested by entering the postman from the invitation or using the request link below.

http://localhost:8080/api/v1/private/get/record/150003

Header :

Api-Token : QJl6eEuAJAfunviF