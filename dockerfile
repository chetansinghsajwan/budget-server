FROM golang:1.24-bookworm

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
    curl \
    git \
    vim \
    && rm -rf /var/lib/apt/lists/*
