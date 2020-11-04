// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Glue Crawler. More information can be found in the [AWS Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html)
//
// ## Example Usage
// ### DynamoDB Target
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewCrawler(ctx, "example", &glue.CrawlerArgs{
// 			DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
// 			Role:         pulumi.Any(aws_iam_role.Example.Arn),
// 			DynamodbTargets: glue.CrawlerDynamodbTargetArray{
// 				&glue.CrawlerDynamodbTargetArgs{
// 					Path: pulumi.String("table-name"),
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
// ### JDBC Target
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewCrawler(ctx, "example", &glue.CrawlerArgs{
// 			DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
// 			Role:         pulumi.Any(aws_iam_role.Example.Arn),
// 			JdbcTargets: glue.CrawlerJdbcTargetArray{
// 				&glue.CrawlerJdbcTargetArgs{
// 					ConnectionName: pulumi.Any(aws_glue_connection.Example.Name),
// 					Path:           pulumi.String(fmt.Sprintf("%v%v", "database-name/", "%")),
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
// ### S3 Target
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewCrawler(ctx, "example", &glue.CrawlerArgs{
// 			DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
// 			Role:         pulumi.Any(aws_iam_role.Example.Arn),
// 			S3Targets: glue.CrawlerS3TargetArray{
// 				&glue.CrawlerS3TargetArgs{
// 					Path: pulumi.String(fmt.Sprintf("%v%v", "s3://", aws_s3_bucket.Example.Bucket)),
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
type Crawler struct {
	pulumi.CustomResourceState

	// The ARN of the crawler
	Arn            pulumi.StringOutput             `pulumi:"arn"`
	CatalogTargets CrawlerCatalogTargetArrayOutput `pulumi:"catalogTargets"`
	// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
	Classifiers pulumi.StringArrayOutput `pulumi:"classifiers"`
	// JSON string of configuration information.
	Configuration pulumi.StringPtrOutput `pulumi:"configuration"`
	// Glue database where results are written.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Description of the crawler.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of nested DynamoDB target arguments. See below.
	DynamodbTargets CrawlerDynamodbTargetArrayOutput `pulumi:"dynamodbTargets"`
	// List of nested JBDC target arguments. See below.
	JdbcTargets CrawlerJdbcTargetArrayOutput `pulumi:"jdbcTargets"`
	// Name of the crawler.
	Name pulumi.StringOutput `pulumi:"name"`
	// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
	Role pulumi.StringOutput `pulumi:"role"`
	// List nested Amazon S3 target arguments. See below.
	S3Targets CrawlerS3TargetArrayOutput `pulumi:"s3Targets"`
	// A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
	Schedule pulumi.StringPtrOutput `pulumi:"schedule"`
	// Policy for the crawler's update and deletion behavior.
	SchemaChangePolicy CrawlerSchemaChangePolicyPtrOutput `pulumi:"schemaChangePolicy"`
	// The name of Security Configuration to be used by the crawler
	SecurityConfiguration pulumi.StringPtrOutput `pulumi:"securityConfiguration"`
	// The table prefix used for catalog tables that are created.
	TablePrefix pulumi.StringPtrOutput `pulumi:"tablePrefix"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewCrawler registers a new resource with the given unique name, arguments, and options.
func NewCrawler(ctx *pulumi.Context,
	name string, args *CrawlerArgs, opts ...pulumi.ResourceOption) (*Crawler, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource Crawler
	err := ctx.RegisterResource("aws:glue/crawler:Crawler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCrawler gets an existing Crawler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCrawler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CrawlerState, opts ...pulumi.ResourceOption) (*Crawler, error) {
	var resource Crawler
	err := ctx.ReadResource("aws:glue/crawler:Crawler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Crawler resources.
type crawlerState struct {
	// The ARN of the crawler
	Arn            *string                `pulumi:"arn"`
	CatalogTargets []CrawlerCatalogTarget `pulumi:"catalogTargets"`
	// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
	Classifiers []string `pulumi:"classifiers"`
	// JSON string of configuration information.
	Configuration *string `pulumi:"configuration"`
	// Glue database where results are written.
	DatabaseName *string `pulumi:"databaseName"`
	// Description of the crawler.
	Description *string `pulumi:"description"`
	// List of nested DynamoDB target arguments. See below.
	DynamodbTargets []CrawlerDynamodbTarget `pulumi:"dynamodbTargets"`
	// List of nested JBDC target arguments. See below.
	JdbcTargets []CrawlerJdbcTarget `pulumi:"jdbcTargets"`
	// Name of the crawler.
	Name *string `pulumi:"name"`
	// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
	Role *string `pulumi:"role"`
	// List nested Amazon S3 target arguments. See below.
	S3Targets []CrawlerS3Target `pulumi:"s3Targets"`
	// A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
	Schedule *string `pulumi:"schedule"`
	// Policy for the crawler's update and deletion behavior.
	SchemaChangePolicy *CrawlerSchemaChangePolicy `pulumi:"schemaChangePolicy"`
	// The name of Security Configuration to be used by the crawler
	SecurityConfiguration *string `pulumi:"securityConfiguration"`
	// The table prefix used for catalog tables that are created.
	TablePrefix *string `pulumi:"tablePrefix"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type CrawlerState struct {
	// The ARN of the crawler
	Arn            pulumi.StringPtrInput
	CatalogTargets CrawlerCatalogTargetArrayInput
	// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
	Classifiers pulumi.StringArrayInput
	// JSON string of configuration information.
	Configuration pulumi.StringPtrInput
	// Glue database where results are written.
	DatabaseName pulumi.StringPtrInput
	// Description of the crawler.
	Description pulumi.StringPtrInput
	// List of nested DynamoDB target arguments. See below.
	DynamodbTargets CrawlerDynamodbTargetArrayInput
	// List of nested JBDC target arguments. See below.
	JdbcTargets CrawlerJdbcTargetArrayInput
	// Name of the crawler.
	Name pulumi.StringPtrInput
	// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
	Role pulumi.StringPtrInput
	// List nested Amazon S3 target arguments. See below.
	S3Targets CrawlerS3TargetArrayInput
	// A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
	Schedule pulumi.StringPtrInput
	// Policy for the crawler's update and deletion behavior.
	SchemaChangePolicy CrawlerSchemaChangePolicyPtrInput
	// The name of Security Configuration to be used by the crawler
	SecurityConfiguration pulumi.StringPtrInput
	// The table prefix used for catalog tables that are created.
	TablePrefix pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (CrawlerState) ElementType() reflect.Type {
	return reflect.TypeOf((*crawlerState)(nil)).Elem()
}

type crawlerArgs struct {
	CatalogTargets []CrawlerCatalogTarget `pulumi:"catalogTargets"`
	// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
	Classifiers []string `pulumi:"classifiers"`
	// JSON string of configuration information.
	Configuration *string `pulumi:"configuration"`
	// Glue database where results are written.
	DatabaseName string `pulumi:"databaseName"`
	// Description of the crawler.
	Description *string `pulumi:"description"`
	// List of nested DynamoDB target arguments. See below.
	DynamodbTargets []CrawlerDynamodbTarget `pulumi:"dynamodbTargets"`
	// List of nested JBDC target arguments. See below.
	JdbcTargets []CrawlerJdbcTarget `pulumi:"jdbcTargets"`
	// Name of the crawler.
	Name *string `pulumi:"name"`
	// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
	Role string `pulumi:"role"`
	// List nested Amazon S3 target arguments. See below.
	S3Targets []CrawlerS3Target `pulumi:"s3Targets"`
	// A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
	Schedule *string `pulumi:"schedule"`
	// Policy for the crawler's update and deletion behavior.
	SchemaChangePolicy *CrawlerSchemaChangePolicy `pulumi:"schemaChangePolicy"`
	// The name of Security Configuration to be used by the crawler
	SecurityConfiguration *string `pulumi:"securityConfiguration"`
	// The table prefix used for catalog tables that are created.
	TablePrefix *string `pulumi:"tablePrefix"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Crawler resource.
type CrawlerArgs struct {
	CatalogTargets CrawlerCatalogTargetArrayInput
	// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
	Classifiers pulumi.StringArrayInput
	// JSON string of configuration information.
	Configuration pulumi.StringPtrInput
	// Glue database where results are written.
	DatabaseName pulumi.StringInput
	// Description of the crawler.
	Description pulumi.StringPtrInput
	// List of nested DynamoDB target arguments. See below.
	DynamodbTargets CrawlerDynamodbTargetArrayInput
	// List of nested JBDC target arguments. See below.
	JdbcTargets CrawlerJdbcTargetArrayInput
	// Name of the crawler.
	Name pulumi.StringPtrInput
	// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
	Role pulumi.StringInput
	// List nested Amazon S3 target arguments. See below.
	S3Targets CrawlerS3TargetArrayInput
	// A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
	Schedule pulumi.StringPtrInput
	// Policy for the crawler's update and deletion behavior.
	SchemaChangePolicy CrawlerSchemaChangePolicyPtrInput
	// The name of Security Configuration to be used by the crawler
	SecurityConfiguration pulumi.StringPtrInput
	// The table prefix used for catalog tables that are created.
	TablePrefix pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (CrawlerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*crawlerArgs)(nil)).Elem()
}
