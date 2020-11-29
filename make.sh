#!/bin/bash

# ./make.sh build

version=1.1.0

GOARCH=amd64

macBuildName=pcu_${version}_darwin_${GOARCH}
winBuildName=pcu_${version}_windows_${GOARCH}
linuxBuildName=pcu_${version}_linux_${GOARCH}

# clear build/
rm -rf pcu_build
mkdir pcu_build

build() {
  # Mac os
  GOOS=darwin GOARCH=${GOARCH} go build cmd/pcu.go
  upx pcu
  mv pcu pcu_build/$macBuildName
  tar -cvf pcu_build/${macBuildName}.tar pcu_build/${macBuildName}
  gzip pcu_build/${macBuildName}.tar
  rm -f pcu_build/${macBuildName}


  # Linux
  GOOS=linux GOARCH=${GOARCH} go build cmd/pcu.go
  upx pcu
  mv pcu pcu_build/$linuxBuildName
  tar -cvf pcu_build/${linuxBuildName}.tar pcu_build/${linuxBuildName}
  gzip pcu_build/${linuxBuildName}.tar
  rm -f pcu_build/${linuxBuildName}


  # Win
  GOOS=windows GOARCH=${GOARCH} go build cmd/pcu.go
  upx pcu.exe
  mv pcu.exe pcu_build/${winBuildName}.exe
  zip -j ${winBuildName}.zip pcu_build/${winBuildName}.exe
  rm -f pcu_build/${winBuildName}.exe
  mv ${winBuildName}.zip pcu_build/${winBuildName}.zip
}

if [ $1 == build ]; then
  build
fi
