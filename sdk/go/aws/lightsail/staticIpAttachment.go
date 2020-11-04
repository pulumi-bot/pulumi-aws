// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a static IP address attachment - relationship between a Lightsail static IP & Lightsail instance.
//
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lightsail"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testStaticIp, err := lightsail.NewStaticIp(ctx, "testStaticIp", nil)
// 		if err != nil {
// 			return err
// 		}
// 		testInstance, err := lightsail.NewInstance(ctx, "testInstance", &lightsail.InstanceArgs{
// 			AvailabilityZone: pulumi.String("us-east-1b"),
// 			BlueprintId:      pulumi.String("string"),
// 			BundleId:         pulumi.String("string"),
// 			KeyPairName:      pulumi.String("some_key_name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lightsail.NewStaticIpAttachment(ctx, "testStaticIpAttachment", &lightsail.StaticIpAttachmentArgs{
// 			StaticIpName: testStaticIp.ID(),
// 			InstanceName: testInstance.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type StaticIpAttachment struct {
	pulumi.CustomResourceState

	// The name of the Lightsail instance to attach the IP to
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The allocated static IP address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The name of the allocated static IP
	StaticIpName pulumi.StringOutput `pulumi:"staticIpName"`
}

// NewStaticIpAttachment registers a new resource with the given unique name, arguments, and options.
func NewStaticIpAttachment(ctx *pulumi.Context,
	name string, args *StaticIpAttachmentArgs, opts ...pulumi.ResourceOption) (*StaticIpAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.StaticIpName == nil {
		return nil, errors.New("invalid value for required argument 'StaticIpName'")
	}
	var resource StaticIpAttachment
	err := ctx.RegisterResource("aws:lightsail/staticIpAttachment:StaticIpAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticIpAttachment gets an existing StaticIpAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticIpAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticIpAttachmentState, opts ...pulumi.ResourceOption) (*StaticIpAttachment, error) {
	var resource StaticIpAttachment
	err := ctx.ReadResource("aws:lightsail/staticIpAttachment:StaticIpAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticIpAttachment resources.
type staticIpAttachmentState struct {
	// The name of the Lightsail instance to attach the IP to
	InstanceName *string `pulumi:"instanceName"`
	// The allocated static IP address
	IpAddress *string `pulumi:"ipAddress"`
	// The name of the allocated static IP
	StaticIpName *string `pulumi:"staticIpName"`
}

type StaticIpAttachmentState struct {
	// The name of the Lightsail instance to attach the IP to
	InstanceName pulumi.StringPtrInput
	// The allocated static IP address
	IpAddress pulumi.StringPtrInput
	// The name of the allocated static IP
	StaticIpName pulumi.StringPtrInput
}

func (StaticIpAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticIpAttachmentState)(nil)).Elem()
}

type staticIpAttachmentArgs struct {
	// The name of the Lightsail instance to attach the IP to
	InstanceName string `pulumi:"instanceName"`
	// The name of the allocated static IP
	StaticIpName string `pulumi:"staticIpName"`
}

// The set of arguments for constructing a StaticIpAttachment resource.
type StaticIpAttachmentArgs struct {
	// The name of the Lightsail instance to attach the IP to
	InstanceName pulumi.StringInput
	// The name of the allocated static IP
	StaticIpName pulumi.StringInput
}

func (StaticIpAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticIpAttachmentArgs)(nil)).Elem()
}
