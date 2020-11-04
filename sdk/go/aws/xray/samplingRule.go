// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package xray

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates and manages an AWS XRay Sampling Rule.
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
// 		_, err := xray.NewSamplingRule(ctx, "example", &xray.SamplingRuleArgs{
// 			Attributes: pulumi.StringMap{
// 				"Hello": pulumi.String("Tris"),
// 			},
// 			FixedRate:     pulumi.Float64(0.05),
// 			Host:          pulumi.String("*"),
// 			HttpMethod:    pulumi.String("*"),
// 			Priority:      pulumi.Int(10000),
// 			ReservoirSize: pulumi.Int(1),
// 			ResourceArn:   pulumi.String("*"),
// 			RuleName:      pulumi.String("example"),
// 			ServiceName:   pulumi.String("*"),
// 			ServiceType:   pulumi.String("*"),
// 			UrlPath:       pulumi.String("*"),
// 			Version:       pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SamplingRule struct {
	pulumi.CustomResourceState

	// The ARN of the sampling rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Matches attributes derived from the request.
	Attributes pulumi.StringMapOutput `pulumi:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate pulumi.Float64Output `pulumi:"fixedRate"`
	// Matches the hostname from a request URL.
	Host pulumi.StringOutput `pulumi:"host"`
	// Matches the HTTP method of a request.
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`
	// The priority of the sampling rule.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize pulumi.IntOutput `pulumi:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// The name of the sampling rule.
	RuleName pulumi.StringPtrOutput `pulumi:"ruleName"`
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
	// Key-value mapping of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Matches the path from a request URL.
	UrlPath pulumi.StringOutput `pulumi:"urlPath"`
	// The version of the sampling rule format (`1` )
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewSamplingRule registers a new resource with the given unique name, arguments, and options.
func NewSamplingRule(ctx *pulumi.Context,
	name string, args *SamplingRuleArgs, opts ...pulumi.ResourceOption) (*SamplingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.FixedRate == nil {
		return nil, errors.New("invalid value for required argument 'FixedRate'")
	}
	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.HttpMethod == nil {
		return nil, errors.New("invalid value for required argument 'HttpMethod'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.ReservoirSize == nil {
		return nil, errors.New("invalid value for required argument 'ReservoirSize'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.ServiceType == nil {
		return nil, errors.New("invalid value for required argument 'ServiceType'")
	}
	if args.UrlPath == nil {
		return nil, errors.New("invalid value for required argument 'UrlPath'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource SamplingRule
	err := ctx.RegisterResource("aws:xray/samplingRule:SamplingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamplingRule gets an existing SamplingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamplingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamplingRuleState, opts ...pulumi.ResourceOption) (*SamplingRule, error) {
	var resource SamplingRule
	err := ctx.ReadResource("aws:xray/samplingRule:SamplingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamplingRule resources.
type samplingRuleState struct {
	// The ARN of the sampling rule.
	Arn *string `pulumi:"arn"`
	// Matches attributes derived from the request.
	Attributes map[string]string `pulumi:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate *float64 `pulumi:"fixedRate"`
	// Matches the hostname from a request URL.
	Host *string `pulumi:"host"`
	// Matches the HTTP method of a request.
	HttpMethod *string `pulumi:"httpMethod"`
	// The priority of the sampling rule.
	Priority *int `pulumi:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize *int `pulumi:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn *string `pulumi:"resourceArn"`
	// The name of the sampling rule.
	RuleName *string `pulumi:"ruleName"`
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName *string `pulumi:"serviceName"`
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType *string `pulumi:"serviceType"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Matches the path from a request URL.
	UrlPath *string `pulumi:"urlPath"`
	// The version of the sampling rule format (`1` )
	Version *int `pulumi:"version"`
}

type SamplingRuleState struct {
	// The ARN of the sampling rule.
	Arn pulumi.StringPtrInput
	// Matches attributes derived from the request.
	Attributes pulumi.StringMapInput
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate pulumi.Float64PtrInput
	// Matches the hostname from a request URL.
	Host pulumi.StringPtrInput
	// Matches the HTTP method of a request.
	HttpMethod pulumi.StringPtrInput
	// The priority of the sampling rule.
	Priority pulumi.IntPtrInput
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize pulumi.IntPtrInput
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn pulumi.StringPtrInput
	// The name of the sampling rule.
	RuleName pulumi.StringPtrInput
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName pulumi.StringPtrInput
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
	// Matches the path from a request URL.
	UrlPath pulumi.StringPtrInput
	// The version of the sampling rule format (`1` )
	Version pulumi.IntPtrInput
}

func (SamplingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*samplingRuleState)(nil)).Elem()
}

type samplingRuleArgs struct {
	// Matches attributes derived from the request.
	Attributes map[string]string `pulumi:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate float64 `pulumi:"fixedRate"`
	// Matches the hostname from a request URL.
	Host string `pulumi:"host"`
	// Matches the HTTP method of a request.
	HttpMethod string `pulumi:"httpMethod"`
	// The priority of the sampling rule.
	Priority int `pulumi:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize int `pulumi:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn string `pulumi:"resourceArn"`
	// The name of the sampling rule.
	RuleName *string `pulumi:"ruleName"`
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName string `pulumi:"serviceName"`
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType string `pulumi:"serviceType"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Matches the path from a request URL.
	UrlPath string `pulumi:"urlPath"`
	// The version of the sampling rule format (`1` )
	Version int `pulumi:"version"`
}

// The set of arguments for constructing a SamplingRule resource.
type SamplingRuleArgs struct {
	// Matches attributes derived from the request.
	Attributes pulumi.StringMapInput
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate pulumi.Float64Input
	// Matches the hostname from a request URL.
	Host pulumi.StringInput
	// Matches the HTTP method of a request.
	HttpMethod pulumi.StringInput
	// The priority of the sampling rule.
	Priority pulumi.IntInput
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize pulumi.IntInput
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn pulumi.StringInput
	// The name of the sampling rule.
	RuleName pulumi.StringPtrInput
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName pulumi.StringInput
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType pulumi.StringInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
	// Matches the path from a request URL.
	UrlPath pulumi.StringInput
	// The version of the sampling rule format (`1` )
	Version pulumi.IntInput
}

func (SamplingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samplingRuleArgs)(nil)).Elem()
}
