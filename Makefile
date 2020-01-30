# Note: project created with: operator-sdk version: "v0.15.1"
# operator-sdk --verbose --git-init --repo=github.com/alanconway/forwarder new forwarder
# cd forwarder
# operator-sdk add api --api-version=logging.openshift.io/v1beta1 --kind=Forwarder

DIR=pkg/apis/logging/v1beta1
SRC=$(wildcard $(DIR)/*_types.go)
GEN=$(DIR)/zz_generated.deepcopy.go
CRDS=deploy/crds/logging.openshift.io_forwarders_crd.yaml deploy/crds/logging.openshift.io_clusterforwarders_crd.yaml

all: build crds

build: gen force
	go build ./...

gen: $(GEN)
$(GEN): $(SRC)
	operator-sdk --verbose generate k8s

crds: $(CRDS)
$(CRDS): $(GEN)
	operator-sdk --verbose generate crds

.PHONY: all build gen crd force

# FIXME(alanconway) 
# doc: # doc/api.html

# TODO(alanconway) choose a better doc generator. Temporarily using:
#     go get github.com/ahmetb/gen-crd-api-reference-docs
# doc/api.html: $(DIR)/* doc/config/*
# 	gen-crd-api-reference-docs -api-dir $(API_PKG) -config doc/config/config.json -template-dir doc/config -out-file doc/api-reference.html


# TODO(alanconway) best way to generate doc? For now:

# TODO(alanconway) this is what the k8s site says they use.
# go get -u github.com/kubernetes-sigs/reference-docs
# go get -u github.com/go-openapi/loads
# go get -u github.com/go-openapi/spec

# TODO(alanconway) this is what openshift uses. Must not be in module mode.
# go get -u k8s.io/kubernetes/cmd/gendocs

# TODO(alanconway) ?
# https://github.com/DapperDox/dapperdox/releases/download/v1.2.2/dapperdox-1.2.2.linux-amd64.tgz
