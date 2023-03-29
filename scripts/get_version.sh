#!/bin/bash

version=$(git describe --tags --abbrev=0 || git rev-parse --short HEAD)

echo "${version}" > version.txt