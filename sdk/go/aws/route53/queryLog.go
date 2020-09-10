// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type QueryLog struct {
	pulumi.CustomResourceState

	CloudwatchLogGroupArn pulumi.StringOutput `pulumi:"cloudwatchLogGroupArn"`
	ZoneId                pulumi.StringOutput `pulumi:"zoneId"`
}

// NewQueryLog registers a new resource with the given unique name, arguments, and options.
func NewQueryLog(ctx *pulumi.Context,
	name string, args *QueryLogArgs, opts ...pulumi.ResourceOption) (*QueryLog, error) {
	if args == nil || args.CloudwatchLogGroupArn == nil {
		return nil, errors.New("missing required argument 'CloudwatchLogGroupArn'")
	}
	if args == nil || args.ZoneId == nil {
		return nil, errors.New("missing required argument 'ZoneId'")
	}
	if args == nil {
		args = &QueryLogArgs{}
	}
	var resource QueryLog
	err := ctx.RegisterResource("aws:route53/queryLog:QueryLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueryLog gets an existing QueryLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueryLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryLogState, opts ...pulumi.ResourceOption) (*QueryLog, error) {
	var resource QueryLog
	err := ctx.ReadResource("aws:route53/queryLog:QueryLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueryLog resources.
type queryLogState struct {
	CloudwatchLogGroupArn *string `pulumi:"cloudwatchLogGroupArn"`
	ZoneId                *string `pulumi:"zoneId"`
}

type QueryLogState struct {
	CloudwatchLogGroupArn pulumi.StringPtrInput
	ZoneId                pulumi.StringPtrInput
}

func (QueryLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryLogState)(nil)).Elem()
}

type queryLogArgs struct {
	CloudwatchLogGroupArn string `pulumi:"cloudwatchLogGroupArn"`
	ZoneId                string `pulumi:"zoneId"`
}

// The set of arguments for constructing a QueryLog resource.
type QueryLogArgs struct {
	CloudwatchLogGroupArn pulumi.StringInput
	ZoneId                pulumi.StringInput
}

func (QueryLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryLogArgs)(nil)).Elem()
}
