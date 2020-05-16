// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `route53getResolverRules` provides details about a set of Route53 Resolver rules.
func GetResolverRules(ctx *pulumi.Context, args *GetResolverRulesArgs, opts ...pulumi.InvokeOption) (*GetResolverRulesResult, error) {
	var rv GetResolverRulesResult
	err := ctx.Invoke("aws:route53/getResolverRules:getResolverRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResolverRules.
type GetResolverRulesArgs struct {
	// When the desired resolver rules are shared with another AWS account, the account ID of the account that the rules are shared with.
	OwnerId *string `pulumi:"ownerId"`
	// The ID of the outbound resolver endpoint for the desired resolver rules.
	ResolverEndpointId *string `pulumi:"resolverEndpointId"`
	// The rule type of the desired resolver rules. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
	RuleType *string `pulumi:"ruleType"`
	// Whether the desired resolver rules are shared and, if so, whether the current account is sharing the rules with another account, or another account is sharing the rules with the current account.
	// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
	ShareStatus *string `pulumi:"shareStatus"`
}

// A collection of values returned by getResolverRules.
type GetResolverRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	OwnerId            *string `pulumi:"ownerId"`
	ResolverEndpointId *string `pulumi:"resolverEndpointId"`
	// The IDs of the matched resolver rules.
	ResolverRuleIds []string `pulumi:"resolverRuleIds"`
	RuleType        *string  `pulumi:"ruleType"`
	ShareStatus     *string  `pulumi:"shareStatus"`
}
