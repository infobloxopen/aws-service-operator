package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Vpc defines the base resource
type Vpc struct {
	metav1.TypeMeta     `json:",inline"`
	metav1.ObjectMeta   `json:"metadata"`
	Spec                VpcSpec                `json:"spec"`
	Status              VpcStatus              `json:"status"`
	Output              VpcOutput              `json:"output"`
	AdditionalResources VpcAdditionalResources `json:"additionalResources"`
}

// VpcSpec defines the Spec resource for Vpc
type VpcSpec struct {
	CloudFormationTemplateName      string `json:"cloudFormationTemplateName"`
	CloudFormationTemplateNamespace string `json:"cloudFormationTemplateNamespace"`
	RollbackCount                   int    `json:"rollbackCount"`
	VpcName                         string `json:"vpcName"`
	VpcCIDR                         string `json:"vpcCIDR"`
	PublicSubnet1CIDR               string `json:"publicSubnet1CIDR"`
	PublicSubnet2CIDR               string `json:"publicSubnet2CIDR"`
	PrivateSubnet1CIDR              string `json:"privateSubnet1CIDR"`
	PrivateSubnet2CIDR              string `json:"privateSubnet2CIDR"`
	Env                             string `json:"env"`
}

// VpcOutput defines the output resource for Vpc
type VpcOutput struct {
	VpnARN                 string `json:"vpnARN"`
	PublicSubnets          string `json:"publicSubnets"`
	PrivateSubnets         string `json:"privateSubnets"`
	PublicSubnet1          string `json:"publicSubnet1"`
	PublicSubnet2          string `json:"publicSubnet2"`
	PrivateSubnet1         string `json:"PrivateSubnet1"`
	PrivateSubnet2         string `json:"privateSubnet2"`
	NoIngressSecurityGroup string `json:"noIngressSecurityGroup"`
}

// VpcStatus holds the status of the Cloudformation template
type VpcStatus struct {
	ResourceStatus       string `json:"resourceStatus"`
	ResourceStatusReason string `json:"resourceStatusReason"`
	StackID              string `json:"stackID"`
}

// VpcAdditionalResources holds the additional resources
type VpcAdditionalResources struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VpcList defines the list attribute for the Vpc type
type VpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Vpc `json:"items"`
}

func init() {
	localSchemeBuilder.Register(addVpcTypes)
}

func addVpcTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Vpc{},
		&VpcList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
