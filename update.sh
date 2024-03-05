git pull
docker stop dss-data
docker rm dss-data
docker rmi dss-data-dss-data:latest
docker-compose up -d