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

import grpc
import invoice.v1.invoice_service_pb2 as invoice_service_pb2
import invoice.v1.invoice_service_pb2_grpc as invoice_service_pb2_grpc
import logging
import sys

from concurrent import futures
from grpc_reflection.v1alpha import reflection
from invoice.v1.invoice_service import InvoiceService

# Configure a logger that won't show Protovalidate logging.
logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)
handler = logging.StreamHandler(sys.stdout)
handler.setFormatter(
    logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")
)
logger.addHandler(handler)

# Server configuration
port = "50051"


def serve():
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=10),
    )
    invoice_service_pb2_grpc.add_InvoiceServiceServicer_to_server(
        InvoiceService(), server
    )
    SERVICE_NAMES = (
        invoice_service_pb2.DESCRIPTOR.services_by_name["InvoiceService"].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)

    server.add_insecure_port("[::]:" + port)
    logger.info("Invoice server started on port " + port)
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
