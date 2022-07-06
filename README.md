
## Docker
 - Create Dockerfile
 - Build Docker image using multi-stage builds

	  ```bash
	    docker build --tag go-ci-cd .
	  ```
 - Run container
      ```bash
	    docker run go-ci-cd
	  ``` 
 
## Useful Docker commands
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

  

