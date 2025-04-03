// Copyright 2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package invoice.v1;

import buf.build.example.protovalidate.ValidationInterceptor;
import build.buf.protovalidate.Validator;
import io.grpc.*;
import io.grpc.protobuf.services.ProtoReflectionServiceV1;
import java.io.IOException;
import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;

/**
 * InvoiceServer provides a class that manages a gRPC server and a main() entry point for its
 * startup.
 */
public class InvoiceServer {

    /** logger is a Logger for use by this class. */
    private static final Logger log = Logger.getLogger(InvoiceServer.class.getName());

    /** gRPC server is an io.grpc.Server. */
    private final Server gRPCServer;

    /** port is the port on which the server should listen. */
    private final int port;

    /**
     * Creates an InvoiceServer
     *
     * @param port - port to listen on
     * @param serverCredentials - ServerCredentials implementation
     */
    public InvoiceServer(int port, ServerCredentials serverCredentials, ServerServiceDefinition service) {
        this.port = port;

        gRPCServer = Grpc.newServerBuilderForPort(port, serverCredentials)
                .addService(service)
                .addService(ProtoReflectionServiceV1.newInstance())
                .build();
    }

    protected void run() throws IOException {
        gRPCServer.start();

        log.info("Server started on port " + port);

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            try {
                InvoiceServer.this.shutdown();
                System.out.println("Server shut down");
            } catch (Exception e) {
                System.err.println("Error stopping server");
                e.printStackTrace(System.err);
            }
        }));
    }

    protected void shutdown() throws InterruptedException {
        if (!gRPCServer.isTerminated()) {
            gRPCServer.shutdown().awaitTermination(1L, TimeUnit.MINUTES);
        }
    }

    private void awaitTermination() throws InterruptedException {
        if (!gRPCServer.isTerminated()) {
            gRPCServer.awaitTermination();
        }
    }

    public static void main(String[] args) throws IOException, InterruptedException {
        final BindableService service = new InvoiceService();
        final ValidationInterceptor interceptor = new ValidationInterceptor(new Validator());
        final ServerServiceDefinition serviceDefinition = ServerInterceptors.intercept(service, interceptor);

        final InvoiceServer invoiceServer =
                new InvoiceServer(50051, InsecureServerCredentials.create(), serviceDefinition);

        invoiceServer.run();
        invoiceServer.awaitTermination();
    }
}
