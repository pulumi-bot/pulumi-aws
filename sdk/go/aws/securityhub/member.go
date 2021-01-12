// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Security Hub member resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/securityhub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewMember(ctx, "exampleMember", &securityhub.MemberArgs{
// 			AccountId: pulumi.String("123456789012"),
// 			Email:     pulumi.String("example@example.com"),
// 			Invite:    pulumi.Bool(true),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAccount,
// 		}))
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
// Security Hub members can be imported using their account ID, e.g.
//
// ```sh
//  $ pulumi import aws:securityhub/member:Member example 123456789012
// ```
type Member struct {
	pulumi.CustomResourceState

	// The ID of the member AWS account.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The email of the member AWS account.
	Email pulumi.StringOutput `pulumi:"email"`
	// Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
	Invite pulumi.BoolPtrOutput `pulumi:"invite"`
	// The ID of the master Security Hub AWS account.
	MasterId pulumi.StringOutput `pulumi:"masterId"`
	// The status of the member account relationship.
	MemberStatus pulumi.StringOutput `pulumi:"memberStatus"`
}

// NewMember registers a new resource with the given unique name, arguments, and options.
func NewMember(ctx *pulumi.Context,
	name string, args *MemberArgs, opts ...pulumi.ResourceOption) (*Member, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	var resource Member
	err := ctx.RegisterResource("aws:securityhub/member:Member", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMember gets an existing Member resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MemberState, opts ...pulumi.ResourceOption) (*Member, error) {
	var resource Member
	err := ctx.ReadResource("aws:securityhub/member:Member", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Member resources.
type memberState struct {
	// The ID of the member AWS account.
	AccountId *string `pulumi:"accountId"`
	// The email of the member AWS account.
	Email *string `pulumi:"email"`
	// Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
	Invite *bool `pulumi:"invite"`
	// The ID of the master Security Hub AWS account.
	MasterId *string `pulumi:"masterId"`
	// The status of the member account relationship.
	MemberStatus *string `pulumi:"memberStatus"`
}

type MemberState struct {
	// The ID of the member AWS account.
	AccountId pulumi.StringPtrInput
	// The email of the member AWS account.
	Email pulumi.StringPtrInput
	// Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
	Invite pulumi.BoolPtrInput
	// The ID of the master Security Hub AWS account.
	MasterId pulumi.StringPtrInput
	// The status of the member account relationship.
	MemberStatus pulumi.StringPtrInput
}

func (MemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*memberState)(nil)).Elem()
}

type memberArgs struct {
	// The ID of the member AWS account.
	AccountId string `pulumi:"accountId"`
	// The email of the member AWS account.
	Email string `pulumi:"email"`
	// Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
	Invite *bool `pulumi:"invite"`
}

// The set of arguments for constructing a Member resource.
type MemberArgs struct {
	// The ID of the member AWS account.
	AccountId pulumi.StringInput
	// The email of the member AWS account.
	Email pulumi.StringInput
	// Boolean whether to invite the account to Security Hub as a member. Defaults to `false`.
	Invite pulumi.BoolPtrInput
}

func (MemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*memberArgs)(nil)).Elem()
}

type MemberInput interface {
	pulumi.Input

	ToMemberOutput() MemberOutput
	ToMemberOutputWithContext(ctx context.Context) MemberOutput
}

func (*Member) ElementType() reflect.Type {
	return reflect.TypeOf((*Member)(nil))
}

func (i *Member) ToMemberOutput() MemberOutput {
	return i.ToMemberOutputWithContext(context.Background())
}

func (i *Member) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberOutput)
}

func (i *Member) ToMemberPtrOutput() MemberPtrOutput {
	return i.ToMemberPtrOutputWithContext(context.Background())
}

func (i *Member) ToMemberPtrOutputWithContext(ctx context.Context) MemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberPtrOutput)
}

type MemberPtrInput interface {
	pulumi.Input

	ToMemberPtrOutput() MemberPtrOutput
	ToMemberPtrOutputWithContext(ctx context.Context) MemberPtrOutput
}

type memberPtrType MemberArgs

func (*memberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil))
}

func (i *memberPtrType) ToMemberPtrOutput() MemberPtrOutput {
	return i.ToMemberPtrOutputWithContext(context.Background())
}

func (i *memberPtrType) ToMemberPtrOutputWithContext(ctx context.Context) MemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberOutput).ToMemberPtrOutput()
}

// MemberArrayInput is an input type that accepts MemberArray and MemberArrayOutput values.
// You can construct a concrete instance of `MemberArrayInput` via:
//
//          MemberArray{ MemberArgs{...} }
type MemberArrayInput interface {
	pulumi.Input

	ToMemberArrayOutput() MemberArrayOutput
	ToMemberArrayOutputWithContext(context.Context) MemberArrayOutput
}

type MemberArray []MemberInput

func (MemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Member)(nil))
}

func (i MemberArray) ToMemberArrayOutput() MemberArrayOutput {
	return i.ToMemberArrayOutputWithContext(context.Background())
}

func (i MemberArray) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberArrayOutput)
}

// MemberMapInput is an input type that accepts MemberMap and MemberMapOutput values.
// You can construct a concrete instance of `MemberMapInput` via:
//
//          MemberMap{ "key": MemberArgs{...} }
type MemberMapInput interface {
	pulumi.Input

	ToMemberMapOutput() MemberMapOutput
	ToMemberMapOutputWithContext(context.Context) MemberMapOutput
}

type MemberMap map[string]MemberInput

func (MemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Member)(nil))
}

func (i MemberMap) ToMemberMapOutput() MemberMapOutput {
	return i.ToMemberMapOutputWithContext(context.Background())
}

func (i MemberMap) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberMapOutput)
}

type MemberOutput struct {
	*pulumi.OutputState
}

func (MemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Member)(nil))
}

func (o MemberOutput) ToMemberOutput() MemberOutput {
	return o
}

func (o MemberOutput) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return o
}

func (o MemberOutput) ToMemberPtrOutput() MemberPtrOutput {
	return o.ToMemberPtrOutputWithContext(context.Background())
}

func (o MemberOutput) ToMemberPtrOutputWithContext(ctx context.Context) MemberPtrOutput {
	return o.ApplyT(func(v Member) *Member {
		return &v
	}).(MemberPtrOutput)
}

type MemberPtrOutput struct {
	*pulumi.OutputState
}

func (MemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil))
}

func (o MemberPtrOutput) ToMemberPtrOutput() MemberPtrOutput {
	return o
}

func (o MemberPtrOutput) ToMemberPtrOutputWithContext(ctx context.Context) MemberPtrOutput {
	return o
}

type MemberArrayOutput struct{ *pulumi.OutputState }

func (MemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Member)(nil))
}

func (o MemberArrayOutput) ToMemberArrayOutput() MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) Index(i pulumi.IntInput) MemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Member {
		return vs[0].([]Member)[vs[1].(int)]
	}).(MemberOutput)
}

type MemberMapOutput struct{ *pulumi.OutputState }

func (MemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Member)(nil))
}

func (o MemberMapOutput) ToMemberMapOutput() MemberMapOutput {
	return o
}

func (o MemberMapOutput) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return o
}

func (o MemberMapOutput) MapIndex(k pulumi.StringInput) MemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Member {
		return vs[0].(map[string]Member)[vs[1].(string)]
	}).(MemberOutput)
}

func init() {
	pulumi.RegisterOutputType(MemberOutput{})
	pulumi.RegisterOutputType(MemberPtrOutput{})
	pulumi.RegisterOutputType(MemberArrayOutput{})
	pulumi.RegisterOutputType(MemberMapOutput{})
}
