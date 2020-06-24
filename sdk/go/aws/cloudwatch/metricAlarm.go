// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CloudWatch Metric Alarm resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = cloudwatch.NewMetricAlarm(ctx, "foobar", &cloudwatch.MetricAlarmArgs{
// 			AlarmDescription:        pulumi.String("This metric monitors ec2 cpu utilization"),
// 			ComparisonOperator:      pulumi.String("GreaterThanOrEqualToThreshold"),
// 			EvaluationPeriods:       pulumi.Int(2),
// 			InsufficientDataActions: []interface{}{},
// 			MetricName:              pulumi.String("CPUUtilization"),
// 			Namespace:               pulumi.String("AWS/EC2"),
// 			Period:                  pulumi.Int(120),
// 			Statistic:               pulumi.String("Average"),
// 			Threshold:               pulumi.Float64(80),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Example in Conjunction with Scaling Policies
//
// ## Example with an Expression
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = cloudwatch.NewMetricAlarm(ctx, "foobar", &cloudwatch.MetricAlarmArgs{
// 			AlarmDescription:        pulumi.String(fmt.Sprintf("%v%v", "Request error rate has exceeded 10", "%")),
// 			ComparisonOperator:      pulumi.String("GreaterThanOrEqualToThreshold"),
// 			EvaluationPeriods:       pulumi.Int(2),
// 			InsufficientDataActions: []interface{}{},
// 			MetricQueries: cloudwatch.MetricAlarmMetricQueryArray{
// 				&cloudwatch.MetricAlarmMetricQueryArgs{
// 					Expression: pulumi.String("m2/m1*100"),
// 					Id:         pulumi.String("e1"),
// 					Label:      pulumi.String("Error Rate"),
// 					ReturnData: pulumi.Bool(true),
// 				},
// 				&cloudwatch.MetricAlarmMetricQueryArgs{
// 					Id: pulumi.String("m1"),
// 					Metric: &cloudwatch.MetricAlarmMetricQueryMetricArgs{
// 						Dimensions: pulumi.Map{
// 							"LoadBalancer": pulumi.String("app/web"),
// 						},
// 						MetricName: pulumi.String("RequestCount"),
// 						Namespace:  pulumi.String("AWS/ApplicationELB"),
// 						Period:     pulumi.Int(120),
// 						Stat:       pulumi.String("Sum"),
// 						Unit:       pulumi.String("Count"),
// 					},
// 				},
// 				&cloudwatch.MetricAlarmMetricQueryArgs{
// 					Id: pulumi.String("m2"),
// 					Metric: &cloudwatch.MetricAlarmMetricQueryMetricArgs{
// 						Dimensions: pulumi.Map{
// 							"LoadBalancer": pulumi.String("app/web"),
// 						},
// 						MetricName: pulumi.String("HTTPCode_ELB_5XX_Count"),
// 						Namespace:  pulumi.String("AWS/ApplicationELB"),
// 						Period:     pulumi.Int(120),
// 						Stat:       pulumi.String("Sum"),
// 						Unit:       pulumi.String("Count"),
// 					},
// 				},
// 			},
// 			Threshold: pulumi.Float64(10),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = cloudwatch.NewMetricAlarm(ctx, "xxAnomalyDetection", &cloudwatch.MetricAlarmArgs{
// 			AlarmDescription:        pulumi.String("This metric monitors ec2 cpu utilization"),
// 			ComparisonOperator:      pulumi.String("GreaterThanUpperThreshold"),
// 			EvaluationPeriods:       pulumi.Int(2),
// 			InsufficientDataActions: []interface{}{},
// 			MetricQueries: cloudwatch.MetricAlarmMetricQueryArray{
// 				&cloudwatch.MetricAlarmMetricQueryArgs{
// 					Expression: pulumi.String("ANOMALY_DETECTION_BAND(m1)"),
// 					Id:         pulumi.String("e1"),
// 					Label:      pulumi.String("CPUUtilization (Expected)"),
// 					ReturnData: pulumi.Bool(true),
// 				},
// 				&cloudwatch.MetricAlarmMetricQueryArgs{
// 					Id: pulumi.String("m1"),
// 					Metric: &cloudwatch.MetricAlarmMetricQueryMetricArgs{
// 						Dimensions: pulumi.Map{
// 							"InstanceId": pulumi.String("i-abc123"),
// 						},
// 						MetricName: pulumi.String("CPUUtilization"),
// 						Namespace:  pulumi.String("AWS/EC2"),
// 						Period:     pulumi.Int(120),
// 						Stat:       pulumi.String("Average"),
// 						Unit:       pulumi.String("Count"),
// 					},
// 					ReturnData: pulumi.Bool(true),
// 				},
// 			},
// 			ThresholdMetricId: pulumi.String("e1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Example of monitoring Healthy Hosts on NLB using Target Group and NLB
//
// > **NOTE:**  You cannot create a metric alarm consisting of both `statistic` and `extendedStatistic` parameters.
// You must choose one or the other
type MetricAlarm struct {
	pulumi.CustomResourceState

	// Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
	ActionsEnabled pulumi.BoolPtrOutput `pulumi:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	AlarmActions pulumi.StringArrayOutput `pulumi:"alarmActions"`
	// The description for the alarm.
	AlarmDescription pulumi.StringPtrOutput `pulumi:"alarmDescription"`
	// The ARN of the cloudwatch metric alarm.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
	ComparisonOperator pulumi.StringOutput `pulumi:"comparisonOperator"`
	// The number of datapoints that must be breaching to trigger the alarm.
	DatapointsToAlarm pulumi.IntPtrOutput `pulumi:"datapointsToAlarm"`
	// The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Dimensions pulumi.StringMapOutput `pulumi:"dimensions"`
	// Used only for alarms
	// based on percentiles. If you specify `ignore`, the alarm state will not
	// change during periods with too few data points to be statistically significant.
	// If you specify `evaluate` or omit this parameter, the alarm will always be
	// evaluated and possibly change state no matter how many data points are available.
	// The following values are supported: `ignore`, and `evaluate`.
	EvaluateLowSampleCountPercentiles pulumi.StringOutput `pulumi:"evaluateLowSampleCountPercentiles"`
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods pulumi.IntOutput `pulumi:"evaluationPeriods"`
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	ExtendedStatistic pulumi.StringPtrOutput `pulumi:"extendedStatistic"`
	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions pulumi.StringArrayOutput `pulumi:"insufficientDataActions"`
	// The name for this metric.
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	MetricName pulumi.StringPtrOutput `pulumi:"metricName"`
	// Enables you to create an alarm based on a metric math expression. You may specify at most 20.
	MetricQueries MetricAlarmMetricQueryArrayOutput `pulumi:"metricQueries"`
	// The descriptive name for the alarm. This name must be unique within the user's AWS account
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	OkActions pulumi.StringArrayOutput `pulumi:"okActions"`
	// The period in seconds over which the specified `stat` is applied.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The statistic to apply to the alarm's associated metric.
	// Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
	Statistic pulumi.StringPtrOutput `pulumi:"statistic"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
	Threshold pulumi.Float64PtrOutput `pulumi:"threshold"`
	// If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
	ThresholdMetricId pulumi.StringPtrOutput `pulumi:"thresholdMetricId"`
	// Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
	TreatMissingData pulumi.StringPtrOutput `pulumi:"treatMissingData"`
	// The unit for this metric.
	Unit pulumi.StringPtrOutput `pulumi:"unit"`
}

// NewMetricAlarm registers a new resource with the given unique name, arguments, and options.
func NewMetricAlarm(ctx *pulumi.Context,
	name string, args *MetricAlarmArgs, opts ...pulumi.ResourceOption) (*MetricAlarm, error) {
	if args == nil || args.ComparisonOperator == nil {
		return nil, errors.New("missing required argument 'ComparisonOperator'")
	}
	if args == nil || args.EvaluationPeriods == nil {
		return nil, errors.New("missing required argument 'EvaluationPeriods'")
	}
	if args == nil {
		args = &MetricAlarmArgs{}
	}
	var resource MetricAlarm
	err := ctx.RegisterResource("aws:cloudwatch/metricAlarm:MetricAlarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricAlarm gets an existing MetricAlarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricAlarmState, opts ...pulumi.ResourceOption) (*MetricAlarm, error) {
	var resource MetricAlarm
	err := ctx.ReadResource("aws:cloudwatch/metricAlarm:MetricAlarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricAlarm resources.
type metricAlarmState struct {
	// Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
	ActionsEnabled *bool `pulumi:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	AlarmActions []string `pulumi:"alarmActions"`
	// The description for the alarm.
	AlarmDescription *string `pulumi:"alarmDescription"`
	// The ARN of the cloudwatch metric alarm.
	Arn *string `pulumi:"arn"`
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
	ComparisonOperator *string `pulumi:"comparisonOperator"`
	// The number of datapoints that must be breaching to trigger the alarm.
	DatapointsToAlarm *int `pulumi:"datapointsToAlarm"`
	// The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Dimensions map[string]string `pulumi:"dimensions"`
	// Used only for alarms
	// based on percentiles. If you specify `ignore`, the alarm state will not
	// change during periods with too few data points to be statistically significant.
	// If you specify `evaluate` or omit this parameter, the alarm will always be
	// evaluated and possibly change state no matter how many data points are available.
	// The following values are supported: `ignore`, and `evaluate`.
	EvaluateLowSampleCountPercentiles *string `pulumi:"evaluateLowSampleCountPercentiles"`
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *int `pulumi:"evaluationPeriods"`
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	ExtendedStatistic *string `pulumi:"extendedStatistic"`
	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions []string `pulumi:"insufficientDataActions"`
	// The name for this metric.
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	MetricName *string `pulumi:"metricName"`
	// Enables you to create an alarm based on a metric math expression. You may specify at most 20.
	MetricQueries []MetricAlarmMetricQuery `pulumi:"metricQueries"`
	// The descriptive name for the alarm. This name must be unique within the user's AWS account
	Name *string `pulumi:"name"`
	// The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Namespace *string `pulumi:"namespace"`
	// The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	OkActions []string `pulumi:"okActions"`
	// The period in seconds over which the specified `stat` is applied.
	Period *int `pulumi:"period"`
	// The statistic to apply to the alarm's associated metric.
	// Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
	Statistic *string `pulumi:"statistic"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
	Threshold *float64 `pulumi:"threshold"`
	// If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
	ThresholdMetricId *string `pulumi:"thresholdMetricId"`
	// Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
	TreatMissingData *string `pulumi:"treatMissingData"`
	// The unit for this metric.
	Unit *string `pulumi:"unit"`
}

type MetricAlarmState struct {
	// Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
	ActionsEnabled pulumi.BoolPtrInput
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	AlarmActions pulumi.StringArrayInput
	// The description for the alarm.
	AlarmDescription pulumi.StringPtrInput
	// The ARN of the cloudwatch metric alarm.
	Arn pulumi.StringPtrInput
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
	ComparisonOperator pulumi.StringPtrInput
	// The number of datapoints that must be breaching to trigger the alarm.
	DatapointsToAlarm pulumi.IntPtrInput
	// The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Dimensions pulumi.StringMapInput
	// Used only for alarms
	// based on percentiles. If you specify `ignore`, the alarm state will not
	// change during periods with too few data points to be statistically significant.
	// If you specify `evaluate` or omit this parameter, the alarm will always be
	// evaluated and possibly change state no matter how many data points are available.
	// The following values are supported: `ignore`, and `evaluate`.
	EvaluateLowSampleCountPercentiles pulumi.StringPtrInput
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods pulumi.IntPtrInput
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	ExtendedStatistic pulumi.StringPtrInput
	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions pulumi.StringArrayInput
	// The name for this metric.
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	MetricName pulumi.StringPtrInput
	// Enables you to create an alarm based on a metric math expression. You may specify at most 20.
	MetricQueries MetricAlarmMetricQueryArrayInput
	// The descriptive name for the alarm. This name must be unique within the user's AWS account
	Name pulumi.StringPtrInput
	// The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Namespace pulumi.StringPtrInput
	// The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	OkActions pulumi.StringArrayInput
	// The period in seconds over which the specified `stat` is applied.
	Period pulumi.IntPtrInput
	// The statistic to apply to the alarm's associated metric.
	// Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
	Statistic pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
	Threshold pulumi.Float64PtrInput
	// If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
	ThresholdMetricId pulumi.StringPtrInput
	// Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
	TreatMissingData pulumi.StringPtrInput
	// The unit for this metric.
	Unit pulumi.StringPtrInput
}

func (MetricAlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricAlarmState)(nil)).Elem()
}

type metricAlarmArgs struct {
	// Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
	ActionsEnabled *bool `pulumi:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	AlarmActions []interface{} `pulumi:"alarmActions"`
	// The description for the alarm.
	AlarmDescription *string `pulumi:"alarmDescription"`
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
	ComparisonOperator string `pulumi:"comparisonOperator"`
	// The number of datapoints that must be breaching to trigger the alarm.
	DatapointsToAlarm *int `pulumi:"datapointsToAlarm"`
	// The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Dimensions map[string]string `pulumi:"dimensions"`
	// Used only for alarms
	// based on percentiles. If you specify `ignore`, the alarm state will not
	// change during periods with too few data points to be statistically significant.
	// If you specify `evaluate` or omit this parameter, the alarm will always be
	// evaluated and possibly change state no matter how many data points are available.
	// The following values are supported: `ignore`, and `evaluate`.
	EvaluateLowSampleCountPercentiles *string `pulumi:"evaluateLowSampleCountPercentiles"`
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods int `pulumi:"evaluationPeriods"`
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	ExtendedStatistic *string `pulumi:"extendedStatistic"`
	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions []interface{} `pulumi:"insufficientDataActions"`
	// The name for this metric.
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	MetricName *string `pulumi:"metricName"`
	// Enables you to create an alarm based on a metric math expression. You may specify at most 20.
	MetricQueries []MetricAlarmMetricQuery `pulumi:"metricQueries"`
	// The descriptive name for the alarm. This name must be unique within the user's AWS account
	Name *string `pulumi:"name"`
	// The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Namespace *string `pulumi:"namespace"`
	// The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	OkActions []interface{} `pulumi:"okActions"`
	// The period in seconds over which the specified `stat` is applied.
	Period *int `pulumi:"period"`
	// The statistic to apply to the alarm's associated metric.
	// Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
	Statistic *string `pulumi:"statistic"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
	Threshold *float64 `pulumi:"threshold"`
	// If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
	ThresholdMetricId *string `pulumi:"thresholdMetricId"`
	// Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
	TreatMissingData *string `pulumi:"treatMissingData"`
	// The unit for this metric.
	Unit *string `pulumi:"unit"`
}

// The set of arguments for constructing a MetricAlarm resource.
type MetricAlarmArgs struct {
	// Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
	ActionsEnabled pulumi.BoolPtrInput
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	AlarmActions pulumi.ArrayInput
	// The description for the alarm.
	AlarmDescription pulumi.StringPtrInput
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
	ComparisonOperator pulumi.StringInput
	// The number of datapoints that must be breaching to trigger the alarm.
	DatapointsToAlarm pulumi.IntPtrInput
	// The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Dimensions pulumi.StringMapInput
	// Used only for alarms
	// based on percentiles. If you specify `ignore`, the alarm state will not
	// change during periods with too few data points to be statistically significant.
	// If you specify `evaluate` or omit this parameter, the alarm will always be
	// evaluated and possibly change state no matter how many data points are available.
	// The following values are supported: `ignore`, and `evaluate`.
	EvaluateLowSampleCountPercentiles pulumi.StringPtrInput
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods pulumi.IntInput
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	ExtendedStatistic pulumi.StringPtrInput
	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions pulumi.ArrayInput
	// The name for this metric.
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	MetricName pulumi.StringPtrInput
	// Enables you to create an alarm based on a metric math expression. You may specify at most 20.
	MetricQueries MetricAlarmMetricQueryArrayInput
	// The descriptive name for the alarm. This name must be unique within the user's AWS account
	Name pulumi.StringPtrInput
	// The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
	// See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
	Namespace pulumi.StringPtrInput
	// The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	OkActions pulumi.ArrayInput
	// The period in seconds over which the specified `stat` is applied.
	Period pulumi.IntPtrInput
	// The statistic to apply to the alarm's associated metric.
	// Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
	Statistic pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
	Threshold pulumi.Float64PtrInput
	// If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
	ThresholdMetricId pulumi.StringPtrInput
	// Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
	TreatMissingData pulumi.StringPtrInput
	// The unit for this metric.
	Unit pulumi.StringPtrInput
}

func (MetricAlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricAlarmArgs)(nil)).Elem()
}
