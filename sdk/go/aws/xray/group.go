// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package xray

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates and manages an AWS XRay Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/xray"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := xray.NewGroup(ctx, "example", &xray.GroupArgs{
// 			FilterExpression: pulumi.String("responsetime > 5"),
// 			GroupName:        pulumi.String("example"),
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
// XRay Groups can be imported using the ARN, e.g.
//
// ```sh
//  $ pulumi import aws:xray/group:Group example arn:aws:xray:us-west-2:1234567890:group/example-group/TNGX7SW5U6QY36T4ZMOUA3HVLBYCZTWDIOOXY3CJAXTHSS3YCWUA
// ```
type Group struct {
	pulumi.CustomResourceState

	// The ARN of the Group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The filter expression defining criteria by which to group traces. more info can be found in official [docs](https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html).
	FilterExpression pulumi.StringOutput `pulumi:"filterExpression"`
	// The name of the group.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Key-value mapping of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FilterExpression == nil {
		return nil, errors.New("invalid value for required argument 'FilterExpression'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	var resource Group
	err := ctx.RegisterResource("aws:xray/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("aws:xray/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The ARN of the Group.
	Arn *string `pulumi:"arn"`
	// The filter expression defining criteria by which to group traces. more info can be found in official [docs](https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html).
	FilterExpression *string `pulumi:"filterExpression"`
	// The name of the group.
	GroupName *string `pulumi:"groupName"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type GroupState struct {
	// The ARN of the Group.
	Arn pulumi.StringPtrInput
	// The filter expression defining criteria by which to group traces. more info can be found in official [docs](https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html).
	FilterExpression pulumi.StringPtrInput
	// The name of the group.
	GroupName pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The filter expression defining criteria by which to group traces. more info can be found in official [docs](https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html).
	FilterExpression string `pulumi:"filterExpression"`
	// The name of the group.
	GroupName string `pulumi:"groupName"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The filter expression defining criteria by which to group traces. more info can be found in official [docs](https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html).
	FilterExpression pulumi.StringInput
	// The name of the group.
	GroupName pulumi.StringInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil)).Elem()
}

func (i Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupOutput)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
