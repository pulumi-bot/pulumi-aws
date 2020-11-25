// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Route53 query logging configuration resource.
//
// > **NOTE:** There are restrictions on the configuration of query logging. Notably,
// the CloudWatch log group must be in the `us-east-1` region,
// a permissive CloudWatch log resource policy must be in place, and
// the Route53 hosted zone must be public.
// See [Configuring Logging for DNS Queries](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html?console_help=true#query-logs-configuring) for additional details.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/providers"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := providers.Newaws(ctx, "us_east_1", &providers.awsArgs{
// 			Region: pulumi.String("us-east-1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		awsRoute53ExampleCom, err := cloudwatch.NewLogGroup(ctx, "awsRoute53ExampleCom", &cloudwatch.LogGroupArgs{
// 			RetentionInDays: pulumi.Int(30),
// 		}, pulumi.Provider(aws.Us-east-1))
// 		if err != nil {
// 			return err
// 		}
// 		route53_query_logging_policyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"logs:CreateLogStream",
// 						"logs:PutLogEvents",
// 					},
// 					Resources: []string{
// 						"arn:aws:logs:*:*:log-group:/aws/route53/*",
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Identifiers: []string{
// 								"route53.amazonaws.com",
// 							},
// 							Type: "Service",
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewLogResourcePolicy(ctx, "route53_query_logging_policyLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
// 			PolicyDocument: pulumi.String(route53_query_logging_policyPolicyDocument.Json),
// 			PolicyName:     pulumi.String("route53-query-logging-policy"),
// 		}, pulumi.Provider(aws.Us-east-1))
// 		if err != nil {
// 			return err
// 		}
// 		exampleComZone, err := route53.NewZone(ctx, "exampleComZone", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewQueryLog(ctx, "exampleComQueryLog", &route53.QueryLogArgs{
// 			CloudwatchLogGroupArn: awsRoute53ExampleCom.Arn,
// 			ZoneId:                exampleComZone.ZoneId,
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			route53_query_logging_policyLogResourcePolicy,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Route53 query logging configurations can be imported using their ID, e.g.
//
// ```sh
//  $ pulumi import aws:route53/queryLog:QueryLog example_com xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// ```
type QueryLog struct {
	pulumi.CustomResourceState

	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn pulumi.StringOutput `pulumi:"cloudwatchLogGroupArn"`
	// Route53 hosted zone ID to enable query logs.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewQueryLog registers a new resource with the given unique name, arguments, and options.
func NewQueryLog(ctx *pulumi.Context,
	name string, args *QueryLogArgs, opts ...pulumi.ResourceOption) (*QueryLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudwatchLogGroupArn == nil {
		return nil, errors.New("invalid value for required argument 'CloudwatchLogGroupArn'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource QueryLog
	err := ctx.RegisterResource("aws:route53/queryLog:QueryLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueryLog gets an existing QueryLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueryLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryLogState, opts ...pulumi.ResourceOption) (*QueryLog, error) {
	var resource QueryLog
	err := ctx.ReadResource("aws:route53/queryLog:QueryLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueryLog resources.
type queryLogState struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn *string `pulumi:"cloudwatchLogGroupArn"`
	// Route53 hosted zone ID to enable query logs.
	ZoneId *string `pulumi:"zoneId"`
}

type QueryLogState struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn pulumi.StringPtrInput
	// Route53 hosted zone ID to enable query logs.
	ZoneId pulumi.StringPtrInput
}

func (QueryLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryLogState)(nil)).Elem()
}

type queryLogArgs struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn string `pulumi:"cloudwatchLogGroupArn"`
	// Route53 hosted zone ID to enable query logs.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a QueryLog resource.
type QueryLogArgs struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn pulumi.StringInput
	// Route53 hosted zone ID to enable query logs.
	ZoneId pulumi.StringInput
}

func (QueryLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryLogArgs)(nil)).Elem()
}

type QueryLogInput interface {
	pulumi.Input

	ToQueryLogOutput() QueryLogOutput
	ToQueryLogOutputWithContext(ctx context.Context) QueryLogOutput
}

func (QueryLog) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryLog)(nil)).Elem()
}

func (i QueryLog) ToQueryLogOutput() QueryLogOutput {
	return i.ToQueryLogOutputWithContext(context.Background())
}

func (i QueryLog) ToQueryLogOutputWithContext(ctx context.Context) QueryLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryLogOutput)
}

type QueryLogOutput struct {
	*pulumi.OutputState
}

func (QueryLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryLogOutput)(nil)).Elem()
}

func (o QueryLogOutput) ToQueryLogOutput() QueryLogOutput {
	return o
}

func (o QueryLogOutput) ToQueryLogOutputWithContext(ctx context.Context) QueryLogOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(QueryLogOutput{})
}
