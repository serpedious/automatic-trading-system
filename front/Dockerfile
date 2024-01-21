FROM node:16-alpine3.15
WORKDIR /app
COPY package*.json ./

# Install Python
RUN apk update && \
    apk upgrade && \
    apk add --no-cache python3 make gcc g++ && \
    yarn install && \
    rm -rf /var/cache/apk/*

COPY . .
EXPOSE 3000
CMD ["yarn", "run", "dev"]