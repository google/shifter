# Removing Existing Local Development Containers
docker container rm -f shifter_ui_development 
docker container rm -f shifter_server_development 
docker container prune -f

# Removing Existing Local Development Images
docker image rm --force local.images.shifter.cloud/shifter-ui:latest
docker image rm --force local.images.shifter.cloud/shifter:latest
docker image prune -f

# Remove Docker Network
docker network rm shifter-network
docker network create shifter-network


# Building Local Development Images for Shifter UI and Shifter Server
docker build  --no-cache -t local.images.shifter.cloud/shifter-ui:latest -f "./shifter-ui/Dockerfile" ./shifter-ui/
docker build  --no-cache -t local.images.shifter.cloud/shifter:latest -f "./shifter/Dockerfile" ./shifter/

# Docker Compose Up - Run It!
docker-compose -f docker-compose-development.yml up --force-recreate