# Copyright 2020-2025 Buf Technologies, Inc.
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

import datetime
import grpc
import invoice.v1.invoice_service_pb2 as invoice_service_pb2
import invoice.v1.invoice_service_pb2_grpc as invoice_service_pb2_grpc
import protovalidate
import unittest
import uuid

from concurrent import futures
from dataclasses import dataclass
from grpc_reflection.v1alpha import reflection
from grpc_status import rpc_status

from invoice.v1 import invoice_pb2
from invoice.v1.invoice_service import InvoiceService
from protovalidate.internal import field_path
from validation.interceptor import ValidationInterceptor


@dataclass
class ViolationSpec:
    constraint_id: str
    field_path: str
    message: str


class InvoiceServerTest(unittest.TestCase):
    def setUp(self):
        self.port = "50052"
        self.server = self.start_server()
        self.client = self.create_client()

    def tearDown(self):
        self.server.stop(None)

    def test_a_valid_invoice_can_be_created(self):
        req = invoice_service_pb2.CreateInvoiceRequest()
        req.invoice.CopyFrom(valid_invoice())
        response = self.client.CreateInvoice(req)
        assert hasattr(response, "invoice_id")
        assert response.invoice_id
        assert hasattr(response, "version")
        assert response.version == 1

    def test_invoice_id_must_be_a_uuid(self):
        invoice = valid_invoice()
        invoice.invoice_id = ""
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "string.uuid_empty",
                    "invoice.invoice_id",
                    "value is empty, which is not a valid UUID",
                ),
            ],
        )

    def test_account_id_must_be_a_uuid(self):
        invoice = valid_invoice()
        invoice.account_id = ""
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "string.uuid_empty",
                    "invoice.account_id",
                    "value is empty, which is not a valid UUID",
                ),
            ],
        )

    def test_invoice_date_must_not_have_a_time_component(self):
        invoice = valid_invoice()
        invoice.invoice_date = datetime.datetime.now()
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "invoice_date.no_time",
                    "invoice.invoice_date",
                    "time component should be zero",
                ),
            ],
        )

    def test_at_least_one_item_must_be_provided(self):
        invoice = valid_invoice()
        while len(invoice.line_items) > 0:
            invoice.line_items.pop(0)
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "repeated.min_items",
                    "invoice.line_items",
                    "value must contain at least 1 item(s)",
                ),
            ],
        )

    def test_two_line_items_cannot_have_the_same_identifier(self):
        invoice = valid_invoice()
        invoice.line_items[0].line_item_id = invoice.line_items[1].line_item_id
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "line_items.unique_line_item_id",
                    "invoice.line_items",
                    "all line_item_id values must be unique",
                ),
            ],
        )

    def test_two_line_items_cannot_have_the_same_product_id_and_unit_price(self):
        invoice = valid_invoice()
        invoice.line_items[0].product_id = invoice.line_items[1].product_id
        invoice.line_items[0].unit_price = invoice.line_items[1].unit_price
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "line_items.logically_unique",
                    "invoice.line_items",
                    "line items must be unique combinations of product_id and unit_price",
                ),
            ],
        )

    def test_line_item_a_line_item_cannot_have_the_same_line_item_id_and_product_id(
        self,
    ):
        invoice = valid_invoice()
        invoice.line_items[0].line_item_id = invoice.line_items[0].product_id
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "line_item_id.not.product_id",
                    "invoice.line_items[0]",
                    "line_item_id can't also be a product_id",
                ),
            ],
        )

    def test_line_item_line_item_id_must_be_a_uuid(self):
        invoice = valid_invoice()
        invoice.line_items[0].line_item_id = "this is not a uuid"
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "string.uuid",
                    "invoice.line_items[0].line_item_id",
                    "value must be a valid UUID",
                ),
            ],
        )

    def test_line_item_product_id_must_be_a_uuid(self):
        invoice = valid_invoice()
        invoice.line_items[0].product_id = "this is not a uuid"
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "string.uuid",
                    "invoice.line_items[0].product_id",
                    "value must be a valid UUID",
                ),
            ],
        )

    def test_line_item_quantity_must_be_greater_than_0(self):
        invoice = valid_invoice()
        invoice.line_items[0].quantity = 0
        self.check_violations(
            invoice,
            [
                ViolationSpec(
                    "uint64.gt",
                    "invoice.line_items[0].quantity",
                    "value must be greater than 0",
                ),
            ],
        )

    def test_line_item_unit_price_can_be_zero(self):
        req = invoice_service_pb2.CreateInvoiceRequest()
        invoice = valid_invoice()
        invoice.line_items[0].unit_price = 0
        req.invoice.CopyFrom(valid_invoice())
        response = self.client.CreateInvoice(req)
        assert hasattr(response, "invoice_id")
        assert response.invoice_id
        assert hasattr(response, "version")
        assert response.version == 1

    def test_line_item_unit_price_can_be_nonzero(self):
        req = invoice_service_pb2.CreateInvoiceRequest()
        invoice = valid_invoice()
        invoice.line_items[0].unit_price = 1
        req.invoice.CopyFrom(valid_invoice())
        response = self.client.CreateInvoice(req)
        assert hasattr(response, "invoice_id")
        assert response.invoice_id
        assert hasattr(response, "version")
        assert response.version == 1

    def check_violations(self, invoice, violation_specs):
        all_violations = None

        try:
            req = invoice_service_pb2.CreateInvoiceRequest()
            req.invoice.CopyFrom(invoice)
            self.client.CreateInvoice(req)
        except grpc.RpcError as err:
            status = rpc_status.from_call(err)
            assert status.code == grpc.StatusCode.INVALID_ARGUMENT.value[0]
            for detail in status.details:
                assert detail.Is(protovalidate.Violations.DESCRIPTOR)
                violations = protovalidate.Violations()
                detail.Unpack(violations)
                all_violations = violations.violations

        assert all_violations is not None
        assert len(violation_specs) == len(all_violations)
        for index, spec in enumerate(violation_specs):
            violation = all_violations[index]
            assert spec.constraint_id == violation.constraint_id
            assert spec.message == violation.message
            assert spec.field_path == field_path.string(violation.field)

    def start_server(self):
        server = grpc.server(
            futures.ThreadPoolExecutor(max_workers=10),
            interceptors=(ValidationInterceptor(),),
        )

        invoice_service_pb2_grpc.add_InvoiceServiceServicer_to_server(
            InvoiceService(), server
        )
        SERVICE_NAMES = (
            invoice_service_pb2.DESCRIPTOR.services_by_name["InvoiceService"].full_name,
            reflection.SERVICE_NAME,
        )
        reflection.enable_server_reflection(SERVICE_NAMES, server)

        server.add_insecure_port("[::]:" + self.port)
        server.start()
        return server

    def create_client(self):
        channel = grpc.insecure_channel("localhost:" + self.port)
        return invoice_service_pb2_grpc.InvoiceServiceStub(channel)


def valid_invoice():
    today_start = datetime.datetime.now().replace(
        hour=0, minute=0, second=0, microsecond=0
    )

    invoice = invoice_pb2.Invoice()
    invoice.invoice_id = str(uuid.uuid4())
    invoice.account_id = str(uuid.uuid4())
    invoice.invoice_date = today_start
    line_item_0 = invoice_pb2.LineItem()
    line_item_0.line_item_id = str(uuid.uuid4())
    line_item_0.product_id = str(uuid.uuid4())
    line_item_0.quantity = 1
    line_item_0.unit_price = 1
    invoice.line_items.append(line_item_0)
    line_item_1 = invoice_pb2.LineItem()
    line_item_1.line_item_id = str(uuid.uuid4())
    line_item_1.product_id = str(uuid.uuid4())
    line_item_1.quantity = 1
    line_item_1.unit_price = 1
    invoice.line_items.append(line_item_1)
    line_item_2 = invoice_pb2.LineItem()
    line_item_2.line_item_id = str(uuid.uuid4())
    line_item_2.product_id = str(uuid.uuid4())
    line_item_2.quantity = 1
    line_item_2.unit_price = 1
    invoice.line_items.append(line_item_2)

    return invoice
