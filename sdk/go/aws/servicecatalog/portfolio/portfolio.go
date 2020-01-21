// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package portfolio

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a Service Catalog Portfolio.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/servicecatalog_portfolio.html.markdown.
type Portfolio struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Description of the portfolio
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the portfolio.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the person or organization who owns the portfolio.
	ProviderName pulumi.StringPtrOutput `pulumi:"providerName"`
	// Tags to apply to the connection.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewPortfolio registers a new resource with the given unique name, arguments, and options.
func NewPortfolio(ctx *pulumi.Context,
	name string, args *PortfolioArgs, opts ...pulumi.ResourceOption) (*Portfolio, error) {
	if args == nil {
		args = &PortfolioArgs{}
	}
	var resource Portfolio
	err := ctx.RegisterResource("aws:servicecatalog/portfolio:Portfolio", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortfolio gets an existing Portfolio resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortfolio(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortfolioState, opts ...pulumi.ResourceOption) (*Portfolio, error) {
	var resource Portfolio
	err := ctx.ReadResource("aws:servicecatalog/portfolio:Portfolio", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Portfolio resources.
type portfolioState struct {
	Arn *string `pulumi:"arn"`
	CreatedTime *string `pulumi:"createdTime"`
	// Description of the portfolio
	Description *string `pulumi:"description"`
	// The name of the portfolio.
	Name *string `pulumi:"name"`
	// Name of the person or organization who owns the portfolio.
	ProviderName *string `pulumi:"providerName"`
	// Tags to apply to the connection.
	Tags map[string]interface{} `pulumi:"tags"`
}

type PortfolioState struct {
	Arn pulumi.StringPtrInput
	CreatedTime pulumi.StringPtrInput
	// Description of the portfolio
	Description pulumi.StringPtrInput
	// The name of the portfolio.
	Name pulumi.StringPtrInput
	// Name of the person or organization who owns the portfolio.
	ProviderName pulumi.StringPtrInput
	// Tags to apply to the connection.
	Tags pulumi.MapInput
}

func (PortfolioState) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioState)(nil)).Elem()
}

type portfolioArgs struct {
	// Description of the portfolio
	Description *string `pulumi:"description"`
	// The name of the portfolio.
	Name *string `pulumi:"name"`
	// Name of the person or organization who owns the portfolio.
	ProviderName *string `pulumi:"providerName"`
	// Tags to apply to the connection.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Portfolio resource.
type PortfolioArgs struct {
	// Description of the portfolio
	Description pulumi.StringPtrInput
	// The name of the portfolio.
	Name pulumi.StringPtrInput
	// Name of the person or organization who owns the portfolio.
	ProviderName pulumi.StringPtrInput
	// Tags to apply to the connection.
	Tags pulumi.MapInput
}

func (PortfolioArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioArgs)(nil)).Elem()
}

