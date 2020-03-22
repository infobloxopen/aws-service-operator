/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/awslabs/aws-service-operator/pkg/client/clientset/versioned/typed/service-operator.aws/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeServiceoperatorV1alpha1 struct {
	*testing.Fake
}

func (c *FakeServiceoperatorV1alpha1) CloudFormationTemplates(namespace string) v1alpha1.CloudFormationTemplateInterface {
	return &FakeCloudFormationTemplates{c, namespace}
}

func (c *FakeServiceoperatorV1alpha1) DynamoDBs(namespace string) v1alpha1.DynamoDBInterface {
	return &FakeDynamoDBs{c, namespace}
}

func (c *FakeServiceoperatorV1alpha1) ECRRepositories(namespace string) v1alpha1.ECRRepositoryInterface {
	return &FakeECRRepositories{c, namespace}
}

func (c *FakeServiceoperatorV1alpha1) ElastiCaches(namespace string) v1alpha1.ElastiCacheInterface {
	return &FakeElastiCaches{c, namespace}
}

func (c *FakeServiceoperatorV1alpha1) S3Buckets(namespace string) v1alpha1.S3BucketInterface {
	return &FakeS3Buckets{c, namespace}
}

func (c *FakeServiceoperatorV1alpha1) SNSSubscriptions(namespace string) v1alpha1.SNSSubscriptionInterface {
	return &FakeSNSSubscriptions{c, namespace}
}

func (c *FakeServiceoperatorV1alpha1) SNSTopics(namespace string) v1alpha1.SNSTopicInterface {
	return &FakeSNSTopics{c, namespace}
}

func (c *FakeServiceoperatorV1alpha1) SQSQueues(namespace string) v1alpha1.SQSQueueInterface {
	return &FakeSQSQueues{c, namespace}
}

func (c *FakeServiceoperatorV1alpha1) Vpcs(namespace string) v1alpha1.VpcInterface {
	return &FakeVpcs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeServiceoperatorV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
