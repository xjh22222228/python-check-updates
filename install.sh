#!/bin/bash

set -e

if [ "$(uname -m)" != "x86_64" ]; then
	echo "Error: Unsupported architecture $(uname -m). Only x64 binaries are available." 1>&2
	exit 1
fi

if ! command -v tar >/dev/null; then
	echo "Error: tar is required to install." 1>&2
	exit 1
fi

goos=$(uname)
version=1.1.0

install() {
  file_name=pcu_${version}_${goos}_amd64
  suffix=.tar.gz
  download_name="${file_name}${suffix}"
  pcu_uri=https://github.com/xjh22222228/python-check-updates/releases/download/v${version}/${download_name}

  echo $pcu_uri

  # Remove current pkg
  rm -f "$download_name"

  curl "$pcu_uri" -OL --progress --retry 2

  tar -xvf "${download_name}"

  if [ $? -ne 0 ]; then
    rm -f "${download_name}"
    echo "Installation failed"
    exit 1
  fi

  chmod +x pcu_build/$file_name
  rm -f "${download_name}"

  cp pcu_build/$file_name /usr/local/bin/pcu
  rm -rf pcu_build
}

if [[ $goos == "Darwin" ]]; then
  goos=darwin
else
  goos=linux
fi

install


echo "Python-check-updates was installed successfully"

echo "Run pcu -h"
