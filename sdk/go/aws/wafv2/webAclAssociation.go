// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a WAFv2 Web ACL Association.
//
// > **NOTE on associating a WAFv2 Web ACL with a Cloudfront distribution:** Do not use this resource to associate a WAFv2 Web ACL with a Cloudfront Distribution. The [AWS API call backing this resource](https://docs.aws.amazon.com/waf/latest/APIReference/API_AssociateWebACL.html) notes that you should use the [`webAclId`][2] property on the [`cloudfrontDistribution`][2] instead.
//
// [1]: https://docs.aws.amazon.com/waf/latest/APIReference/API_AssociateWebACL.html
// [2]: https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#web_acl_id
type WebAclAssociation struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn pulumi.StringOutput `pulumi:"webAclArn"`
}

// NewWebAclAssociation registers a new resource with the given unique name, arguments, and options.
func NewWebAclAssociation(ctx *pulumi.Context,
	name string, args *WebAclAssociationArgs, opts ...pulumi.ResourceOption) (*WebAclAssociation, error) {
	if args == nil || args.ResourceArn == nil {
		return nil, errors.New("missing required argument 'ResourceArn'")
	}
	if args == nil || args.WebAclArn == nil {
		return nil, errors.New("missing required argument 'WebAclArn'")
	}
	if args == nil {
		args = &WebAclAssociationArgs{}
	}
	var resource WebAclAssociation
	err := ctx.RegisterResource("aws:wafv2/webAclAssociation:WebAclAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAclAssociation gets an existing WebAclAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAclAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclAssociationState, opts ...pulumi.ResourceOption) (*WebAclAssociation, error) {
	var resource WebAclAssociation
	err := ctx.ReadResource("aws:wafv2/webAclAssociation:WebAclAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAclAssociation resources.
type webAclAssociationState struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
	ResourceArn *string `pulumi:"resourceArn"`
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn *string `pulumi:"webAclArn"`
}

type WebAclAssociationState struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
	ResourceArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn pulumi.StringPtrInput
}

func (WebAclAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclAssociationState)(nil)).Elem()
}

type webAclAssociationArgs struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
	ResourceArn string `pulumi:"resourceArn"`
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn string `pulumi:"webAclArn"`
}

// The set of arguments for constructing a WebAclAssociation resource.
type WebAclAssociationArgs struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer or an Amazon API Gateway stage.
	ResourceArn pulumi.StringInput
	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn pulumi.StringInput
}

func (WebAclAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclAssociationArgs)(nil)).Elem()
}
