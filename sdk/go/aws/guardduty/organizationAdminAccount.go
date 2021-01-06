// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a GuardDuty Organization Admin Account. The AWS account utilizing this resource must be an Organizations primary account. More information about Organizations support in GuardDuty can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/guardduty"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
// 			AwsServiceAccessPrincipals: pulumi.StringArray{
// 				pulumi.String("guardduty.amazonaws.com"),
// 			},
// 			FeatureSet: pulumi.String("ALL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = guardduty.NewDetector(ctx, "exampleDetector", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = guardduty.NewOrganizationAdminAccount(ctx, "exampleOrganizationAdminAccount", &guardduty.OrganizationAdminAccountArgs{
// 			AdminAccountId: pulumi.String("123456789012"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleOrganization,
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
// GuardDuty Organization Admin Account can be imported using the AWS account ID, e.g.
//
// ```sh
//  $ pulumi import aws:guardduty/organizationAdminAccount:OrganizationAdminAccount example 123456789012
// ```
type OrganizationAdminAccount struct {
	pulumi.CustomResourceState

	// AWS account identifier to designate as a delegated administrator for GuardDuty.
	AdminAccountId pulumi.StringOutput `pulumi:"adminAccountId"`
}

// NewOrganizationAdminAccount registers a new resource with the given unique name, arguments, and options.
func NewOrganizationAdminAccount(ctx *pulumi.Context,
	name string, args *OrganizationAdminAccountArgs, opts ...pulumi.ResourceOption) (*OrganizationAdminAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AdminAccountId'")
	}
	var resource OrganizationAdminAccount
	err := ctx.RegisterResource("aws:guardduty/organizationAdminAccount:OrganizationAdminAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationAdminAccount gets an existing OrganizationAdminAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationAdminAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationAdminAccountState, opts ...pulumi.ResourceOption) (*OrganizationAdminAccount, error) {
	var resource OrganizationAdminAccount
	err := ctx.ReadResource("aws:guardduty/organizationAdminAccount:OrganizationAdminAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationAdminAccount resources.
type organizationAdminAccountState struct {
	// AWS account identifier to designate as a delegated administrator for GuardDuty.
	AdminAccountId *string `pulumi:"adminAccountId"`
}

type OrganizationAdminAccountState struct {
	// AWS account identifier to designate as a delegated administrator for GuardDuty.
	AdminAccountId pulumi.StringPtrInput
}

func (OrganizationAdminAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationAdminAccountState)(nil)).Elem()
}

type organizationAdminAccountArgs struct {
	// AWS account identifier to designate as a delegated administrator for GuardDuty.
	AdminAccountId string `pulumi:"adminAccountId"`
}

// The set of arguments for constructing a OrganizationAdminAccount resource.
type OrganizationAdminAccountArgs struct {
	// AWS account identifier to designate as a delegated administrator for GuardDuty.
	AdminAccountId pulumi.StringInput
}

func (OrganizationAdminAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationAdminAccountArgs)(nil)).Elem()
}

type OrganizationAdminAccountInput interface {
	pulumi.Input

	ToOrganizationAdminAccountOutput() OrganizationAdminAccountOutput
	ToOrganizationAdminAccountOutputWithContext(ctx context.Context) OrganizationAdminAccountOutput
}

func (*OrganizationAdminAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationAdminAccount)(nil))
}

func (i *OrganizationAdminAccount) ToOrganizationAdminAccountOutput() OrganizationAdminAccountOutput {
	return i.ToOrganizationAdminAccountOutputWithContext(context.Background())
}

func (i *OrganizationAdminAccount) ToOrganizationAdminAccountOutputWithContext(ctx context.Context) OrganizationAdminAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountOutput)
}

func (i *OrganizationAdminAccount) ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput {
	return i.ToOrganizationAdminAccountPtrOutputWithContext(context.Background())
}

func (i *OrganizationAdminAccount) ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountPtrOutput)
}

type OrganizationAdminAccountPtrInput interface {
	pulumi.Input

	ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput
	ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput
}

type organizationAdminAccountPtrType OrganizationAdminAccountArgs

func (*organizationAdminAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationAdminAccount)(nil))
}

func (i *organizationAdminAccountPtrType) ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput {
	return i.ToOrganizationAdminAccountPtrOutputWithContext(context.Background())
}

func (i *organizationAdminAccountPtrType) ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountOutput).ToOrganizationAdminAccountPtrOutput()
}

// OrganizationAdminAccountArrayInput is an input type that accepts OrganizationAdminAccountArray and OrganizationAdminAccountArrayOutput values.
// You can construct a concrete instance of `OrganizationAdminAccountArrayInput` via:
//
//          OrganizationAdminAccountArray{ OrganizationAdminAccountArgs{...} }
type OrganizationAdminAccountArrayInput interface {
	pulumi.Input

	ToOrganizationAdminAccountArrayOutput() OrganizationAdminAccountArrayOutput
	ToOrganizationAdminAccountArrayOutputWithContext(context.Context) OrganizationAdminAccountArrayOutput
}

type OrganizationAdminAccountArray []OrganizationAdminAccountInput

func (OrganizationAdminAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrganizationAdminAccount)(nil))
}

func (i OrganizationAdminAccountArray) ToOrganizationAdminAccountArrayOutput() OrganizationAdminAccountArrayOutput {
	return i.ToOrganizationAdminAccountArrayOutputWithContext(context.Background())
}

func (i OrganizationAdminAccountArray) ToOrganizationAdminAccountArrayOutputWithContext(ctx context.Context) OrganizationAdminAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountArrayOutput)
}

// OrganizationAdminAccountMapInput is an input type that accepts OrganizationAdminAccountMap and OrganizationAdminAccountMapOutput values.
// You can construct a concrete instance of `OrganizationAdminAccountMapInput` via:
//
//          OrganizationAdminAccountMap{ "key": OrganizationAdminAccountArgs{...} }
type OrganizationAdminAccountMapInput interface {
	pulumi.Input

	ToOrganizationAdminAccountMapOutput() OrganizationAdminAccountMapOutput
	ToOrganizationAdminAccountMapOutputWithContext(context.Context) OrganizationAdminAccountMapOutput
}

type OrganizationAdminAccountMap map[string]OrganizationAdminAccountInput

func (OrganizationAdminAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OrganizationAdminAccount)(nil))
}

func (i OrganizationAdminAccountMap) ToOrganizationAdminAccountMapOutput() OrganizationAdminAccountMapOutput {
	return i.ToOrganizationAdminAccountMapOutputWithContext(context.Background())
}

func (i OrganizationAdminAccountMap) ToOrganizationAdminAccountMapOutputWithContext(ctx context.Context) OrganizationAdminAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationAdminAccountMapOutput)
}

type OrganizationAdminAccountOutput struct {
	*pulumi.OutputState
}

func (OrganizationAdminAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationAdminAccount)(nil))
}

func (o OrganizationAdminAccountOutput) ToOrganizationAdminAccountOutput() OrganizationAdminAccountOutput {
	return o
}

func (o OrganizationAdminAccountOutput) ToOrganizationAdminAccountOutputWithContext(ctx context.Context) OrganizationAdminAccountOutput {
	return o
}

func (o OrganizationAdminAccountOutput) ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput {
	return o.ToOrganizationAdminAccountPtrOutputWithContext(context.Background())
}

func (o OrganizationAdminAccountOutput) ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput {
	return o.ApplyT(func(v OrganizationAdminAccount) *OrganizationAdminAccount {
		return &v
	}).(OrganizationAdminAccountPtrOutput)
}

type OrganizationAdminAccountPtrOutput struct {
	*pulumi.OutputState
}

func (OrganizationAdminAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationAdminAccount)(nil))
}

func (o OrganizationAdminAccountPtrOutput) ToOrganizationAdminAccountPtrOutput() OrganizationAdminAccountPtrOutput {
	return o
}

func (o OrganizationAdminAccountPtrOutput) ToOrganizationAdminAccountPtrOutputWithContext(ctx context.Context) OrganizationAdminAccountPtrOutput {
	return o
}

type OrganizationAdminAccountArrayOutput struct{ *pulumi.OutputState }

func (OrganizationAdminAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrganizationAdminAccount)(nil))
}

func (o OrganizationAdminAccountArrayOutput) ToOrganizationAdminAccountArrayOutput() OrganizationAdminAccountArrayOutput {
	return o
}

func (o OrganizationAdminAccountArrayOutput) ToOrganizationAdminAccountArrayOutputWithContext(ctx context.Context) OrganizationAdminAccountArrayOutput {
	return o
}

func (o OrganizationAdminAccountArrayOutput) Index(i pulumi.IntInput) OrganizationAdminAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OrganizationAdminAccount {
		return vs[0].([]OrganizationAdminAccount)[vs[1].(int)]
	}).(OrganizationAdminAccountOutput)
}

type OrganizationAdminAccountMapOutput struct{ *pulumi.OutputState }

func (OrganizationAdminAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OrganizationAdminAccount)(nil))
}

func (o OrganizationAdminAccountMapOutput) ToOrganizationAdminAccountMapOutput() OrganizationAdminAccountMapOutput {
	return o
}

func (o OrganizationAdminAccountMapOutput) ToOrganizationAdminAccountMapOutputWithContext(ctx context.Context) OrganizationAdminAccountMapOutput {
	return o
}

func (o OrganizationAdminAccountMapOutput) MapIndex(k pulumi.StringInput) OrganizationAdminAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OrganizationAdminAccount {
		return vs[0].(map[string]OrganizationAdminAccount)[vs[1].(string)]
	}).(OrganizationAdminAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationAdminAccountOutput{})
	pulumi.RegisterOutputType(OrganizationAdminAccountPtrOutput{})
	pulumi.RegisterOutputType(OrganizationAdminAccountArrayOutput{})
	pulumi.RegisterOutputType(OrganizationAdminAccountMapOutput{})
}
