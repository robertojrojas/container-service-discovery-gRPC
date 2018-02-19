package io.robertojrojas.container_service_discovery_grpc.name;

import io.grpc.Server;
import io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.NettyServerBuilder;
import io.netty.handler.ssl.ClientAuth;
import io.netty.handler.ssl.SslContextBuilder;
import io.netty.handler.ssl.SslProvider;

import java.io.File;
import java.io.IOException;
import java.net.InetSocketAddress;
import java.util.logging.Logger;
import java.util.LinkedHashSet;
import java.util.Map;
import java.util.Arrays;

/**
 * Server that manages startup/shutdown of a Name Server with TLS enabled.
 */
public class NameServerMain {
    private static final Logger logger = Logger.getLogger(NameServerMain.class.getName());
    private static final String NAME_SERVER_HOST = "NAME_SERVER_HOST";
    private static final String NAME_SERVER_PORT = "NAME_SERVER_PORT";
    private static final String  NAME_SERVER_CERT = "NAME_SERVER_CERT";
    private static final String  NAME_SERVER_PRIVATE_KEY = "NAME_SERVER_PRIVATE_KEY";
    private static final String  NAME_SERVER_CLIENT_CERT = "NAME_SERVER_CLIENT_CERT"; 
    private static final LinkedHashSet<String> requireEnvVars = new LinkedHashSet<String>(Arrays.asList(new String[]{
        NAME_SERVER_HOST, NAME_SERVER_PORT, NAME_SERVER_CERT, NAME_SERVER_PRIVATE_KEY,
        NAME_SERVER_CLIENT_CERT
    }));

    private Server server;
    private final String host;
    private final int port;
    private final String certChainFilePath;
    private final String privateKeyFilePath;
    private final String clientCertChainFilePath;


    public NameServerMain(String host,
                               int port,
                               String certChainFilePath,
                               String privateKeyFilePath,
                               String clientCertChainFilePath) {
        this.host = host;
        this.port = port;
        this.certChainFilePath = certChainFilePath;
        this.privateKeyFilePath = privateKeyFilePath;
        this.clientCertChainFilePath = clientCertChainFilePath;
    }

    private SslContextBuilder getSslContextBuilder() {
        SslContextBuilder sslClientContextBuilder = SslContextBuilder.forServer(new File(certChainFilePath),
                new File(privateKeyFilePath));
        if (clientCertChainFilePath != null) {
            sslClientContextBuilder.trustManager(new File(clientCertChainFilePath));
            sslClientContextBuilder.clientAuth(ClientAuth.OPTIONAL);
        }
        return GrpcSslContexts.configure(sslClientContextBuilder,
                SslProvider.OPENSSL);
    }

    private void start() throws IOException {
        server = NettyServerBuilder.forAddress(new InetSocketAddress(host, port))
                .addService(new NameServerImpl())
                .sslContext(getSslContextBuilder().build())
                .build()
                .start();
        logger.info("Server started, listening on " + port);
        Runtime.getRuntime().addShutdownHook(new Thread() {
            @Override
            public void run() {
                // Use stderr here since the logger may have been reset by its JVM shutdown hook.
                System.err.println("*** shutting down gRPC server since JVM is shutting down");
                NameServerMain.this.stop();
                System.err.println("*** server shut down");
            }
        });
    }

    private void stop() {
        if (server != null) {
            server.shutdown();
        }
    }

    /**
     * Await termination on the main thread since the grpc library uses daemon threads.
     */
    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    private static void checkEnvirontmentVars(Map <String, String> env) {
       if (!env.keySet().containsAll(requireEnvVars)) {
            System.out.println(
                "USAGE: NameServerMain - The following environment variables are required: " );
            for (String entry: requireEnvVars) {
                System.out.println("   - " + entry);
            }    
            System.exit(0);
       }
    }

    /**
     * Main launches the server from the command line.
     */
    public static void main(String[] args) throws IOException, InterruptedException {

        NameServerMain.checkEnvirontmentVars(System.getenv());

        String host = System.getenv(NAME_SERVER_HOST);
        int port    = Integer.parseInt(System.getenv(NAME_SERVER_PORT));
        String certChainFilePath = System.getenv(NAME_SERVER_CERT);
        String privateKeyFilePath = System.getenv(NAME_SERVER_PRIVATE_KEY);
        String clientCertChainFilePath = null; //System.getenv(NAME_SERVER_CLIENT_CERT); 

        final NameServerMain server = new NameServerMain(host,
                port,
                certChainFilePath,
                privateKeyFilePath,
                clientCertChainFilePath);
        server.start();
        server.blockUntilShutdown();
    }
}