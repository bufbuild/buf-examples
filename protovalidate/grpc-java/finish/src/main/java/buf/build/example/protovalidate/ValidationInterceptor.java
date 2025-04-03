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

package buf.build.example.protovalidate;

import build.buf.protovalidate.ValidationResult;
import build.buf.protovalidate.Validator;
import build.buf.protovalidate.exceptions.ValidationException;
import com.google.protobuf.Any;
import com.google.protobuf.Message;
import com.google.rpc.Code;
import com.google.rpc.Status;
import io.grpc.*;
import io.grpc.protobuf.StatusProto;

public class ValidationInterceptor implements ServerInterceptor {

    private final Validator validator;

    public ValidationInterceptor(Validator validator) {
        this.validator = validator;
    }

    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(
            ServerCall<ReqT, RespT> call, Metadata headers, ServerCallHandler<ReqT, RespT> next) {
        ServerCall.Listener<ReqT> listener = next.startCall(call, headers);
        return new RequestValidationServerCallListener<>(validator, call, listener);
    }

    private static class RequestValidationServerCallListener<ReqT, RespT>
            extends ForwardingServerCallListener.SimpleForwardingServerCallListener<ReqT> {

        private final Validator validator;
        private final ServerCall<ReqT, RespT> call;

        protected RequestValidationServerCallListener(
                Validator validator, ServerCall<ReqT, RespT> call, ServerCall.Listener<ReqT> delegate) {
            super(delegate);
            this.validator = validator;
            this.call = call;
        }

        @Override
        public void onHalfClose() {
            if (this.call.isReady()) {
                super.onHalfClose();
            }
        }

        @Override
        public void onMessage(ReqT message) {
            if (!(message instanceof Message)) {
                throw new IllegalArgumentException(
                        "Message is of type " + message.getClass() + ", not a " + Message.class.getName());
            }

            try {
                ValidationResult validationResult = validator.validate((Message) message);

                if (validationResult.isSuccess()) {
                    super.onMessage(message);
                } else {
                    Status status = com.google.rpc.Status.newBuilder()
                            .setCode(Code.INVALID_ARGUMENT.getNumber())
                            .setMessage(
                                    Code.INVALID_ARGUMENT.getValueDescriptor().getName())
                            .addDetails(Any.pack(validationResult.toProto()))
                            .build();

                    StatusRuntimeException sre = StatusProto.toStatusRuntimeException(status);

                    call.close(sre.getStatus(), sre.getTrailers());
                }

            } catch (ValidationException e) {
                Status status = com.google.rpc.Status.newBuilder()
                        .setCode(Code.INTERNAL.getNumber())
                        .setMessage(e.getMessage())
                        .build();

                throw (StatusProto.toStatusRuntimeException(status));
            }
        }
    }
}
