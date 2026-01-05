// Copyright 2020-2026 Buf Technologies, Inc.
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

import static org.junit.jupiter.api.Assertions.*;

import buf.build.example.protovalidate.ValidationInterceptor;
import build.buf.protovalidate.ValidatorFactory;
import build.buf.validate.FieldPath;
import build.buf.validate.FieldPathElement;
import build.buf.validate.Violation;
import build.buf.validate.Violations;
import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Timestamp;
import com.google.rpc.Status;
import invoice.v1.gen.*;
import io.grpc.*;
import io.grpc.StatusRuntimeException;
import io.grpc.protobuf.StatusProto;
import io.grpc.protobuf.services.ProtoReflectionServiceV1;
import java.net.InetSocketAddress;
import java.time.LocalDate;
import java.time.ZoneOffset;
import java.time.ZonedDateTime;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.UUID;
import java.util.stream.IntStream;
import org.junit.jupiter.api.*;

public class InvoiceServerTest {
    private static class ViolationSpec {
        public final String ruleId;
        public final String fieldPath;
        public final String message;

        public ViolationSpec(String ruleId, String fieldPath, String message) {
            this.ruleId = ruleId;
            this.fieldPath = fieldPath;
            this.message = message;
        }
    }

    private static Server server;
    private static InvoiceServiceGrpc.InvoiceServiceBlockingStub invoiceClient;

    @BeforeAll
    public static void startServer() {
        ValidationInterceptor interceptor =
                new ValidationInterceptor(ValidatorFactory.newBuilder().build());
        ServerServiceDefinition serviceDefinition = ServerInterceptors.intercept(new InvoiceService(), interceptor);
        server = Grpc.newServerBuilderForPort(0, InsecureServerCredentials.create())
                .addService(ProtoReflectionServiceV1.newInstance())
                .addService(serviceDefinition)
                .build();

        try {
            server.start();
        } catch (Exception e) {
            throw new RuntimeException("Couldn't start server for testing", e);
        }
    }

    @BeforeAll
    public static void createClient() {
        int port = ((InetSocketAddress) server.getListenSockets().get(0)).getPort();
        ManagedChannel channel = Grpc.newChannelBuilder("localhost:" + port, InsecureChannelCredentials.create())
                .build();
        invoiceClient = InvoiceServiceGrpc.newBlockingStub(channel);
    }

    @AfterAll
    public static void shutdownServer() {
        if (server != null) {
            try {
                server.shutdown();
            } catch (Exception e) {
                throw new RuntimeException(
                        "Couldn't stop server after testing. You may need to stop it forcefully.", e);
            }
        }
    }

    @Test
    @DisplayName("A valid invoice passes validation")
    public void testValidInvoice() {
        Invoice invoice = newValidInvoice();
        CreateInvoiceRequest req =
                CreateInvoiceRequest.newBuilder().setInvoice(invoice).build();
        assertDoesNotThrow(() -> invoiceClient.createInvoice(req));
    }

    @Test
    @DisplayName("InvoiceId is required")
    public void testInvoiceIdIsRequired() {
        Invoice invoice = Invoice.newBuilder()
                .mergeFrom(newValidInvoice())
                .setInvoiceId("")
                .build();
        CreateInvoiceRequest req =
                CreateInvoiceRequest.newBuilder().setInvoice(invoice).build();

        StatusRuntimeException exception =
                assertThrows(StatusRuntimeException.class, () -> invoiceClient.createInvoice(req));
        checkStatusRuntimeException(
                exception,
                Collections.singletonList(new ViolationSpec(
                        "string.uuid_empty", "invoice.invoice_id", "value is empty, which is not a valid UUID")));
    }

    @Test
    @DisplayName("Two line items cannot have the same product_id and unit price")
    public void testTwoLineItemsCannotHaveTheSameProductIdAndUnitPrice() {
        Invoice template = newValidInvoice();
        Invoice invoice = Invoice.newBuilder()
                .mergeFrom(template)
                .setLineItems(
                        1,
                        LineItem.newBuilder()
                                .mergeFrom(template.getLineItems(1))
                                .setProductId(template.getLineItems(0).getProductId())
                                .setUnitPrice(template.getLineItems(0).getUnitPrice()))
                .build();

        CreateInvoiceRequest req =
                CreateInvoiceRequest.newBuilder().setInvoice(invoice).build();

        StatusRuntimeException exception =
                assertThrows(StatusRuntimeException.class, () -> invoiceClient.createInvoice(req));
        checkStatusRuntimeException(
                exception,
                Collections.singletonList(new ViolationSpec(
                        "line_items.logically_unique",
                        "invoice.line_items",
                        "line items must be unique combinations of product_id and unit_price")));
    }

    private Invoice newValidInvoice() {
        LocalDate now = LocalDate.now();
        ZonedDateTime today = now.atStartOfDay(ZoneOffset.UTC);

        return Invoice.newBuilder()
                .setInvoiceId(UUID.randomUUID().toString())
                .setAccountId(UUID.randomUUID().toString())
                .setInvoiceDate(Timestamp.newBuilder()
                        .setSeconds(today.toEpochSecond())
                        .setNanos(0)
                        .build())
                .addAllLineItems(Arrays.asList(
                        LineItem.newBuilder()
                                .setLineItemId(UUID.randomUUID().toString())
                                .setProductId(UUID.randomUUID().toString())
                                .setQuantity(1L)
                                .setUnitPrice(1L)
                                .build(),
                        LineItem.newBuilder()
                                .setLineItemId(UUID.randomUUID().toString())
                                .setProductId(UUID.randomUUID().toString())
                                .setQuantity(1L)
                                .setUnitPrice(1L)
                                .build(),
                        LineItem.newBuilder()
                                .setLineItemId(UUID.randomUUID().toString())
                                .setProductId(UUID.randomUUID().toString())
                                .setQuantity(1L)
                                .setUnitPrice(1L)
                                .build()))
                .build();
    }

    private void checkStatusRuntimeException(StatusRuntimeException exception, List<ViolationSpec> violationSpecs) {
        // We should have messages from our advice
        Status status = StatusProto.fromThrowable(exception);
        assertNotNull(status);

        Violations violations = status.getDetailsList().stream()
                .filter(it -> it.is(Violations.class))
                .map(it -> {
                    try {
                        return it.unpack(Violations.class);
                    } catch (InvalidProtocolBufferException e) {
                        throw new RuntimeException("Couldn't unpack Violations", e);
                    }
                })
                .findFirst()
                .orElseThrow(() -> new RuntimeException("Couldn't unpack Violations"));

        List<Violation> violationList = violations.getViolationsList();
        assertEquals(
                violationSpecs.size(),
                violationList.size(),
                "Expected " + violationSpecs.size() + " violations, but got " + violationList.size());

        IntStream.range(0, violationSpecs.size()).forEach(i -> {
            ViolationSpec violationSpec = violationSpecs.get(i);
            Violation violation = violationList.get(i);

            assertEquals(
                    violationSpec.ruleId,
                    violation.getRuleId(),
                    "Expected " + violationSpec.ruleId + " but got " + violation.getRuleId());
            assertEquals(
                    violationSpec.fieldPath,
                    fieldPathString(violation.getField()),
                    "Expected " + violationSpec.fieldPath + " but got " + fieldPathString(violation.getField()));
            assertEquals(
                    violationSpec.message,
                    violation.getMessage(),
                    "Expected " + violationSpec.message + " but got " + violation.getMessage());
        });
    }

    static String fieldPathString(FieldPath fieldPath) {
        StringBuilder builder = new StringBuilder();
        for (FieldPathElement element : fieldPath.getElementsList()) {
            if (builder.length() > 0) {
                builder.append(".");
            }
            builder.append(element.getFieldName());
            switch (element.getSubscriptCase()) {
                case INDEX:
                    builder.append("[");
                    builder.append(element.getIndex());
                    builder.append("]");
                    break;
                case BOOL_KEY:
                    if (element.getBoolKey()) {
                        builder.append("[true]");
                    } else {
                        builder.append("[false]");
                    }
                    break;
                case INT_KEY:
                    builder.append("[");
                    builder.append(element.getIntKey());
                    builder.append("]");
                    break;
                case UINT_KEY:
                    builder.append("[");
                    builder.append(element.getUintKey());
                    builder.append("]");
                    break;
                case STRING_KEY:
                    builder.append("[\"");
                    builder.append(element.getStringKey().replace("\\", "\\\\").replace("\"", "\\\""));
                    builder.append("\"]");
                    break;
                case SUBSCRIPT_NOT_SET:
                    break;
            }
        }
        return builder.toString();
    }
}
