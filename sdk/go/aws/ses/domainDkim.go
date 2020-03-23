// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ses

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SES domain DKIM generation resource.
//
// Domain ownership needs to be confirmed first using [sesDomainIdentity Resource](https://www.terraform.io/docs/providers/aws/r/ses_domain_identity.html)
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_domain_dkim.html.markdown.
type DomainDkim struct {
	pulumi.CustomResourceState

	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.
	// Find out more about verifying domains in Amazon SES
	// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	DkimTokens pulumi.StringArrayOutput `pulumi:"dkimTokens"`
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringOutput `pulumi:"domain"`
}

// NewDomainDkim registers a new resource with the given unique name, arguments, and options.
func NewDomainDkim(ctx *pulumi.Context,
	name string, args *DomainDkimArgs, opts ...pulumi.ResourceOption) (*DomainDkim, error) {
	if args == nil || args.Domain == nil {
		return nil, errors.New("missing required argument 'Domain'")
	}
	if args == nil {
		args = &DomainDkimArgs{}
	}
	var resource DomainDkim
	err := ctx.RegisterResource("aws:ses/domainDkim:DomainDkim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainDkim gets an existing DomainDkim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainDkim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainDkimState, opts ...pulumi.ResourceOption) (*DomainDkim, error) {
	var resource DomainDkim
	err := ctx.ReadResource("aws:ses/domainDkim:DomainDkim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainDkim resources.
type domainDkimState struct {
	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.
	// Find out more about verifying domains in Amazon SES
	// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	DkimTokens []string `pulumi:"dkimTokens"`
	// Verified domain name to generate DKIM tokens for.
	Domain *string `pulumi:"domain"`
}

type DomainDkimState struct {
	// DKIM tokens generated by SES.
	// These tokens should be used to create CNAME records used to verify SES Easy DKIM.
	// See below for an example of how this might be achieved
	// when the domain is hosted in Route 53 and managed by this provider.
	// Find out more about verifying domains in Amazon SES
	// in the [AWS SES docs](http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	DkimTokens pulumi.StringArrayInput
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringPtrInput
}

func (DomainDkimState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainDkimState)(nil)).Elem()
}

type domainDkimArgs struct {
	// Verified domain name to generate DKIM tokens for.
	Domain string `pulumi:"domain"`
}

// The set of arguments for constructing a DomainDkim resource.
type DomainDkimArgs struct {
	// Verified domain name to generate DKIM tokens for.
	Domain pulumi.StringInput
}

func (DomainDkimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainDkimArgs)(nil)).Elem()
}

