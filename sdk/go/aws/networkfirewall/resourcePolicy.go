// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Network Firewall Resource Policy Resource for a rule group or firewall policy.
//
// ## Example Usage
// ### For a Firewall Policy resource
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action":   "network-firewall:ListFirewallPolicies",
// 					"Effect":   "Allow",
// 					"Resource": aws_networkfirewall_firewall_policy.Example.Arn,
// 					"Principal": map[string]interface{}{
// 						"AWS": "arn:aws:iam::123456789012:root",
// 					},
// 				},
// 			},
// 			"Version": "2012-10-17",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err := networkfirewall.NewResourcePolicy(ctx, "example", &networkfirewall.ResourcePolicyArgs{
// 			ResourceArn: pulumi.Any(aws_networkfirewall_firewall_policy.Example.Arn),
// 			Policy:      pulumi.String(json0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### For a Rule Group resource
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/networkfirewall"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action":   "network-firewall:ListRuleGroups",
// 					"Effect":   "Allow",
// 					"Resource": aws_networkfirewall_rule_group.Example.Arn,
// 					"Principal": map[string]interface{}{
// 						"AWS": "arn:aws:iam::123456789012:root",
// 					},
// 				},
// 			},
// 			"Version": "2012-10-17",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err := networkfirewall.NewResourcePolicy(ctx, "example", &networkfirewall.ResourcePolicyArgs{
// 			ResourceArn: pulumi.Any(aws_networkfirewall_rule_group.Example.Arn),
// 			Policy:      pulumi.String(json0),
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
// Network Firewall Resource Policies can be imported using the `resource_arn` e.g.
//
// ```sh
//  $ pulumi import aws:networkfirewall/resourcePolicy:ResourcePolicy example aws_networkfirewall_rule_group.example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
// ```
type ResourcePolicy struct {
	pulumi.CustomResourceState

	Policy pulumi.StringOutput `pulumi:"policy"`
	// The Amazon Resource Name (ARN) of the rule group or firewall policy.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewResourcePolicy registers a new resource with the given unique name, arguments, and options.
func NewResourcePolicy(ctx *pulumi.Context,
	name string, args *ResourcePolicyArgs, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	var resource ResourcePolicy
	err := ctx.RegisterResource("aws:networkfirewall/resourcePolicy:ResourcePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePolicy gets an existing ResourcePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePolicyState, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	var resource ResourcePolicy
	err := ctx.ReadResource("aws:networkfirewall/resourcePolicy:ResourcePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePolicy resources.
type resourcePolicyState struct {
	Policy *string `pulumi:"policy"`
	// The Amazon Resource Name (ARN) of the rule group or firewall policy.
	ResourceArn *string `pulumi:"resourceArn"`
}

type ResourcePolicyState struct {
	Policy pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the rule group or firewall policy.
	ResourceArn pulumi.StringPtrInput
}

func (ResourcePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyState)(nil)).Elem()
}

type resourcePolicyArgs struct {
	Policy string `pulumi:"policy"`
	// The Amazon Resource Name (ARN) of the rule group or firewall policy.
	ResourceArn string `pulumi:"resourceArn"`
}

// The set of arguments for constructing a ResourcePolicy resource.
type ResourcePolicyArgs struct {
	Policy pulumi.StringInput
	// The Amazon Resource Name (ARN) of the rule group or firewall policy.
	ResourceArn pulumi.StringInput
}

func (ResourcePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyArgs)(nil)).Elem()
}

type ResourcePolicyInput interface {
	pulumi.Input

	ToResourcePolicyOutput() ResourcePolicyOutput
	ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput
}

func (*ResourcePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePolicy)(nil))
}

func (i *ResourcePolicy) ToResourcePolicyOutput() ResourcePolicyOutput {
	return i.ToResourcePolicyOutputWithContext(context.Background())
}

func (i *ResourcePolicy) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyOutput)
}

type ResourcePolicyOutput struct {
	*pulumi.OutputState
}

func (ResourcePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePolicy)(nil))
}

func (o ResourcePolicyOutput) ToResourcePolicyOutput() ResourcePolicyOutput {
	return o
}

func (o ResourcePolicyOutput) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourcePolicyOutput{})
}
