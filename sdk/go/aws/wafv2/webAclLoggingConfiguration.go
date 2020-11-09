// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a WAFv2 Web ACL Logging Configuration resource.
//
// > **Note:** To start logging from a WAFv2 Web ACL, an Amazon Kinesis Data Firehose (e.g. [`kinesis.FirehoseDeliveryStream` resource](https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream.html) must also be created with a PUT source (not a stream) and in the region that you are operating.
// If you are capturing logs for Amazon CloudFront, always create the firehose in US East (N. Virginia).
// Be sure to give the data firehose a name that starts with the prefix `aws-waf-logs-`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafv2.NewWebAclLoggingConfiguration(ctx, "example", &wafv2.WebAclLoggingConfigurationArgs{
// 			LogDestinationConfigs: pulumi.StringArray{
// 				pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
// 			},
// 			ResourceArn: pulumi.Any(aws_wafv2_web_acl.Example.Arn),
// 			RedactedFields: wafv2.WebAclLoggingConfigurationRedactedFieldArray{
// 				&wafv2.WebAclLoggingConfigurationRedactedFieldArgs{
// 					SingleHeader: &wafv2.WebAclLoggingConfigurationRedactedFieldSingleHeaderArgs{
// 						Name: pulumi.String("user-agent"),
// 					},
// 				},
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
// WAFv2 Web ACL Logging Configurations can be imported using the WAFv2 Web ACL ARN e.g.
//
// ```sh
//  $ pulumi import aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration example arn:aws:wafv2:us-west-2:123456789012:regional/webacl/test-logs/a1b2c3d4-5678-90ab-cdef
// ```
type WebAclLoggingConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs pulumi.StringArrayOutput `pulumi:"logDestinationConfigs"`
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields WebAclLoggingConfigurationRedactedFieldArrayOutput `pulumi:"redactedFields"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewWebAclLoggingConfiguration registers a new resource with the given unique name, arguments, and options.
func NewWebAclLoggingConfiguration(ctx *pulumi.Context,
	name string, args *WebAclLoggingConfigurationArgs, opts ...pulumi.ResourceOption) (*WebAclLoggingConfiguration, error) {
	if args == nil || args.LogDestinationConfigs == nil {
		return nil, errors.New("missing required argument 'LogDestinationConfigs'")
	}
	if args == nil || args.ResourceArn == nil {
		return nil, errors.New("missing required argument 'ResourceArn'")
	}
	if args == nil {
		args = &WebAclLoggingConfigurationArgs{}
	}
	var resource WebAclLoggingConfiguration
	err := ctx.RegisterResource("aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAclLoggingConfiguration gets an existing WebAclLoggingConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAclLoggingConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclLoggingConfigurationState, opts ...pulumi.ResourceOption) (*WebAclLoggingConfiguration, error) {
	var resource WebAclLoggingConfiguration
	err := ctx.ReadResource("aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAclLoggingConfiguration resources.
type webAclLoggingConfigurationState struct {
	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs []string `pulumi:"logDestinationConfigs"`
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields []WebAclLoggingConfigurationRedactedField `pulumi:"redactedFields"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn *string `pulumi:"resourceArn"`
}

type WebAclLoggingConfigurationState struct {
	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs pulumi.StringArrayInput
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields WebAclLoggingConfigurationRedactedFieldArrayInput
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn pulumi.StringPtrInput
}

func (WebAclLoggingConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclLoggingConfigurationState)(nil)).Elem()
}

type webAclLoggingConfigurationArgs struct {
	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs []string `pulumi:"logDestinationConfigs"`
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields []WebAclLoggingConfigurationRedactedField `pulumi:"redactedFields"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn string `pulumi:"resourceArn"`
}

// The set of arguments for constructing a WebAclLoggingConfiguration resource.
type WebAclLoggingConfigurationArgs struct {
	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs pulumi.StringArrayInput
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields WebAclLoggingConfigurationRedactedFieldArrayInput
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn pulumi.StringInput
}

func (WebAclLoggingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclLoggingConfigurationArgs)(nil)).Elem()
}
