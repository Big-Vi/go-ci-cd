
## Docker
 - Create Dockerfile and Dockerfile.production
 - Build Docker production image using multi-stage builds
 - Add Docker compose and Makefile for easy of use
 - Since the Go app and Postgres running in separate container, shared network needs to be created. The name of the network can be referenced rather than the IP which is not permanent.
 
### Useful Docker commands
 - To list images

      ```bash
        docker image ls
      ```
 
 - To tag image
 
	  ```bash
	    docker image tag go-ci-cd:latest go-ci-cd:v1.0
	  ```
 
 - To remove the tagged image
 
	  ```bash
	    docker image rm go-ci-cd:v1.0
	  ```
    
 - To publish port of container add --publish flag and add --detach flag to run in detached mode

      ```bash
	    docker run --detach --publish 8000:8000 go-ci-cd
	  ```
 - To list all containers
      ```bash
	    docker ps -all
	  ``` 
  
  - To stop, restart & remove container

     ```bash
	    docker stop <CONTAINER_ID>
        docker restart <CONTAINER_ID>
        docker rm <CONTAINER_ID>
	 ```

## Push Docker image to Amazon ECR using CLI
 - To login to ecr
     ```bash
        aws ecr get-login-password --region ap-southeast-2 | docker login --username AWS --password-stdin <ACCOUNT_ID>.dkr.ecr.ap-southeast-2.amazonaws.com
     ```
 - Tag your local Docker image 
     ```bash
            docker tag <IMAGE_ID> <ACCOUNT_ID>.dkr.ecr.ap-southeast-2.amazonaws.com/golang
     ```
 - Before pushing the image to ECR, create repository named as golang in AWS.
     ```bash
            docker push <ACCOUNT_ID>.dkr.ecr.ap-southeast-2.amazonaws.com/golang
     ```
 
