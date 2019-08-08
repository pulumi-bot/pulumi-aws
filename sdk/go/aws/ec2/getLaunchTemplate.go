// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a Launch Template.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/launch_template.html.markdown.
func LookupLaunchTemplate(ctx *pulumi.Context, args *GetLaunchTemplateArgs) (*GetLaunchTemplateResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:ec2/getLaunchTemplate:getLaunchTemplate", inputs)
	if err != nil {
		return nil, err
	}
	return &GetLaunchTemplateResult{
		Arn: outputs["arn"],
		BlockDeviceMappings: outputs["blockDeviceMappings"],
		CreditSpecifications: outputs["creditSpecifications"],
		DefaultVersion: outputs["defaultVersion"],
		Description: outputs["description"],
		DisableApiTermination: outputs["disableApiTermination"],
		EbsOptimized: outputs["ebsOptimized"],
		ElasticGpuSpecifications: outputs["elasticGpuSpecifications"],
		IamInstanceProfiles: outputs["iamInstanceProfiles"],
		ImageId: outputs["imageId"],
		InstanceInitiatedShutdownBehavior: outputs["instanceInitiatedShutdownBehavior"],
		InstanceMarketOptions: outputs["instanceMarketOptions"],
		InstanceType: outputs["instanceType"],
		KernelId: outputs["kernelId"],
		KeyName: outputs["keyName"],
		LatestVersion: outputs["latestVersion"],
		Monitorings: outputs["monitorings"],
		Name: outputs["name"],
		NetworkInterfaces: outputs["networkInterfaces"],
		Placements: outputs["placements"],
		RamDiskId: outputs["ramDiskId"],
		SecurityGroupNames: outputs["securityGroupNames"],
		TagSpecifications: outputs["tagSpecifications"],
		Tags: outputs["tags"],
		UserData: outputs["userData"],
		VpcSecurityGroupIds: outputs["vpcSecurityGroupIds"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getLaunchTemplate.
type GetLaunchTemplateArgs struct {
	// The name of the launch template.
	Name interface{}
	Tags interface{}
}

// A collection of values returned by getLaunchTemplate.
type GetLaunchTemplateResult struct {
	// Amazon Resource Name (ARN) of the launch template.
	Arn interface{}
	// Specify volumes to attach to the instance besides the volumes specified by the AMI.
	BlockDeviceMappings interface{}
	// Customize the credit specification of the instance. See Credit
	// Specification below for more details.
	CreditSpecifications interface{}
	// The default version of the launch template.
	DefaultVersion interface{}
	// Description of the launch template.
	Description interface{}
	// If `true`, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination interface{}
	// If `true`, the launched EC2 instance will be EBS-optimized.
	EbsOptimized interface{}
	// The elastic GPU to attach to the instance. See Elastic GPU
	// below for more details.
	ElasticGpuSpecifications interface{}
	// The IAM Instance Profile to launch the instance with. See Instance Profile
	// below for more details.
	IamInstanceProfiles interface{}
	// The AMI from which to launch the instance.
	ImageId interface{}
	// Shutdown behavior for the instance. Can be `stop` or `terminate`.
	// (Default: `stop`).
	InstanceInitiatedShutdownBehavior interface{}
	// The market (purchasing) option for the instance.
	// below for details.
	InstanceMarketOptions interface{}
	// The type of the instance.
	InstanceType interface{}
	// The kernel ID.
	KernelId interface{}
	// The key name to use for the instance.
	KeyName interface{}
	// The latest version of the launch template.
	LatestVersion interface{}
	// The monitoring option for the instance.
	Monitorings interface{}
	Name interface{}
	// Customize network interfaces to be attached at instance boot time. See Network
	// Interfaces below for more details.
	NetworkInterfaces interface{}
	// The placement of the instance.
	Placements interface{}
	// The ID of the RAM disk.
	RamDiskId interface{}
	// A list of security group names to associate with. If you are creating Instances in a VPC, use
	// `vpcSecurityGroupIds` instead.
	SecurityGroupNames interface{}
	// The tags to apply to the resources during launch.
	TagSpecifications interface{}
	// (Optional) A mapping of tags to assign to the launch template.
	Tags interface{}
	// The Base64-encoded user data to provide when launching the instance.
	UserData interface{}
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
