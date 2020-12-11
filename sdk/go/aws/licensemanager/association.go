// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a License Manager association.
//
// > **Note:** License configurations can also be associated with launch templates by specifying the `licenseSpecifications` block for an `ec2.LaunchTemplate`.
//
// ## Import
//
// License configurations can be imported in the form `resource_arn,license_configuration_arn`, e.g.
//
// ```sh
//  $ pulumi import aws:licensemanager/association:Association example arn:aws:ec2:eu-west-1:123456789012:image/ami-123456789abcdef01,arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef
// ```
type Association struct {
	pulumi.CustomResourceState

	// ARN of the license configuration.
	LicenseConfigurationArn pulumi.StringOutput `pulumi:"licenseConfigurationArn"`
	// ARN of the resource associated with the license configuration.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewAssociation registers a new resource with the given unique name, arguments, and options.
func NewAssociation(ctx *pulumi.Context,
	name string, args *AssociationArgs, opts ...pulumi.ResourceOption) (*Association, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LicenseConfigurationArn == nil {
		return nil, errors.New("invalid value for required argument 'LicenseConfigurationArn'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	var resource Association
	err := ctx.RegisterResource("aws:licensemanager/association:Association", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssociation gets an existing Association resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssociationState, opts ...pulumi.ResourceOption) (*Association, error) {
	var resource Association
	err := ctx.ReadResource("aws:licensemanager/association:Association", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Association resources.
type associationState struct {
	// ARN of the license configuration.
	LicenseConfigurationArn *string `pulumi:"licenseConfigurationArn"`
	// ARN of the resource associated with the license configuration.
	ResourceArn *string `pulumi:"resourceArn"`
}

type AssociationState struct {
	// ARN of the license configuration.
	LicenseConfigurationArn pulumi.StringPtrInput
	// ARN of the resource associated with the license configuration.
	ResourceArn pulumi.StringPtrInput
}

func (AssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*associationState)(nil)).Elem()
}

type associationArgs struct {
	// ARN of the license configuration.
	LicenseConfigurationArn string `pulumi:"licenseConfigurationArn"`
	// ARN of the resource associated with the license configuration.
	ResourceArn string `pulumi:"resourceArn"`
}

// The set of arguments for constructing a Association resource.
type AssociationArgs struct {
	// ARN of the license configuration.
	LicenseConfigurationArn pulumi.StringInput
	// ARN of the resource associated with the license configuration.
	ResourceArn pulumi.StringInput
}

func (AssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*associationArgs)(nil)).Elem()
}

type AssociationInput interface {
	pulumi.Input

	ToAssociationOutput() AssociationOutput
	ToAssociationOutputWithContext(ctx context.Context) AssociationOutput
}

type AssociationPtrInput interface {
	pulumi.Input

	ToAssociationPtrOutput() AssociationPtrOutput
	ToAssociationPtrOutputWithContext(ctx context.Context) AssociationPtrOutput
}

func (Association) ElementType() reflect.Type {
	return reflect.TypeOf((*Association)(nil)).Elem()
}

func (i Association) ToAssociationOutput() AssociationOutput {
	return i.ToAssociationOutputWithContext(context.Background())
}

func (i Association) ToAssociationOutputWithContext(ctx context.Context) AssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationOutput)
}

func (i Association) ToAssociationPtrOutput() AssociationPtrOutput {
	return i.ToAssociationPtrOutputWithContext(context.Background())
}

func (i Association) ToAssociationPtrOutputWithContext(ctx context.Context) AssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationPtrOutput)
}

type AssociationOutput struct {
	*pulumi.OutputState
}

func (AssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationOutput)(nil)).Elem()
}

func (o AssociationOutput) ToAssociationOutput() AssociationOutput {
	return o
}

func (o AssociationOutput) ToAssociationOutputWithContext(ctx context.Context) AssociationOutput {
	return o
}

type AssociationPtrOutput struct {
	*pulumi.OutputState
}

func (AssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Association)(nil)).Elem()
}

func (o AssociationPtrOutput) ToAssociationPtrOutput() AssociationPtrOutput {
	return o
}

func (o AssociationPtrOutput) ToAssociationPtrOutputWithContext(ctx context.Context) AssociationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AssociationOutput{})
	pulumi.RegisterOutputType(AssociationPtrOutput{})
}
