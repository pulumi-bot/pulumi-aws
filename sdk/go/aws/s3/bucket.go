// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a S3 bucket resource.
//
// ## Example Usage
// ### Private Bucket w/ Tags
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 			Tags: pulumi.Map{
// 				"Environment": pulumi.String("Dev"),
// 				"Name":        pulumi.String("My bucket"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using CORS
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Acl: pulumi.String("public-read"),
// 			CorsRules: s3.BucketCorsRuleArray{
// 				&s3.BucketCorsRuleArgs{
// 					AllowedHeaders: pulumi.StringArray{
// 						pulumi.String("*"),
// 					},
// 					AllowedMethods: pulumi.StringArray{
// 						pulumi.String("PUT"),
// 						pulumi.String("POST"),
// 					},
// 					AllowedOrigins: pulumi.StringArray{
// 						pulumi.String("https://s3-website-test.mydomain.com"),
// 					},
// 					ExposeHeaders: pulumi.StringArray{
// 						pulumi.String("ETag"),
// 					},
// 					MaxAgeSeconds: pulumi.Int(3000),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using versioning
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 			Versioning: &s3.BucketVersioningArgs{
// 				Enabled: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Enable Logging
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		logBucket, err := s3.NewBucket(ctx, "logBucket", &s3.BucketArgs{
// 			Acl: pulumi.String("log-delivery-write"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 			Loggings: s3.BucketLoggingArray{
// 				&s3.BucketLoggingArgs{
// 					TargetBucket: logBucket.ID(),
// 					TargetPrefix: pulumi.String("log/"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using object lifecycle
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 			LifecycleRules: s3.BucketLifecycleRuleArray{
// 				&s3.BucketLifecycleRuleArgs{
// 					Enabled: pulumi.Bool(true),
// 					Expiration: &s3.BucketLifecycleRuleExpirationArgs{
// 						Days: pulumi.Int(90),
// 					},
// 					Id:     pulumi.String("log"),
// 					Prefix: pulumi.String("log/"),
// 					Tags: pulumi.Map{
// 						"autoclean": pulumi.String("true"),
// 						"rule":      pulumi.String("log"),
// 					},
// 					Transition: pulumi.MapArray{
// 						pulumi.Map{
// 							"days":         pulumi.Float64(30),
// 							"storageClass": pulumi.String("STANDARD_IA"),
// 						},
// 						pulumi.Map{
// 							"days":         pulumi.Float64(60),
// 							"storageClass": pulumi.String("GLACIER"),
// 						},
// 					},
// 				},
// 				&s3.BucketLifecycleRuleArgs{
// 					Enabled: pulumi.Bool(true),
// 					Expiration: &s3.BucketLifecycleRuleExpirationArgs{
// 						Date: pulumi.String("2016-01-12"),
// 					},
// 					Id:     pulumi.String("tmp"),
// 					Prefix: pulumi.String("tmp/"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucket(ctx, "versioningBucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 			LifecycleRules: s3.BucketLifecycleRuleArray{
// 				&s3.BucketLifecycleRuleArgs{
// 					Enabled: pulumi.Bool(true),
// 					NoncurrentVersionExpiration: &s3.BucketLifecycleRuleNoncurrentVersionExpirationArgs{
// 						Days: pulumi.Int(90),
// 					},
// 					NoncurrentVersionTransition: pulumi.MapArray{
// 						pulumi.Map{
// 							"days":         pulumi.Float64(30),
// 							"storageClass": pulumi.String("STANDARD_IA"),
// 						},
// 						pulumi.Map{
// 							"days":         pulumi.Float64(60),
// 							"storageClass": pulumi.String("GLACIER"),
// 						},
// 					},
// 					Prefix: pulumi.String("config/"),
// 				},
// 			},
// 			Versioning: &s3.BucketVersioningArgs{
// 				Enabled: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Enable Default Server Side Encryption
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		mykey, err := kms.NewKey(ctx, "mykey", &kms.KeyArgs{
// 			DeletionWindowInDays: pulumi.Int(10),
// 			Description:          pulumi.String("This key is used to encrypt bucket objects"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucket(ctx, "mybucket", &s3.BucketArgs{
// 			ServerSideEncryptionConfiguration: &s3.BucketServerSideEncryptionConfigurationArgs{
// 				Rule: &s3.BucketServerSideEncryptionConfigurationRuleArgs{
// 					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
// 						KmsMasterKeyId: mykey.Arn,
// 						SseAlgorithm:   pulumi.String("aws:kms"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using ACL policy grants
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		currentUser, err := aws.GetCanonicalUserId(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Grants: s3.BucketGrantArray{
// 				&s3.BucketGrantArgs{
// 					Id: pulumi.String(currentUser.Id),
// 					Permissions: pulumi.StringArray{
// 						pulumi.String("FULL_CONTROL"),
// 					},
// 					Type: pulumi.String("CanonicalUser"),
// 				},
// 				&s3.BucketGrantArgs{
// 					Permissions: pulumi.StringArray{
// 						pulumi.String("READ"),
// 						pulumi.String("WRITE"),
// 					},
// 					Type: pulumi.String("Group"),
// 					Uri:  pulumi.String("http://acs.amazonaws.com/groups/s3/LogDelivery"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Bucket struct {
	pulumi.CustomResourceState

	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus pulumi.StringOutput `pulumi:"accelerationStatus"`
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
	BucketDomainName pulumi.StringOutput `pulumi:"bucketDomainName"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrOutput `pulumi:"bucketPrefix"`
	// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
	BucketRegionalDomainName pulumi.StringOutput `pulumi:"bucketRegionalDomainName"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules BucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants BucketGrantArrayOutput `pulumi:"grants"`
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId pulumi.StringOutput `pulumi:"hostedZoneId"`
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings BucketLoggingArrayOutput `pulumi:"loggings"`
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration BucketObjectLockConfigurationPtrOutput `pulumi:"objectLockConfiguration"`
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), the provider may view the policy as constantly changing in a `pulumi up / preview / update`. In this case, please make sure you use the verbose/specific version of the policy.
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region pulumi.StringOutput `pulumi:"region"`
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration BucketReplicationConfigurationPtrOutput `pulumi:"replicationConfiguration"`
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer pulumi.StringOutput `pulumi:"requestPayer"`
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration BucketServerSideEncryptionConfigurationPtrOutput `pulumi:"serverSideEncryptionConfiguration"`
	// A mapping of tags to assign to the bucket.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning BucketVersioningOutput `pulumi:"versioning"`
	// A website object (documented below).
	Website BucketWebsitePtrOutput `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain pulumi.StringOutput `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringOutput `pulumi:"websiteEndpoint"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}
	var resource Bucket
	err := ctx.RegisterResource("aws:s3/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("aws:s3/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus *string `pulumi:"accelerationStatus"`
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl *string `pulumi:"acl"`
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn *string `pulumi:"arn"`
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket *string `pulumi:"bucket"`
	// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
	BucketDomainName *string `pulumi:"bucketDomainName"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
	BucketRegionalDomainName *string `pulumi:"bucketRegionalDomainName"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants []BucketGrant `pulumi:"grants"`
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId *string `pulumi:"hostedZoneId"`
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings []BucketLogging `pulumi:"loggings"`
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration *BucketObjectLockConfiguration `pulumi:"objectLockConfiguration"`
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), the provider may view the policy as constantly changing in a `pulumi up / preview / update`. In this case, please make sure you use the verbose/specific version of the policy.
	Policy *string `pulumi:"policy"`
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region *string `pulumi:"region"`
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration *BucketReplicationConfiguration `pulumi:"replicationConfiguration"`
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer *string `pulumi:"requestPayer"`
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration *BucketServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	// A mapping of tags to assign to the bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning *BucketVersioning `pulumi:"versioning"`
	// A website object (documented below).
	Website *BucketWebsite `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

type BucketState struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus pulumi.StringPtrInput
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl pulumi.StringPtrInput
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn pulumi.StringPtrInput
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket pulumi.StringPtrInput
	// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
	BucketDomainName pulumi.StringPtrInput
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrInput
	// The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
	BucketRegionalDomainName pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules BucketCorsRuleArrayInput
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants BucketGrantArrayInput
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId pulumi.StringPtrInput
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings BucketLoggingArrayInput
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration BucketObjectLockConfigurationPtrInput
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), the provider may view the policy as constantly changing in a `pulumi up / preview / update`. In this case, please make sure you use the verbose/specific version of the policy.
	Policy pulumi.StringPtrInput
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region pulumi.StringPtrInput
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration BucketReplicationConfigurationPtrInput
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer pulumi.StringPtrInput
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration BucketServerSideEncryptionConfigurationPtrInput
	// A mapping of tags to assign to the bucket.
	Tags pulumi.MapInput
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning BucketVersioningPtrInput
	// A website object (documented below).
	Website BucketWebsitePtrInput
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus *string `pulumi:"accelerationStatus"`
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl interface{} `pulumi:"acl"`
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn *string `pulumi:"arn"`
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket *string `pulumi:"bucket"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants []BucketGrant `pulumi:"grants"`
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId *string `pulumi:"hostedZoneId"`
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings []BucketLogging `pulumi:"loggings"`
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration *BucketObjectLockConfiguration `pulumi:"objectLockConfiguration"`
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), the provider may view the policy as constantly changing in a `pulumi up / preview / update`. In this case, please make sure you use the verbose/specific version of the policy.
	Policy interface{} `pulumi:"policy"`
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region *string `pulumi:"region"`
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration *BucketReplicationConfiguration `pulumi:"replicationConfiguration"`
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer *string `pulumi:"requestPayer"`
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration *BucketServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	// A mapping of tags to assign to the bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning *BucketVersioning `pulumi:"versioning"`
	// A website object (documented below).
	Website *BucketWebsite `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
	AccelerationStatus pulumi.StringPtrInput
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".  Conflicts with `grant`.
	Acl pulumi.Input
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn pulumi.StringPtrInput
	// The name of the bucket. If omitted, this provider will assign a random, unique name.
	Bucket pulumi.StringPtrInput
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules BucketCorsRuleArrayInput
	// A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// An [ACL policy grant](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#sample-acl) (documented below). Conflicts with `acl`.
	Grants BucketGrantArrayInput
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId pulumi.StringPtrInput
	// A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
	Loggings BucketLoggingArrayInput
	// A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
	ObjectLockConfiguration BucketObjectLockConfigurationPtrInput
	// A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), the provider may view the policy as constantly changing in a `pulumi up / preview / update`. In this case, please make sure you use the verbose/specific version of the policy.
	Policy pulumi.Input
	// If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
	Region pulumi.StringPtrInput
	// A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
	ReplicationConfiguration BucketReplicationConfigurationPtrInput
	// Specifies who should bear the cost of Amazon S3 data transfer.
	// Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
	// the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// developer guide for more information.
	RequestPayer pulumi.StringPtrInput
	// A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
	ServerSideEncryptionConfiguration BucketServerSideEncryptionConfigurationPtrInput
	// A mapping of tags to assign to the bucket.
	Tags pulumi.MapInput
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning BucketVersioningPtrInput
	// A website object (documented below).
	Website BucketWebsitePtrInput
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}
