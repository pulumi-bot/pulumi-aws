// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Elasticsearch Domain.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elasticsearch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elasticsearch.NewDomain(ctx, "example", &elasticsearch.DomainArgs{
// 			ClusterConfig: &elasticsearch.DomainClusterConfigArgs{
// 				InstanceType: pulumi.String("r4.large.elasticsearch"),
// 			},
// 			ElasticsearchVersion: pulumi.String("1.5"),
// 			SnapshotOptions: &elasticsearch.DomainSnapshotOptionsArgs{
// 				AutomatedSnapshotStartHour: pulumi.Int(23),
// 			},
// 			Tags: pulumi.StringMap{
// 				"Domain": pulumi.String("TestDomain"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Access Policy
//
// > See also: `elasticsearch.DomainPolicy` resource
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elasticsearch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		domain := "tf-test"
// 		if param := cfg.Get("domain"); param != "" {
// 			domain = param
// 		}
// 		currentRegion, err := aws.GetRegion(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticsearch.NewDomain(ctx, "example", &elasticsearch.DomainArgs{
// 			AccessPolicies: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"es:*\",\n", "      \"Principal\": \"*\",\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"arn:aws:es:", currentRegion.Name, ":", currentCallerIdentity.AccountId, ":domain/", domain, "/*\",\n", "      \"Condition\": {\n", "        \"IpAddress\": {\"aws:SourceIp\": [\"66.193.100.22/32\"]}\n", "      }\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Log Publishing to CloudWatch Logs
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elasticsearch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewLogResourcePolicy(ctx, "exampleLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
// 			PolicyName:     pulumi.String("example"),
// 			PolicyDocument: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"es.amazonaws.com\"\n", "      },\n", "      \"Action\": [\n", "        \"logs:PutLogEvents\",\n", "        \"logs:PutLogEventsBatch\",\n", "        \"logs:CreateLogStream\"\n", "      ],\n", "      \"Resource\": \"arn:aws:logs:*\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticsearch.NewDomain(ctx, "exampleDomain", &elasticsearch.DomainArgs{
// 			LogPublishingOptions: elasticsearch.DomainLogPublishingOptionArray{
// 				&elasticsearch.DomainLogPublishingOptionArgs{
// 					CloudwatchLogGroupArn: exampleLogGroup.Arn,
// 					LogType:               pulumi.String("INDEX_SLOW_LOGS"),
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
//
// ## Import
//
// Elasticsearch domains can be imported using the `domain_name`, e.g.
//
// ```sh
//  $ pulumi import aws:elasticsearch/domain:Domain example domain_name
// ```
type Domain struct {
	pulumi.CustomResourceState

	// IAM policy document specifying the access policies for the domain
	AccessPolicies pulumi.StringOutput `pulumi:"accessPolicies"`
	// Key-value string pairs to specify advanced configuration options.
	// Note that the values for these configuration options must be strings (wrapped in quotes) or they
	// may be wrong and cause a perpetual diff, causing this provider to want to recreate your Elasticsearch
	// domain on every apply.
	AdvancedOptions pulumi.StringMapOutput `pulumi:"advancedOptions"`
	// Options for [fine-grained access control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html). See below for more details.
	AdvancedSecurityOptions DomainAdvancedSecurityOptionsOutput `pulumi:"advancedSecurityOptions"`
	// Amazon Resource Name (ARN) of the domain.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Cluster configuration of the domain, see below.
	ClusterConfig  DomainClusterConfigOutput     `pulumi:"clusterConfig"`
	CognitoOptions DomainCognitoOptionsPtrOutput `pulumi:"cognitoOptions"`
	// Domain endpoint HTTP(S) related options. See below.
	DomainEndpointOptions DomainDomainEndpointOptionsOutput `pulumi:"domainEndpointOptions"`
	// Unique identifier for the domain.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// Name of the domain.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
	EbsOptions DomainEbsOptionsOutput `pulumi:"ebsOptions"`
	// The version of Elasticsearch to deploy. Defaults to `1.5`
	ElasticsearchVersion pulumi.StringPtrOutput `pulumi:"elasticsearchVersion"`
	// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
	EncryptAtRest DomainEncryptAtRestOutput `pulumi:"encryptAtRest"`
	// Domain-specific endpoint used to submit index, search, and data upload requests.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Domain-specific endpoint for kibana without https scheme.
	// * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnetIds` were created inside.
	// * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
	KibanaEndpoint pulumi.StringOutput `pulumi:"kibanaEndpoint"`
	// Options for publishing slow  and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource.
	LogPublishingOptions DomainLogPublishingOptionArrayOutput `pulumi:"logPublishingOptions"`
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption DomainNodeToNodeEncryptionOutput `pulumi:"nodeToNodeEncryption"`
	// Snapshot related options, see below.
	SnapshotOptions DomainSnapshotOptionsPtrOutput `pulumi:"snapshotOptions"`
	// A map of tags to assign to the resource
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
	VpcOptions DomainVpcOptionsPtrOutput `pulumi:"vpcOptions"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		args = &DomainArgs{}
	}

	var resource Domain
	err := ctx.RegisterResource("aws:elasticsearch/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("aws:elasticsearch/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies *string `pulumi:"accessPolicies"`
	// Key-value string pairs to specify advanced configuration options.
	// Note that the values for these configuration options must be strings (wrapped in quotes) or they
	// may be wrong and cause a perpetual diff, causing this provider to want to recreate your Elasticsearch
	// domain on every apply.
	AdvancedOptions map[string]string `pulumi:"advancedOptions"`
	// Options for [fine-grained access control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html). See below for more details.
	AdvancedSecurityOptions *DomainAdvancedSecurityOptions `pulumi:"advancedSecurityOptions"`
	// Amazon Resource Name (ARN) of the domain.
	Arn *string `pulumi:"arn"`
	// Cluster configuration of the domain, see below.
	ClusterConfig  *DomainClusterConfig  `pulumi:"clusterConfig"`
	CognitoOptions *DomainCognitoOptions `pulumi:"cognitoOptions"`
	// Domain endpoint HTTP(S) related options. See below.
	DomainEndpointOptions *DomainDomainEndpointOptions `pulumi:"domainEndpointOptions"`
	// Unique identifier for the domain.
	DomainId *string `pulumi:"domainId"`
	// Name of the domain.
	DomainName *string `pulumi:"domainName"`
	// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
	EbsOptions *DomainEbsOptions `pulumi:"ebsOptions"`
	// The version of Elasticsearch to deploy. Defaults to `1.5`
	ElasticsearchVersion *string `pulumi:"elasticsearchVersion"`
	// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
	EncryptAtRest *DomainEncryptAtRest `pulumi:"encryptAtRest"`
	// Domain-specific endpoint used to submit index, search, and data upload requests.
	Endpoint *string `pulumi:"endpoint"`
	// Domain-specific endpoint for kibana without https scheme.
	// * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnetIds` were created inside.
	// * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
	KibanaEndpoint *string `pulumi:"kibanaEndpoint"`
	// Options for publishing slow  and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource.
	LogPublishingOptions []DomainLogPublishingOption `pulumi:"logPublishingOptions"`
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption *DomainNodeToNodeEncryption `pulumi:"nodeToNodeEncryption"`
	// Snapshot related options, see below.
	SnapshotOptions *DomainSnapshotOptions `pulumi:"snapshotOptions"`
	// A map of tags to assign to the resource
	Tags map[string]string `pulumi:"tags"`
	// VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
	VpcOptions *DomainVpcOptions `pulumi:"vpcOptions"`
}

type DomainState struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies pulumi.StringPtrInput
	// Key-value string pairs to specify advanced configuration options.
	// Note that the values for these configuration options must be strings (wrapped in quotes) or they
	// may be wrong and cause a perpetual diff, causing this provider to want to recreate your Elasticsearch
	// domain on every apply.
	AdvancedOptions pulumi.StringMapInput
	// Options for [fine-grained access control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html). See below for more details.
	AdvancedSecurityOptions DomainAdvancedSecurityOptionsPtrInput
	// Amazon Resource Name (ARN) of the domain.
	Arn pulumi.StringPtrInput
	// Cluster configuration of the domain, see below.
	ClusterConfig  DomainClusterConfigPtrInput
	CognitoOptions DomainCognitoOptionsPtrInput
	// Domain endpoint HTTP(S) related options. See below.
	DomainEndpointOptions DomainDomainEndpointOptionsPtrInput
	// Unique identifier for the domain.
	DomainId pulumi.StringPtrInput
	// Name of the domain.
	DomainName pulumi.StringPtrInput
	// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
	EbsOptions DomainEbsOptionsPtrInput
	// The version of Elasticsearch to deploy. Defaults to `1.5`
	ElasticsearchVersion pulumi.StringPtrInput
	// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
	EncryptAtRest DomainEncryptAtRestPtrInput
	// Domain-specific endpoint used to submit index, search, and data upload requests.
	Endpoint pulumi.StringPtrInput
	// Domain-specific endpoint for kibana without https scheme.
	// * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnetIds` were created inside.
	// * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
	KibanaEndpoint pulumi.StringPtrInput
	// Options for publishing slow  and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource.
	LogPublishingOptions DomainLogPublishingOptionArrayInput
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption DomainNodeToNodeEncryptionPtrInput
	// Snapshot related options, see below.
	SnapshotOptions DomainSnapshotOptionsPtrInput
	// A map of tags to assign to the resource
	Tags pulumi.StringMapInput
	// VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
	VpcOptions DomainVpcOptionsPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies interface{} `pulumi:"accessPolicies"`
	// Key-value string pairs to specify advanced configuration options.
	// Note that the values for these configuration options must be strings (wrapped in quotes) or they
	// may be wrong and cause a perpetual diff, causing this provider to want to recreate your Elasticsearch
	// domain on every apply.
	AdvancedOptions map[string]string `pulumi:"advancedOptions"`
	// Options for [fine-grained access control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html). See below for more details.
	AdvancedSecurityOptions *DomainAdvancedSecurityOptions `pulumi:"advancedSecurityOptions"`
	// Cluster configuration of the domain, see below.
	ClusterConfig  *DomainClusterConfig  `pulumi:"clusterConfig"`
	CognitoOptions *DomainCognitoOptions `pulumi:"cognitoOptions"`
	// Domain endpoint HTTP(S) related options. See below.
	DomainEndpointOptions *DomainDomainEndpointOptions `pulumi:"domainEndpointOptions"`
	// Name of the domain.
	DomainName *string `pulumi:"domainName"`
	// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
	EbsOptions *DomainEbsOptions `pulumi:"ebsOptions"`
	// The version of Elasticsearch to deploy. Defaults to `1.5`
	ElasticsearchVersion *string `pulumi:"elasticsearchVersion"`
	// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
	EncryptAtRest *DomainEncryptAtRest `pulumi:"encryptAtRest"`
	// Options for publishing slow  and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource.
	LogPublishingOptions []DomainLogPublishingOption `pulumi:"logPublishingOptions"`
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption *DomainNodeToNodeEncryption `pulumi:"nodeToNodeEncryption"`
	// Snapshot related options, see below.
	SnapshotOptions *DomainSnapshotOptions `pulumi:"snapshotOptions"`
	// A map of tags to assign to the resource
	Tags map[string]string `pulumi:"tags"`
	// VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
	VpcOptions *DomainVpcOptions `pulumi:"vpcOptions"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies pulumi.Input
	// Key-value string pairs to specify advanced configuration options.
	// Note that the values for these configuration options must be strings (wrapped in quotes) or they
	// may be wrong and cause a perpetual diff, causing this provider to want to recreate your Elasticsearch
	// domain on every apply.
	AdvancedOptions pulumi.StringMapInput
	// Options for [fine-grained access control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html). See below for more details.
	AdvancedSecurityOptions DomainAdvancedSecurityOptionsPtrInput
	// Cluster configuration of the domain, see below.
	ClusterConfig  DomainClusterConfigPtrInput
	CognitoOptions DomainCognitoOptionsPtrInput
	// Domain endpoint HTTP(S) related options. See below.
	DomainEndpointOptions DomainDomainEndpointOptionsPtrInput
	// Name of the domain.
	DomainName pulumi.StringPtrInput
	// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
	EbsOptions DomainEbsOptionsPtrInput
	// The version of Elasticsearch to deploy. Defaults to `1.5`
	ElasticsearchVersion pulumi.StringPtrInput
	// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
	EncryptAtRest DomainEncryptAtRestPtrInput
	// Options for publishing slow  and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource.
	LogPublishingOptions DomainLogPublishingOptionArrayInput
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption DomainNodeToNodeEncryptionPtrInput
	// Snapshot related options, see below.
	SnapshotOptions DomainSnapshotOptionsPtrInput
	// A map of tags to assign to the resource
	Tags pulumi.StringMapInput
	// VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
	VpcOptions DomainVpcOptionsPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

type DomainPtrInput interface {
	pulumi.Input

	ToDomainPtrOutput() DomainPtrOutput
	ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput
}

func (Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil)).Elem()
}

func (i Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

func (i Domain) ToDomainPtrOutput() DomainPtrOutput {
	return i.ToDomainPtrOutputWithContext(context.Background())
}

func (i Domain) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPtrOutput)
}

type DomainOutput struct {
	*pulumi.OutputState
}

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainOutput)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

type DomainPtrOutput struct {
	*pulumi.OutputState
}

func (DomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainPtrOutput) ToDomainPtrOutput() DomainPtrOutput {
	return o
}

func (o DomainPtrOutput) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainPtrOutput{})
}
