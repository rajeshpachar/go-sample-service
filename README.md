# go-sample-service

docker build -t go-sample-service .

docker run  -d -p 3333:3333 sync-sample-service
docker ps
docker logs e8b69c552245

curl -f -v localhost:3333