// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package adminAccount

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to associate/disassociate an AWS Firewall Manager administrator account. This operation must be performed in the `us-east-1` region.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/fms_admin_account.html.markdown.
type AdminAccount struct {
	pulumi.CustomResourceState

	// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
}

// NewAdminAccount registers a new resource with the given unique name, arguments, and options.
func NewAdminAccount(ctx *pulumi.Context,
	name string, args *AdminAccountArgs, opts ...pulumi.ResourceOption) (*AdminAccount, error) {
	if args == nil {
		args = &AdminAccountArgs{}
	}
	var resource AdminAccount
	err := ctx.RegisterResource("aws:fms/adminAccount:AdminAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdminAccount gets an existing AdminAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdminAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdminAccountState, opts ...pulumi.ResourceOption) (*AdminAccount, error) {
	var resource AdminAccount
	err := ctx.ReadResource("aws:fms/adminAccount:AdminAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdminAccount resources.
type adminAccountState struct {
	// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
	AccountId *string `pulumi:"accountId"`
}

type AdminAccountState struct {
	// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
	AccountId pulumi.StringPtrInput
}

func (AdminAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*adminAccountState)(nil)).Elem()
}

type adminAccountArgs struct {
	// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
	AccountId *string `pulumi:"accountId"`
}

// The set of arguments for constructing a AdminAccount resource.
type AdminAccountArgs struct {
	// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
	AccountId pulumi.StringPtrInput
}

func (AdminAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adminAccountArgs)(nil)).Elem()
}

