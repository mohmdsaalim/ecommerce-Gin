docker ps ---> list running containers
docker exec -it ecommerce_postgres psql -U ecommerce_user -d ecommerce_db ---> postgres terminal
docker exec -it ecommerce_redis redis-cli ---> redis CLI

imp     docker-compose up -d ---> start project in background 
imp     docker-compose down ---> stop project

docker logs -f ecommerce_postgres ---> follow logs






DOCKER COMMANDS – SINGLE PAGE (WITH USE CASES)

docker --version
# Check Docker is installed

docker info
# Check Docker status and system info

docker ps
# List running containers

docker ps -a
# List all containers (running + stopped)

docker start <container_name_or_id>
# Start a stopped container

docker stop <container_name_or_id>
# Stop a running container

docker restart <container_name_or_id>
# Restart a container

docker rm <container_name_or_id>
# Remove a stopped container

docker rm -f <container_name_or_id>
# Force stop and remove a container

docker rm -f $(docker ps -aq)
# Stop and remove ALL containers (dev cleanup)

docker logs <container_name_or_id>
# View container logs

docker logs -f <container_name_or_id>
# Follow logs live

docker logs --tail=50 <container_name_or_id>
# View last 50 log lines

docker exec -it <container_name_or_id> sh
# Open shell inside container

docker exec -it <container_name_or_id> bash
# Open bash shell (if available)

docker exec -it ecommerce_postgres psql -U ecommerce_user -d ecommerce_db
# Open PostgreSQL terminal inside container

docker exec -it ecommerce_redis redis-cli
# Open Redis CLI inside container

docker images
# List docker images

docker rmi <image_id>
# Remove docker image

docker image prune
# Remove unused images

docker volume ls
# List volumes (DB data stored here)

docker volume inspect <volume_name>
# Inspect volume details

docker volume prune
# Remove unused volumes (deletes DB data)

docker network ls
# List docker networks

docker network inspect <network_name>
# Inspect network details

docker-compose up
# Start project services

docker-compose up -d
# Start project services in background

docker-compose down
# Stop project services

docker-compose down -v
# Stop services and remove volumes (DB reset)

docker-compose up -d --build
# Rebuild and restart services

docker-compose logs
# View logs of all services

docker-compose logs -f
# Follow logs live

docker-compose logs postgres
# View logs of postgres service

docker-compose ps
# List services in compose project

Request
 → Routes
   → Controller
     → Service
       → Repository
         → PostgreSQL




                imp NOTES. daily use
--------------------------------------------
 
# to see the postgresDB inside the docker
 // docker exec -it ecommerce_postgres psql -U ecommerce_user -d ecommerce_db 

# to stop the docker 
 // docker compose down

# start the docker 
 // docker compose up -d

# check the container are running 
 // docker ps
