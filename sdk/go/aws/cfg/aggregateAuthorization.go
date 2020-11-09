// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Config Aggregate Authorization
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cfg"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cfg.NewAggregateAuthorization(ctx, "example", &cfg.AggregateAuthorizationArgs{
// 			AccountId: pulumi.String("123456789012"),
// 			Region:    pulumi.String("eu-west-2"),
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
// Config aggregate authorizations can be imported using `account_id:region`, e.g.
//
// ```sh
//  $ pulumi import aws:cfg/aggregateAuthorization:AggregateAuthorization example 123456789012:us-east-1
// ```
type AggregateAuthorization struct {
	pulumi.CustomResourceState

	// Account ID
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The ARN of the authorization
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Region
	Region pulumi.StringOutput `pulumi:"region"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAggregateAuthorization registers a new resource with the given unique name, arguments, and options.
func NewAggregateAuthorization(ctx *pulumi.Context,
	name string, args *AggregateAuthorizationArgs, opts ...pulumi.ResourceOption) (*AggregateAuthorization, error) {
	if args == nil || args.AccountId == nil {
		return nil, errors.New("missing required argument 'AccountId'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil {
		args = &AggregateAuthorizationArgs{}
	}
	var resource AggregateAuthorization
	err := ctx.RegisterResource("aws:cfg/aggregateAuthorization:AggregateAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAggregateAuthorization gets an existing AggregateAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAggregateAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AggregateAuthorizationState, opts ...pulumi.ResourceOption) (*AggregateAuthorization, error) {
	var resource AggregateAuthorization
	err := ctx.ReadResource("aws:cfg/aggregateAuthorization:AggregateAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AggregateAuthorization resources.
type aggregateAuthorizationState struct {
	// Account ID
	AccountId *string `pulumi:"accountId"`
	// The ARN of the authorization
	Arn *string `pulumi:"arn"`
	// Region
	Region *string `pulumi:"region"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AggregateAuthorizationState struct {
	// Account ID
	AccountId pulumi.StringPtrInput
	// The ARN of the authorization
	Arn pulumi.StringPtrInput
	// Region
	Region pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AggregateAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateAuthorizationState)(nil)).Elem()
}

type aggregateAuthorizationArgs struct {
	// Account ID
	AccountId string `pulumi:"accountId"`
	// Region
	Region string `pulumi:"region"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AggregateAuthorization resource.
type AggregateAuthorizationArgs struct {
	// Account ID
	AccountId pulumi.StringInput
	// Region
	Region pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AggregateAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateAuthorizationArgs)(nil)).Elem()
}
