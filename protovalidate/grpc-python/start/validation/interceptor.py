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

import protovalidate

from google.protobuf import any_pb2
from google.rpc import code_pb2, status_pb2
from grpc_interceptor import ServerInterceptor
from grpc_status import rpc_status


class ValidationInterceptor(ServerInterceptor):
    def intercept(
        self,
        method,
        request_or_iterator,
        context,
        method_name,
    ):
        if hasattr(request_or_iterator, "DESCRIPTOR"):
            try:
                protovalidate.validate(request_or_iterator)
            except protovalidate.ValidationError as e:
                detail = any_pb2.Any()
                detail.Pack(e.to_proto())
                rich_status = status_pb2.Status(
                    code=code_pb2.INVALID_ARGUMENT,
                    message="Validation Failure",
                    details=[detail],
                )
                status = rpc_status.to_status(rich_status)
                context.abort_with_status(status)

            return method(request_or_iterator, context)
        else:
            return method(request_or_iterator, context)
