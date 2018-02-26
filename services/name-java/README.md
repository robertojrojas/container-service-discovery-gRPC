First [install protoc](https://github.com/google/protobuf/blob/master/README.md)
Then install [Maven](https://maven.apache.org/install.html)

To generate the Java gRPC sources:

`protoc -I ../proto --java_out ./src/main/java ../proto/name.proto --grpc_out ./src/main/java --plugin=protoc-gen-grpc=./grpc_java_plugin/protoc-gen-grpc-java-1.10.0-osx-x86_64.exe`


To build and run the Java gRPC server:

`java -cp ./target/container_service_discovery_grpc-1.0.2.jar io.robertojrojas.container_service_discovery_grpc.name.NameServerMain`

`mvn clean package exec:java -Dexec.mainClass=io.robertojrojas.container_service_discovery_grpc.name.NameServerMain -Dexec.args="$HOME"`
  
