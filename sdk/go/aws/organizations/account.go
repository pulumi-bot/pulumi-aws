// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a member account in the current organization.
// 
// > **Note:** Account management must be done from the organization's master account.
// 
// !> **WARNING:** Deleting this resource will only remove an AWS account from an organization. This provider will not close the account. The member account must be prepared to be a standalone account beforehand. See the [AWS Organizations documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_remove.html) for more information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/organizations_account.html.markdown.
type Account struct {
	s *pulumi.ResourceState
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOpt) (*Account, error) {
	if args == nil || args.Email == nil {
		return nil, errors.New("missing required argument 'Email'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["email"] = nil
		inputs["iamUserAccessToBilling"] = nil
		inputs["name"] = nil
		inputs["parentId"] = nil
		inputs["roleName"] = nil
		inputs["tags"] = nil
	} else {
		inputs["email"] = args.Email
		inputs["iamUserAccessToBilling"] = args.IamUserAccessToBilling
		inputs["name"] = args.Name
		inputs["parentId"] = args.ParentId
		inputs["roleName"] = args.RoleName
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	inputs["joinedMethod"] = nil
	inputs["joinedTimestamp"] = nil
	inputs["status"] = nil
	s, err := ctx.RegisterResource("aws:organizations/account:Account", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Account{s: s}, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AccountState, opts ...pulumi.ResourceOpt) (*Account, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["email"] = state.Email
		inputs["iamUserAccessToBilling"] = state.IamUserAccessToBilling
		inputs["joinedMethod"] = state.JoinedMethod
		inputs["joinedTimestamp"] = state.JoinedTimestamp
		inputs["name"] = state.Name
		inputs["parentId"] = state.ParentId
		inputs["roleName"] = state.RoleName
		inputs["status"] = state.Status
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:organizations/account:Account", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Account{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Account) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Account) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN for this account.
func (r *Account) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
func (r *Account) Email() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["email"])
}

// If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
func (r *Account) IamUserAccessToBilling() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["iamUserAccessToBilling"])
}

func (r *Account) JoinedMethod() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["joinedMethod"])
}

func (r *Account) JoinedTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["joinedTimestamp"])
}

// A friendly name for the member account.
func (r *Account) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
func (r *Account) ParentId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["parentId"])
}

// The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so this provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless [`ignoreChanges`](https://www.terraform.io/docs/configuration/resources.html#ignore_changes) is used.
func (r *Account) RoleName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["roleName"])
}

func (r *Account) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Key-value mapping of resource tags.
func (r *Account) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Account resources.
type AccountState struct {
	// The ARN for this account.
	Arn interface{}
	// The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	Email interface{}
	// If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
	IamUserAccessToBilling interface{}
	JoinedMethod interface{}
	JoinedTimestamp interface{}
	// A friendly name for the member account.
	Name interface{}
	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	ParentId interface{}
	// The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so this provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless [`ignoreChanges`](https://www.terraform.io/docs/configuration/resources.html#ignore_changes) is used.
	RoleName interface{}
	Status interface{}
	// Key-value mapping of resource tags.
	Tags interface{}
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
	Email interface{}
	// If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
	IamUserAccessToBilling interface{}
	// A friendly name for the member account.
	Name interface{}
	// Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
	ParentId interface{}
	// The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so this provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless [`ignoreChanges`](https://www.terraform.io/docs/configuration/resources.html#ignore_changes) is used.
	RoleName interface{}
	// Key-value mapping of resource tags.
	Tags interface{}
}
