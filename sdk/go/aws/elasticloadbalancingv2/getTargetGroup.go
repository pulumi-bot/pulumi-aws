// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Deprecated: aws.elasticloadbalancingv2.getTargetGroup has been deprecated in favor of aws.lb.getTargetGroup
func LookupTargetGroup(ctx *pulumi.Context, args *LookupTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupTargetGroupResult, error) {
	var rv LookupTargetGroupResult
	err := ctx.Invoke("aws:elasticloadbalancingv2/getTargetGroup:getTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTargetGroup.
type LookupTargetGroupArgs struct {
	Arn  *string           `pulumi:"arn"`
	Name *string           `pulumi:"name"`
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getTargetGroup.
type LookupTargetGroupResult struct {
	Arn                 string                    `pulumi:"arn"`
	ArnSuffix           string                    `pulumi:"arnSuffix"`
	DeregistrationDelay int                       `pulumi:"deregistrationDelay"`
	HealthCheck         GetTargetGroupHealthCheck `pulumi:"healthCheck"`
	// The provider-assigned unique ID for this managed resource.
	Id                             string                   `pulumi:"id"`
	LambdaMultiValueHeadersEnabled bool                     `pulumi:"lambdaMultiValueHeadersEnabled"`
	LoadBalancingAlgorithmType     string                   `pulumi:"loadBalancingAlgorithmType"`
	Name                           string                   `pulumi:"name"`
	Port                           int                      `pulumi:"port"`
	Protocol                       string                   `pulumi:"protocol"`
	ProxyProtocolV2                bool                     `pulumi:"proxyProtocolV2"`
	SlowStart                      int                      `pulumi:"slowStart"`
	Stickiness                     GetTargetGroupStickiness `pulumi:"stickiness"`
	Tags                           map[string]string        `pulumi:"tags"`
	TargetType                     string                   `pulumi:"targetType"`
	VpcId                          string                   `pulumi:"vpcId"`
}
