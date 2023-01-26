# This Dockerfile will: 
# set up a Golang environment(using Golang Docker Image)
# install Ignite
# copy the local checkers game repository
# with the leaderboard included into the image and
# start a checkers blockchain
# expose api endpoints

# Use Golang Docker Image
FROM golang:1.18

# Set working dir
WORKDIR /home/leaderboard2

# Install tools
RUN apt-get install curl git

# Install ignite
RUN curl -L https://get.ignite.com/cli@v0.17.3! | bash

# Install dependencies
RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# Copy local repository
ADD checkers /home/leaderboard2

# Install more dependencies
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint
RUN go get golang.org/x/crypto/ssh/terminal@v0.0.0-20220411220226-7b82a4e95df4
RUN git clone https://github.com/regen-network/cosmos-proto && cd cosmos-proto/protoc-gen-gocosmos && go install

# Copy Config.yml
ARG configyml
COPY ${configyml} /home/leaderboard2/config.yml

# Copy Tests
COPY ./test.sh /home/leaderboard2
RUN chmod +x ./test.sh
ARG testchain
COPY ${testchain} /home/leaderboard2/createandplaygame.sh
RUN chmod +x ./createandplaygame.sh

# copy entrypoint script
COPY ./run-leaderboard.sh /home/leaderboard2
RUN chmod +x ./run-leaderboard.sh

ENTRYPOINT ["/bin/bash", "-c", "/home/leaderboard2/run-leaderboard.sh"]
EXPOSE 26657 1317 4500

# build with ./build-images.sh
