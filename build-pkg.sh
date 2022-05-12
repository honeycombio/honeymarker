#!/bin/bash

# Build deb or rpm packages for honeymarker.
set -e

function usage() {
    echo "Usage: build-pkg.sh -v <version> -t <package_type>"
    exit 2
}

while getopts "v:t:m:" opt; do
    case "$opt" in
    v)
        version=$OPTARG
        ;;
    t)
        pkg_type=$OPTARG
        ;;
    m)
        arch=$OPTARG
        ;;
    esac
done

if [ -z "$version" ] || [ -z "$pkg_type" ] || [ -z "$arch" ]; then
    usage
fi

PACKAGE_DIR=~/packages/${arch}
mkdir -p ${PACKAGE_DIR}
fpm -s dir -n honeymarker \
    -m "Honeycomb <solutions@honeycomb.io>" \
    -p ${PACKAGE_DIR} \
    -v $version \
    -t $pkg_type \
    -a $arch \
    ~/binaries/honeymarker-linux-${arch}=/usr/bin/honeymarker
