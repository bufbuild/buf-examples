ARG goVersion

FROM golang:${goVersion} AS builder
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
WORKDIR /build
COPY . ./
RUN go build -a -installsuffix nocgo -o /plugin .

FROM scratch
ARG pluginOwner pluginName pluginVersion
LABEL "build.buf.plugins.runtime_library_versions.0.name"="${pluginOwner}/${pluginName}"
LABEL "build.buf.plugins.runtime_library_versions.0.version"="${pluginVersion}"
COPY --from=builder /plugin /
ENTRYPOINT ["/plugin"]