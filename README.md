# go-webapi-demo

  This repository is having go web service with two apis api/encrypt and api/decrypt. You could create docker container and run this command.

## Prerequisites

  You need linux machines with Docker and curl installed.

## Docker build

   We are using docker multistage file to create small size docker images. Go to the root directory and run docker build command to create docker image.

   ```
   docker build -t web .
   ```

   Check web docker image is available in your machine.
   ```
   docker images
   ```
   Run the docker image using below command.
   ```
   docker run -d -p 80:8080 web
   ```
## Validate the api request

   Run the below curl command to encrypt the data.
   ```
   curl -X POST -H "Content-Type: application/json" \
      -d '{ "data": "{plain_text}" }' http://localhost/api/encrypt
   ```
   It will response with encrypted content.
   ```
   {"encrypted":"{encrypted_text}"}
   ```
   Run the below command to decrypt the data and it will response with original decrypted data.
   ```
   curl -X POST -H "Content-Type: application/json" \
       -d '{"encrypted":"{encrypted_text}"}' \
        http://localhost:8080/api/decrypt
   
   {"decrypted":"{plain_text}"}
   ```
