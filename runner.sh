#!/bin/zsh

# Build Image
docker build -t golist:latest .

# Run container without volume mounting
#docker run --name golist -p 8080:8080 golist:latest

# Create Volume
docker volume create datalist

# Run container with pre-created volume
#docker run -it --rm --volume datalist:/data -p 8080:8080 golist:latest
docker run --name golist --volume datalist:/data -p 8080:8080 golist:latest
