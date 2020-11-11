OS_NAME=$(shell uname -s)

OPERATOR_SDK_VERSION=v1.1.0

curl -LO https://github.com/operator-framework/operator-sdk/releases/download/$OPERATOR_SDK_VERSION/operator-sdk-$OPERATOR_SDK_VERSION-x86_64-linux-gnu
chmod +x operator-sdk-$OPERATOR_SDK_VERSION-x86_64-linux-gnu && sudo mkdir -p /usr/local/bin/ && sudo cp operator-sdk-$OPERATOR_SDK_VERSION-x86_64-linux-gnu /usr/local/bin/operator-sdk && rm operator-sdk-$OPERATOR_SDK_VERSION-x86_64-linux-gnu


operator-sdk version

OPM_VERSION=v1.13.8

ifeq ($(OS_NAME), Linux)
    OPM_URL=https://github.com/operator-framework/operator-registry/releases/download/$(OPM_VERSION)/linux-amd64-opm
else ifeq ($(OS_NAME), Darwin)
    OPM_URL=https://github.com/operator-framework/operator-registry/releases/download/$(OPM_VERSION)/darwin-amd64-opm
endif


wget -nv $(OPM_URL) -O $(GOPATH)/bin/opm || (echo "wget returned $$? trying to fetch opm. please install opm and try again"; exit 1)
chmod +x $(GOPATH)/bin/opm
