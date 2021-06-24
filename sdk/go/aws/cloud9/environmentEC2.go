// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloud9

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud9 EC2 Development Environment.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloud9"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloud9.NewEnvironmentEC2(ctx, "example", &cloud9.EnvironmentEC2Args{
// 			InstanceType: pulumi.String("t2.micro"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Get the URL of the Cloud9 environment after creation:
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloud9"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := cloud9.NewEnvironmentEC2(ctx, "example", &cloud9.EnvironmentEC2Args{
// 			InstanceType: pulumi.String("t2.micro"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cloud9Url", example.ID().ApplyT(func(id string) (string, error) {
// 			return fmt.Sprintf("%v%v%v%v", "https://", _var.Region, ".console.aws.amazon.com/cloud9/ide/", id), nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
//
// Allocate a static IP to the Cloud9 environment:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloud9"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := cloud9.NewEnvironmentEC2(ctx, "example", &cloud9.EnvironmentEC2Args{
// 			InstanceType: pulumi.String("t2.micro"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		cloud9Eip, err := ec2.NewEip(ctx, "cloud9Eip", &ec2.EipArgs{
// 			Instance: cloud9Instance.ApplyT(func(cloud9Instance ec2.GetInstanceResult) (string, error) {
// 				return cloud9Instance.Id, nil
// 			}).(pulumi.StringOutput),
// 			Vpc: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cloud9PublicIp", cloud9Eip.PublicIp)
// 		return nil
// 	})
// }
// ```
type EnvironmentEC2 struct {
	pulumi.CustomResourceState

	// The ARN of the environment.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes pulumi.IntPtrOutput `pulumi:"automaticStopTimeMinutes"`
	// The description of the environment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The type of instance to connect to the environment, e.g. `t2.micro`.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The name of the environment.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn pulumi.StringOutput `pulumi:"ownerArn"`
	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of the environment (e.g. `ssh` or `ec2`)
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnvironmentEC2 registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentEC2(ctx *pulumi.Context,
	name string, args *EnvironmentEC2Args, opts ...pulumi.ResourceOption) (*EnvironmentEC2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	var resource EnvironmentEC2
	err := ctx.RegisterResource("aws:cloud9/environmentEC2:EnvironmentEC2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentEC2 gets an existing EnvironmentEC2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentEC2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentEC2State, opts ...pulumi.ResourceOption) (*EnvironmentEC2, error) {
	var resource EnvironmentEC2
	err := ctx.ReadResource("aws:cloud9/environmentEC2:EnvironmentEC2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentEC2 resources.
type environmentEC2State struct {
	// The ARN of the environment.
	Arn *string `pulumi:"arn"`
	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes *int `pulumi:"automaticStopTimeMinutes"`
	// The description of the environment.
	Description *string `pulumi:"description"`
	// The type of instance to connect to the environment, e.g. `t2.micro`.
	InstanceType *string `pulumi:"instanceType"`
	// The name of the environment.
	Name *string `pulumi:"name"`
	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn *string `pulumi:"ownerArn"`
	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of the environment (e.g. `ssh` or `ec2`)
	Type *string `pulumi:"type"`
}

type EnvironmentEC2State struct {
	// The ARN of the environment.
	Arn pulumi.StringPtrInput
	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes pulumi.IntPtrInput
	// The description of the environment.
	Description pulumi.StringPtrInput
	// The type of instance to connect to the environment, e.g. `t2.micro`.
	InstanceType pulumi.StringPtrInput
	// The name of the environment.
	Name pulumi.StringPtrInput
	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn pulumi.StringPtrInput
	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The type of the environment (e.g. `ssh` or `ec2`)
	Type pulumi.StringPtrInput
}

func (EnvironmentEC2State) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentEC2State)(nil)).Elem()
}

type environmentEC2Args struct {
	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes *int `pulumi:"automaticStopTimeMinutes"`
	// The description of the environment.
	Description *string `pulumi:"description"`
	// The type of instance to connect to the environment, e.g. `t2.micro`.
	InstanceType string `pulumi:"instanceType"`
	// The name of the environment.
	Name *string `pulumi:"name"`
	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn *string `pulumi:"ownerArn"`
	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a EnvironmentEC2 resource.
type EnvironmentEC2Args struct {
	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes pulumi.IntPtrInput
	// The description of the environment.
	Description pulumi.StringPtrInput
	// The type of instance to connect to the environment, e.g. `t2.micro`.
	InstanceType pulumi.StringInput
	// The name of the environment.
	Name pulumi.StringPtrInput
	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn pulumi.StringPtrInput
	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (EnvironmentEC2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentEC2Args)(nil)).Elem()
}

type EnvironmentEC2Input interface {
	pulumi.Input

	ToEnvironmentEC2Output() EnvironmentEC2Output
	ToEnvironmentEC2OutputWithContext(ctx context.Context) EnvironmentEC2Output
}

func (*EnvironmentEC2) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentEC2)(nil))
}

func (i *EnvironmentEC2) ToEnvironmentEC2Output() EnvironmentEC2Output {
	return i.ToEnvironmentEC2OutputWithContext(context.Background())
}

func (i *EnvironmentEC2) ToEnvironmentEC2OutputWithContext(ctx context.Context) EnvironmentEC2Output {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEC2Output)
}

func (i *EnvironmentEC2) ToEnvironmentEC2PtrOutput() EnvironmentEC2PtrOutput {
	return i.ToEnvironmentEC2PtrOutputWithContext(context.Background())
}

func (i *EnvironmentEC2) ToEnvironmentEC2PtrOutputWithContext(ctx context.Context) EnvironmentEC2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEC2PtrOutput)
}

type EnvironmentEC2PtrInput interface {
	pulumi.Input

	ToEnvironmentEC2PtrOutput() EnvironmentEC2PtrOutput
	ToEnvironmentEC2PtrOutputWithContext(ctx context.Context) EnvironmentEC2PtrOutput
}

type environmentEC2PtrType EnvironmentEC2Args

func (*environmentEC2PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentEC2)(nil))
}

func (i *environmentEC2PtrType) ToEnvironmentEC2PtrOutput() EnvironmentEC2PtrOutput {
	return i.ToEnvironmentEC2PtrOutputWithContext(context.Background())
}

func (i *environmentEC2PtrType) ToEnvironmentEC2PtrOutputWithContext(ctx context.Context) EnvironmentEC2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEC2PtrOutput)
}

// EnvironmentEC2ArrayInput is an input type that accepts EnvironmentEC2Array and EnvironmentEC2ArrayOutput values.
// You can construct a concrete instance of `EnvironmentEC2ArrayInput` via:
//
//          EnvironmentEC2Array{ EnvironmentEC2Args{...} }
type EnvironmentEC2ArrayInput interface {
	pulumi.Input

	ToEnvironmentEC2ArrayOutput() EnvironmentEC2ArrayOutput
	ToEnvironmentEC2ArrayOutputWithContext(context.Context) EnvironmentEC2ArrayOutput
}

type EnvironmentEC2Array []EnvironmentEC2Input

func (EnvironmentEC2Array) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EnvironmentEC2)(nil))
}

func (i EnvironmentEC2Array) ToEnvironmentEC2ArrayOutput() EnvironmentEC2ArrayOutput {
	return i.ToEnvironmentEC2ArrayOutputWithContext(context.Background())
}

func (i EnvironmentEC2Array) ToEnvironmentEC2ArrayOutputWithContext(ctx context.Context) EnvironmentEC2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEC2ArrayOutput)
}

// EnvironmentEC2MapInput is an input type that accepts EnvironmentEC2Map and EnvironmentEC2MapOutput values.
// You can construct a concrete instance of `EnvironmentEC2MapInput` via:
//
//          EnvironmentEC2Map{ "key": EnvironmentEC2Args{...} }
type EnvironmentEC2MapInput interface {
	pulumi.Input

	ToEnvironmentEC2MapOutput() EnvironmentEC2MapOutput
	ToEnvironmentEC2MapOutputWithContext(context.Context) EnvironmentEC2MapOutput
}

type EnvironmentEC2Map map[string]EnvironmentEC2Input

func (EnvironmentEC2Map) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EnvironmentEC2)(nil))
}

func (i EnvironmentEC2Map) ToEnvironmentEC2MapOutput() EnvironmentEC2MapOutput {
	return i.ToEnvironmentEC2MapOutputWithContext(context.Background())
}

func (i EnvironmentEC2Map) ToEnvironmentEC2MapOutputWithContext(ctx context.Context) EnvironmentEC2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEC2MapOutput)
}

type EnvironmentEC2Output struct {
	*pulumi.OutputState
}

func (EnvironmentEC2Output) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentEC2)(nil))
}

func (o EnvironmentEC2Output) ToEnvironmentEC2Output() EnvironmentEC2Output {
	return o
}

func (o EnvironmentEC2Output) ToEnvironmentEC2OutputWithContext(ctx context.Context) EnvironmentEC2Output {
	return o
}

func (o EnvironmentEC2Output) ToEnvironmentEC2PtrOutput() EnvironmentEC2PtrOutput {
	return o.ToEnvironmentEC2PtrOutputWithContext(context.Background())
}

func (o EnvironmentEC2Output) ToEnvironmentEC2PtrOutputWithContext(ctx context.Context) EnvironmentEC2PtrOutput {
	return o.ApplyT(func(v EnvironmentEC2) *EnvironmentEC2 {
		return &v
	}).(EnvironmentEC2PtrOutput)
}

type EnvironmentEC2PtrOutput struct {
	*pulumi.OutputState
}

func (EnvironmentEC2PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentEC2)(nil))
}

func (o EnvironmentEC2PtrOutput) ToEnvironmentEC2PtrOutput() EnvironmentEC2PtrOutput {
	return o
}

func (o EnvironmentEC2PtrOutput) ToEnvironmentEC2PtrOutputWithContext(ctx context.Context) EnvironmentEC2PtrOutput {
	return o
}

type EnvironmentEC2ArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentEC2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentEC2)(nil))
}

func (o EnvironmentEC2ArrayOutput) ToEnvironmentEC2ArrayOutput() EnvironmentEC2ArrayOutput {
	return o
}

func (o EnvironmentEC2ArrayOutput) ToEnvironmentEC2ArrayOutputWithContext(ctx context.Context) EnvironmentEC2ArrayOutput {
	return o
}

func (o EnvironmentEC2ArrayOutput) Index(i pulumi.IntInput) EnvironmentEC2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentEC2 {
		return vs[0].([]EnvironmentEC2)[vs[1].(int)]
	}).(EnvironmentEC2Output)
}

type EnvironmentEC2MapOutput struct{ *pulumi.OutputState }

func (EnvironmentEC2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EnvironmentEC2)(nil))
}

func (o EnvironmentEC2MapOutput) ToEnvironmentEC2MapOutput() EnvironmentEC2MapOutput {
	return o
}

func (o EnvironmentEC2MapOutput) ToEnvironmentEC2MapOutputWithContext(ctx context.Context) EnvironmentEC2MapOutput {
	return o
}

func (o EnvironmentEC2MapOutput) MapIndex(k pulumi.StringInput) EnvironmentEC2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EnvironmentEC2 {
		return vs[0].(map[string]EnvironmentEC2)[vs[1].(string)]
	}).(EnvironmentEC2Output)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentEC2Output{})
	pulumi.RegisterOutputType(EnvironmentEC2PtrOutput{})
	pulumi.RegisterOutputType(EnvironmentEC2ArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentEC2MapOutput{})
}
