// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a member account in the current organization.
//
// > **Note:** Account management must be done from the organization's master account.
//
// !> **WARNING:** Deleting this resource will only remove an AWS account from an organization. This provider will not close the account. The member account must be prepared to be a standalone account beforehand. See the [AWS Organizations documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_remove.html) for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = organizations.NewAccount(ctx, "account", &organizations.AccountArgs{
// 			Email: pulumi.String("john@doe.org"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Account struct {
	pulumi.CustomResourceState

	// The ARN for this account.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	Email pulumi.StringOutput `pulumi:"email"`
	// If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
	IamUserAccessToBilling pulumi.StringPtrOutput `pulumi:"iamUserAccessToBilling"`
	JoinedMethod           pulumi.StringOutput    `pulumi:"joinedMethod"`
	JoinedTimestamp        pulumi.StringOutput    `pulumi:"joinedTimestamp"`
	// A friendly name for the member account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so this provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) is used.
	RoleName pulumi.StringPtrOutput `pulumi:"roleName"`
	Status   pulumi.StringOutput    `pulumi:"status"`
	// Key-value mapping of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil || args.Email == nil {
		return nil, errors.New("missing required argument 'Email'")
	}
	if args == nil {
		args = &AccountArgs{}
	}
	var resource Account
	err := ctx.RegisterResource("aws:organizations/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("aws:organizations/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The ARN for this account.
	Arn *string `pulumi:"arn"`
	// The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	Email *string `pulumi:"email"`
	// If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
	IamUserAccessToBilling *string `pulumi:"iamUserAccessToBilling"`
	JoinedMethod           *string `pulumi:"joinedMethod"`
	JoinedTimestamp        *string `pulumi:"joinedTimestamp"`
	// A friendly name for the member account.
	Name *string `pulumi:"name"`
	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	ParentId *string `pulumi:"parentId"`
	// The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so this provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) is used.
	RoleName *string `pulumi:"roleName"`
	Status   *string `pulumi:"status"`
	// Key-value mapping of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

type AccountState struct {
	// The ARN for this account.
	Arn pulumi.StringPtrInput
	// The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	Email pulumi.StringPtrInput
	// If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
	IamUserAccessToBilling pulumi.StringPtrInput
	JoinedMethod           pulumi.StringPtrInput
	JoinedTimestamp        pulumi.StringPtrInput
	// A friendly name for the member account.
	Name pulumi.StringPtrInput
	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	ParentId pulumi.StringPtrInput
	// The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so this provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) is used.
	RoleName pulumi.StringPtrInput
	Status   pulumi.StringPtrInput
	// Key-value mapping of resource tags.
	Tags pulumi.StringMapInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	Email string `pulumi:"email"`
	// If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
	IamUserAccessToBilling *string `pulumi:"iamUserAccessToBilling"`
	// A friendly name for the member account.
	Name *string `pulumi:"name"`
	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	ParentId *string `pulumi:"parentId"`
	// The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so this provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) is used.
	RoleName *string `pulumi:"roleName"`
	// Key-value mapping of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	Email pulumi.StringInput
	// If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
	IamUserAccessToBilling pulumi.StringPtrInput
	// A friendly name for the member account.
	Name pulumi.StringPtrInput
	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	ParentId pulumi.StringPtrInput
	// The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so this provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) is used.
	RoleName pulumi.StringPtrInput
	// Key-value mapping of resource tags.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}
