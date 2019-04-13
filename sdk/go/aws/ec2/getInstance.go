// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an Amazon EC2 Instance for use in other
// resources.
func LookupInstance(ctx *pulumi.Context, args *GetInstanceArgs) (*GetInstanceResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["filters"] = args.Filters
		inputs["getPasswordData"] = args.GetPasswordData
		inputs["getUserData"] = args.GetUserData
		inputs["instanceId"] = args.InstanceId
		inputs["instanceTags"] = args.InstanceTags
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:ec2/getInstance:getInstance", inputs)
	if err != nil {
		return nil, err
	}
	return &GetInstanceResult{
		Ami: outputs["ami"],
		Arn: outputs["arn"],
		AssociatePublicIpAddress: outputs["associatePublicIpAddress"],
		AvailabilityZone: outputs["availabilityZone"],
		CreditSpecifications: outputs["creditSpecifications"],
		DisableApiTermination: outputs["disableApiTermination"],
		EbsBlockDevices: outputs["ebsBlockDevices"],
		EbsOptimized: outputs["ebsOptimized"],
		EphemeralBlockDevices: outputs["ephemeralBlockDevices"],
		Filters: outputs["filters"],
		GetPasswordData: outputs["getPasswordData"],
		GetUserData: outputs["getUserData"],
		HostId: outputs["hostId"],
		IamInstanceProfile: outputs["iamInstanceProfile"],
		InstanceId: outputs["instanceId"],
		InstanceState: outputs["instanceState"],
		InstanceTags: outputs["instanceTags"],
		InstanceType: outputs["instanceType"],
		KeyName: outputs["keyName"],
		Monitoring: outputs["monitoring"],
		NetworkInterfaceId: outputs["networkInterfaceId"],
		PasswordData: outputs["passwordData"],
		PlacementGroup: outputs["placementGroup"],
		PrivateDns: outputs["privateDns"],
		PrivateIp: outputs["privateIp"],
		PublicDns: outputs["publicDns"],
		PublicIp: outputs["publicIp"],
		RootBlockDevices: outputs["rootBlockDevices"],
		SecurityGroups: outputs["securityGroups"],
		SourceDestCheck: outputs["sourceDestCheck"],
		SubnetId: outputs["subnetId"],
		Tags: outputs["tags"],
		Tenancy: outputs["tenancy"],
		UserData: outputs["userData"],
		UserDataBase64: outputs["userDataBase64"],
		VpcSecurityGroupIds: outputs["vpcSecurityGroupIds"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getInstance.
type GetInstanceArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [describe-instances in the AWS CLI reference][1].
	Filters interface{}
	// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	GetPasswordData interface{}
	// Retrieve Base64 encoded User Data contents into the `user_data_base64` attribute. A SHA-1 hash of the User Data contents will always be present in the `user_data` attribute. Defaults to `false`.
	GetUserData interface{}
	// Specify the exact Instance ID with which to populate the data source.
	InstanceId interface{}
	// A mapping of tags, each pair of which must
	// exactly match a pair on the desired Instance.
	InstanceTags interface{}
	Tags interface{}
}

// A collection of values returned by getInstance.
type GetInstanceResult struct {
	// The ID of the AMI used to launch the instance.
	Ami interface{}
	// The ARN of the instance.
	Arn interface{}
	// Whether or not the Instance is associated with a public IP address or not (Boolean).
	AssociatePublicIpAddress interface{}
	// The availability zone of the Instance.
	AvailabilityZone interface{}
	// The credit specification of the Instance.
	CreditSpecifications interface{}
	DisableApiTermination interface{}
	// The EBS block device mappings of the Instance.
	EbsBlockDevices interface{}
	// Whether the Instance is EBS optimized or not (Boolean).
	EbsOptimized interface{}
	// The ephemeral block device mappings of the Instance.
	EphemeralBlockDevices interface{}
	Filters interface{}
	GetPasswordData interface{}
	GetUserData interface{}
	// The Id of the dedicated host the instance will be assigned to.
	HostId interface{}
	// The name of the instance profile associated with the Instance.
	IamInstanceProfile interface{}
	InstanceId interface{}
	InstanceState interface{}
	InstanceTags interface{}
	// The type of the Instance.
	InstanceType interface{}
	// The key name of the Instance.
	KeyName interface{}
	// Whether detailed monitoring is enabled or disabled for the Instance (Boolean).
	Monitoring interface{}
	// The ID of the network interface that was created with the Instance.
	NetworkInterfaceId interface{}
	// Base-64 encoded encrypted password data for the instance.
	// Useful for getting the administrator password for instances running Microsoft Windows.
	// This attribute is only exported if `get_password_data` is true.
	// See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	PasswordData interface{}
	// The placement group of the Instance.
	PlacementGroup interface{}
	// The private DNS name assigned to the Instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC.
	PrivateDns interface{}
	// The private IP address assigned to the Instance.
	PrivateIp interface{}
	// The public DNS name assigned to the Instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC.
	PublicDns interface{}
	// The public IP address assigned to the Instance, if applicable. **NOTE**: If you are using an [`aws_eip`](https://www.terraform.io/docs/providers/aws/r/eip.html) with your instance, you should refer to the EIP's address directly and not use `public_ip`, as this field will change after the EIP is attached.
	PublicIp interface{}
	// The root block device mappings of the Instance
	RootBlockDevices interface{}
	// The associated security groups.
	SecurityGroups interface{}
	// Whether the network interface performs source/destination checking (Boolean).
	SourceDestCheck interface{}
	// The VPC subnet ID.
	SubnetId interface{}
	// A mapping of tags assigned to the Instance.
	Tags interface{}
	// The tenancy of the instance: `dedicated`, `default`, `host`.
	Tenancy interface{}
	// SHA-1 hash of User Data supplied to the Instance.
	UserData interface{}
	// Base64 encoded contents of User Data supplied to the Instance. Valid UTF-8 contents can be decoded with the [`base64decode` function](https://www.terraform.io/docs/configuration/functions/base64decode.html). This attribute is only exported if `get_user_data` is true.
	UserDataBase64 interface{}
	// The associated security groups in a non-default VPC.
	VpcSecurityGroupIds interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
