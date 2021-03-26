FROM node:14.4.0-alpine as build-stage
WORKDIR /app
COPY package*.json ./

RUN apk update && \
    apk upgrade && \
    apk add --no-cache make gcc g++ python && \
    yarn install && \
    rm -rf /var/cache/apk/*

COPY . .
EXPOSE 3000
CMD ["yarn", "run", "dev"]