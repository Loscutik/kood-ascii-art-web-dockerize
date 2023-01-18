echo "Building a docker image: (docker build -t ascii-art-dockerize .)"
echo "**********************"
docker build --rm -t ascii-art-dockerize .
echo "**********************"
read -p "Press enter to continue"

echo "!"
echo "Print the list of images: (docker images)"
echo "**********************"
docker images
echo "**********************"
read -p "Press enter to continue"

echo "!"
echo "Starting the container using the docker image: (docker run -dp 8080:8080 --name=ascii-container ascii-art-dockerize)"
echo "**********************"
docker run -dp 8080:8080 --name=ascii-container ascii-art-dockerize
echo "**********************"
echo "Please open http://localhost:8080 in browser"
echo "**********************"
read -p "Press enter to continue"

echo "!"
echo "Print the list of containers: (docker ps -a)"
echo "**********************"
docker ps -a
echo "**********************"
read -p "Press enter to continue"

echo "!"
echo "Print Docker image metadata: (docker inspect --format='{{json .Config.Labels}}' ascii-container)"
echo "**********************"
docker inspect --format='{{json .Config.Labels}}' ascii-container
echo "**********************"
read -p "Press enter to continue"

echo "!"
echo "Connect to container''s filesystem: (docker exec -it ascii-container /bin/bash)"
echo "You can run ls -l  to see the file system. Type 'exit' to qiut"
echo "**********************"
docker exec -it ascii-container /bin/sh
echo "**********************"
read -p "Press enter to continue"

echo "!"
echo "Stop and delete the container (docker rm -f ascii-container)"
echo "**********************"
docker rm -f ascii-container
echo "**********************"

echo "!"
echo "Delete images (docker image prune and )"
echo "**********************"
docker image prune -f && docker image rm ascii-art-dockerize:latest golang:1.19-alpine alpine:latest
echo "**********************"



echo "**********************"
