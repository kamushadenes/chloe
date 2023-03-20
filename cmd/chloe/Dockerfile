FROM ubuntu:latest

# Install any additional packages you need.
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy Golang binary from the host machine to the container.
COPY ./chloe /usr/local/bin/

# Set the working directory to /app.
WORKDIR /app

# Set the entry point to the Golang binary.
ENTRYPOINT ["/usr/local/bin/chloe"]