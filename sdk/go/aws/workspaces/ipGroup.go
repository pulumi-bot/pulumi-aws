// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package workspaces

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IP access control group in AWS WorkSpaces Service
type IpGroup struct {
	pulumi.CustomResourceState

	// The description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the IP group.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	Rules IpGroupRuleArrayOutput `pulumi:"rules"`
	Tags  pulumi.MapOutput       `pulumi:"tags"`
}

// NewIpGroup registers a new resource with the given unique name, arguments, and options.
func NewIpGroup(ctx *pulumi.Context,
	name string, args *IpGroupArgs, opts ...pulumi.ResourceOption) (*IpGroup, error) {
	if args == nil {
		args = &IpGroupArgs{}
	}
	var resource IpGroup
	err := ctx.RegisterResource("aws:workspaces/ipGroup:IpGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpGroup gets an existing IpGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpGroupState, opts ...pulumi.ResourceOption) (*IpGroup, error) {
	var resource IpGroup
	err := ctx.ReadResource("aws:workspaces/ipGroup:IpGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpGroup resources.
type ipGroupState struct {
	// The description.
	Description *string `pulumi:"description"`
	// The name of the IP group.
	Name *string `pulumi:"name"`
	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	Rules []IpGroupRule          `pulumi:"rules"`
	Tags  map[string]interface{} `pulumi:"tags"`
}

type IpGroupState struct {
	// The description.
	Description pulumi.StringPtrInput
	// The name of the IP group.
	Name pulumi.StringPtrInput
	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	Rules IpGroupRuleArrayInput
	Tags  pulumi.MapInput
}

func (IpGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipGroupState)(nil)).Elem()
}

type ipGroupArgs struct {
	// The description.
	Description *string `pulumi:"description"`
	// The name of the IP group.
	Name *string `pulumi:"name"`
	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	Rules []IpGroupRuleArgs      `pulumi:"rules"`
	Tags  map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a IpGroup resource.
type IpGroupArgs struct {
	// The description.
	Description pulumi.StringPtrInput
	// The name of the IP group.
	Name pulumi.StringPtrInput
	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	Rules IpGroupRuleArgsArrayInput
	Tags  pulumi.MapInput
}

func (IpGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipGroupArgs)(nil)).Elem()
}
