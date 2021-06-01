# ocha-url
Test Link : https://gist.github.com/williamn/cfad86ab218101e0c5d7be89226c5c85

### Points 
- net/http vs frameworks
- bigcache vs redis
- Folder Structure vs Clean Code Architecture

### Generate Doc
```shell 
swag init -g cmd/ocha/main.go -o ./api
```

Docker Compose Up
docker-compose -f deployments/docker-compose.yml up
