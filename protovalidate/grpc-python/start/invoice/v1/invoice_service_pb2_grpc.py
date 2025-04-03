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

# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from invoice.v1 import invoice_service_pb2 as invoice_dot_v1_dot_invoice__service__pb2


class InvoiceServiceStub(object):
    """InvoiceService is a simple CRUD service for managing invoices.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateInvoice = channel.unary_unary(
                '/invoice.v1.InvoiceService/CreateInvoice',
                request_serializer=invoice_dot_v1_dot_invoice__service__pb2.CreateInvoiceRequest.SerializeToString,
                response_deserializer=invoice_dot_v1_dot_invoice__service__pb2.CreateInvoiceResponse.FromString,
                _registered_method=True)


class InvoiceServiceServicer(object):
    """InvoiceService is a simple CRUD service for managing invoices.
    """

    def CreateInvoice(self, request, context):
        """CreateInvoice creates a new invoice.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_InvoiceServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateInvoice': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateInvoice,
                    request_deserializer=invoice_dot_v1_dot_invoice__service__pb2.CreateInvoiceRequest.FromString,
                    response_serializer=invoice_dot_v1_dot_invoice__service__pb2.CreateInvoiceResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'invoice.v1.InvoiceService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('invoice.v1.InvoiceService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class InvoiceService(object):
    """InvoiceService is a simple CRUD service for managing invoices.
    """

    @staticmethod
    def CreateInvoice(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/invoice.v1.InvoiceService/CreateInvoice',
            invoice_dot_v1_dot_invoice__service__pb2.CreateInvoiceRequest.SerializeToString,
            invoice_dot_v1_dot_invoice__service__pb2.CreateInvoiceResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
