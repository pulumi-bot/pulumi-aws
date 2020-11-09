// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CloudFormation Stack resource.
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
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudformation.NewStack(ctx, "network", &cloudformation.StackArgs{
// 			Parameters: pulumi.StringMap{
// 				"VPCCidr": pulumi.String("10.0.0.0/16"),
// 			},
// 			TemplateBody: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Parameters\" : {\n", "    \"VPCCidr\" : {\n", "      \"Type\" : \"String\",\n", "      \"Default\" : \"10.0.0.0/16\",\n", "      \"Description\" : \"Enter the CIDR block for the VPC. Default is 10.0.0.0/16.\"\n", "    }\n", "  },\n", "  \"Resources\" : {\n", "    \"myVpc\": {\n", "      \"Type\" : \"AWS::EC2::VPC\",\n", "      \"Properties\" : {\n", "        \"CidrBlock\" : { \"Ref\" : \"VPCCidr\" },\n", "        \"Tags\" : [\n", "          {\"Key\": \"Name\", \"Value\": \"Primary_CF_VPC\"}\n", "        ]\n", "      }\n", "    }\n", "  }\n", "}\n", "\n")),
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
// Cloudformation Stacks can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:cloudformation/stack:Stack stack networking-stack
// ```
type Stack struct {
	pulumi.CustomResourceState

	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayOutput `pulumi:"capabilities"`
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback pulumi.BoolPtrOutput `pulumi:"disableRollback"`
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn pulumi.StringPtrOutput `pulumi:"iamRoleArn"`
	// Stack name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns pulumi.StringArrayOutput `pulumi:"notificationArns"`
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure pulumi.StringPtrOutput `pulumi:"onFailure"`
	// A map of outputs from the stack.
	Outputs pulumi.StringMapOutput `pulumi:"outputs"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody pulumi.StringOutput `pulumi:"policyBody"`
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl pulumi.StringPtrOutput `pulumi:"policyUrl"`
	// A list of tags to associate with this stack.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody pulumi.StringOutput `pulumi:"templateBody"`
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl pulumi.StringPtrOutput `pulumi:"templateUrl"`
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes pulumi.IntPtrOutput `pulumi:"timeoutInMinutes"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	if args == nil {
		args = &StackArgs{}
	}
	var resource Stack
	err := ctx.RegisterResource("aws:cloudformation/stack:Stack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackState, opts ...pulumi.ResourceOption) (*Stack, error) {
	var resource Stack
	err := ctx.ReadResource("aws:cloudformation/stack:Stack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stack resources.
type stackState struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities []string `pulumi:"capabilities"`
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback *bool `pulumi:"disableRollback"`
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// Stack name.
	Name *string `pulumi:"name"`
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns []string `pulumi:"notificationArns"`
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure *string `pulumi:"onFailure"`
	// A map of outputs from the stack.
	Outputs map[string]string `pulumi:"outputs"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]string `pulumi:"parameters"`
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody *string `pulumi:"policyBody"`
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl *string `pulumi:"policyUrl"`
	// A list of tags to associate with this stack.
	Tags map[string]string `pulumi:"tags"`
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody *string `pulumi:"templateBody"`
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl *string `pulumi:"templateUrl"`
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes *int `pulumi:"timeoutInMinutes"`
}

type StackState struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayInput
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback pulumi.BoolPtrInput
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn pulumi.StringPtrInput
	// Stack name.
	Name pulumi.StringPtrInput
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns pulumi.StringArrayInput
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure pulumi.StringPtrInput
	// A map of outputs from the stack.
	Outputs pulumi.StringMapInput
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.StringMapInput
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody pulumi.StringPtrInput
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl pulumi.StringPtrInput
	// A list of tags to associate with this stack.
	Tags pulumi.StringMapInput
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody pulumi.StringPtrInput
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl pulumi.StringPtrInput
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes pulumi.IntPtrInput
}

func (StackState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackState)(nil)).Elem()
}

type stackArgs struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities []string `pulumi:"capabilities"`
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback *bool `pulumi:"disableRollback"`
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// Stack name.
	Name *string `pulumi:"name"`
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns []string `pulumi:"notificationArns"`
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure *string `pulumi:"onFailure"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]string `pulumi:"parameters"`
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody *string `pulumi:"policyBody"`
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl *string `pulumi:"policyUrl"`
	// A list of tags to associate with this stack.
	Tags map[string]string `pulumi:"tags"`
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody *string `pulumi:"templateBody"`
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl *string `pulumi:"templateUrl"`
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes *int `pulumi:"timeoutInMinutes"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayInput
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback pulumi.BoolPtrInput
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn pulumi.StringPtrInput
	// Stack name.
	Name pulumi.StringPtrInput
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns pulumi.StringArrayInput
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure pulumi.StringPtrInput
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.StringMapInput
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody pulumi.StringPtrInput
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl pulumi.StringPtrInput
	// A list of tags to associate with this stack.
	Tags pulumi.StringMapInput
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody pulumi.StringPtrInput
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl pulumi.StringPtrInput
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes pulumi.IntPtrInput
}

func (StackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackArgs)(nil)).Elem()
}
