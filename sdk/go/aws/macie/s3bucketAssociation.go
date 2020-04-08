// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Associates an S3 resource with Amazon Macie for monitoring and data classification.
//
// > **NOTE:** Before using Amazon Macie for the first time it must be enabled manually. Instructions are [here](https://docs.aws.amazon.com/macie/latest/userguide/macie-setting-up.html#macie-setting-up-enable).
type S3BucketAssociation struct {
	pulumi.CustomResourceState

	// The name of the S3 bucket that you want to associate with Amazon Macie.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// The configuration of how Amazon Macie classifies the S3 objects.
	ClassificationType S3BucketAssociationClassificationTypeOutput `pulumi:"classificationType"`
	// The ID of the Amazon Macie member account whose S3 resources you want to associate with Macie. If `memberAccountId` isn't specified, the action associates specified S3 resources with Macie for the current master account.
	MemberAccountId pulumi.StringPtrOutput `pulumi:"memberAccountId"`
	// Object key prefix identifying one or more S3 objects to which the association applies.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
}

// NewS3BucketAssociation registers a new resource with the given unique name, arguments, and options.
func NewS3BucketAssociation(ctx *pulumi.Context,
	name string, args *S3BucketAssociationArgs, opts ...pulumi.ResourceOption) (*S3BucketAssociation, error) {
	if args == nil || args.BucketName == nil {
		return nil, errors.New("missing required argument 'BucketName'")
	}
	if args == nil {
		args = &S3BucketAssociationArgs{}
	}
	var resource S3BucketAssociation
	err := ctx.RegisterResource("aws:macie/s3BucketAssociation:S3BucketAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetS3BucketAssociation gets an existing S3BucketAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetS3BucketAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *S3BucketAssociationState, opts ...pulumi.ResourceOption) (*S3BucketAssociation, error) {
	var resource S3BucketAssociation
	err := ctx.ReadResource("aws:macie/s3BucketAssociation:S3BucketAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering S3BucketAssociation resources.
type s3bucketAssociationState struct {
	// The name of the S3 bucket that you want to associate with Amazon Macie.
	BucketName *string `pulumi:"bucketName"`
	// The configuration of how Amazon Macie classifies the S3 objects.
	ClassificationType *S3BucketAssociationClassificationType `pulumi:"classificationType"`
	// The ID of the Amazon Macie member account whose S3 resources you want to associate with Macie. If `memberAccountId` isn't specified, the action associates specified S3 resources with Macie for the current master account.
	MemberAccountId *string `pulumi:"memberAccountId"`
	// Object key prefix identifying one or more S3 objects to which the association applies.
	Prefix *string `pulumi:"prefix"`
}

type S3BucketAssociationState struct {
	// The name of the S3 bucket that you want to associate with Amazon Macie.
	BucketName pulumi.StringPtrInput
	// The configuration of how Amazon Macie classifies the S3 objects.
	ClassificationType S3BucketAssociationClassificationTypePtrInput
	// The ID of the Amazon Macie member account whose S3 resources you want to associate with Macie. If `memberAccountId` isn't specified, the action associates specified S3 resources with Macie for the current master account.
	MemberAccountId pulumi.StringPtrInput
	// Object key prefix identifying one or more S3 objects to which the association applies.
	Prefix pulumi.StringPtrInput
}

func (S3BucketAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*s3bucketAssociationState)(nil)).Elem()
}

type s3bucketAssociationArgs struct {
	// The name of the S3 bucket that you want to associate with Amazon Macie.
	BucketName string `pulumi:"bucketName"`
	// The configuration of how Amazon Macie classifies the S3 objects.
	ClassificationType *S3BucketAssociationClassificationTypeArgs `pulumi:"classificationType"`
	// The ID of the Amazon Macie member account whose S3 resources you want to associate with Macie. If `memberAccountId` isn't specified, the action associates specified S3 resources with Macie for the current master account.
	MemberAccountId *string `pulumi:"memberAccountId"`
	// Object key prefix identifying one or more S3 objects to which the association applies.
	Prefix *string `pulumi:"prefix"`
}

// The set of arguments for constructing a S3BucketAssociation resource.
type S3BucketAssociationArgs struct {
	// The name of the S3 bucket that you want to associate with Amazon Macie.
	BucketName pulumi.StringInput
	// The configuration of how Amazon Macie classifies the S3 objects.
	ClassificationType S3BucketAssociationClassificationTypeArgsPtrInput
	// The ID of the Amazon Macie member account whose S3 resources you want to associate with Macie. If `memberAccountId` isn't specified, the action associates specified S3 resources with Macie for the current master account.
	MemberAccountId pulumi.StringPtrInput
	// Object key prefix identifying one or more S3 objects to which the association applies.
	Prefix pulumi.StringPtrInput
}

func (S3BucketAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*s3bucketAssociationArgs)(nil)).Elem()
}
