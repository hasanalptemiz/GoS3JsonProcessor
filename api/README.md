## Getting Started

To run this project locally, follow these steps:

1. Start the Docker containers by running the following command:

docker compose up -d


2. Once the containers are up and running, obtain the IP address of the `my-mongodb` container using the following command:


docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' my-mongodb


3. Update the environment variables in the `.env` file to match your configuration:

API_PORT=<your_api_port>
MONGO_URI=<your_mongodb_uri>
MONGO_DB=S3JsonProcessor


4. Set up the API token securely. Instead of directly including it in the README, it's recommended to manage secrets securely, such as using environment variables or a secrets management tool.

5. After setting up the environment variables, you can test the API endpoints. Access the Postman collection provided via email or use the following request link:

http://localhost:8080/api/v1/private/get/record/150003


Remember to include the API token in the request header:

Header:
Api-Token: <your_api_token>


By following these steps, you can set up and run the project locally. Make sure to securely manage your environment variables and avoid exposing sensitive information in your code or README files. If you encounter any issues, refer to the project documentation or seek assistance from the project maintainers.

