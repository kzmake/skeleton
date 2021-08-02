# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.191.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.16, 1.15
ARG VARIANT="1.16"
FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT} as base

# [Choice] Node.js version: none, lts, 16, 14, 12, 10
ARG NODE_VERSION="lts"
RUN if [ "${NODE_VERSION}" = "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

FROM base

RUN apt-get -qq update && \
    apt-get -qq install -y zsh fish && \
    apt-get -qq install -y build-essential git wget curl && \
    apt-get -qq install -y protobuf-compiler && \
    rm -rf /var/lib/apt/lists/*

# install docker
ARG DOCKER_VERSION=20.10.7
RUN curl -Lo docker.tgz "https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz" && \
    tar --extract --file docker.tgz --strip-components 1 --directory /usr/local/bin/ && rm docker.tgz

# install kind
ARG KIND_VERSION=0.11.1
RUN curl -Lo kind https://kind.sigs.k8s.io/dl/v${KIND_VERSION}/kind-linux-amd64 && \
    chmod +x kind && mv kind /usr/local/bin/kind

# install kubectl
ARG KUBECTL_VERSION=1.21.3
RUN curl -Lo kubectl https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
    chmod +x kubectl && mv kubectl /usr/local/bin/kubectl

# install helm
RUN curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

# install skaffold
ARG SKAFFOLD_VERSION=1.28.0
RUN curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v${SKAFFOLD_VERSION}/skaffold-linux-amd64 && \
    chmod +x skaffold && mv skaffold /usr/local/bin/skaffold

# install k9s
RUN curl -sS https://webinstall.dev/k9s | /bin/bash && ln -s /root/.local/bin/k9s /usr/local/bin/k9s

# install dapr
ARG DAPR_VERSION=1.3.0
RUN curl -sS https://raw.githubusercontent.com/dapr/cli/master/install/install.sh | /bin/bash -s ${DAPR_VERSION}

# install tools
COPY Makefile /tmp/Makefile
RUN cd /tmp && make install
RUN chmod -R a+w /go/pkg

USER $USER_UID

# for kube config
ENV KUBECONFIG /kube/config
RUN mkdir -p /kube && chmod a+w /kube

WORKDIR /app
