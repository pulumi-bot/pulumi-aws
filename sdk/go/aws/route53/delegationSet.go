// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a [Route53 Delegation Set](https://docs.aws.amazon.com/Route53/latest/APIReference/API-actions-by-function.html#actions-by-function-reusable-delegation-sets) resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := route53.NewDelegationSet(ctx, "main", &route53.DelegationSetArgs{
// 			ReferenceName: pulumi.String("DynDNS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewZone(ctx, "primary", &route53.ZoneArgs{
// 			DelegationSetId: main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewZone(ctx, "secondary", &route53.ZoneArgs{
// 			DelegationSetId: main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DelegationSet struct {
	pulumi.CustomResourceState

	// A list of authoritative name servers for the hosted zone
	// (effectively a list of NS records).
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName pulumi.StringPtrOutput `pulumi:"referenceName"`
}

// NewDelegationSet registers a new resource with the given unique name, arguments, and options.
func NewDelegationSet(ctx *pulumi.Context,
	name string, args *DelegationSetArgs, opts ...pulumi.ResourceOption) (*DelegationSet, error) {
	if args == nil {
		args = &DelegationSetArgs{}
	}
	var resource DelegationSet
	err := ctx.RegisterResource("aws:route53/delegationSet:DelegationSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDelegationSet gets an existing DelegationSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDelegationSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DelegationSetState, opts ...pulumi.ResourceOption) (*DelegationSet, error) {
	var resource DelegationSet
	err := ctx.ReadResource("aws:route53/delegationSet:DelegationSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DelegationSet resources.
type delegationSetState struct {
	// A list of authoritative name servers for the hosted zone
	// (effectively a list of NS records).
	NameServers []string `pulumi:"nameServers"`
	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName *string `pulumi:"referenceName"`
}

type DelegationSetState struct {
	// A list of authoritative name servers for the hosted zone
	// (effectively a list of NS records).
	NameServers pulumi.StringArrayInput
	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName pulumi.StringPtrInput
}

func (DelegationSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*delegationSetState)(nil)).Elem()
}

type delegationSetArgs struct {
	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName *string `pulumi:"referenceName"`
}

// The set of arguments for constructing a DelegationSet resource.
type DelegationSetArgs struct {
	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName pulumi.StringPtrInput
}

func (DelegationSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*delegationSetArgs)(nil)).Elem()
}

type DelegationSetInput interface {
	pulumi.Input

	ToDelegationSetOutput() DelegationSetOutput
	ToDelegationSetOutputWithContext(ctx context.Context) DelegationSetOutput
}

func (DelegationSet) ElementType() reflect.Type {
	return reflect.TypeOf((*DelegationSet)(nil)).Elem()
}

func (i DelegationSet) ToDelegationSetOutput() DelegationSetOutput {
	return i.ToDelegationSetOutputWithContext(context.Background())
}

func (i DelegationSet) ToDelegationSetOutputWithContext(ctx context.Context) DelegationSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelegationSetOutput)
}

type DelegationSetOutput struct {
	*pulumi.OutputState
}

func (DelegationSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DelegationSetOutput)(nil)).Elem()
}

func (o DelegationSetOutput) ToDelegationSetOutput() DelegationSetOutput {
	return o
}

func (o DelegationSetOutput) ToDelegationSetOutputWithContext(ctx context.Context) DelegationSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DelegationSetOutput{})
}
