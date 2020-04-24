// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the ID of an Amazon EC2 Instance for use in other
// resources.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("aws:ec2/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [describe-instances in the AWS CLI reference][1].
	Filters []GetInstanceFilter `pulumi:"filters"`
	// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `passwordData` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	GetPasswordData *bool `pulumi:"getPasswordData"`
	// Retrieve Base64 encoded User Data contents into the `userDataBase64` attribute. A SHA-1 hash of the User Data contents will always be present in the `userData` attribute. Defaults to `false`.
	GetUserData *bool `pulumi:"getUserData"`
	// Specify the exact Instance ID with which to populate the data source.
	InstanceId *string `pulumi:"instanceId"`
	// A mapping of tags, each pair of which must
	// exactly match a pair on the desired Instance.
	InstanceTags map[string]interface{} `pulumi:"instanceTags"`
	// A mapping of tags assigned to the Instance.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	// The ID of the AMI used to launch the instance.
	Ami string `pulumi:"ami"`
	// The ARN of the instance.
	Arn string `pulumi:"arn"`
	// Whether or not the Instance is associated with a public IP address or not (Boolean).
	AssociatePublicIpAddress bool `pulumi:"associatePublicIpAddress"`
	// The availability zone of the Instance.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The credit specification of the Instance.
	CreditSpecifications  []GetInstanceCreditSpecification `pulumi:"creditSpecifications"`
	DisableApiTermination bool                             `pulumi:"disableApiTermination"`
	// The EBS block device mappings of the Instance.
	EbsBlockDevices []GetInstanceEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Whether the Instance is EBS optimized or not (Boolean).
	EbsOptimized bool `pulumi:"ebsOptimized"`
	// The ephemeral block device mappings of the Instance.
	EphemeralBlockDevices []GetInstanceEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	Filters               []GetInstanceFilter               `pulumi:"filters"`
	GetPasswordData       *bool                             `pulumi:"getPasswordData"`
	GetUserData           *bool                             `pulumi:"getUserData"`
	// The Id of the dedicated host the instance will be assigned to.
	HostId string `pulumi:"hostId"`
	// The name of the instance profile associated with the Instance.
	IamInstanceProfile string `pulumi:"iamInstanceProfile"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
	// The state of the instance. One of: `pending`, `running`, `shutting-down`, `terminated`, `stopping`, `stopped`. See [Instance Lifecycle](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html) for more information.
	InstanceState string                 `pulumi:"instanceState"`
	InstanceTags  map[string]interface{} `pulumi:"instanceTags"`
	// The type of the Instance.
	InstanceType string `pulumi:"instanceType"`
	// The key name of the Instance.
	KeyName string `pulumi:"keyName"`
	// The metadata options of the Instance.
	MetadataOptions []GetInstanceMetadataOption `pulumi:"metadataOptions"`
	// Whether detailed monitoring is enabled or disabled for the Instance (Boolean).
	Monitoring bool `pulumi:"monitoring"`
	// The ID of the network interface that was created with the Instance.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// Base-64 encoded encrypted password data for the instance.
	// Useful for getting the administrator password for instances running Microsoft Windows.
	// This attribute is only exported if `getPasswordData` is true.
	// See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	PasswordData string `pulumi:"passwordData"`
	// The placement group of the Instance.
	PlacementGroup string `pulumi:"placementGroup"`
	// The private DNS name assigned to the Instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC.
	PrivateDns string `pulumi:"privateDns"`
	// The private IP address assigned to the Instance.
	PrivateIp string `pulumi:"privateIp"`
	// The public DNS name assigned to the Instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC.
	PublicDns string `pulumi:"publicDns"`
	// The public IP address assigned to the Instance, if applicable. **NOTE**: If you are using an [`ec2.Eip`](https://www.terraform.io/docs/providers/aws/r/eip.html) with your instance, you should refer to the EIP's address directly and not use `publicIp`, as this field will change after the EIP is attached.
	PublicIp string `pulumi:"publicIp"`
	// The root block device mappings of the Instance
	RootBlockDevices []GetInstanceRootBlockDevice `pulumi:"rootBlockDevices"`
	// The associated security groups.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Whether the network interface performs source/destination checking (Boolean).
	SourceDestCheck bool `pulumi:"sourceDestCheck"`
	// The VPC subnet ID.
	SubnetId string `pulumi:"subnetId"`
	// A mapping of tags assigned to the Instance.
	Tags map[string]interface{} `pulumi:"tags"`
	// The tenancy of the instance: `dedicated`, `default`, `host`.
	Tenancy string `pulumi:"tenancy"`
	// SHA-1 hash of User Data supplied to the Instance.
	UserData string `pulumi:"userData"`
	// Base64 encoded contents of User Data supplied to the Instance. This attribute is only exported if `getUserData` is true.
	UserDataBase64 string `pulumi:"userDataBase64"`
	// The associated security groups in a non-default VPC.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}
