// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package worklink

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// WorkLink Website Certificate Authority can be imported using `FLEET-ARN,WEBSITE-CA-ID`, e.g.
//
// ```sh
//  $ pulumi import aws:worklink/websiteCertificateAuthorityAssociation:WebsiteCertificateAuthorityAssociation example arn:aws:worklink::123456789012:fleet/example,abcdefghijk
// ```
type WebsiteCertificateAuthorityAssociation struct {
	pulumi.CustomResourceState

	// The root certificate of the Certificate Authority.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The certificate name to display.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The ARN of the fleet.
	FleetArn pulumi.StringOutput `pulumi:"fleetArn"`
	// A unique identifier for the Certificate Authority.
	WebsiteCaId pulumi.StringOutput `pulumi:"websiteCaId"`
}

// NewWebsiteCertificateAuthorityAssociation registers a new resource with the given unique name, arguments, and options.
func NewWebsiteCertificateAuthorityAssociation(ctx *pulumi.Context,
	name string, args *WebsiteCertificateAuthorityAssociationArgs, opts ...pulumi.ResourceOption) (*WebsiteCertificateAuthorityAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	if args.FleetArn == nil {
		return nil, errors.New("invalid value for required argument 'FleetArn'")
	}
	var resource WebsiteCertificateAuthorityAssociation
	err := ctx.RegisterResource("aws:worklink/websiteCertificateAuthorityAssociation:WebsiteCertificateAuthorityAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebsiteCertificateAuthorityAssociation gets an existing WebsiteCertificateAuthorityAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebsiteCertificateAuthorityAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebsiteCertificateAuthorityAssociationState, opts ...pulumi.ResourceOption) (*WebsiteCertificateAuthorityAssociation, error) {
	var resource WebsiteCertificateAuthorityAssociation
	err := ctx.ReadResource("aws:worklink/websiteCertificateAuthorityAssociation:WebsiteCertificateAuthorityAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebsiteCertificateAuthorityAssociation resources.
type websiteCertificateAuthorityAssociationState struct {
	// The root certificate of the Certificate Authority.
	Certificate *string `pulumi:"certificate"`
	// The certificate name to display.
	DisplayName *string `pulumi:"displayName"`
	// The ARN of the fleet.
	FleetArn *string `pulumi:"fleetArn"`
	// A unique identifier for the Certificate Authority.
	WebsiteCaId *string `pulumi:"websiteCaId"`
}

type WebsiteCertificateAuthorityAssociationState struct {
	// The root certificate of the Certificate Authority.
	Certificate pulumi.StringPtrInput
	// The certificate name to display.
	DisplayName pulumi.StringPtrInput
	// The ARN of the fleet.
	FleetArn pulumi.StringPtrInput
	// A unique identifier for the Certificate Authority.
	WebsiteCaId pulumi.StringPtrInput
}

func (WebsiteCertificateAuthorityAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*websiteCertificateAuthorityAssociationState)(nil)).Elem()
}

type websiteCertificateAuthorityAssociationArgs struct {
	// The root certificate of the Certificate Authority.
	Certificate string `pulumi:"certificate"`
	// The certificate name to display.
	DisplayName *string `pulumi:"displayName"`
	// The ARN of the fleet.
	FleetArn string `pulumi:"fleetArn"`
}

// The set of arguments for constructing a WebsiteCertificateAuthorityAssociation resource.
type WebsiteCertificateAuthorityAssociationArgs struct {
	// The root certificate of the Certificate Authority.
	Certificate pulumi.StringInput
	// The certificate name to display.
	DisplayName pulumi.StringPtrInput
	// The ARN of the fleet.
	FleetArn pulumi.StringInput
}

func (WebsiteCertificateAuthorityAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*websiteCertificateAuthorityAssociationArgs)(nil)).Elem()
}

type WebsiteCertificateAuthorityAssociationInput interface {
	pulumi.Input

	ToWebsiteCertificateAuthorityAssociationOutput() WebsiteCertificateAuthorityAssociationOutput
	ToWebsiteCertificateAuthorityAssociationOutputWithContext(ctx context.Context) WebsiteCertificateAuthorityAssociationOutput
}

type WebsiteCertificateAuthorityAssociationPtrInput interface {
	pulumi.Input

	ToWebsiteCertificateAuthorityAssociationPtrOutput() WebsiteCertificateAuthorityAssociationPtrOutput
	ToWebsiteCertificateAuthorityAssociationPtrOutputWithContext(ctx context.Context) WebsiteCertificateAuthorityAssociationPtrOutput
}

func (WebsiteCertificateAuthorityAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*WebsiteCertificateAuthorityAssociation)(nil)).Elem()
}

func (i WebsiteCertificateAuthorityAssociation) ToWebsiteCertificateAuthorityAssociationOutput() WebsiteCertificateAuthorityAssociationOutput {
	return i.ToWebsiteCertificateAuthorityAssociationOutputWithContext(context.Background())
}

func (i WebsiteCertificateAuthorityAssociation) ToWebsiteCertificateAuthorityAssociationOutputWithContext(ctx context.Context) WebsiteCertificateAuthorityAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebsiteCertificateAuthorityAssociationOutput)
}

func (i WebsiteCertificateAuthorityAssociation) ToWebsiteCertificateAuthorityAssociationPtrOutput() WebsiteCertificateAuthorityAssociationPtrOutput {
	return i.ToWebsiteCertificateAuthorityAssociationPtrOutputWithContext(context.Background())
}

func (i WebsiteCertificateAuthorityAssociation) ToWebsiteCertificateAuthorityAssociationPtrOutputWithContext(ctx context.Context) WebsiteCertificateAuthorityAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebsiteCertificateAuthorityAssociationPtrOutput)
}

type WebsiteCertificateAuthorityAssociationOutput struct {
	*pulumi.OutputState
}

func (WebsiteCertificateAuthorityAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebsiteCertificateAuthorityAssociationOutput)(nil)).Elem()
}

func (o WebsiteCertificateAuthorityAssociationOutput) ToWebsiteCertificateAuthorityAssociationOutput() WebsiteCertificateAuthorityAssociationOutput {
	return o
}

func (o WebsiteCertificateAuthorityAssociationOutput) ToWebsiteCertificateAuthorityAssociationOutputWithContext(ctx context.Context) WebsiteCertificateAuthorityAssociationOutput {
	return o
}

type WebsiteCertificateAuthorityAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (WebsiteCertificateAuthorityAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebsiteCertificateAuthorityAssociation)(nil)).Elem()
}

func (o WebsiteCertificateAuthorityAssociationPtrOutput) ToWebsiteCertificateAuthorityAssociationPtrOutput() WebsiteCertificateAuthorityAssociationPtrOutput {
	return o
}

func (o WebsiteCertificateAuthorityAssociationPtrOutput) ToWebsiteCertificateAuthorityAssociationPtrOutputWithContext(ctx context.Context) WebsiteCertificateAuthorityAssociationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebsiteCertificateAuthorityAssociationOutput{})
	pulumi.RegisterOutputType(WebsiteCertificateAuthorityAssociationPtrOutput{})
}
