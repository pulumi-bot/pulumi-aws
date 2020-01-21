// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package originAccessIdentity

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Amazon CloudFront origin access identity.
// 
// For information about CloudFront distributions, see the
// [Amazon CloudFront Developer Guide][1]. For more information on generating
// origin access identities, see
// [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content][2].
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudfront_origin_access_identity.html.markdown.
type OriginAccessIdentity struct {
	pulumi.CustomResourceState

	// Internal value used by CloudFront to allow future
	// updates to the origin access identity.
	CallerReference pulumi.StringOutput `pulumi:"callerReference"`
	// A shortcut to the full path for the
	// origin access identity to use in CloudFront, see below.
	CloudfrontAccessIdentityPath pulumi.StringOutput `pulumi:"cloudfrontAccessIdentityPath"`
	// An optional comment for the origin access identity.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The current version of the origin access identity's information.
	// For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A pre-generated ARN for use in S3 bucket policies (see below).
	// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
	// E2QWRUHAPOMQZL`.
	IamArn pulumi.StringOutput `pulumi:"iamArn"`
	// The Amazon S3 canonical user ID for the origin
	// access identity, which you use when giving the origin access identity read
	// permission to an object in Amazon S3.
	S3CanonicalUserId pulumi.StringOutput `pulumi:"s3CanonicalUserId"`
}

// NewOriginAccessIdentity registers a new resource with the given unique name, arguments, and options.
func NewOriginAccessIdentity(ctx *pulumi.Context,
	name string, args *OriginAccessIdentityArgs, opts ...pulumi.ResourceOption) (*OriginAccessIdentity, error) {
	if args == nil {
		args = &OriginAccessIdentityArgs{}
	}
	var resource OriginAccessIdentity
	err := ctx.RegisterResource("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOriginAccessIdentity gets an existing OriginAccessIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOriginAccessIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginAccessIdentityState, opts ...pulumi.ResourceOption) (*OriginAccessIdentity, error) {
	var resource OriginAccessIdentity
	err := ctx.ReadResource("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OriginAccessIdentity resources.
type originAccessIdentityState struct {
	// Internal value used by CloudFront to allow future
	// updates to the origin access identity.
	CallerReference *string `pulumi:"callerReference"`
	// A shortcut to the full path for the
	// origin access identity to use in CloudFront, see below.
	CloudfrontAccessIdentityPath *string `pulumi:"cloudfrontAccessIdentityPath"`
	// An optional comment for the origin access identity.
	Comment *string `pulumi:"comment"`
	// The current version of the origin access identity's information.
	// For example: `E2QWRUHAPOMQZL`.
	Etag *string `pulumi:"etag"`
	// A pre-generated ARN for use in S3 bucket policies (see below).
	// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
	// E2QWRUHAPOMQZL`.
	IamArn *string `pulumi:"iamArn"`
	// The Amazon S3 canonical user ID for the origin
	// access identity, which you use when giving the origin access identity read
	// permission to an object in Amazon S3.
	S3CanonicalUserId *string `pulumi:"s3CanonicalUserId"`
}

type OriginAccessIdentityState struct {
	// Internal value used by CloudFront to allow future
	// updates to the origin access identity.
	CallerReference pulumi.StringPtrInput
	// A shortcut to the full path for the
	// origin access identity to use in CloudFront, see below.
	CloudfrontAccessIdentityPath pulumi.StringPtrInput
	// An optional comment for the origin access identity.
	Comment pulumi.StringPtrInput
	// The current version of the origin access identity's information.
	// For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringPtrInput
	// A pre-generated ARN for use in S3 bucket policies (see below).
	// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
	// E2QWRUHAPOMQZL`.
	IamArn pulumi.StringPtrInput
	// The Amazon S3 canonical user ID for the origin
	// access identity, which you use when giving the origin access identity read
	// permission to an object in Amazon S3.
	S3CanonicalUserId pulumi.StringPtrInput
}

func (OriginAccessIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*originAccessIdentityState)(nil)).Elem()
}

type originAccessIdentityArgs struct {
	// An optional comment for the origin access identity.
	Comment *string `pulumi:"comment"`
}

// The set of arguments for constructing a OriginAccessIdentity resource.
type OriginAccessIdentityArgs struct {
	// An optional comment for the origin access identity.
	Comment pulumi.StringPtrInput
}

func (OriginAccessIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originAccessIdentityArgs)(nil)).Elem()
}

