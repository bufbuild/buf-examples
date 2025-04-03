# Copyright 2025 Buf Technologies, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import invoice.v1.invoice_service_pb2 as invoice_service_pb2
import invoice.v1.invoice_service_pb2_grpc as invoice_service_pb2_grpc


class InvoiceService(invoice_service_pb2_grpc.InvoiceServiceServicer):
    def CreateInvoice(self, request, context):
        inv = request.invoice

        response = invoice_service_pb2.CreateInvoiceResponse()
        response.invoice_id = inv.invoice_id
        response.version = 1

        return response
