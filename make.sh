#!/bin/bash

# ./make.sh build

version=1.0.0

GOARCH=amd64

macBuildName=pcu_${version}_darwin_${GOARCH}
winBuildName=pcu_${version}_windows_${GOARCH}
linuxBuildName=pcu_${version}_linux_${GOARCH}

# clear build/
rm -rf build
mkdir build

build() {
  # Mac os
  $(
    GOOS=darwin GOARCH=${GOARCH} go build cmd/pcu.go \
    && mv pcu build/$macBuildName \
    && tar -cvf build/${macBuildName}.tar build/${macBuildName} \
    && gzip build/${macBuildName}.tar \
    && rm -f build/${macBuildName}
  )

  # Linux
  $(
    GOOS=linux GOARCH=${GOARCH} go build cmd/pcu.go \
    && mv pcu build/$linuxBuildName \
    && tar -cvf build/${linuxBuildName}.tar build/${linuxBuildName} \
    && gzip build/${linuxBuildName}.tar \
    && rm -f build/${linuxBuildName}
  )

  # Win
  $(
    GOOS=windows GOARCH=${GOARCH} go build cmd/pcu.go \
    && mv pcu.exe build/${winBuildName}.exe \
    && zip -j ${winBuildName}.zip build/${winBuildName}.exe \
    && rm -f build/${winBuildName}.exe \
    && mv ${winBuildName}.zip build/${winBuildName}.zip
  )
}

if [ $1 == build ]; then
  build
fi
