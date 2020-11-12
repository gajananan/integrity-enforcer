# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This repo is build in Travis-ci by default;
# Override this variable in local env.
TRAVIS_BUILD ?= 1


.PHONY: fmt lint test coverage build build-images


############################################################
# format section
############################################################

# All available format: format-go format-protos format-python
# Default value will run all formats, override these make target with your requirements:
#    eg: fmt: format-go format-protos
fmt: format-go


############################################################
# check section
############################################################

check: lint

# All available linters: lint-dockerfiles lint-scripts lint-yaml lint-copyright-banner lint-go lint-python lint-helm lint-markdown lint-sass lint-typescript lint-protos
# Default value will run all linters, override these make target with your requirements:
#    eg: lint: lint-go lint-yaml
lint: lint-all


############################################################
# test section
############################################################

test:
	@go test ${TESTARGS} `go list ./... | grep -v test/e2e`

############################################################
# coverage section
############################################################

coverage:
	@build/common/scripts/codecov.sh


############################################################
# build section
############################################################

build:


############################################################
# images section
############################################################

build-images:
	./develop/scripts/build_images.sh


############################################################
# bundle section
############################################################

build-bundle:
	- ./develop/scripts/build_bundle.sh

############################################################
# clean section
############################################################
clean::

############################################################
# check copyright section
############################################################
copyright-check:
	./build/copyright-check.sh $(TRAVIS_BRANCH)


############################################################
# e2e test section
############################################################
.PHONY: kind-bootstrap-cluster
kind-bootstrap-cluster: kind-create-cluster install-crds kind-deploy-controller install-resources

.PHONY: kind-bootstrap-cluster-dev
kind-bootstrap-cluster-dev: kind-create-cluster install-crds install-resources

check-env:
ifndef DOCKER_USER
	$(error DOCKER_USER is undefined)
endif
ifndef DOCKER_PASS
	$(error DOCKER_PASS is undefined)
endif

kind-deploy-controller: check-env
	@echo installing config policy controller

kind-create-cluster:
	@echo "creating cluster"
	kind create cluster --name test-managed
	kind get kubeconfig --name test-managed > $(PWD)/kubeconfig_managed

kind-delete-cluster:
	kind delete cluster --name test-managed

install-crds:
	@echo installing crds

install-resources:
	@echo creating namespaces

e2e-test:
	${GOPATH}/bin/ginkgo -v --slowSpecThreshold=10 test/e2e

############################################################
# e2e test coverage
############################################################
build-instrumented:
	go test -covermode=atomic -coverpkg=github.com/open-cluster-management/$(IMG)... -c -tags e2e ./cmd/manager -o build/_output/bin/$(IMG)-instrumented

run-instrumented:
	WATCH_NAMESPACE="managed" ./build/_output/bin/$(IMG)-instrumented -test.run "^TestRunMain$$" -test.coverprofile=coverage_e2e.out &>/dev/null &

stop-instrumented:
	ps -ef | grep 'config-po' | grep -v grep | awk '{print $$2}' | xargs kill

coverage-merge:
	@echo merging the coverage report
	gocovmerge $(PWD)/coverage_* >> coverage.out
	cat coverage.out
