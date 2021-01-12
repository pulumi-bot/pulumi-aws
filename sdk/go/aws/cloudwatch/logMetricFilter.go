// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CloudWatch Log Metric Filter resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		dada, err := cloudwatch.NewLogGroup(ctx, "dada", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewLogMetricFilter(ctx, "yada", &cloudwatch.LogMetricFilterArgs{
// 			Pattern:      pulumi.String(""),
// 			LogGroupName: dada.Name,
// 			MetricTransformation: &cloudwatch.LogMetricFilterMetricTransformationArgs{
// 				Name:      pulumi.String("EventCount"),
// 				Namespace: pulumi.String("YourNamespace"),
// 				Value:     pulumi.String("1"),
// 			},
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
// CloudWatch Log Metric Filter can be imported using the `log_group_name:name`, e.g.
//
// ```sh
//  $ pulumi import aws:cloudwatch/logMetricFilter:LogMetricFilter test /aws/lambda/function:test
// ```
type LogMetricFilter struct {
	pulumi.CustomResourceState

	// The name of the log group to associate the metric filter with.
	LogGroupName pulumi.StringOutput `pulumi:"logGroupName"`
	// A block defining collection of information needed to define how metric data gets emitted. See below.
	MetricTransformation LogMetricFilterMetricTransformationOutput `pulumi:"metricTransformation"`
	// A name for the metric filter.
	Name pulumi.StringOutput `pulumi:"name"`
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern pulumi.StringOutput `pulumi:"pattern"`
}

// NewLogMetricFilter registers a new resource with the given unique name, arguments, and options.
func NewLogMetricFilter(ctx *pulumi.Context,
	name string, args *LogMetricFilterArgs, opts ...pulumi.ResourceOption) (*LogMetricFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogGroupName == nil {
		return nil, errors.New("invalid value for required argument 'LogGroupName'")
	}
	if args.MetricTransformation == nil {
		return nil, errors.New("invalid value for required argument 'MetricTransformation'")
	}
	if args.Pattern == nil {
		return nil, errors.New("invalid value for required argument 'Pattern'")
	}
	var resource LogMetricFilter
	err := ctx.RegisterResource("aws:cloudwatch/logMetricFilter:LogMetricFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogMetricFilter gets an existing LogMetricFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogMetricFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogMetricFilterState, opts ...pulumi.ResourceOption) (*LogMetricFilter, error) {
	var resource LogMetricFilter
	err := ctx.ReadResource("aws:cloudwatch/logMetricFilter:LogMetricFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogMetricFilter resources.
type logMetricFilterState struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName *string `pulumi:"logGroupName"`
	// A block defining collection of information needed to define how metric data gets emitted. See below.
	MetricTransformation *LogMetricFilterMetricTransformation `pulumi:"metricTransformation"`
	// A name for the metric filter.
	Name *string `pulumi:"name"`
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern *string `pulumi:"pattern"`
}

type LogMetricFilterState struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName pulumi.StringPtrInput
	// A block defining collection of information needed to define how metric data gets emitted. See below.
	MetricTransformation LogMetricFilterMetricTransformationPtrInput
	// A name for the metric filter.
	Name pulumi.StringPtrInput
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern pulumi.StringPtrInput
}

func (LogMetricFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logMetricFilterState)(nil)).Elem()
}

type logMetricFilterArgs struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName string `pulumi:"logGroupName"`
	// A block defining collection of information needed to define how metric data gets emitted. See below.
	MetricTransformation LogMetricFilterMetricTransformation `pulumi:"metricTransformation"`
	// A name for the metric filter.
	Name *string `pulumi:"name"`
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern string `pulumi:"pattern"`
}

// The set of arguments for constructing a LogMetricFilter resource.
type LogMetricFilterArgs struct {
	// The name of the log group to associate the metric filter with.
	LogGroupName pulumi.StringInput
	// A block defining collection of information needed to define how metric data gets emitted. See below.
	MetricTransformation LogMetricFilterMetricTransformationInput
	// A name for the metric filter.
	Name pulumi.StringPtrInput
	// A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
	// for extracting metric data out of ingested log events.
	Pattern pulumi.StringInput
}

func (LogMetricFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logMetricFilterArgs)(nil)).Elem()
}

type LogMetricFilterInput interface {
	pulumi.Input

	ToLogMetricFilterOutput() LogMetricFilterOutput
	ToLogMetricFilterOutputWithContext(ctx context.Context) LogMetricFilterOutput
}

func (*LogMetricFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMetricFilter)(nil))
}

func (i *LogMetricFilter) ToLogMetricFilterOutput() LogMetricFilterOutput {
	return i.ToLogMetricFilterOutputWithContext(context.Background())
}

func (i *LogMetricFilter) ToLogMetricFilterOutputWithContext(ctx context.Context) LogMetricFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricFilterOutput)
}

func (i *LogMetricFilter) ToLogMetricFilterPtrOutput() LogMetricFilterPtrOutput {
	return i.ToLogMetricFilterPtrOutputWithContext(context.Background())
}

func (i *LogMetricFilter) ToLogMetricFilterPtrOutputWithContext(ctx context.Context) LogMetricFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricFilterPtrOutput)
}

type LogMetricFilterPtrInput interface {
	pulumi.Input

	ToLogMetricFilterPtrOutput() LogMetricFilterPtrOutput
	ToLogMetricFilterPtrOutputWithContext(ctx context.Context) LogMetricFilterPtrOutput
}

type logMetricFilterPtrType LogMetricFilterArgs

func (*logMetricFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMetricFilter)(nil))
}

func (i *logMetricFilterPtrType) ToLogMetricFilterPtrOutput() LogMetricFilterPtrOutput {
	return i.ToLogMetricFilterPtrOutputWithContext(context.Background())
}

func (i *logMetricFilterPtrType) ToLogMetricFilterPtrOutputWithContext(ctx context.Context) LogMetricFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricFilterOutput).ToLogMetricFilterPtrOutput()
}

// LogMetricFilterArrayInput is an input type that accepts LogMetricFilterArray and LogMetricFilterArrayOutput values.
// You can construct a concrete instance of `LogMetricFilterArrayInput` via:
//
//          LogMetricFilterArray{ LogMetricFilterArgs{...} }
type LogMetricFilterArrayInput interface {
	pulumi.Input

	ToLogMetricFilterArrayOutput() LogMetricFilterArrayOutput
	ToLogMetricFilterArrayOutputWithContext(context.Context) LogMetricFilterArrayOutput
}

type LogMetricFilterArray []LogMetricFilterInput

func (LogMetricFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogMetricFilter)(nil))
}

func (i LogMetricFilterArray) ToLogMetricFilterArrayOutput() LogMetricFilterArrayOutput {
	return i.ToLogMetricFilterArrayOutputWithContext(context.Background())
}

func (i LogMetricFilterArray) ToLogMetricFilterArrayOutputWithContext(ctx context.Context) LogMetricFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricFilterArrayOutput)
}

// LogMetricFilterMapInput is an input type that accepts LogMetricFilterMap and LogMetricFilterMapOutput values.
// You can construct a concrete instance of `LogMetricFilterMapInput` via:
//
//          LogMetricFilterMap{ "key": LogMetricFilterArgs{...} }
type LogMetricFilterMapInput interface {
	pulumi.Input

	ToLogMetricFilterMapOutput() LogMetricFilterMapOutput
	ToLogMetricFilterMapOutputWithContext(context.Context) LogMetricFilterMapOutput
}

type LogMetricFilterMap map[string]LogMetricFilterInput

func (LogMetricFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogMetricFilter)(nil))
}

func (i LogMetricFilterMap) ToLogMetricFilterMapOutput() LogMetricFilterMapOutput {
	return i.ToLogMetricFilterMapOutputWithContext(context.Background())
}

func (i LogMetricFilterMap) ToLogMetricFilterMapOutputWithContext(ctx context.Context) LogMetricFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricFilterMapOutput)
}

type LogMetricFilterOutput struct {
	*pulumi.OutputState
}

func (LogMetricFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMetricFilter)(nil))
}

func (o LogMetricFilterOutput) ToLogMetricFilterOutput() LogMetricFilterOutput {
	return o
}

func (o LogMetricFilterOutput) ToLogMetricFilterOutputWithContext(ctx context.Context) LogMetricFilterOutput {
	return o
}

func (o LogMetricFilterOutput) ToLogMetricFilterPtrOutput() LogMetricFilterPtrOutput {
	return o.ToLogMetricFilterPtrOutputWithContext(context.Background())
}

func (o LogMetricFilterOutput) ToLogMetricFilterPtrOutputWithContext(ctx context.Context) LogMetricFilterPtrOutput {
	return o.ApplyT(func(v LogMetricFilter) *LogMetricFilter {
		return &v
	}).(LogMetricFilterPtrOutput)
}

type LogMetricFilterPtrOutput struct {
	*pulumi.OutputState
}

func (LogMetricFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMetricFilter)(nil))
}

func (o LogMetricFilterPtrOutput) ToLogMetricFilterPtrOutput() LogMetricFilterPtrOutput {
	return o
}

func (o LogMetricFilterPtrOutput) ToLogMetricFilterPtrOutputWithContext(ctx context.Context) LogMetricFilterPtrOutput {
	return o
}

type LogMetricFilterArrayOutput struct{ *pulumi.OutputState }

func (LogMetricFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogMetricFilter)(nil))
}

func (o LogMetricFilterArrayOutput) ToLogMetricFilterArrayOutput() LogMetricFilterArrayOutput {
	return o
}

func (o LogMetricFilterArrayOutput) ToLogMetricFilterArrayOutputWithContext(ctx context.Context) LogMetricFilterArrayOutput {
	return o
}

func (o LogMetricFilterArrayOutput) Index(i pulumi.IntInput) LogMetricFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogMetricFilter {
		return vs[0].([]LogMetricFilter)[vs[1].(int)]
	}).(LogMetricFilterOutput)
}

type LogMetricFilterMapOutput struct{ *pulumi.OutputState }

func (LogMetricFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogMetricFilter)(nil))
}

func (o LogMetricFilterMapOutput) ToLogMetricFilterMapOutput() LogMetricFilterMapOutput {
	return o
}

func (o LogMetricFilterMapOutput) ToLogMetricFilterMapOutputWithContext(ctx context.Context) LogMetricFilterMapOutput {
	return o
}

func (o LogMetricFilterMapOutput) MapIndex(k pulumi.StringInput) LogMetricFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogMetricFilter {
		return vs[0].(map[string]LogMetricFilter)[vs[1].(string)]
	}).(LogMetricFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(LogMetricFilterOutput{})
	pulumi.RegisterOutputType(LogMetricFilterPtrOutput{})
	pulumi.RegisterOutputType(LogMetricFilterArrayOutput{})
	pulumi.RegisterOutputType(LogMetricFilterMapOutput{})
}
