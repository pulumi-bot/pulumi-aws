// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to designate the active SES receipt rule set
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewActiveReceiptRuleSet(ctx, "main", &ses.ActiveReceiptRuleSetArgs{
// 			RuleSetName: pulumi.String("primary-rules"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ActiveReceiptRuleSet struct {
	pulumi.CustomResourceState

	// The name of the rule set
	RuleSetName pulumi.StringOutput `pulumi:"ruleSetName"`
}

// NewActiveReceiptRuleSet registers a new resource with the given unique name, arguments, and options.
func NewActiveReceiptRuleSet(ctx *pulumi.Context,
	name string, args *ActiveReceiptRuleSetArgs, opts ...pulumi.ResourceOption) (*ActiveReceiptRuleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RuleSetName == nil {
		return nil, errors.New("invalid value for required argument 'RuleSetName'")
	}
	var resource ActiveReceiptRuleSet
	err := ctx.RegisterResource("aws:ses/activeReceiptRuleSet:ActiveReceiptRuleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActiveReceiptRuleSet gets an existing ActiveReceiptRuleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveReceiptRuleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActiveReceiptRuleSetState, opts ...pulumi.ResourceOption) (*ActiveReceiptRuleSet, error) {
	var resource ActiveReceiptRuleSet
	err := ctx.ReadResource("aws:ses/activeReceiptRuleSet:ActiveReceiptRuleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActiveReceiptRuleSet resources.
type activeReceiptRuleSetState struct {
	// The name of the rule set
	RuleSetName *string `pulumi:"ruleSetName"`
}

type ActiveReceiptRuleSetState struct {
	// The name of the rule set
	RuleSetName pulumi.StringPtrInput
}

func (ActiveReceiptRuleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*activeReceiptRuleSetState)(nil)).Elem()
}

type activeReceiptRuleSetArgs struct {
	// The name of the rule set
	RuleSetName string `pulumi:"ruleSetName"`
}

// The set of arguments for constructing a ActiveReceiptRuleSet resource.
type ActiveReceiptRuleSetArgs struct {
	// The name of the rule set
	RuleSetName pulumi.StringInput
}

func (ActiveReceiptRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activeReceiptRuleSetArgs)(nil)).Elem()
}

type ActiveReceiptRuleSetInput interface {
	pulumi.Input

	ToActiveReceiptRuleSetOutput() ActiveReceiptRuleSetOutput
	ToActiveReceiptRuleSetOutputWithContext(ctx context.Context) ActiveReceiptRuleSetOutput
}

func (*ActiveReceiptRuleSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveReceiptRuleSet)(nil))
}

func (i *ActiveReceiptRuleSet) ToActiveReceiptRuleSetOutput() ActiveReceiptRuleSetOutput {
	return i.ToActiveReceiptRuleSetOutputWithContext(context.Background())
}

func (i *ActiveReceiptRuleSet) ToActiveReceiptRuleSetOutputWithContext(ctx context.Context) ActiveReceiptRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveReceiptRuleSetOutput)
}

func (i *ActiveReceiptRuleSet) ToActiveReceiptRuleSetPtrOutput() ActiveReceiptRuleSetPtrOutput {
	return i.ToActiveReceiptRuleSetPtrOutputWithContext(context.Background())
}

func (i *ActiveReceiptRuleSet) ToActiveReceiptRuleSetPtrOutputWithContext(ctx context.Context) ActiveReceiptRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveReceiptRuleSetPtrOutput)
}

type ActiveReceiptRuleSetPtrInput interface {
	pulumi.Input

	ToActiveReceiptRuleSetPtrOutput() ActiveReceiptRuleSetPtrOutput
	ToActiveReceiptRuleSetPtrOutputWithContext(ctx context.Context) ActiveReceiptRuleSetPtrOutput
}

type activeReceiptRuleSetPtrType ActiveReceiptRuleSetArgs

func (*activeReceiptRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveReceiptRuleSet)(nil))
}

func (i *activeReceiptRuleSetPtrType) ToActiveReceiptRuleSetPtrOutput() ActiveReceiptRuleSetPtrOutput {
	return i.ToActiveReceiptRuleSetPtrOutputWithContext(context.Background())
}

func (i *activeReceiptRuleSetPtrType) ToActiveReceiptRuleSetPtrOutputWithContext(ctx context.Context) ActiveReceiptRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveReceiptRuleSetPtrOutput)
}

// ActiveReceiptRuleSetArrayInput is an input type that accepts ActiveReceiptRuleSetArray and ActiveReceiptRuleSetArrayOutput values.
// You can construct a concrete instance of `ActiveReceiptRuleSetArrayInput` via:
//
//          ActiveReceiptRuleSetArray{ ActiveReceiptRuleSetArgs{...} }
type ActiveReceiptRuleSetArrayInput interface {
	pulumi.Input

	ToActiveReceiptRuleSetArrayOutput() ActiveReceiptRuleSetArrayOutput
	ToActiveReceiptRuleSetArrayOutputWithContext(context.Context) ActiveReceiptRuleSetArrayOutput
}

type ActiveReceiptRuleSetArray []ActiveReceiptRuleSetInput

func (ActiveReceiptRuleSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ActiveReceiptRuleSet)(nil))
}

func (i ActiveReceiptRuleSetArray) ToActiveReceiptRuleSetArrayOutput() ActiveReceiptRuleSetArrayOutput {
	return i.ToActiveReceiptRuleSetArrayOutputWithContext(context.Background())
}

func (i ActiveReceiptRuleSetArray) ToActiveReceiptRuleSetArrayOutputWithContext(ctx context.Context) ActiveReceiptRuleSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveReceiptRuleSetArrayOutput)
}

// ActiveReceiptRuleSetMapInput is an input type that accepts ActiveReceiptRuleSetMap and ActiveReceiptRuleSetMapOutput values.
// You can construct a concrete instance of `ActiveReceiptRuleSetMapInput` via:
//
//          ActiveReceiptRuleSetMap{ "key": ActiveReceiptRuleSetArgs{...} }
type ActiveReceiptRuleSetMapInput interface {
	pulumi.Input

	ToActiveReceiptRuleSetMapOutput() ActiveReceiptRuleSetMapOutput
	ToActiveReceiptRuleSetMapOutputWithContext(context.Context) ActiveReceiptRuleSetMapOutput
}

type ActiveReceiptRuleSetMap map[string]ActiveReceiptRuleSetInput

func (ActiveReceiptRuleSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ActiveReceiptRuleSet)(nil))
}

func (i ActiveReceiptRuleSetMap) ToActiveReceiptRuleSetMapOutput() ActiveReceiptRuleSetMapOutput {
	return i.ToActiveReceiptRuleSetMapOutputWithContext(context.Background())
}

func (i ActiveReceiptRuleSetMap) ToActiveReceiptRuleSetMapOutputWithContext(ctx context.Context) ActiveReceiptRuleSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveReceiptRuleSetMapOutput)
}

type ActiveReceiptRuleSetOutput struct {
	*pulumi.OutputState
}

func (ActiveReceiptRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveReceiptRuleSet)(nil))
}

func (o ActiveReceiptRuleSetOutput) ToActiveReceiptRuleSetOutput() ActiveReceiptRuleSetOutput {
	return o
}

func (o ActiveReceiptRuleSetOutput) ToActiveReceiptRuleSetOutputWithContext(ctx context.Context) ActiveReceiptRuleSetOutput {
	return o
}

func (o ActiveReceiptRuleSetOutput) ToActiveReceiptRuleSetPtrOutput() ActiveReceiptRuleSetPtrOutput {
	return o.ToActiveReceiptRuleSetPtrOutputWithContext(context.Background())
}

func (o ActiveReceiptRuleSetOutput) ToActiveReceiptRuleSetPtrOutputWithContext(ctx context.Context) ActiveReceiptRuleSetPtrOutput {
	return o.ApplyT(func(v ActiveReceiptRuleSet) *ActiveReceiptRuleSet {
		return &v
	}).(ActiveReceiptRuleSetPtrOutput)
}

type ActiveReceiptRuleSetPtrOutput struct {
	*pulumi.OutputState
}

func (ActiveReceiptRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveReceiptRuleSet)(nil))
}

func (o ActiveReceiptRuleSetPtrOutput) ToActiveReceiptRuleSetPtrOutput() ActiveReceiptRuleSetPtrOutput {
	return o
}

func (o ActiveReceiptRuleSetPtrOutput) ToActiveReceiptRuleSetPtrOutputWithContext(ctx context.Context) ActiveReceiptRuleSetPtrOutput {
	return o
}

type ActiveReceiptRuleSetArrayOutput struct{ *pulumi.OutputState }

func (ActiveReceiptRuleSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActiveReceiptRuleSet)(nil))
}

func (o ActiveReceiptRuleSetArrayOutput) ToActiveReceiptRuleSetArrayOutput() ActiveReceiptRuleSetArrayOutput {
	return o
}

func (o ActiveReceiptRuleSetArrayOutput) ToActiveReceiptRuleSetArrayOutputWithContext(ctx context.Context) ActiveReceiptRuleSetArrayOutput {
	return o
}

func (o ActiveReceiptRuleSetArrayOutput) Index(i pulumi.IntInput) ActiveReceiptRuleSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActiveReceiptRuleSet {
		return vs[0].([]ActiveReceiptRuleSet)[vs[1].(int)]
	}).(ActiveReceiptRuleSetOutput)
}

type ActiveReceiptRuleSetMapOutput struct{ *pulumi.OutputState }

func (ActiveReceiptRuleSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ActiveReceiptRuleSet)(nil))
}

func (o ActiveReceiptRuleSetMapOutput) ToActiveReceiptRuleSetMapOutput() ActiveReceiptRuleSetMapOutput {
	return o
}

func (o ActiveReceiptRuleSetMapOutput) ToActiveReceiptRuleSetMapOutputWithContext(ctx context.Context) ActiveReceiptRuleSetMapOutput {
	return o
}

func (o ActiveReceiptRuleSetMapOutput) MapIndex(k pulumi.StringInput) ActiveReceiptRuleSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ActiveReceiptRuleSet {
		return vs[0].(map[string]ActiveReceiptRuleSet)[vs[1].(string)]
	}).(ActiveReceiptRuleSetOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveReceiptRuleSetOutput{})
	pulumi.RegisterOutputType(ActiveReceiptRuleSetPtrOutput{})
	pulumi.RegisterOutputType(ActiveReceiptRuleSetArrayOutput{})
	pulumi.RegisterOutputType(ActiveReceiptRuleSetMapOutput{})
}
