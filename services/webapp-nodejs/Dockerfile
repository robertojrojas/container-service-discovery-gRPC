FROM node:boron

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY package.json .
# For npm@5 or later, copy package-lock.json as well
# COPY package.json package-lock.json .

RUN npm install

# Bundle app source
COPY . .

ENV NAME_PROTO_PATH      "/app/proto/name.proto"
ENV NAME_SERVICE_HOST    "localhost"
ENV NAME_SERVICE_PORT    "8090"
ENV GREETER_PROTO_PATH   "/app/proto/greeter.proto"
ENV GREETER_SERVICE_HOST "localhost"
ENV GREETER_SERVICE_PORT "8091"
ENV WEBAPP_SERVER_PORT   "8080"
ENV WEBAPP_SERVER_HOST   "0.0.0.0"
ENV SERVER_CERT          "/app/ssl-certs/server.crt"
ENV SERVER_PRIVATE_CERT  "/app/ssl-certs/client.crt"
ENV CLIENT_KEY           "/app/ssl-certs/client.key"

EXPOSE 8080
CMD [ "npm", "start" ]

