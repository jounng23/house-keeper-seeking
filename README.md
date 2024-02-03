# Microservice: Housekeeper Seeking Application

## Setup Instructions

1. Make the **start_services.sh** file executable by running the following command in your terminal (for linux):

```
chmod +x start_services.sh
```

2. Make sure you have **Go**, **docker** and **docker-compose** installed on your system.

3. Open a terminal and navigate to the root directory of the repository.

4. Run the following command to start mongoDB container and all of three Go services simultaneously (use Gitbash terminal if using windows):

```
sh start.sh
```
Please wait for a while until all commands are executed completely. This will start all three services in the background.

5. Use the following cURL command to test booking a housekeeper via API:

```
curl --location --request POST 'localhost:3000/api/v1/job/book-house-keeper' \
--header 'Content-Type: application/json' \
--data-raw '{
    "client_info": {
        "id": "1",
        "name": "Nguyen Van X",
        "phone_number": "0987654321"
    },
    "booking_date": "2024-01-01 08:30:00"
}'
```

6. Watch the response and check if data is correct. To ensure database is correct, please connect MongoDB via MONGODB_URI in .env.example.
