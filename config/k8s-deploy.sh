#!/bin/bash

SSL_CERTS_DIR="../services/ssl-certs"
PROTO_DIR="../services/proto"

kubectl create ns service-discovery-grpc
kubectl create secret generic service-discovery-grpc-secrets           \
     --from-file=name-server.crt=${SSL_CERTS_DIR}/name-server.crt       \
     --from-file=name-server.pem=${SSL_CERTS_DIR}/name-server.pem       \
     --from-file=greeter-server.crt=${SSL_CERTS_DIR}/greeter-server.crt \
     --from-file=greeter-server.pem=${SSL_CERTS_DIR}/greeter-server.pem \
     --from-file=client.crt=${SSL_CERTS_DIR}/client.crt \
     --from-file=client.key=${SSL_CERTS_DIR}/client.key \
     --dry-run -o yaml > service-discovery-grpc-secrets.yaml
kubectl apply -n service-discovery-grpc -f service-discovery-grpc-secrets.yaml 

kubectl create configmap service-discovery-grpc-configmap \
     --from-file=name.proto=${PROTO_DIR}/name.proto \
     --from-file=greeter.proto=${PROTO_DIR}/greeter.proto \
     --dry-run -o yaml > service-discovery-grpc-configmap.yaml
kubectl apply -n service-discovery-grpc -f service-discovery-grpc-configmap.yaml 

kubectl apply -n service-discovery-grpc -f k8s-config.yml
