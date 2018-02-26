FROM anapsix/alpine-java:8_jdk
MAINTAINER Roberto J. Rojas <robertojrojas@gmail.com>

ADD target/container_service_discovery_grpc-1.0.2.jar /app/container_service_discovery_grpc.jar

ENV NAME_SERVER_HOST "localhost"
ENV NAME_SERVER_PORT "8090"
ENV NAME_SERVER_CERT "/app/ssl-certs/name-server.crt"
ENV NAME_SERVER_PRIVATE_KEY "/app/ssl-certs/name-server.pem"
ENV NAME_SERVER_CLIENT_CERT ""

EXPOSE 8090

ENTRYPOINT ["java", "-cp",  "/app/container_service_discovery_grpc.jar", "io.robertojrojas.container_service_discovery_grpc.name.NameServerMain"]
