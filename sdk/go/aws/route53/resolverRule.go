// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Route53 Resolver rule.
//
// ## Example Usage
// ### System rule
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
// 		_, err := route53.NewResolverRule(ctx, "sys", &route53.ResolverRuleArgs{
// 			DomainName: pulumi.String("subdomain.example.com"),
// 			RuleType:   pulumi.String("SYSTEM"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Forward rule
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
// 		_, err := route53.NewResolverRule(ctx, "fwd", &route53.ResolverRuleArgs{
// 			DomainName:         pulumi.String("example.com"),
// 			RuleType:           pulumi.String("FORWARD"),
// 			ResolverEndpointId: pulumi.Any(aws_route53_resolver_endpoint.Foo.Id),
// 			TargetIps: route53.ResolverRuleTargetIpArray{
// 				&route53.ResolverRuleTargetIpArgs{
// 					Ip: pulumi.String("123.45.67.89"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Prod"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ResolverRule struct {
	pulumi.CustomResourceState

	// The ARN (Amazon Resource Name) for the resolver rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// DNS queries for this domain name are forwarded to the IP addresses that are specified using `targetIp`.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	Name pulumi.StringOutput `pulumi:"name"`
	// When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `targetIp`.
	// This argument should only be specified for `FORWARD` type rules.
	ResolverEndpointId pulumi.StringPtrOutput `pulumi:"resolverEndpointId"`
	// The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
	RuleType pulumi.StringOutput `pulumi:"ruleType"`
	// Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
	// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
	ShareStatus pulumi.StringOutput `pulumi:"shareStatus"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
	// This argument should only be specified for `FORWARD` type rules.
	TargetIps ResolverRuleTargetIpArrayOutput `pulumi:"targetIps"`
}

// NewResolverRule registers a new resource with the given unique name, arguments, and options.
func NewResolverRule(ctx *pulumi.Context,
	name string, args *ResolverRuleArgs, opts ...pulumi.ResourceOption) (*ResolverRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.RuleType == nil {
		return nil, errors.New("invalid value for required argument 'RuleType'")
	}
	var resource ResolverRule
	err := ctx.RegisterResource("aws:route53/resolverRule:ResolverRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverRule gets an existing ResolverRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverRuleState, opts ...pulumi.ResourceOption) (*ResolverRule, error) {
	var resource ResolverRule
	err := ctx.ReadResource("aws:route53/resolverRule:ResolverRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverRule resources.
type resolverRuleState struct {
	// The ARN (Amazon Resource Name) for the resolver rule.
	Arn *string `pulumi:"arn"`
	// DNS queries for this domain name are forwarded to the IP addresses that are specified using `targetIp`.
	DomainName *string `pulumi:"domainName"`
	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	Name *string `pulumi:"name"`
	// When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
	OwnerId *string `pulumi:"ownerId"`
	// The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `targetIp`.
	// This argument should only be specified for `FORWARD` type rules.
	ResolverEndpointId *string `pulumi:"resolverEndpointId"`
	// The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
	RuleType *string `pulumi:"ruleType"`
	// Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
	// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
	ShareStatus *string `pulumi:"shareStatus"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
	// This argument should only be specified for `FORWARD` type rules.
	TargetIps []ResolverRuleTargetIp `pulumi:"targetIps"`
}

type ResolverRuleState struct {
	// The ARN (Amazon Resource Name) for the resolver rule.
	Arn pulumi.StringPtrInput
	// DNS queries for this domain name are forwarded to the IP addresses that are specified using `targetIp`.
	DomainName pulumi.StringPtrInput
	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	Name pulumi.StringPtrInput
	// When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
	OwnerId pulumi.StringPtrInput
	// The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `targetIp`.
	// This argument should only be specified for `FORWARD` type rules.
	ResolverEndpointId pulumi.StringPtrInput
	// The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
	RuleType pulumi.StringPtrInput
	// Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
	// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
	ShareStatus pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
	// This argument should only be specified for `FORWARD` type rules.
	TargetIps ResolverRuleTargetIpArrayInput
}

func (ResolverRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleState)(nil)).Elem()
}

type resolverRuleArgs struct {
	// DNS queries for this domain name are forwarded to the IP addresses that are specified using `targetIp`.
	DomainName string `pulumi:"domainName"`
	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	Name *string `pulumi:"name"`
	// The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `targetIp`.
	// This argument should only be specified for `FORWARD` type rules.
	ResolverEndpointId *string `pulumi:"resolverEndpointId"`
	// The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
	RuleType string `pulumi:"ruleType"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
	// This argument should only be specified for `FORWARD` type rules.
	TargetIps []ResolverRuleTargetIp `pulumi:"targetIps"`
}

// The set of arguments for constructing a ResolverRule resource.
type ResolverRuleArgs struct {
	// DNS queries for this domain name are forwarded to the IP addresses that are specified using `targetIp`.
	DomainName pulumi.StringInput
	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	Name pulumi.StringPtrInput
	// The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `targetIp`.
	// This argument should only be specified for `FORWARD` type rules.
	ResolverEndpointId pulumi.StringPtrInput
	// The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
	RuleType pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
	// This argument should only be specified for `FORWARD` type rules.
	TargetIps ResolverRuleTargetIpArrayInput
}

func (ResolverRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleArgs)(nil)).Elem()
}

type ResolverRuleInput interface {
	pulumi.Input

	ToResolverRuleOutput() ResolverRuleOutput
	ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput
}

func (ResolverRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverRule)(nil)).Elem()
}

func (i ResolverRule) ToResolverRuleOutput() ResolverRuleOutput {
	return i.ToResolverRuleOutputWithContext(context.Background())
}

func (i ResolverRule) ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleOutput)
}

type ResolverRuleOutput struct {
	*pulumi.OutputState
}

func (ResolverRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverRuleOutput)(nil)).Elem()
}

func (o ResolverRuleOutput) ToResolverRuleOutput() ResolverRuleOutput {
	return o
}

func (o ResolverRuleOutput) ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResolverRuleOutput{})
}
