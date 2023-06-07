package docker

// This go file contains the Dockerfile templates for different languages

var golang string = `# base image. Change it to latest version of golang
FROM golang:1.14.2-alpine3.11
# set working directory
WORKDIR /go/src/app
# copy go.mod and go.sum files
COPY go.mod go.sum ./
# download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download
# copy the source from the current directory to the Working Directory inside the container
COPY . .
# build the executable
RUN go build -o main .
# running the executable
CMD ["/go/src/app/main"]`

var node string = `# base image. Change it to latest version of node
FROM node:12.16.1-alpine3.9
# set working directory
WORKDIR /usr/src/app
# copy package.json and package-lock.json
COPY package*.json ./
# install dependencies
RUN npm install
# copy the source from the current directory to the Working Directory inside the container
COPY . .
# expose port 3000
EXPOSE 3000
# start app
CMD [ "npm", "start" ]`

var python string = `# base image. Change it to latest version of python
FROM python:3.8.2-alpine3.11
# set working directory
WORKDIR /usr/src/app
# copy requirements.txt
COPY requirements.txt ./
# install dependencies
RUN pip install --no-cache-dir -r requirements.txt
# copy the source from the current directory to the Working Directory inside the container
COPY . .
# expose port 5000
EXPOSE 5000
# start app
CMD [ "python", "./app.py" ]`

var ruby string = `# base image. Change it to latest version of ruby
FROM ruby:2.7.1-alpine3.11
# set working directory
WORKDIR /usr/src/app
# copy Gemfile and Gemfile.lock
COPY Gemfile Gemfile.lock ./
# install dependencies
RUN bundle install
# copy the source from the current directory to the Working Directory inside the container
COPY . .
# expose port 3000
EXPOSE 3000
# start app
CMD [ "ruby", "app.rb" ]`

var java string = `# base image. Change it to latest version of java
FROM openjdk:8-jdk-alpine
# set volume point to /tmp
VOLUME /tmp
# set argument for jar file
ARG JAR_FILE
# copy jar file to app.jar
COPY ${JAR_FILE} app.jar
# run the jar file
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/app.jar"]`
