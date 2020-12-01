// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAFv2 IP Set Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafv2.NewIpSet(ctx, "example", &wafv2.IpSetArgs{
// 			Addresses: pulumi.StringArray{
// 				pulumi.String("1.2.3.4/32"),
// 				pulumi.String("5.6.7.8/32"),
// 			},
// 			Description:      pulumi.String("Example IP set"),
// 			IpAddressVersion: pulumi.String("IPV4"),
// 			Scope:            pulumi.String("REGIONAL"),
// 			Tags: pulumi.StringMap{
// 				"Tag1": pulumi.String("Value1"),
// 				"Tag2": pulumi.String("Value2"),
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
// WAFv2 IP Sets can be imported using `ID/name/scope`
//
// ```sh
//  $ pulumi import aws:wafv2/ipSet:IpSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
// ```
type IpSet struct {
	pulumi.CustomResourceState

	// Contains an array of strings that specify one or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports all address ranges for IP versions IPv4 and IPv6.
	Addresses pulumi.StringArrayOutput `pulumi:"addresses"`
	// The Amazon Resource Name (ARN) that identifies the cluster.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A friendly description of the IP set.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion pulumi.StringOutput `pulumi:"ipAddressVersion"`
	LockToken        pulumi.StringOutput `pulumi:"lockToken"`
	// A friendly name of the IP set.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope pulumi.StringOutput `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewIpSet registers a new resource with the given unique name, arguments, and options.
func NewIpSet(ctx *pulumi.Context,
	name string, args *IpSetArgs, opts ...pulumi.ResourceOption) (*IpSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAddressVersion == nil {
		return nil, errors.New("invalid value for required argument 'IpAddressVersion'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource IpSet
	err := ctx.RegisterResource("aws:wafv2/ipSet:IpSet", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:wafv2/ipSet:IpSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpSet resources.
type ipSetState struct {
	// Contains an array of strings that specify one or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports all address ranges for IP versions IPv4 and IPv6.
	Addresses []string `pulumi:"addresses"`
	// The Amazon Resource Name (ARN) that identifies the cluster.
	Arn *string `pulumi:"arn"`
	// A friendly description of the IP set.
	Description *string `pulumi:"description"`
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion *string `pulumi:"ipAddressVersion"`
	LockToken        *string `pulumi:"lockToken"`
	// A friendly name of the IP set.
	Name *string `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope *string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
}

type IpSetState struct {
	// Contains an array of strings that specify one or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports all address ranges for IP versions IPv4 and IPv6.
	Addresses pulumi.StringArrayInput
	// The Amazon Resource Name (ARN) that identifies the cluster.
	Arn pulumi.StringPtrInput
	// A friendly description of the IP set.
	Description pulumi.StringPtrInput
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion pulumi.StringPtrInput
	LockToken        pulumi.StringPtrInput
	// A friendly name of the IP set.
	Name pulumi.StringPtrInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope pulumi.StringPtrInput
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.StringMapInput
}

func (IpSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetState)(nil)).Elem()
}

type ipSetArgs struct {
	// Contains an array of strings that specify one or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports all address ranges for IP versions IPv4 and IPv6.
	Addresses []string `pulumi:"addresses"`
	// A friendly description of the IP set.
	Description *string `pulumi:"description"`
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion string `pulumi:"ipAddressVersion"`
	// A friendly name of the IP set.
	Name *string `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IpSet resource.
type IpSetArgs struct {
	// Contains an array of strings that specify one or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports all address ranges for IP versions IPv4 and IPv6.
	Addresses pulumi.StringArrayInput
	// A friendly description of the IP set.
	Description pulumi.StringPtrInput
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion pulumi.StringInput
	// A friendly name of the IP set.
	Name pulumi.StringPtrInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope pulumi.StringInput
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.StringMapInput
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
