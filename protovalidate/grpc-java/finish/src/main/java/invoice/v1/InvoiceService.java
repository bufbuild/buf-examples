// Copyright 2020-2025 Buf Technologies, Inc.
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

import invoice.v1.gen.CreateInvoiceRequest;
import invoice.v1.gen.CreateInvoiceResponse;
import invoice.v1.gen.Invoice;
import invoice.v1.gen.InvoiceServiceGrpc;
import io.grpc.stub.StreamObserver;

public class InvoiceService extends InvoiceServiceGrpc.InvoiceServiceImplBase {

    @Override
    public void createInvoice(CreateInvoiceRequest request, StreamObserver<CreateInvoiceResponse> responseObserver) {
        Invoice invoice = request.getInvoice();

        // Handle the request: persist the invoice, or maybe place it on a stream
        // or queue where workers would handle its creation.

        // Create our response model.
        CreateInvoiceResponse response = CreateInvoiceResponse.newBuilder()
                .setInvoiceId(invoice.getInvoiceId())
                .setVersion(1L)
                .build();

        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
