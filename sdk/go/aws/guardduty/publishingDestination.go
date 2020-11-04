// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage a GuardDuty PublishingDestination. Requires an existing GuardDuty Detector.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/guardduty"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kms"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		currentRegion, err := aws.GetRegion(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		gdBucket, err := s3.NewBucket(ctx, "gdBucket", &s3.BucketArgs{
// 			Acl:          pulumi.String("private"),
// 			ForceDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		kmsPol, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Sid: "Allow GuardDuty to encrypt findings",
// 					Actions: []string{
// 						"kms:GenerateDataKey",
// 					},
// 					Resources: []string{
// 						fmt.Sprintf("%v%v%v%v%v", "arn:aws:kms:", currentRegion.Name, ":", currentCallerIdentity.AccountId, ":key/*"),
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "Service",
// 							Identifiers: []string{
// 								"guardduty.amazonaws.com",
// 							},
// 						},
// 					},
// 				},
// 				iam.GetPolicyDocumentStatement{
// 					Sid: "Allow all users to modify/delete key (test only)",
// 					Actions: []string{
// 						"kms:*",
// 					},
// 					Resources: []string{
// 						fmt.Sprintf("%v%v%v%v%v", "arn:aws:kms:", currentRegion.Name, ":", currentCallerIdentity.AccountId, ":key/*"),
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "AWS",
// 							Identifiers: []string{
// 								fmt.Sprintf("%v%v%v", "arn:aws:iam::", currentCallerIdentity.AccountId, ":root"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		testGd, err := guardduty.NewDetector(ctx, "testGd", &guardduty.DetectorArgs{
// 			Enable: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		gdBucketPolicy, err := s3.NewBucketPolicy(ctx, "gdBucketPolicy", &s3.BucketPolicyArgs{
// 			Bucket: gdBucket.ID(),
// 			Policy: bucketPol.ApplyT(func(bucketPol iam.GetPolicyDocumentResult) (string, error) {
// 				return bucketPol.Json, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		gdKey, err := kms.NewKey(ctx, "gdKey", &kms.KeyArgs{
// 			Description:          pulumi.String("Temporary key for AccTest of TF"),
// 			DeletionWindowInDays: pulumi.Int(7),
// 			Policy:               pulumi.String(kmsPol.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = guardduty.NewPublishingDestination(ctx, "test", &guardduty.PublishingDestinationArgs{
// 			DetectorId:     testGd.ID(),
// 			DestinationArn: gdBucket.Arn,
// 			KmsKeyArn:      gdKey.Arn,
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			gdBucketPolicy,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// > **Note:** Please do not use this simple example for Bucket-Policy and KMS Key Policy in a production environment. It is much too open for such a use-case. Refer to the AWS documentation here: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_exportfindings.html
type PublishingDestination struct {
	pulumi.CustomResourceState

	// The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
	DestinationArn pulumi.StringOutput `pulumi:"destinationArn"`
	// Currently there is only "S3" available as destination type which is also the default value
	DestinationType pulumi.StringPtrOutput `pulumi:"destinationType"`
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`
	// The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
	KmsKeyArn pulumi.StringOutput `pulumi:"kmsKeyArn"`
}

// NewPublishingDestination registers a new resource with the given unique name, arguments, and options.
func NewPublishingDestination(ctx *pulumi.Context,
	name string, args *PublishingDestinationArgs, opts ...pulumi.ResourceOption) (*PublishingDestination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DestinationArn == nil {
		return nil, errors.New("invalid value for required argument 'DestinationArn'")
	}
	if args.DetectorId == nil {
		return nil, errors.New("invalid value for required argument 'DetectorId'")
	}
	if args.KmsKeyArn == nil {
		return nil, errors.New("invalid value for required argument 'KmsKeyArn'")
	}
	var resource PublishingDestination
	err := ctx.RegisterResource("aws:guardduty/publishingDestination:PublishingDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublishingDestination gets an existing PublishingDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublishingDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublishingDestinationState, opts ...pulumi.ResourceOption) (*PublishingDestination, error) {
	var resource PublishingDestination
	err := ctx.ReadResource("aws:guardduty/publishingDestination:PublishingDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublishingDestination resources.
type publishingDestinationState struct {
	// The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
	DestinationArn *string `pulumi:"destinationArn"`
	// Currently there is only "S3" available as destination type which is also the default value
	DestinationType *string `pulumi:"destinationType"`
	// The detector ID of the GuardDuty.
	DetectorId *string `pulumi:"detectorId"`
	// The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
}

type PublishingDestinationState struct {
	// The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
	DestinationArn pulumi.StringPtrInput
	// Currently there is only "S3" available as destination type which is also the default value
	DestinationType pulumi.StringPtrInput
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringPtrInput
	// The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
	KmsKeyArn pulumi.StringPtrInput
}

func (PublishingDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*publishingDestinationState)(nil)).Elem()
}

type publishingDestinationArgs struct {
	// The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
	DestinationArn string `pulumi:"destinationArn"`
	// Currently there is only "S3" available as destination type which is also the default value
	DestinationType *string `pulumi:"destinationType"`
	// The detector ID of the GuardDuty.
	DetectorId string `pulumi:"detectorId"`
	// The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
	KmsKeyArn string `pulumi:"kmsKeyArn"`
}

// The set of arguments for constructing a PublishingDestination resource.
type PublishingDestinationArgs struct {
	// The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
	DestinationArn pulumi.StringInput
	// Currently there is only "S3" available as destination type which is also the default value
	DestinationType pulumi.StringPtrInput
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringInput
	// The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
	KmsKeyArn pulumi.StringInput
}

func (PublishingDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publishingDestinationArgs)(nil)).Elem()
}
