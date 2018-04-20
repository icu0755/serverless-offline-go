FROM golang:latest
RUN apt-get update && apt-get install -y libmagickwand-dev && rm -r /var/lib/apt/lists/*
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
