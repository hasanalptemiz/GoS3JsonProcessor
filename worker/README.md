## Getting Started

To run this project locally, follow these steps:

1. Start the Docker containers by running the following command:

docker compose up -d


2. Once the containers are up and running, obtain the IP address of the `my-mongodb` container using the following command:

docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' my-mongodb


3. Update the environment variables in the `.env` file to match your configuration:

MONGO_URI=mongodb://admin:password@<mongodb_container_ip>:27017
MONGO_DB=S3JsonProcessor


4. Configure the following environment variables based on your AWS S3 settings:

ACCESS_KEY=<your_aws_s3_access_key>
SECRET_KEY=<your_aws_s3_secret_key>
REGION=<your_aws_s3_region>
BUCKET_NAME=<your_aws_s3_bucket_name>

**Note**: Instead of directly including your AWS S3 access key and secret key here, it's recommended to securely manage these credentials using environment variables or a secrets management tool.


5. Update the worker code in `main.go` to specify the object keys statically:

objectKeys := []string{"products-1.jsonl", "products-2.jsonl", "products-3.jsonl", "products-4.jsonl"}

docker-compose logs my-go-worker


This revised README provides clear instructions on setting up the project without including any sensitive information directly in the file. Users can safely configure their environment variables and run the project without violating GitHub's security policies. If you have any further questions or need additional assistance, feel free to ask!




