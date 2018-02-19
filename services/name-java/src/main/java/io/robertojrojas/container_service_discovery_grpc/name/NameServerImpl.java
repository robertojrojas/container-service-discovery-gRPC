package io.robertojrojas.container_service_discovery_grpc.name;

import io.grpc.stub.StreamObserver;

public class NameServerImpl extends NameServiceGrpc.NameServiceImplBase {

    public static String[] names = {"Bjarne Stroustrup", "James Gosling", "Ryan Dahl", "Rob Pike"};

    @Override
    public void generate(Name.NameRequest request, 
         StreamObserver<Name.NameResponse> responseObserver) {
        
        Name.NameResponse response = null;
        if (request.getRandomName()) {
            response = Name.NameResponse.newBuilder().setName(randomName()).build();
        } else {
            response = Name.NameResponse.newBuilder().setName(names[1]).build();
        }
        
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    public String randomName() {
       return names[((int) (Math.random()*(names.length - 0))) + 0];
    }
}