// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a CloudFormation StackSet. StackSets allow CloudFormation templates to be easily deployed across multiple accounts and regions via StackSet Instances (`cloudformation.StackSetInstance` resource). Additional information about StackSets can be found in the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html).
//
// > **NOTE:** All template parameters, including those with a `Default`, must be configured or ignored with the `lifecycle` configuration block `ignoreChanges` argument.
//
// > **NOTE:** All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignoreChanges` argument.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudformation"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"sts:AssumeRole",
// 					},
// 					Effect: "Allow",
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Identifiers: []string{
// 								"cloudformation.amazonaws.com",
// 							},
// 							Type: "Service",
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		aWSCloudFormationStackSetAdministrationRole, err := iam.NewRole(ctx, "aWSCloudFormationStackSetAdministrationRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := cloudformation.NewStackSet(ctx, "example", &cloudformation.StackSetArgs{
// 			AdministrationRoleArn: aWSCloudFormationStackSetAdministrationRole.Arn,
// 			Parameters: pulumi.StringMap{
// 				"VPCCidr": pulumi.String("10.0.0.0/16"),
// 			},
// 			TemplateBody: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Parameters\" : {\n", "    \"VPCCidr\" : {\n", "      \"Type\" : \"String\",\n", "      \"Default\" : \"10.0.0.0/16\",\n", "      \"Description\" : \"Enter the CIDR block for the VPC. Default is 10.0.0.0/16.\"\n", "    }\n", "  },\n", "  \"Resources\" : {\n", "    \"myVpc\": {\n", "      \"Type\" : \"AWS::EC2::VPC\",\n", "      \"Properties\" : {\n", "        \"CidrBlock\" : { \"Ref\" : \"VPCCidr\" },\n", "        \"Tags\" : [\n", "          {\"Key\": \"Name\", \"Value\": \"Primary_CF_VPC\"}\n", "        ]\n", "      }\n", "    }\n", "  }\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy", &iam.RolePolicyArgs{
// 			Policy: aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.ApplyT(func(aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument iam.GetPolicyDocumentResult) (string, error) {
// 				return aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.Json, nil
// 			}).(pulumi.StringOutput),
// 			Role: aWSCloudFormationStackSetAdministrationRole.Name,
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
// CloudFormation StackSets can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:cloudformation/stackSet:StackSet example example
// ```
type StackSet struct {
	pulumi.CustomResourceState

	// Amazon Resource Number (ARN) of the IAM Role in the administrator account.
	AdministrationRoleArn pulumi.StringOutput `pulumi:"administrationRoleArn"`
	// Amazon Resource Name (ARN) of the StackSet.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
	Capabilities pulumi.StringArrayOutput `pulumi:"capabilities"`
	// Description of the StackSet.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
	ExecutionRoleName pulumi.StringPtrOutput `pulumi:"executionRoleName"`
	// Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignoreChanges` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignoreChanges` argument.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Unique identifier of the StackSet.
	StackSetId pulumi.StringOutput `pulumi:"stackSetId"`
	// Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `templateUrl`.
	TemplateBody pulumi.StringOutput `pulumi:"templateBody"`
	// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `templateBody`.
	TemplateUrl pulumi.StringPtrOutput `pulumi:"templateUrl"`
}

// NewStackSet registers a new resource with the given unique name, arguments, and options.
func NewStackSet(ctx *pulumi.Context,
	name string, args *StackSetArgs, opts ...pulumi.ResourceOption) (*StackSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdministrationRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'AdministrationRoleArn'")
	}
	var resource StackSet
	err := ctx.RegisterResource("aws:cloudformation/stackSet:StackSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStackSet gets an existing StackSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStackSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackSetState, opts ...pulumi.ResourceOption) (*StackSet, error) {
	var resource StackSet
	err := ctx.ReadResource("aws:cloudformation/stackSet:StackSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StackSet resources.
type stackSetState struct {
	// Amazon Resource Number (ARN) of the IAM Role in the administrator account.
	AdministrationRoleArn *string `pulumi:"administrationRoleArn"`
	// Amazon Resource Name (ARN) of the StackSet.
	Arn *string `pulumi:"arn"`
	// A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
	Capabilities []string `pulumi:"capabilities"`
	// Description of the StackSet.
	Description *string `pulumi:"description"`
	// Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
	ExecutionRoleName *string `pulumi:"executionRoleName"`
	// Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
	Name *string `pulumi:"name"`
	// Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignoreChanges` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignoreChanges` argument.
	Parameters map[string]string `pulumi:"parameters"`
	// Unique identifier of the StackSet.
	StackSetId *string `pulumi:"stackSetId"`
	// Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
	Tags map[string]string `pulumi:"tags"`
	// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `templateUrl`.
	TemplateBody *string `pulumi:"templateBody"`
	// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `templateBody`.
	TemplateUrl *string `pulumi:"templateUrl"`
}

type StackSetState struct {
	// Amazon Resource Number (ARN) of the IAM Role in the administrator account.
	AdministrationRoleArn pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the StackSet.
	Arn pulumi.StringPtrInput
	// A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
	Capabilities pulumi.StringArrayInput
	// Description of the StackSet.
	Description pulumi.StringPtrInput
	// Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
	ExecutionRoleName pulumi.StringPtrInput
	// Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
	Name pulumi.StringPtrInput
	// Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignoreChanges` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignoreChanges` argument.
	Parameters pulumi.StringMapInput
	// Unique identifier of the StackSet.
	StackSetId pulumi.StringPtrInput
	// Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
	Tags pulumi.StringMapInput
	// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `templateUrl`.
	TemplateBody pulumi.StringPtrInput
	// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `templateBody`.
	TemplateUrl pulumi.StringPtrInput
}

func (StackSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackSetState)(nil)).Elem()
}

type stackSetArgs struct {
	// Amazon Resource Number (ARN) of the IAM Role in the administrator account.
	AdministrationRoleArn string `pulumi:"administrationRoleArn"`
	// A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
	Capabilities []string `pulumi:"capabilities"`
	// Description of the StackSet.
	Description *string `pulumi:"description"`
	// Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
	ExecutionRoleName *string `pulumi:"executionRoleName"`
	// Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
	Name *string `pulumi:"name"`
	// Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignoreChanges` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignoreChanges` argument.
	Parameters map[string]string `pulumi:"parameters"`
	// Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
	Tags map[string]string `pulumi:"tags"`
	// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `templateUrl`.
	TemplateBody *string `pulumi:"templateBody"`
	// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `templateBody`.
	TemplateUrl *string `pulumi:"templateUrl"`
}

// The set of arguments for constructing a StackSet resource.
type StackSetArgs struct {
	// Amazon Resource Number (ARN) of the IAM Role in the administrator account.
	AdministrationRoleArn pulumi.StringInput
	// A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
	Capabilities pulumi.StringArrayInput
	// Description of the StackSet.
	Description pulumi.StringPtrInput
	// Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
	ExecutionRoleName pulumi.StringPtrInput
	// Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
	Name pulumi.StringPtrInput
	// Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignoreChanges` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignoreChanges` argument.
	Parameters pulumi.StringMapInput
	// Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
	Tags pulumi.StringMapInput
	// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `templateUrl`.
	TemplateBody pulumi.StringPtrInput
	// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `templateBody`.
	TemplateUrl pulumi.StringPtrInput
}

func (StackSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackSetArgs)(nil)).Elem()
}

type StackSetInput interface {
	pulumi.Input

	ToStackSetOutput() StackSetOutput
	ToStackSetOutputWithContext(ctx context.Context) StackSetOutput
}

func (*StackSet) ElementType() reflect.Type {
	return reflect.TypeOf((*StackSet)(nil))
}

func (i *StackSet) ToStackSetOutput() StackSetOutput {
	return i.ToStackSetOutputWithContext(context.Background())
}

func (i *StackSet) ToStackSetOutputWithContext(ctx context.Context) StackSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSetOutput)
}

func (i *StackSet) ToStackSetPtrOutput() StackSetPtrOutput {
	return i.ToStackSetPtrOutputWithContext(context.Background())
}

func (i *StackSet) ToStackSetPtrOutputWithContext(ctx context.Context) StackSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSetPtrOutput)
}

type StackSetPtrInput interface {
	pulumi.Input

	ToStackSetPtrOutput() StackSetPtrOutput
	ToStackSetPtrOutputWithContext(ctx context.Context) StackSetPtrOutput
}

type stackSetPtrType StackSetArgs

func (*stackSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StackSet)(nil))
}

func (i *stackSetPtrType) ToStackSetPtrOutput() StackSetPtrOutput {
	return i.ToStackSetPtrOutputWithContext(context.Background())
}

func (i *stackSetPtrType) ToStackSetPtrOutputWithContext(ctx context.Context) StackSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSetOutput).ToStackSetPtrOutput()
}

// StackSetArrayInput is an input type that accepts StackSetArray and StackSetArrayOutput values.
// You can construct a concrete instance of `StackSetArrayInput` via:
//
//          StackSetArray{ StackSetArgs{...} }
type StackSetArrayInput interface {
	pulumi.Input

	ToStackSetArrayOutput() StackSetArrayOutput
	ToStackSetArrayOutputWithContext(context.Context) StackSetArrayOutput
}

type StackSetArray []StackSetInput

func (StackSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StackSet)(nil))
}

func (i StackSetArray) ToStackSetArrayOutput() StackSetArrayOutput {
	return i.ToStackSetArrayOutputWithContext(context.Background())
}

func (i StackSetArray) ToStackSetArrayOutputWithContext(ctx context.Context) StackSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSetArrayOutput)
}

// StackSetMapInput is an input type that accepts StackSetMap and StackSetMapOutput values.
// You can construct a concrete instance of `StackSetMapInput` via:
//
//          StackSetMap{ "key": StackSetArgs{...} }
type StackSetMapInput interface {
	pulumi.Input

	ToStackSetMapOutput() StackSetMapOutput
	ToStackSetMapOutputWithContext(context.Context) StackSetMapOutput
}

type StackSetMap map[string]StackSetInput

func (StackSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StackSet)(nil))
}

func (i StackSetMap) ToStackSetMapOutput() StackSetMapOutput {
	return i.ToStackSetMapOutputWithContext(context.Background())
}

func (i StackSetMap) ToStackSetMapOutputWithContext(ctx context.Context) StackSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSetMapOutput)
}

type StackSetOutput struct {
	*pulumi.OutputState
}

func (StackSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackSet)(nil))
}

func (o StackSetOutput) ToStackSetOutput() StackSetOutput {
	return o
}

func (o StackSetOutput) ToStackSetOutputWithContext(ctx context.Context) StackSetOutput {
	return o
}

func (o StackSetOutput) ToStackSetPtrOutput() StackSetPtrOutput {
	return o.ToStackSetPtrOutputWithContext(context.Background())
}

func (o StackSetOutput) ToStackSetPtrOutputWithContext(ctx context.Context) StackSetPtrOutput {
	return o.ApplyT(func(v StackSet) *StackSet {
		return &v
	}).(StackSetPtrOutput)
}

type StackSetPtrOutput struct {
	*pulumi.OutputState
}

func (StackSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackSet)(nil))
}

func (o StackSetPtrOutput) ToStackSetPtrOutput() StackSetPtrOutput {
	return o
}

func (o StackSetPtrOutput) ToStackSetPtrOutputWithContext(ctx context.Context) StackSetPtrOutput {
	return o
}

type StackSetArrayOutput struct{ *pulumi.OutputState }

func (StackSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StackSet)(nil))
}

func (o StackSetArrayOutput) ToStackSetArrayOutput() StackSetArrayOutput {
	return o
}

func (o StackSetArrayOutput) ToStackSetArrayOutputWithContext(ctx context.Context) StackSetArrayOutput {
	return o
}

func (o StackSetArrayOutput) Index(i pulumi.IntInput) StackSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StackSet {
		return vs[0].([]StackSet)[vs[1].(int)]
	}).(StackSetOutput)
}

type StackSetMapOutput struct{ *pulumi.OutputState }

func (StackSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StackSet)(nil))
}

func (o StackSetMapOutput) ToStackSetMapOutput() StackSetMapOutput {
	return o
}

func (o StackSetMapOutput) ToStackSetMapOutputWithContext(ctx context.Context) StackSetMapOutput {
	return o
}

func (o StackSetMapOutput) MapIndex(k pulumi.StringInput) StackSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StackSet {
		return vs[0].(map[string]StackSet)[vs[1].(string)]
	}).(StackSetOutput)
}

func init() {
	pulumi.RegisterOutputType(StackSetOutput{})
	pulumi.RegisterOutputType(StackSetPtrOutput{})
	pulumi.RegisterOutputType(StackSetArrayOutput{})
	pulumi.RegisterOutputType(StackSetMapOutput{})
}
