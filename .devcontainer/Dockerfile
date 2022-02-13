# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.205.2/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.17-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends protobuf-compiler

ARG PB_VERSION=3.19.4
RUN ARCH="$(uname -m | sed -e 's/aarch64$/aarch_64/')" && \
    curl -fsSLO https://github.com/protocolbuffers/protobuf/releases/download/v${PB_VERSION}/protoc-${PB_VERSION}-linux-${ARCH}.zip && \
    unzip -j protoc-${PB_VERSION}-linux-${ARCH}.zip bin/protoc -d /usr/local/bin && \
    chmod +x /usr/local/bin/protoc && \
    rm protoc-${PB_VERSION}-linux-${ARCH}.zip

USER vscode
ARG PB_GEN_GO_VERSION=1.26
ARG PB_GEN_GO_GRPC_VERSION=1.1
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PB_GEN_GO_VERSION} && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PB_GEN_GO_GRPC_VERSION}
