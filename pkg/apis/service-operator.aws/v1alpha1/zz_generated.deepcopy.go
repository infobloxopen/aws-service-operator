// +build !ignore_autogenerated

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplate) DeepCopyInto(out *CloudFormationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Data = in.Data
	out.Status = in.Status
	out.Output = in.Output
	out.AdditionalResources = in.AdditionalResources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplate.
func (in *CloudFormationTemplate) DeepCopy() *CloudFormationTemplate {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudFormationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateAdditionalResources) DeepCopyInto(out *CloudFormationTemplateAdditionalResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateAdditionalResources.
func (in *CloudFormationTemplateAdditionalResources) DeepCopy() *CloudFormationTemplateAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateData) DeepCopyInto(out *CloudFormationTemplateData) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateData.
func (in *CloudFormationTemplateData) DeepCopy() *CloudFormationTemplateData {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateList) DeepCopyInto(out *CloudFormationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudFormationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateList.
func (in *CloudFormationTemplateList) DeepCopy() *CloudFormationTemplateList {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudFormationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateOutput) DeepCopyInto(out *CloudFormationTemplateOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateOutput.
func (in *CloudFormationTemplateOutput) DeepCopy() *CloudFormationTemplateOutput {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateStatus) DeepCopyInto(out *CloudFormationTemplateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateStatus.
func (in *CloudFormationTemplateStatus) DeepCopy() *CloudFormationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDB) DeepCopyInto(out *DynamoDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	out.Output = in.Output
	in.AdditionalResources.DeepCopyInto(&out.AdditionalResources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDB.
func (in *DynamoDB) DeepCopy() *DynamoDB {
	if in == nil {
		return nil
	}
	out := new(DynamoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDBAdditionalResources) DeepCopyInto(out *DynamoDBAdditionalResources) {
	*out = *in
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDBAdditionalResources.
func (in *DynamoDBAdditionalResources) DeepCopy() *DynamoDBAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(DynamoDBAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDBHashAttribute) DeepCopyInto(out *DynamoDBHashAttribute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDBHashAttribute.
func (in *DynamoDBHashAttribute) DeepCopy() *DynamoDBHashAttribute {
	if in == nil {
		return nil
	}
	out := new(DynamoDBHashAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDBList) DeepCopyInto(out *DynamoDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamoDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDBList.
func (in *DynamoDBList) DeepCopy() *DynamoDBList {
	if in == nil {
		return nil
	}
	out := new(DynamoDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDBOutput) DeepCopyInto(out *DynamoDBOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDBOutput.
func (in *DynamoDBOutput) DeepCopy() *DynamoDBOutput {
	if in == nil {
		return nil
	}
	out := new(DynamoDBOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDBRangeAttribute) DeepCopyInto(out *DynamoDBRangeAttribute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDBRangeAttribute.
func (in *DynamoDBRangeAttribute) DeepCopy() *DynamoDBRangeAttribute {
	if in == nil {
		return nil
	}
	out := new(DynamoDBRangeAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDBSpec) DeepCopyInto(out *DynamoDBSpec) {
	*out = *in
	out.RangeAttribute = in.RangeAttribute
	out.HashAttribute = in.HashAttribute
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDBSpec.
func (in *DynamoDBSpec) DeepCopy() *DynamoDBSpec {
	if in == nil {
		return nil
	}
	out := new(DynamoDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDBStatus) DeepCopyInto(out *DynamoDBStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDBStatus.
func (in *DynamoDBStatus) DeepCopy() *DynamoDBStatus {
	if in == nil {
		return nil
	}
	out := new(DynamoDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRRepository) DeepCopyInto(out *ECRRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	out.Output = in.Output
	out.AdditionalResources = in.AdditionalResources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRRepository.
func (in *ECRRepository) DeepCopy() *ECRRepository {
	if in == nil {
		return nil
	}
	out := new(ECRRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ECRRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRRepositoryAdditionalResources) DeepCopyInto(out *ECRRepositoryAdditionalResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRRepositoryAdditionalResources.
func (in *ECRRepositoryAdditionalResources) DeepCopy() *ECRRepositoryAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(ECRRepositoryAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRRepositoryList) DeepCopyInto(out *ECRRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ECRRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRRepositoryList.
func (in *ECRRepositoryList) DeepCopy() *ECRRepositoryList {
	if in == nil {
		return nil
	}
	out := new(ECRRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ECRRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRRepositoryOutput) DeepCopyInto(out *ECRRepositoryOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRRepositoryOutput.
func (in *ECRRepositoryOutput) DeepCopy() *ECRRepositoryOutput {
	if in == nil {
		return nil
	}
	out := new(ECRRepositoryOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRRepositorySpec) DeepCopyInto(out *ECRRepositorySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRRepositorySpec.
func (in *ECRRepositorySpec) DeepCopy() *ECRRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(ECRRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRRepositoryStatus) DeepCopyInto(out *ECRRepositoryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRRepositoryStatus.
func (in *ECRRepositoryStatus) DeepCopy() *ECRRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(ECRRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastiCache) DeepCopyInto(out *ElastiCache) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	out.Output = in.Output
	in.AdditionalResources.DeepCopyInto(&out.AdditionalResources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastiCache.
func (in *ElastiCache) DeepCopy() *ElastiCache {
	if in == nil {
		return nil
	}
	out := new(ElastiCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElastiCache) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastiCacheAdditionalResources) DeepCopyInto(out *ElastiCacheAdditionalResources) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastiCacheAdditionalResources.
func (in *ElastiCacheAdditionalResources) DeepCopy() *ElastiCacheAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(ElastiCacheAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastiCacheList) DeepCopyInto(out *ElastiCacheList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElastiCache, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastiCacheList.
func (in *ElastiCacheList) DeepCopy() *ElastiCacheList {
	if in == nil {
		return nil
	}
	out := new(ElastiCacheList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElastiCacheList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastiCacheOutput) DeepCopyInto(out *ElastiCacheOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastiCacheOutput.
func (in *ElastiCacheOutput) DeepCopy() *ElastiCacheOutput {
	if in == nil {
		return nil
	}
	out := new(ElastiCacheOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastiCacheSpec) DeepCopyInto(out *ElastiCacheSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastiCacheSpec.
func (in *ElastiCacheSpec) DeepCopy() *ElastiCacheSpec {
	if in == nil {
		return nil
	}
	out := new(ElastiCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastiCacheStatus) DeepCopyInto(out *ElastiCacheStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastiCacheStatus.
func (in *ElastiCacheStatus) DeepCopy() *ElastiCacheStatus {
	if in == nil {
		return nil
	}
	out := new(ElastiCacheStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bucket) DeepCopyInto(out *S3Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	out.Output = in.Output
	in.AdditionalResources.DeepCopyInto(&out.AdditionalResources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bucket.
func (in *S3Bucket) DeepCopy() *S3Bucket {
	if in == nil {
		return nil
	}
	out := new(S3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketAdditionalResources) DeepCopyInto(out *S3BucketAdditionalResources) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketAdditionalResources.
func (in *S3BucketAdditionalResources) DeepCopy() *S3BucketAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(S3BucketAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketList) DeepCopyInto(out *S3BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketList.
func (in *S3BucketList) DeepCopy() *S3BucketList {
	if in == nil {
		return nil
	}
	out := new(S3BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketLogging) DeepCopyInto(out *S3BucketLogging) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketLogging.
func (in *S3BucketLogging) DeepCopy() *S3BucketLogging {
	if in == nil {
		return nil
	}
	out := new(S3BucketLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketOutput) DeepCopyInto(out *S3BucketOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketOutput.
func (in *S3BucketOutput) DeepCopy() *S3BucketOutput {
	if in == nil {
		return nil
	}
	out := new(S3BucketOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketSpec) DeepCopyInto(out *S3BucketSpec) {
	*out = *in
	out.Logging = in.Logging
	out.Website = in.Website
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketSpec.
func (in *S3BucketSpec) DeepCopy() *S3BucketSpec {
	if in == nil {
		return nil
	}
	out := new(S3BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketStatus) DeepCopyInto(out *S3BucketStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketStatus.
func (in *S3BucketStatus) DeepCopy() *S3BucketStatus {
	if in == nil {
		return nil
	}
	out := new(S3BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketWebsite) DeepCopyInto(out *S3BucketWebsite) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketWebsite.
func (in *S3BucketWebsite) DeepCopy() *S3BucketWebsite {
	if in == nil {
		return nil
	}
	out := new(S3BucketWebsite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSSubscription) DeepCopyInto(out *SNSSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	out.Output = in.Output
	out.AdditionalResources = in.AdditionalResources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSSubscription.
func (in *SNSSubscription) DeepCopy() *SNSSubscription {
	if in == nil {
		return nil
	}
	out := new(SNSSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SNSSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSSubscriptionAdditionalResources) DeepCopyInto(out *SNSSubscriptionAdditionalResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSSubscriptionAdditionalResources.
func (in *SNSSubscriptionAdditionalResources) DeepCopy() *SNSSubscriptionAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(SNSSubscriptionAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSSubscriptionList) DeepCopyInto(out *SNSSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SNSSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSSubscriptionList.
func (in *SNSSubscriptionList) DeepCopy() *SNSSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SNSSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SNSSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSSubscriptionOutput) DeepCopyInto(out *SNSSubscriptionOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSSubscriptionOutput.
func (in *SNSSubscriptionOutput) DeepCopy() *SNSSubscriptionOutput {
	if in == nil {
		return nil
	}
	out := new(SNSSubscriptionOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSSubscriptionSpec) DeepCopyInto(out *SNSSubscriptionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSSubscriptionSpec.
func (in *SNSSubscriptionSpec) DeepCopy() *SNSSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SNSSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSSubscriptionStatus) DeepCopyInto(out *SNSSubscriptionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSSubscriptionStatus.
func (in *SNSSubscriptionStatus) DeepCopy() *SNSSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SNSSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSTopic) DeepCopyInto(out *SNSTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	out.Output = in.Output
	out.AdditionalResources = in.AdditionalResources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSTopic.
func (in *SNSTopic) DeepCopy() *SNSTopic {
	if in == nil {
		return nil
	}
	out := new(SNSTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SNSTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSTopicAdditionalResources) DeepCopyInto(out *SNSTopicAdditionalResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSTopicAdditionalResources.
func (in *SNSTopicAdditionalResources) DeepCopy() *SNSTopicAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(SNSTopicAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSTopicList) DeepCopyInto(out *SNSTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SNSTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSTopicList.
func (in *SNSTopicList) DeepCopy() *SNSTopicList {
	if in == nil {
		return nil
	}
	out := new(SNSTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SNSTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSTopicOutput) DeepCopyInto(out *SNSTopicOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSTopicOutput.
func (in *SNSTopicOutput) DeepCopy() *SNSTopicOutput {
	if in == nil {
		return nil
	}
	out := new(SNSTopicOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSTopicSpec) DeepCopyInto(out *SNSTopicSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSTopicSpec.
func (in *SNSTopicSpec) DeepCopy() *SNSTopicSpec {
	if in == nil {
		return nil
	}
	out := new(SNSTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SNSTopicStatus) DeepCopyInto(out *SNSTopicStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SNSTopicStatus.
func (in *SNSTopicStatus) DeepCopy() *SNSTopicStatus {
	if in == nil {
		return nil
	}
	out := new(SNSTopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQSQueue) DeepCopyInto(out *SQSQueue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	out.Output = in.Output
	in.AdditionalResources.DeepCopyInto(&out.AdditionalResources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQSQueue.
func (in *SQSQueue) DeepCopy() *SQSQueue {
	if in == nil {
		return nil
	}
	out := new(SQSQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQSQueue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQSQueueAdditionalResources) DeepCopyInto(out *SQSQueueAdditionalResources) {
	*out = *in
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQSQueueAdditionalResources.
func (in *SQSQueueAdditionalResources) DeepCopy() *SQSQueueAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(SQSQueueAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQSQueueList) DeepCopyInto(out *SQSQueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SQSQueue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQSQueueList.
func (in *SQSQueueList) DeepCopy() *SQSQueueList {
	if in == nil {
		return nil
	}
	out := new(SQSQueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQSQueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQSQueueOutput) DeepCopyInto(out *SQSQueueOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQSQueueOutput.
func (in *SQSQueueOutput) DeepCopy() *SQSQueueOutput {
	if in == nil {
		return nil
	}
	out := new(SQSQueueOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQSQueueSpec) DeepCopyInto(out *SQSQueueSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQSQueueSpec.
func (in *SQSQueueSpec) DeepCopy() *SQSQueueSpec {
	if in == nil {
		return nil
	}
	out := new(SQSQueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQSQueueStatus) DeepCopyInto(out *SQSQueueStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQSQueueStatus.
func (in *SQSQueueStatus) DeepCopy() *SQSQueueStatus {
	if in == nil {
		return nil
	}
	out := new(SQSQueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vpc) DeepCopyInto(out *Vpc) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	out.Output = in.Output
	out.AdditionalResources = in.AdditionalResources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vpc.
func (in *Vpc) DeepCopy() *Vpc {
	if in == nil {
		return nil
	}
	out := new(Vpc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vpc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcAdditionalResources) DeepCopyInto(out *VpcAdditionalResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcAdditionalResources.
func (in *VpcAdditionalResources) DeepCopy() *VpcAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(VpcAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcList) DeepCopyInto(out *VpcList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vpc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcList.
func (in *VpcList) DeepCopy() *VpcList {
	if in == nil {
		return nil
	}
	out := new(VpcList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcOutput) DeepCopyInto(out *VpcOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcOutput.
func (in *VpcOutput) DeepCopy() *VpcOutput {
	if in == nil {
		return nil
	}
	out := new(VpcOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcSpec) DeepCopyInto(out *VpcSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcSpec.
func (in *VpcSpec) DeepCopy() *VpcSpec {
	if in == nil {
		return nil
	}
	out := new(VpcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcStatus) DeepCopyInto(out *VpcStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcStatus.
func (in *VpcStatus) DeepCopy() *VpcStatus {
	if in == nil {
		return nil
	}
	out := new(VpcStatus)
	in.DeepCopyInto(out)
	return out
}
