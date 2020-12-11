// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates Security Hub custom action.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/securityhub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewActionTarget(ctx, "exampleActionTarget", &securityhub.ActionTargetArgs{
// 			Identifier:  pulumi.String("SendToChat"),
// 			Description: pulumi.String("This is custom action sends selected findings to chat"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAccount,
// 		}))
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
// Security Hub custom action can be imported using the action target ARN e.g.
//
// ```sh
//  $ pulumi import aws:securityhub/actionTarget:ActionTarget example arn:aws:securityhub:eu-west-1:312940875350:action/custom/a
// ```
type ActionTarget struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the Security Hub custom action target.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the custom action target.
	Description pulumi.StringOutput `pulumi:"description"`
	// The ID for the custom action target.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// The description for the custom action target.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewActionTarget registers a new resource with the given unique name, arguments, and options.
func NewActionTarget(ctx *pulumi.Context,
	name string, args *ActionTargetArgs, opts ...pulumi.ResourceOption) (*ActionTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	var resource ActionTarget
	err := ctx.RegisterResource("aws:securityhub/actionTarget:ActionTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionTarget gets an existing ActionTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionTargetState, opts ...pulumi.ResourceOption) (*ActionTarget, error) {
	var resource ActionTarget
	err := ctx.ReadResource("aws:securityhub/actionTarget:ActionTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionTarget resources.
type actionTargetState struct {
	// Amazon Resource Name (ARN) of the Security Hub custom action target.
	Arn *string `pulumi:"arn"`
	// The name of the custom action target.
	Description *string `pulumi:"description"`
	// The ID for the custom action target.
	Identifier *string `pulumi:"identifier"`
	// The description for the custom action target.
	Name *string `pulumi:"name"`
}

type ActionTargetState struct {
	// Amazon Resource Name (ARN) of the Security Hub custom action target.
	Arn pulumi.StringPtrInput
	// The name of the custom action target.
	Description pulumi.StringPtrInput
	// The ID for the custom action target.
	Identifier pulumi.StringPtrInput
	// The description for the custom action target.
	Name pulumi.StringPtrInput
}

func (ActionTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionTargetState)(nil)).Elem()
}

type actionTargetArgs struct {
	// The name of the custom action target.
	Description string `pulumi:"description"`
	// The ID for the custom action target.
	Identifier string `pulumi:"identifier"`
	// The description for the custom action target.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ActionTarget resource.
type ActionTargetArgs struct {
	// The name of the custom action target.
	Description pulumi.StringInput
	// The ID for the custom action target.
	Identifier pulumi.StringInput
	// The description for the custom action target.
	Name pulumi.StringPtrInput
}

func (ActionTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionTargetArgs)(nil)).Elem()
}

type ActionTargetInput interface {
	pulumi.Input

	ToActionTargetOutput() ActionTargetOutput
	ToActionTargetOutputWithContext(ctx context.Context) ActionTargetOutput
}

type ActionTargetPtrInput interface {
	pulumi.Input

	ToActionTargetPtrOutput() ActionTargetPtrOutput
	ToActionTargetPtrOutputWithContext(ctx context.Context) ActionTargetPtrOutput
}

func (ActionTarget) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionTarget)(nil)).Elem()
}

func (i ActionTarget) ToActionTargetOutput() ActionTargetOutput {
	return i.ToActionTargetOutputWithContext(context.Background())
}

func (i ActionTarget) ToActionTargetOutputWithContext(ctx context.Context) ActionTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionTargetOutput)
}

func (i ActionTarget) ToActionTargetPtrOutput() ActionTargetPtrOutput {
	return i.ToActionTargetPtrOutputWithContext(context.Background())
}

func (i ActionTarget) ToActionTargetPtrOutputWithContext(ctx context.Context) ActionTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionTargetPtrOutput)
}

type ActionTargetOutput struct {
	*pulumi.OutputState
}

func (ActionTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionTargetOutput)(nil)).Elem()
}

func (o ActionTargetOutput) ToActionTargetOutput() ActionTargetOutput {
	return o
}

func (o ActionTargetOutput) ToActionTargetOutputWithContext(ctx context.Context) ActionTargetOutput {
	return o
}

type ActionTargetPtrOutput struct {
	*pulumi.OutputState
}

func (ActionTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionTarget)(nil)).Elem()
}

func (o ActionTargetPtrOutput) ToActionTargetPtrOutput() ActionTargetPtrOutput {
	return o
}

func (o ActionTargetPtrOutput) ToActionTargetPtrOutputWithContext(ctx context.Context) ActionTargetPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ActionTargetOutput{})
	pulumi.RegisterOutputType(ActionTargetPtrOutput{})
}
