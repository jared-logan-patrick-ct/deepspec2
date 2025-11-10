FROM mcr.microsoft.com/devcontainers/go:2-1.24-bookworm

# Install Node.js v24 (architecture-aware)
RUN ARCH=$(dpkg --print-architecture) && \
    if [ "$ARCH" = "amd64" ]; then \
        NODE_ARCH="x64"; \
    elif [ "$ARCH" = "arm64" ]; then \
        NODE_ARCH="arm64"; \
    else \
        echo "Unsupported architecture: $ARCH" && exit 1; \
    fi && \
    curl -LO https://nodejs.org/dist/v24.11.0/node-v24.11.0-linux-${NODE_ARCH}.tar.xz && \
    tar -xf node-v24.11.0-linux-${NODE_ARCH}.tar.xz && \
    mv node-v24.11.0-linux-${NODE_ARCH} /usr/share/node && \
    rm node-v24.11.0-linux-${NODE_ARCH}.tar.xz

# Add Node.js to PATH globally
ENV PATH="/usr/share/node/bin:${PATH}"

# Install Google Cloud SDK for the vscode user
# Based on official installation instructions from https://cloud.google.com/sdk/docs/install
USER vscode

# Install Google Cloud SDK
RUN curl -sSL https://sdk.cloud.google.com | bash

# Add gcloud and Node.js to PATH for vscode user
RUN echo 'export PATH="/home/vscode/google-cloud-sdk/bin:${PATH}"' >> /home/vscode/.bashrc
RUN echo 'export PATH="/home/vscode/google-cloud-sdk/bin:${PATH}"' >> /home/vscode/.zshrc
RUN echo 'export PATH="/usr/share/node/bin:${PATH}"' >> /home/vscode/.bashrc
RUN echo 'export PATH="/usr/share/node/bin:${PATH}"' >> /home/vscode/.zshrc

# Set PATH for current session
ENV PATH="/home/vscode/google-cloud-sdk/bin:/usr/share/node/bin:${PATH}"
