// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package athena

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Athena Workgroup.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/athena"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = athena.NewWorkgroup(ctx, "example", &athena.WorkgroupArgs{
// 			Configuration: &athena.WorkgroupConfigurationArgs{
// 				EnforceWorkgroupConfiguration:   pulumi.Bool(true),
// 				PublishCloudwatchMetricsEnabled: pulumi.Bool(true),
// 				ResultConfiguration: &athena.WorkgroupConfigurationResultConfigurationArgs{
// 					EncryptionConfiguration: &athena.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs{
// 						EncryptionOption: pulumi.String("SSE_KMS"),
// 						KmsKeyArn:        pulumi.String(aws_kms_key.Example.Arn),
// 					},
// 					OutputLocation: pulumi.String("s3://{aws_s3_bucket.example.bucket}/output/"),
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
type Workgroup struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the workgroup
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration WorkgroupConfigurationPtrOutput `pulumi:"configuration"`
	// Description of the workgroup.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// Name of the workgroup.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Key-value map of resource tags for the workgroup.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewWorkgroup registers a new resource with the given unique name, arguments, and options.
func NewWorkgroup(ctx *pulumi.Context,
	name string, args *WorkgroupArgs, opts ...pulumi.ResourceOption) (*Workgroup, error) {
	if args == nil {
		args = &WorkgroupArgs{}
	}
	var resource Workgroup
	err := ctx.RegisterResource("aws:athena/workgroup:Workgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkgroup gets an existing Workgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkgroupState, opts ...pulumi.ResourceOption) (*Workgroup, error) {
	var resource Workgroup
	err := ctx.ReadResource("aws:athena/workgroup:Workgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workgroup resources.
type workgroupState struct {
	// Amazon Resource Name (ARN) of the workgroup
	Arn *string `pulumi:"arn"`
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration *WorkgroupConfiguration `pulumi:"configuration"`
	// Description of the workgroup.
	Description *string `pulumi:"description"`
	// The option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Name of the workgroup.
	Name *string `pulumi:"name"`
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State *string `pulumi:"state"`
	// Key-value map of resource tags for the workgroup.
	Tags map[string]string `pulumi:"tags"`
}

type WorkgroupState struct {
	// Amazon Resource Name (ARN) of the workgroup
	Arn pulumi.StringPtrInput
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration WorkgroupConfigurationPtrInput
	// Description of the workgroup.
	Description pulumi.StringPtrInput
	// The option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy pulumi.BoolPtrInput
	// Name of the workgroup.
	Name pulumi.StringPtrInput
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State pulumi.StringPtrInput
	// Key-value map of resource tags for the workgroup.
	Tags pulumi.StringMapInput
}

func (WorkgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workgroupState)(nil)).Elem()
}

type workgroupArgs struct {
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration *WorkgroupConfiguration `pulumi:"configuration"`
	// Description of the workgroup.
	Description *string `pulumi:"description"`
	// The option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Name of the workgroup.
	Name *string `pulumi:"name"`
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State *string `pulumi:"state"`
	// Key-value map of resource tags for the workgroup.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Workgroup resource.
type WorkgroupArgs struct {
	// Configuration block with various settings for the workgroup. Documented below.
	Configuration WorkgroupConfigurationPtrInput
	// Description of the workgroup.
	Description pulumi.StringPtrInput
	// The option to delete the workgroup and its contents even if the workgroup contains any named queries.
	ForceDestroy pulumi.BoolPtrInput
	// Name of the workgroup.
	Name pulumi.StringPtrInput
	// State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
	State pulumi.StringPtrInput
	// Key-value map of resource tags for the workgroup.
	Tags pulumi.StringMapInput
}

func (WorkgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workgroupArgs)(nil)).Elem()
}
