FROM anapsix/alpine-java:8_jdk
MAINTAINER Roberto J. Rojas <robertojrojas@gmail.com>

ADD target/container_service_discovery_grpc-1.0.0.jar /app/container_service_discovery_grpc.jar

EXPOSE 8090

ENTRYPOINT ["java", "-cp",  "/app/container_service_discovery_grpc.jar", "io.robertojrojas.container_service_discovery_grpc.name.NameServerMain"]
