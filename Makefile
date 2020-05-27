commitSHA := $(shell git describe --dirty --always)
dateStr := $(shell date +%s)
repo ?= github.com/awslab/aws-service-operator

.PHONY: build
build:
	go build -ldflags "-X main.commit=$(commitSHA) -X main.date=$(dateStr)" ./cmd/aws-service-operator

.PHONY: release
release:
	goreleaser --rm-dist

.PHONY: dev-release
dev-release:
	goreleaser --rm-dist --snapshot --skip-publish

.PHONY: test
test:
	go test -v -cover -race $(repo)/...

.PHONY: tag
tag:
	git tag -a ${VERSION} -s
	git push awslabs --tags

.PHONY: install-aws-codegen
install-aws-codegen:
	${MAKE} -C code-generation install

.PHONY: aws-codegen
aws-codegen:
	aws-service-operator-codegen process

.PHONY: k8s-codegen
k8s-codegen:
	./hack/update-codegen.sh

.PHONY: codegen
codegen: aws-codegen k8s-codegen

.PHONY: rebuild
rebuild: codegen build

loadcfts:
	kubectl apply -f examples/cloudformationtemplates

loadvpc:
	kubectl apply -f .vpc.yaml

.id:
	git config user.email | awk -F@ '{print $$1}' > .id

s3: .id
	# TODO: make our own namespaces
	#kubectl create ns `cat .id` || true
	# TOOD - Move the examples to a namespace away from default
	#kubectl -n `cat .id` apply -f examples/s3bucket.yaml
	kubectl apply -f examples/s3bucket.yaml

vpc: .id
	# TODO: make our own namespaces
	#kubectl create ns `cat .id` || true
	# TOOD - Move the examples to a namespace away from default
	#kubectl -n `cat .id` apply -f examples/s3bucket.yaml
	kubectl apply -f examples/vpc.yaml

status:
	# TOOD - Move the examples to a namespace away from default
	#kubectl -n `cat .id` get s3buckets
	#kubectl -n `cat .id` describe s3bucket private
	kubectl get vpcs
	kubectl describe vpc test-aws-operator-seizadi

delete:
	# TOOD - Move the examples to a namespace away from default
	#kubectl -n `cat .id` delete s3bucket private
	kubectl delete vpc test-aws-operator-seizadi

s3-status:
	# TOOD - Move the examples to a namespace away from default
	#kubectl -n `cat .id` get s3buckets
	#kubectl -n `cat .id` describe s3bucket private
	kubectl get s3buckets
	kubectl describe s3bucket test.aws-operator.seizadi.infoblox.com

s3-delete:
	# TOOD - Move the examples to a namespace away from default
	#kubectl -n `cat .id` delete s3bucket private
	kubectl delete s3bucket test.aws-operator.seizadi.infoblox.com
