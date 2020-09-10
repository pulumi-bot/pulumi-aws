// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Activity struct {
	pulumi.CustomResourceState

	CreationDate pulumi.StringOutput    `pulumi:"creationDate"`
	Name         pulumi.StringOutput    `pulumi:"name"`
	Tags         pulumi.StringMapOutput `pulumi:"tags"`
}

// NewActivity registers a new resource with the given unique name, arguments, and options.
func NewActivity(ctx *pulumi.Context,
	name string, args *ActivityArgs, opts ...pulumi.ResourceOption) (*Activity, error) {
	if args == nil {
		args = &ActivityArgs{}
	}
	var resource Activity
	err := ctx.RegisterResource("aws:sfn/activity:Activity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActivity gets an existing Activity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivityState, opts ...pulumi.ResourceOption) (*Activity, error) {
	var resource Activity
	err := ctx.ReadResource("aws:sfn/activity:Activity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Activity resources.
type activityState struct {
	CreationDate *string           `pulumi:"creationDate"`
	Name         *string           `pulumi:"name"`
	Tags         map[string]string `pulumi:"tags"`
}

type ActivityState struct {
	CreationDate pulumi.StringPtrInput
	Name         pulumi.StringPtrInput
	Tags         pulumi.StringMapInput
}

func (ActivityState) ElementType() reflect.Type {
	return reflect.TypeOf((*activityState)(nil)).Elem()
}

type activityArgs struct {
	Name *string           `pulumi:"name"`
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Activity resource.
type ActivityArgs struct {
	Name pulumi.StringPtrInput
	Tags pulumi.StringMapInput
}

func (ActivityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activityArgs)(nil)).Elem()
}
