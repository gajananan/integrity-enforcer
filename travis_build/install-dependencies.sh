#!/bin/bash

OS_NAME=$(uname -s)

OPERATOR_SDK_VERSION=v1.1.0

curl -LO https://github.com/operator-framework/operator-sdk/releases/download/$OPERATOR_SDK_VERSION/operator-sdk-$OPERATOR_SDK_VERSION-x86_64-linux-gnu
chmod +x operator-sdk-$OPERATOR_SDK_VERSION-x86_64-linux-gnu && sudo mkdir -p /usr/local/bin/ && sudo cp operator-sdk-$OPERATOR_SDK_VERSION-x86_64-linux-gnu /usr/local/bin/operator-sdk && rm operator-sdk-$OPERATOR_SDK_VERSION-x86_64-linux-gnu

operator-sdk version

OPM_VERSION=v1.15.1

if [[ "$OS_NAME" == "Linux" ]]; then
    OPM_URL=https://github.com/operator-framework/operator-registry/releases/download/$OPM_VERSION/linux-amd64-opm
elif [[ "$OS_NAME" == "Darwin" ]]; then
    OPM_URL=https://github.com/operator-framework/operator-registry/releases/download/$OPM_VERSION/darwin-amd64-opm
fi

echo $GOPATH
wget -nv $OPM_URL -O $GOPATH/bin/opm
chmod +x $GOPATH/bin/opm

$GOPATH/bin/opm version


if [[ "$OS_NAME" == "Linux" ]]; then
    curl -s "https://raw.githubusercontent.com/\
             kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
fi

chmod +x ./kustomize
mv ./kustomize $GOPATH/bin/kustomize

echo "Finished setting up."
