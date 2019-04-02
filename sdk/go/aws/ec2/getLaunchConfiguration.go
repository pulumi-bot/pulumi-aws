// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a Launch Configuration.
func LookupLaunchConfiguration(ctx *pulumi.Context, args *GetLaunchConfigurationArgs) (*GetLaunchConfigurationResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("aws:ec2/getLaunchConfiguration:getLaunchConfiguration", inputs)
	if err != nil {
		return nil, err
	}
	return &GetLaunchConfigurationResult{
		AssociatePublicIpAddress: outputs["associatePublicIpAddress"],
		EbsBlockDevices: outputs["ebsBlockDevices"],
		EbsOptimized: outputs["ebsOptimized"],
		EnableMonitoring: outputs["enableMonitoring"],
		EphemeralBlockDevices: outputs["ephemeralBlockDevices"],
		IamInstanceProfile: outputs["iamInstanceProfile"],
		ImageId: outputs["imageId"],
		InstanceType: outputs["instanceType"],
		KeyName: outputs["keyName"],
		Name: outputs["name"],
		PlacementTenancy: outputs["placementTenancy"],
		RootBlockDevices: outputs["rootBlockDevices"],
		SecurityGroups: outputs["securityGroups"],
		SpotPrice: outputs["spotPrice"],
		UserData: outputs["userData"],
		VpcClassicLinkId: outputs["vpcClassicLinkId"],
		VpcClassicLinkSecurityGroups: outputs["vpcClassicLinkSecurityGroups"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getLaunchConfiguration.
type GetLaunchConfigurationArgs struct {
	// The name of the launch configuration.
	Name interface{}
}

// A collection of values returned by getLaunchConfiguration.
type GetLaunchConfigurationResult struct {
	// Whether a Public IP address is associated with the instance.
	AssociatePublicIpAddress interface{}
	// The EBS Block Devices attached to the instance.
	EbsBlockDevices interface{}
	// Whether the launched EC2 instance will be EBS-optimized.
	EbsOptimized interface{}
	// Whether Detailed Monitoring is Enabled.
	EnableMonitoring interface{}
	// The Ephemeral volumes on the instance.
	EphemeralBlockDevices interface{}
	// The IAM Instance Profile to associate with launched instances.
	IamInstanceProfile interface{}
	// The EC2 Image ID of the instance.
	ImageId interface{}
	// The Instance Type of the instance to launch.
	InstanceType interface{}
	// The Key Name that should be used for the instance.
	KeyName interface{}
	// The Name of the launch configuration.
	Name interface{}
	// The Tenancy of the instance.
	PlacementTenancy interface{}
	// The Root Block Device of the instance.
	RootBlockDevices interface{}
	// A list of associated Security Group IDS.
	SecurityGroups interface{}
	// The Price to use for reserving Spot instances.
	SpotPrice interface{}
	// The User Data of the instance.
	UserData interface{}
	// The ID of a ClassicLink-enabled VPC.
	VpcClassicLinkId interface{}
	// The IDs of one or more Security Groups for the specified ClassicLink-enabled VPC.
	VpcClassicLinkSecurityGroups interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
