# Removing Existing Local Development Containers
docker container rm -f shifter_server_development shifter_ui_development
docker container prune -f shifter_server_development shifter_ui_development
# Removing Existing Local Development Images
docker image rm --force local.images.shifter.cloud/shifter-ui:latest
docker image rm --force local.images.shifter.cloud/shifter:latest
# Building Local Development Images for Shifter UI and Shifter Server
docker build  --no-cache -t local.images.shifter.cloud/shifter-ui:latest -f "./shifter-ui/Dockerfile" ./shifter-ui/
docker build  --no-cache -t local.images.shifter.cloud/shifter:latest -f "./shifter/Dockerfile" ./shifter/
# Docker Compose Up - Run It!
docker-compose -f docker-compose-development.yml up