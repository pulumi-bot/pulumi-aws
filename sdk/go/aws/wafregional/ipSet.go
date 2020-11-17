// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Regional IPSet Resource for use with Application Load Balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafregional.NewIpSet(ctx, "ipset", &wafregional.IpSetArgs{
// 			IpSetDescriptors: wafregional.IpSetIpSetDescriptorArray{
// 				&wafregional.IpSetIpSetDescriptorArgs{
// 					Type:  pulumi.String("IPV4"),
// 					Value: pulumi.String("192.0.7.0/24"),
// 				},
// 				&wafregional.IpSetIpSetDescriptorArgs{
// 					Type:  pulumi.String("IPV4"),
// 					Value: pulumi.String("10.16.16.0/16"),
// 				},
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
// WAF Regional IPSets can be imported using their ID, e.g.
//
// ```sh
//  $ pulumi import aws:wafregional/ipSet:IpSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type IpSet struct {
	pulumi.CustomResourceState

	// The ARN of the WAF IPSet.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	IpSetDescriptors IpSetIpSetDescriptorArrayOutput `pulumi:"ipSetDescriptors"`
	// The name or description of the IPSet.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewIpSet registers a new resource with the given unique name, arguments, and options.
func NewIpSet(ctx *pulumi.Context,
	name string, args *IpSetArgs, opts ...pulumi.ResourceOption) (*IpSet, error) {
	if args == nil {
		args = &IpSetArgs{}
	}
	var resource IpSet
	err := ctx.RegisterResource("aws:wafregional/ipSet:IpSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpSet gets an existing IpSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpSetState, opts ...pulumi.ResourceOption) (*IpSet, error) {
	var resource IpSet
	err := ctx.ReadResource("aws:wafregional/ipSet:IpSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpSet resources.
type ipSetState struct {
	// The ARN of the WAF IPSet.
	Arn *string `pulumi:"arn"`
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	IpSetDescriptors []IpSetIpSetDescriptor `pulumi:"ipSetDescriptors"`
	// The name or description of the IPSet.
	Name *string `pulumi:"name"`
}

type IpSetState struct {
	// The ARN of the WAF IPSet.
	Arn pulumi.StringPtrInput
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	IpSetDescriptors IpSetIpSetDescriptorArrayInput
	// The name or description of the IPSet.
	Name pulumi.StringPtrInput
}

func (IpSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetState)(nil)).Elem()
}

type ipSetArgs struct {
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	IpSetDescriptors []IpSetIpSetDescriptor `pulumi:"ipSetDescriptors"`
	// The name or description of the IPSet.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a IpSet resource.
type IpSetArgs struct {
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	IpSetDescriptors IpSetIpSetDescriptorArrayInput
	// The name or description of the IPSet.
	Name pulumi.StringPtrInput
}

func (IpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetArgs)(nil)).Elem()
}

type IpSetInput interface {
	pulumi.Input

	ToIpSetOutput() IpSetOutput
	ToIpSetOutputWithContext(ctx context.Context) IpSetOutput
}

func (IpSet) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSet)(nil)).Elem()
}

func (i IpSet) ToIpSetOutput() IpSetOutput {
	return i.ToIpSetOutputWithContext(context.Background())
}

func (i IpSet) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetOutput)
}

type IpSetOutput struct {
	*pulumi.OutputState
}

func (IpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSetOutput)(nil)).Elem()
}

func (o IpSetOutput) ToIpSetOutput() IpSetOutput {
	return o
}

func (o IpSetOutput) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IpSetOutput{})
}
