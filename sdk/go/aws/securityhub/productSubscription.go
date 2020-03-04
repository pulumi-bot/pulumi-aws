// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package securityhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Subscribes to a Security Hub product.
// 
// > **NOTE:** This AWS service is in Preview and may change before General Availability release. Backwards compatibility is not guaranteed between AWS Provider releases.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/securityhub_product_subscription.markdown.
type ProductSubscription struct {
	pulumi.CustomResourceState

	// The ARN of a resource that represents your subscription to the product that generates the findings that you want to import into Security Hub.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN of the product that generates findings that you want to import into Security Hub - see below.
	ProductArn pulumi.StringOutput `pulumi:"productArn"`
}

// NewProductSubscription registers a new resource with the given unique name, arguments, and options.
func NewProductSubscription(ctx *pulumi.Context,
	name string, args *ProductSubscriptionArgs, opts ...pulumi.ResourceOption) (*ProductSubscription, error) {
	if args == nil || args.ProductArn == nil {
		return nil, errors.New("missing required argument 'ProductArn'")
	}
	if args == nil {
		args = &ProductSubscriptionArgs{}
	}
	var resource ProductSubscription
	err := ctx.RegisterResource("aws:securityhub/productSubscription:ProductSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductSubscription gets an existing ProductSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductSubscriptionState, opts ...pulumi.ResourceOption) (*ProductSubscription, error) {
	var resource ProductSubscription
	err := ctx.ReadResource("aws:securityhub/productSubscription:ProductSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductSubscription resources.
type productSubscriptionState struct {
	// The ARN of a resource that represents your subscription to the product that generates the findings that you want to import into Security Hub.
	Arn *string `pulumi:"arn"`
	// The ARN of the product that generates findings that you want to import into Security Hub - see below.
	ProductArn *string `pulumi:"productArn"`
}

type ProductSubscriptionState struct {
	// The ARN of a resource that represents your subscription to the product that generates the findings that you want to import into Security Hub.
	Arn pulumi.StringPtrInput
	// The ARN of the product that generates findings that you want to import into Security Hub - see below.
	ProductArn pulumi.StringPtrInput
}

func (ProductSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*productSubscriptionState)(nil)).Elem()
}

type productSubscriptionArgs struct {
	// The ARN of the product that generates findings that you want to import into Security Hub - see below.
	ProductArn string `pulumi:"productArn"`
}

// The set of arguments for constructing a ProductSubscription resource.
type ProductSubscriptionArgs struct {
	// The ARN of the product that generates findings that you want to import into Security Hub - see below.
	ProductArn pulumi.StringInput
}

func (ProductSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productSubscriptionArgs)(nil)).Elem()
}

