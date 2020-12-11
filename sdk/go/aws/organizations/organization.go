// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// The AWS organization can be imported by using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:organizations/organization:Organization my_org o-1234567
// ```
type Organization struct {
	pulumi.CustomResourceState

	// List of organization accounts including the master account. For a list excluding the master account, see the `nonMasterAccounts` attribute. All elements have these attributes:
	Accounts OrganizationAccountArrayOutput `pulumi:"accounts"`
	// ARN of the root
	Arn pulumi.StringOutput `pulumi:"arn"`
	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `featureSet` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
	AwsServiceAccessPrincipals pulumi.StringArrayOutput `pulumi:"awsServiceAccessPrincipals"`
	// List of Organizations policy types to enable in the Organization Root. Organization must have `featureSet` set to `ALL`. For additional information about valid policy types (e.g. `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY`, and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
	EnabledPolicyTypes pulumi.StringArrayOutput `pulumi:"enabledPolicyTypes"`
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet pulumi.StringPtrOutput `pulumi:"featureSet"`
	// ARN of the master account
	MasterAccountArn pulumi.StringOutput `pulumi:"masterAccountArn"`
	// Email address of the master account
	MasterAccountEmail pulumi.StringOutput `pulumi:"masterAccountEmail"`
	// Identifier of the master account
	MasterAccountId pulumi.StringOutput `pulumi:"masterAccountId"`
	// List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
	NonMasterAccounts OrganizationNonMasterAccountArrayOutput `pulumi:"nonMasterAccounts"`
	// List of organization roots. All elements have these attributes:
	Roots OrganizationRootArrayOutput `pulumi:"roots"`
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		args = &OrganizationArgs{}
	}

	var resource Organization
	err := ctx.RegisterResource("aws:organizations/organization:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("aws:organizations/organization:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type organizationState struct {
	// List of organization accounts including the master account. For a list excluding the master account, see the `nonMasterAccounts` attribute. All elements have these attributes:
	Accounts []OrganizationAccount `pulumi:"accounts"`
	// ARN of the root
	Arn *string `pulumi:"arn"`
	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `featureSet` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
	AwsServiceAccessPrincipals []string `pulumi:"awsServiceAccessPrincipals"`
	// List of Organizations policy types to enable in the Organization Root. Organization must have `featureSet` set to `ALL`. For additional information about valid policy types (e.g. `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY`, and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
	EnabledPolicyTypes []string `pulumi:"enabledPolicyTypes"`
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet *string `pulumi:"featureSet"`
	// ARN of the master account
	MasterAccountArn *string `pulumi:"masterAccountArn"`
	// Email address of the master account
	MasterAccountEmail *string `pulumi:"masterAccountEmail"`
	// Identifier of the master account
	MasterAccountId *string `pulumi:"masterAccountId"`
	// List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
	NonMasterAccounts []OrganizationNonMasterAccount `pulumi:"nonMasterAccounts"`
	// List of organization roots. All elements have these attributes:
	Roots []OrganizationRoot `pulumi:"roots"`
}

type OrganizationState struct {
	// List of organization accounts including the master account. For a list excluding the master account, see the `nonMasterAccounts` attribute. All elements have these attributes:
	Accounts OrganizationAccountArrayInput
	// ARN of the root
	Arn pulumi.StringPtrInput
	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `featureSet` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
	AwsServiceAccessPrincipals pulumi.StringArrayInput
	// List of Organizations policy types to enable in the Organization Root. Organization must have `featureSet` set to `ALL`. For additional information about valid policy types (e.g. `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY`, and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
	EnabledPolicyTypes pulumi.StringArrayInput
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet pulumi.StringPtrInput
	// ARN of the master account
	MasterAccountArn pulumi.StringPtrInput
	// Email address of the master account
	MasterAccountEmail pulumi.StringPtrInput
	// Identifier of the master account
	MasterAccountId pulumi.StringPtrInput
	// List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
	NonMasterAccounts OrganizationNonMasterAccountArrayInput
	// List of organization roots. All elements have these attributes:
	Roots OrganizationRootArrayInput
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `featureSet` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
	AwsServiceAccessPrincipals []string `pulumi:"awsServiceAccessPrincipals"`
	// List of Organizations policy types to enable in the Organization Root. Organization must have `featureSet` set to `ALL`. For additional information about valid policy types (e.g. `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY`, and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
	EnabledPolicyTypes []string `pulumi:"enabledPolicyTypes"`
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet *string `pulumi:"featureSet"`
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `featureSet` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
	AwsServiceAccessPrincipals pulumi.StringArrayInput
	// List of Organizations policy types to enable in the Organization Root. Organization must have `featureSet` set to `ALL`. For additional information about valid policy types (e.g. `BACKUP_POLICY`, `SERVICE_CONTROL_POLICY`, and `TAG_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
	EnabledPolicyTypes pulumi.StringArrayInput
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet pulumi.StringPtrInput
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationInput interface {
	pulumi.Input

	ToOrganizationOutput() OrganizationOutput
	ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput
}

type OrganizationPtrInput interface {
	pulumi.Input

	ToOrganizationPtrOutput() OrganizationPtrOutput
	ToOrganizationPtrOutputWithContext(ctx context.Context) OrganizationPtrOutput
}

func (Organization) ElementType() reflect.Type {
	return reflect.TypeOf((*Organization)(nil)).Elem()
}

func (i Organization) ToOrganizationOutput() OrganizationOutput {
	return i.ToOrganizationOutputWithContext(context.Background())
}

func (i Organization) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationOutput)
}

func (i Organization) ToOrganizationPtrOutput() OrganizationPtrOutput {
	return i.ToOrganizationPtrOutputWithContext(context.Background())
}

func (i Organization) ToOrganizationPtrOutputWithContext(ctx context.Context) OrganizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationPtrOutput)
}

type OrganizationOutput struct {
	*pulumi.OutputState
}

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationOutput)(nil)).Elem()
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

type OrganizationPtrOutput struct {
	*pulumi.OutputState
}

func (OrganizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (o OrganizationPtrOutput) ToOrganizationPtrOutput() OrganizationPtrOutput {
	return o
}

func (o OrganizationPtrOutput) ToOrganizationPtrOutputWithContext(ctx context.Context) OrganizationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationOutput{})
	pulumi.RegisterOutputType(OrganizationPtrOutput{})
}
