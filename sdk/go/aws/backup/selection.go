// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package backup

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages selection conditions for AWS Backup plan resources.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/backup_selection.html.markdown.
type Selection struct {
	pulumi.CustomResourceState

	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn pulumi.StringOutput `pulumi:"iamRoleArn"`
	// The display name of a resource selection document.
	Name pulumi.StringOutput `pulumi:"name"`
	// The backup plan ID to be associated with the selection of resources.
	PlanId pulumi.StringOutput `pulumi:"planId"`
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources pulumi.StringArrayOutput `pulumi:"resources"`
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags SelectionSelectionTagArrayOutput `pulumi:"selectionTags"`
}

// NewSelection registers a new resource with the given unique name, arguments, and options.
func NewSelection(ctx *pulumi.Context,
	name string, args *SelectionArgs, opts ...pulumi.ResourceOption) (*Selection, error) {
	if args == nil || args.IamRoleArn == nil {
		return nil, errors.New("missing required argument 'IamRoleArn'")
	}
	if args == nil || args.PlanId == nil {
		return nil, errors.New("missing required argument 'PlanId'")
	}
	if args == nil {
		args = &SelectionArgs{}
	}
	var resource Selection
	err := ctx.RegisterResource("aws:backup/selection:Selection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelection gets an existing Selection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelectionState, opts ...pulumi.ResourceOption) (*Selection, error) {
	var resource Selection
	err := ctx.ReadResource("aws:backup/selection:Selection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Selection resources.
type selectionState struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// The display name of a resource selection document.
	Name *string `pulumi:"name"`
	// The backup plan ID to be associated with the selection of resources.
	PlanId *string `pulumi:"planId"`
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources []string `pulumi:"resources"`
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags []SelectionSelectionTag `pulumi:"selectionTags"`
}

type SelectionState struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn pulumi.StringPtrInput
	// The display name of a resource selection document.
	Name pulumi.StringPtrInput
	// The backup plan ID to be associated with the selection of resources.
	PlanId pulumi.StringPtrInput
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources pulumi.StringArrayInput
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags SelectionSelectionTagArrayInput
}

func (SelectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*selectionState)(nil)).Elem()
}

type selectionArgs struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn string `pulumi:"iamRoleArn"`
	// The display name of a resource selection document.
	Name *string `pulumi:"name"`
	// The backup plan ID to be associated with the selection of resources.
	PlanId string `pulumi:"planId"`
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources []string `pulumi:"resources"`
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags []SelectionSelectionTag `pulumi:"selectionTags"`
}

// The set of arguments for constructing a Selection resource.
type SelectionArgs struct {
	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn pulumi.StringInput
	// The display name of a resource selection document.
	Name pulumi.StringPtrInput
	// The backup plan ID to be associated with the selection of resources.
	PlanId pulumi.StringInput
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
	Resources pulumi.StringArrayInput
	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags SelectionSelectionTagArrayInput
}

func (SelectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selectionArgs)(nil)).Elem()
}

