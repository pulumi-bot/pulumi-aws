// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage a GuardDuty member.
// 
// > **NOTE:** Currently after using this resource, you must manually accept member account invitations before GuardDuty will begin sending cross-account events. More information for how to accomplish this via the AWS Console or API can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_accounts.html). Terraform implementation of the member acceptance resource can be tracked in [Github](https://github.com/terraform-providers/terraform-provider-aws/issues/2489).
type Member struct {
	s *pulumi.ResourceState
}

// NewMember registers a new resource with the given unique name, arguments, and options.
func NewMember(ctx *pulumi.Context,
	name string, args *MemberArgs, opts ...pulumi.ResourceOpt) (*Member, error) {
	if args == nil || args.AccountId == nil {
		return nil, errors.New("missing required argument 'AccountId'")
	}
	if args == nil || args.DetectorId == nil {
		return nil, errors.New("missing required argument 'DetectorId'")
	}
	if args == nil || args.Email == nil {
		return nil, errors.New("missing required argument 'Email'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accountId"] = nil
		inputs["detectorId"] = nil
		inputs["disableEmailNotification"] = nil
		inputs["email"] = nil
		inputs["invitationMessage"] = nil
		inputs["invite"] = nil
	} else {
		inputs["accountId"] = args.AccountId
		inputs["detectorId"] = args.DetectorId
		inputs["disableEmailNotification"] = args.DisableEmailNotification
		inputs["email"] = args.Email
		inputs["invitationMessage"] = args.InvitationMessage
		inputs["invite"] = args.Invite
	}
	inputs["relationshipStatus"] = nil
	s, err := ctx.RegisterResource("aws:guardduty/member:Member", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Member{s: s}, nil
}

// GetMember gets an existing Member resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MemberState, opts ...pulumi.ResourceOpt) (*Member, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accountId"] = state.AccountId
		inputs["detectorId"] = state.DetectorId
		inputs["disableEmailNotification"] = state.DisableEmailNotification
		inputs["email"] = state.Email
		inputs["invitationMessage"] = state.InvitationMessage
		inputs["invite"] = state.Invite
		inputs["relationshipStatus"] = state.RelationshipStatus
	}
	s, err := ctx.ReadResource("aws:guardduty/member:Member", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Member{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Member) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Member) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// AWS account ID for member account.
func (r *Member) AccountId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accountId"])
}

// The detector ID of the GuardDuty account where you want to create member accounts.
func (r *Member) DetectorId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["detectorId"])
}

// Boolean whether an email notification is sent to the accounts. Defaults to `false`.
func (r *Member) DisableEmailNotification() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disableEmailNotification"])
}

// Email address for member account.
func (r *Member) Email() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["email"])
}

// Message for invitation.
func (r *Member) InvitationMessage() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["invitationMessage"])
}

// Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the Terraform state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
func (r *Member) Invite() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["invite"])
}

// The status of the relationship between the member account and its master account. More information can be found in [Amazon GuardDuty API Reference](https://docs.aws.amazon.com/guardduty/latest/ug/get-members.html).
func (r *Member) RelationshipStatus() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["relationshipStatus"])
}

// Input properties used for looking up and filtering Member resources.
type MemberState struct {
	// AWS account ID for member account.
	AccountId interface{}
	// The detector ID of the GuardDuty account where you want to create member accounts.
	DetectorId interface{}
	// Boolean whether an email notification is sent to the accounts. Defaults to `false`.
	DisableEmailNotification interface{}
	// Email address for member account.
	Email interface{}
	// Message for invitation.
	InvitationMessage interface{}
	// Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the Terraform state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
	Invite interface{}
	// The status of the relationship between the member account and its master account. More information can be found in [Amazon GuardDuty API Reference](https://docs.aws.amazon.com/guardduty/latest/ug/get-members.html).
	RelationshipStatus interface{}
}

// The set of arguments for constructing a Member resource.
type MemberArgs struct {
	// AWS account ID for member account.
	AccountId interface{}
	// The detector ID of the GuardDuty account where you want to create member accounts.
	DetectorId interface{}
	// Boolean whether an email notification is sent to the accounts. Defaults to `false`.
	DisableEmailNotification interface{}
	// Email address for member account.
	Email interface{}
	// Message for invitation.
	InvitationMessage interface{}
	// Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the Terraform state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
	Invite interface{}
}
