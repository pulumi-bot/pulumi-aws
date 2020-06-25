// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Service Discovery Public DNS Namespace resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/servicediscovery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = servicediscovery.NewPublicDnsNamespace(ctx, "example", &servicediscovery.PublicDnsNamespaceArgs{
// 			Description: pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PublicDnsNamespace struct {
	pulumi.CustomResourceState

	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone pulumi.StringOutput `pulumi:"hostedZone"`
	// The name of the namespace.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the namespace.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPublicDnsNamespace registers a new resource with the given unique name, arguments, and options.
func NewPublicDnsNamespace(ctx *pulumi.Context,
	name string, args *PublicDnsNamespaceArgs, opts ...pulumi.ResourceOption) (*PublicDnsNamespace, error) {
	if args == nil {
		args = &PublicDnsNamespaceArgs{}
	}
	var resource PublicDnsNamespace
	err := ctx.RegisterResource("aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicDnsNamespace gets an existing PublicDnsNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicDnsNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicDnsNamespaceState, opts ...pulumi.ResourceOption) (*PublicDnsNamespace, error) {
	var resource PublicDnsNamespace
	err := ctx.ReadResource("aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicDnsNamespace resources.
type publicDnsNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn *string `pulumi:"arn"`
	// The description that you specify for the namespace when you create it.
	Description *string `pulumi:"description"`
	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone *string `pulumi:"hostedZone"`
	// The name of the namespace.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the namespace.
	Tags map[string]string `pulumi:"tags"`
}

type PublicDnsNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringPtrInput
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrInput
	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone pulumi.StringPtrInput
	// The name of the namespace.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the namespace.
	Tags pulumi.StringMapInput
}

func (PublicDnsNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDnsNamespaceState)(nil)).Elem()
}

type publicDnsNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description *string `pulumi:"description"`
	// The name of the namespace.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the namespace.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PublicDnsNamespace resource.
type PublicDnsNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringPtrInput
	// The name of the namespace.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the namespace.
	Tags pulumi.StringMapInput
}

func (PublicDnsNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDnsNamespaceArgs)(nil)).Elem()
}
