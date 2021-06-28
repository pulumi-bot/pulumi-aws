// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an API Gateway VPC Link.
//
// > **Note:** Amazon API Gateway Version 1 VPC Links enable private integrations that connect REST APIs to private resources in a VPC.
// To enable private integration for HTTP APIs, use the `Amazon API Gateway Version 2 VPC Link` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			Internal:         pulumi.Bool(true),
// 			LoadBalancerType: pulumi.String("network"),
// 			SubnetMappings: lb.LoadBalancerSubnetMappingArray{
// 				&lb.LoadBalancerSubnetMappingArgs{
// 					SubnetId: pulumi.String("12345"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewVpcLink(ctx, "exampleVpcLink", &apigateway.VpcLinkArgs{
// 			Description: pulumi.String("example description"),
// 			TargetArn: pulumi.String{
// 				exampleLoadBalancer.Arn,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// API Gateway VPC Link can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/vpcLink:VpcLink example <vpc_link_id>
// ```
type VpcLink struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the VPC link.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name used to label and identify the VPC link.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
}

// NewVpcLink registers a new resource with the given unique name, arguments, and options.
func NewVpcLink(ctx *pulumi.Context,
	name string, args *VpcLinkArgs, opts ...pulumi.ResourceOption) (*VpcLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetArn == nil {
		return nil, errors.New("invalid value for required argument 'TargetArn'")
	}
	var resource VpcLink
	err := ctx.RegisterResource("aws:apigateway/vpcLink:VpcLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcLink gets an existing VpcLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcLinkState, opts ...pulumi.ResourceOption) (*VpcLink, error) {
	var resource VpcLink
	err := ctx.ReadResource("aws:apigateway/vpcLink:VpcLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcLink resources.
type vpcLinkState struct {
	Arn *string `pulumi:"arn"`
	// The description of the VPC link.
	Description *string `pulumi:"description"`
	// The name used to label and identify the VPC link.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn *string `pulumi:"targetArn"`
}

type VpcLinkState struct {
	Arn pulumi.StringPtrInput
	// The description of the VPC link.
	Description pulumi.StringPtrInput
	// The name used to label and identify the VPC link.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn pulumi.StringPtrInput
}

func (VpcLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcLinkState)(nil)).Elem()
}

type vpcLinkArgs struct {
	// The description of the VPC link.
	Description *string `pulumi:"description"`
	// The name used to label and identify the VPC link.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn string `pulumi:"targetArn"`
}

// The set of arguments for constructing a VpcLink resource.
type VpcLinkArgs struct {
	// The description of the VPC link.
	Description pulumi.StringPtrInput
	// The name used to label and identify the VPC link.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn pulumi.StringInput
}

func (VpcLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcLinkArgs)(nil)).Elem()
}

type VpcLinkInput interface {
	pulumi.Input

	ToVpcLinkOutput() VpcLinkOutput
	ToVpcLinkOutputWithContext(ctx context.Context) VpcLinkOutput
}

func (*VpcLink) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcLink)(nil))
}

func (i *VpcLink) ToVpcLinkOutput() VpcLinkOutput {
	return i.ToVpcLinkOutputWithContext(context.Background())
}

func (i *VpcLink) ToVpcLinkOutputWithContext(ctx context.Context) VpcLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcLinkOutput)
}

func (i *VpcLink) ToVpcLinkPtrOutput() VpcLinkPtrOutput {
	return i.ToVpcLinkPtrOutputWithContext(context.Background())
}

func (i *VpcLink) ToVpcLinkPtrOutputWithContext(ctx context.Context) VpcLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcLinkPtrOutput)
}

type VpcLinkPtrInput interface {
	pulumi.Input

	ToVpcLinkPtrOutput() VpcLinkPtrOutput
	ToVpcLinkPtrOutputWithContext(ctx context.Context) VpcLinkPtrOutput
}

type vpcLinkPtrType VpcLinkArgs

func (*vpcLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcLink)(nil))
}

func (i *vpcLinkPtrType) ToVpcLinkPtrOutput() VpcLinkPtrOutput {
	return i.ToVpcLinkPtrOutputWithContext(context.Background())
}

func (i *vpcLinkPtrType) ToVpcLinkPtrOutputWithContext(ctx context.Context) VpcLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcLinkPtrOutput)
}

// VpcLinkArrayInput is an input type that accepts VpcLinkArray and VpcLinkArrayOutput values.
// You can construct a concrete instance of `VpcLinkArrayInput` via:
//
//          VpcLinkArray{ VpcLinkArgs{...} }
type VpcLinkArrayInput interface {
	pulumi.Input

	ToVpcLinkArrayOutput() VpcLinkArrayOutput
	ToVpcLinkArrayOutputWithContext(context.Context) VpcLinkArrayOutput
}

type VpcLinkArray []VpcLinkInput

func (VpcLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VpcLink)(nil))
}

func (i VpcLinkArray) ToVpcLinkArrayOutput() VpcLinkArrayOutput {
	return i.ToVpcLinkArrayOutputWithContext(context.Background())
}

func (i VpcLinkArray) ToVpcLinkArrayOutputWithContext(ctx context.Context) VpcLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcLinkArrayOutput)
}

// VpcLinkMapInput is an input type that accepts VpcLinkMap and VpcLinkMapOutput values.
// You can construct a concrete instance of `VpcLinkMapInput` via:
//
//          VpcLinkMap{ "key": VpcLinkArgs{...} }
type VpcLinkMapInput interface {
	pulumi.Input

	ToVpcLinkMapOutput() VpcLinkMapOutput
	ToVpcLinkMapOutputWithContext(context.Context) VpcLinkMapOutput
}

type VpcLinkMap map[string]VpcLinkInput

func (VpcLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VpcLink)(nil))
}

func (i VpcLinkMap) ToVpcLinkMapOutput() VpcLinkMapOutput {
	return i.ToVpcLinkMapOutputWithContext(context.Background())
}

func (i VpcLinkMap) ToVpcLinkMapOutputWithContext(ctx context.Context) VpcLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcLinkMapOutput)
}

type VpcLinkOutput struct {
	*pulumi.OutputState
}

func (VpcLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcLink)(nil))
}

func (o VpcLinkOutput) ToVpcLinkOutput() VpcLinkOutput {
	return o
}

func (o VpcLinkOutput) ToVpcLinkOutputWithContext(ctx context.Context) VpcLinkOutput {
	return o
}

func (o VpcLinkOutput) ToVpcLinkPtrOutput() VpcLinkPtrOutput {
	return o.ToVpcLinkPtrOutputWithContext(context.Background())
}

func (o VpcLinkOutput) ToVpcLinkPtrOutputWithContext(ctx context.Context) VpcLinkPtrOutput {
	return o.ApplyT(func(v VpcLink) *VpcLink {
		return &v
	}).(VpcLinkPtrOutput)
}

type VpcLinkPtrOutput struct {
	*pulumi.OutputState
}

func (VpcLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcLink)(nil))
}

func (o VpcLinkPtrOutput) ToVpcLinkPtrOutput() VpcLinkPtrOutput {
	return o
}

func (o VpcLinkPtrOutput) ToVpcLinkPtrOutputWithContext(ctx context.Context) VpcLinkPtrOutput {
	return o
}

type VpcLinkArrayOutput struct{ *pulumi.OutputState }

func (VpcLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcLink)(nil))
}

func (o VpcLinkArrayOutput) ToVpcLinkArrayOutput() VpcLinkArrayOutput {
	return o
}

func (o VpcLinkArrayOutput) ToVpcLinkArrayOutputWithContext(ctx context.Context) VpcLinkArrayOutput {
	return o
}

func (o VpcLinkArrayOutput) Index(i pulumi.IntInput) VpcLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcLink {
		return vs[0].([]VpcLink)[vs[1].(int)]
	}).(VpcLinkOutput)
}

type VpcLinkMapOutput struct{ *pulumi.OutputState }

func (VpcLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcLink)(nil))
}

func (o VpcLinkMapOutput) ToVpcLinkMapOutput() VpcLinkMapOutput {
	return o
}

func (o VpcLinkMapOutput) ToVpcLinkMapOutputWithContext(ctx context.Context) VpcLinkMapOutput {
	return o
}

func (o VpcLinkMapOutput) MapIndex(k pulumi.StringInput) VpcLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcLink {
		return vs[0].(map[string]VpcLink)[vs[1].(string)]
	}).(VpcLinkOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcLinkOutput{})
	pulumi.RegisterOutputType(VpcLinkPtrOutput{})
	pulumi.RegisterOutputType(VpcLinkArrayOutput{})
	pulumi.RegisterOutputType(VpcLinkMapOutput{})
}
