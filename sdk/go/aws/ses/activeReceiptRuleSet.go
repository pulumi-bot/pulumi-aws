// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"fmt"
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
