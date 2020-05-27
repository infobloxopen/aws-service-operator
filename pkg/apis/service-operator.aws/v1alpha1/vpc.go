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
	StackName                       string `json:"stackName"`
	VpcCIDR                         string `json:"vpcCIDR"`
	Env                             string `json:"env"`
	JenkinsPeerEnable               string `json:"jenkinsPeerEnable"`
	JenkinsPeerProperties           string `json:"jenkinsPeerProperties"`
	RDSPeerEnable                   string `json:"rdsPeerEnable"`
	RDSPeerProperties               string `json:"rdsPeerProperties"`
	RedisPeerEnable                 string `json:"redisPeerEnable"`
	RedisPeerProperties             string `json:"redisPeerProperties"`
	TIDEPeerEnable                  string `json:"tidePeerEnable"`
	TIDEPeerProperties              string `json:"tidePeerProperties"`
	BastionPeerProperties           string `json:"bastionPeerProperties"`
	WBR1                            string `json:"wbr1"`
	WBR2                            string `json:"wbr2"`
	AzIndexList                     string `json:"azIndexList"`
}

// VpcOutput defines the output resource for Vpc
type VpcOutput struct {
	VpnARN            string `json:"vpnARN"`
	VpcCIDR           string `json:"VpcCIDR"`
	UtilityRouteTable string `json:"utilityRouteTable"`
	PrivateRouteTable string `json:"privateRouteTable"`
	UtilitySubnet1    string `json:"UtilitySubnet1"`
	UtilitySubnet2    string `json:"utilitySubnet2"`
	UtilitySubnet3    string `json:"UtilitySubnet3"`
	PrivateSubnet1    string `json:"privateSubnet1"`
	PrivateSubnet2    string `json:"privateSubnet2"`
	PrivateSubnet3    string `json:"privateSubnet3"`
	QuaggaSubnet1     string `json:"quaggaSubnet1"`
	QuaggaSubnet2     string `json:"quaggaSubnet2"`
	QuaggaSubnet3     string `json:"quaggaSubnet3"`
	JenkinsPCX        string `json:"jenkinsPCX"`
	RDSPCX            string `json:"rdsPCX"`
	RedisPCX          string `json:"redisPCX"`
	TIDEPCX           string `json:"tidePCX"`
	BastionPCX        string `json:"bastionPCX"`
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
