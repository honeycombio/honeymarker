#!/usr/bin/env bash

# Build deb or rpm packages for honeymarker.
set -ex

function usage() {
    echo "Usage: build-pkg.sh -v <version> -t <package_type> -m <arch>"
    exit 2
}

while getopts "v:t:m:" opt; do
    # shellcheck disable=SC2220
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

if [ "$pkg_type" == "deb" ]; then
    # for .deb, remove the leading v from version since debian doesn't permit that
    version=${version#"v"}
fi

PACKAGE_DIR=~/packages/${arch}
mkdir -p "${PACKAGE_DIR}"
fpm -s dir -n honeymarker \
    -m "Honeycomb <solutions@honeycomb.io>" \
    -p "${PACKAGE_DIR}" \
    -v "$version" \
    -t "$pkg_type" \
    -a "$arch" \
    ~/binaries/"honeymarker-linux-${arch}=/usr/bin/honeymarker"
