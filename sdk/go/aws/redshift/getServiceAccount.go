// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the Account ID of the [AWS Redshift Service Account](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
// in a given region for the purpose of allowing Redshift to store audit data in S3.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := redshift.GetServiceAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			ForceDestroy: pulumi.Bool(true),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"Version\": \"2008-10-17\",\n", "	\"Statement\": [\n", "		{\n", "        			\"Sid\": \"Put bucket policy needed for audit logging\",\n", "        			\"Effect\": \"Allow\",\n", "        			\"Principal\": {\n", "						\"AWS\": \"", main.Arn, "\"\n", "        			},\n", "        			\"Action\": \"s3:PutObject\",\n", "        			\"Resource\": \"arn:aws:s3:::tf-redshift-logging-test-bucket/*\"\n", "        		},\n", "        		{\n", "        			\"Sid\": \"Get bucket policy needed for audit logging \",\n", "        			\"Effect\": \"Allow\",\n", "        			\"Principal\": {\n", "						\"AWS\": \"", main.Arn, "\"\n", "        			},\n", "        			\"Action\": \"s3:GetBucketAcl\",\n", "        			\"Resource\": \"arn:aws:s3:::tf-redshift-logging-test-bucket\"\n", "        		}\n", "	]\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetServiceAccount(ctx *pulumi.Context, args *GetServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetServiceAccountResult, error) {
	var rv GetServiceAccountResult
	err := ctx.Invoke("aws:redshift/getServiceAccount:getServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceAccount.
type GetServiceAccountArgs struct {
	// Name of the region whose AWS Redshift account ID is desired.
	// Defaults to the region from the AWS provider configuration.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getServiceAccount.
type GetServiceAccountResult struct {
	// The ARN of the AWS Redshift service account in the selected region.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Region *string `pulumi:"region"`
}
