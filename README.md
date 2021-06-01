# Ocha
**Requirement  Link :** https://gist.github.com/williamn/cfad86ab218101e0c5d7be89226c5c85

### Libraries Used



- [x] **Gin** - Web Server
- [x] **BigCache** - in-memory storage
- [x] **Swagger (Swag)** - API Documentation
- [x] **govalidator** - Payload Validation



### Commands

####  Generate API Document (Swagger)
```shell 
swag init -g cmd/ocha/main.go -o ./api  
```  

- Swagger URL : http://localhost:8080/swagger/index.html

####  Build & Execute
```shell 
make build-release && ./build/bin/ocha  
```  

####  Docker Compose
```shell 
make docker-run
```  
*-- or --*
```shell 
docker-compose -f deployments/docker-compose.yml up
```  

####  Unit testing
```shell 
make unit
```  
*-- or --*

```shell 
go test --tags=unit ./...
```  
####  Integration testing
```shell 
make integration
```  
*-- or --*
```shell 
go test --tags=integration ./...
```  
 
