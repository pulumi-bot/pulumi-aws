// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a way to set SNS SMS preferences.
type SmsPreferences struct {
	s *pulumi.ResourceState
}

// NewSmsPreferences registers a new resource with the given unique name, arguments, and options.
func NewSmsPreferences(ctx *pulumi.Context,
	name string, args *SmsPreferencesArgs, opts ...pulumi.ResourceOpt) (*SmsPreferences, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["defaultSenderId"] = nil
		inputs["defaultSmsType"] = nil
		inputs["deliveryStatusIamRoleArn"] = nil
		inputs["deliveryStatusSuccessSamplingRate"] = nil
		inputs["monthlySpendLimit"] = nil
		inputs["usageReportS3Bucket"] = nil
	} else {
		inputs["defaultSenderId"] = args.DefaultSenderId
		inputs["defaultSmsType"] = args.DefaultSmsType
		inputs["deliveryStatusIamRoleArn"] = args.DeliveryStatusIamRoleArn
		inputs["deliveryStatusSuccessSamplingRate"] = args.DeliveryStatusSuccessSamplingRate
		inputs["monthlySpendLimit"] = args.MonthlySpendLimit
		inputs["usageReportS3Bucket"] = args.UsageReportS3Bucket
	}
	s, err := ctx.RegisterResource("aws:sns/smsPreferences:SmsPreferences", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SmsPreferences{s: s}, nil
}

// GetSmsPreferences gets an existing SmsPreferences resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmsPreferences(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SmsPreferencesState, opts ...pulumi.ResourceOpt) (*SmsPreferences, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["defaultSenderId"] = state.DefaultSenderId
		inputs["defaultSmsType"] = state.DefaultSmsType
		inputs["deliveryStatusIamRoleArn"] = state.DeliveryStatusIamRoleArn
		inputs["deliveryStatusSuccessSamplingRate"] = state.DeliveryStatusSuccessSamplingRate
		inputs["monthlySpendLimit"] = state.MonthlySpendLimit
		inputs["usageReportS3Bucket"] = state.UsageReportS3Bucket
	}
	s, err := ctx.ReadResource("aws:sns/smsPreferences:SmsPreferences", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SmsPreferences{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SmsPreferences) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SmsPreferences) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A string, such as your business brand, that is displayed as the sender on the receiving device.
func (r *SmsPreferences) DefaultSenderId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultSenderId"])
}

// The type of SMS message that you will send by default. Possible values are: Promotional, Transactional
func (r *SmsPreferences) DefaultSmsType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultSmsType"])
}

// The ARN of the IAM role that allows Amazon SNS to write logs about SMS deliveries in CloudWatch Logs.
func (r *SmsPreferences) DeliveryStatusIamRoleArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["deliveryStatusIamRoleArn"])
}

// The percentage of successful SMS deliveries for which Amazon SNS will write logs in CloudWatch Logs. The value must be between 0 and 100.
func (r *SmsPreferences) DeliveryStatusSuccessSamplingRate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["deliveryStatusSuccessSamplingRate"])
}

// The maximum amount in USD that you are willing to spend each month to send SMS messages.
func (r *SmsPreferences) MonthlySpendLimit() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["monthlySpendLimit"])
}

// The name of the Amazon S3 bucket to receive daily SMS usage reports from Amazon SNS.
func (r *SmsPreferences) UsageReportS3Bucket() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["usageReportS3Bucket"])
}

// Input properties used for looking up and filtering SmsPreferences resources.
type SmsPreferencesState struct {
	// A string, such as your business brand, that is displayed as the sender on the receiving device.
	DefaultSenderId interface{}
	// The type of SMS message that you will send by default. Possible values are: Promotional, Transactional
	DefaultSmsType interface{}
	// The ARN of the IAM role that allows Amazon SNS to write logs about SMS deliveries in CloudWatch Logs.
	DeliveryStatusIamRoleArn interface{}
	// The percentage of successful SMS deliveries for which Amazon SNS will write logs in CloudWatch Logs. The value must be between 0 and 100.
	DeliveryStatusSuccessSamplingRate interface{}
	// The maximum amount in USD that you are willing to spend each month to send SMS messages.
	MonthlySpendLimit interface{}
	// The name of the Amazon S3 bucket to receive daily SMS usage reports from Amazon SNS.
	UsageReportS3Bucket interface{}
}

// The set of arguments for constructing a SmsPreferences resource.
type SmsPreferencesArgs struct {
	// A string, such as your business brand, that is displayed as the sender on the receiving device.
	DefaultSenderId interface{}
	// The type of SMS message that you will send by default. Possible values are: Promotional, Transactional
	DefaultSmsType interface{}
	// The ARN of the IAM role that allows Amazon SNS to write logs about SMS deliveries in CloudWatch Logs.
	DeliveryStatusIamRoleArn interface{}
	// The percentage of successful SMS deliveries for which Amazon SNS will write logs in CloudWatch Logs. The value must be between 0 and 100.
	DeliveryStatusSuccessSamplingRate interface{}
	// The maximum amount in USD that you are willing to spend each month to send SMS messages.
	MonthlySpendLimit interface{}
	// The name of the Amazon S3 bucket to receive daily SMS usage reports from Amazon SNS.
	UsageReportS3Bucket interface{}
}
