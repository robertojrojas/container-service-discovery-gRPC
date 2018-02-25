FROM alpine:3.4
MAINTAINER Roberto J. Rojas <robertojrojas@gmail.com>

ADD greeter-go.bin /usr/bin/greeter-go

ENV GREETER_SERVER_HOST "localhost"
ENV GREETER_SERVER_PORT "8091"
ENV GREETER_SERVER_CERT "/app/ssl-certs/server.crt"
ENV GREETER_SERVER_PRIVATE_KEY "/app/ssl-certs/server.pem"

EXPOSE 8091

ENTRYPOINT ["greeter-go"]