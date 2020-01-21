// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package elasticsearch

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an AWS Elasticsearch Domain.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elasticsearch_domain.html.markdown.
type Domain struct {
	pulumi.CustomResourceState

	// IAM policy document specifying the access policies for the domain
	AccessPolicies pulumi.StringOutput `pulumi:"accessPolicies"`
	// Key-value string pairs to specify advanced configuration options.
	// Note that the values for these configuration options must be strings (wrapped in quotes) or they
	// may be wrong and cause a perpetual diff, causing this provider to want to recreate your Elasticsearch
	// domain on every apply.
	AdvancedOptions pulumi.MapOutput `pulumi:"advancedOptions"`
	// Amazon Resource Name (ARN) of the domain.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Cluster configuration of the domain, see below.
	ClusterConfig DomainClusterConfigOutput `pulumi:"clusterConfig"`
	CognitoOptions DomainCognitoOptionsPtrOutput `pulumi:"cognitoOptions"`
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
	// Options for publishing slow logs to CloudWatch Logs.
	LogPublishingOptions DomainLogPublishingOptionArrayOutput `pulumi:"logPublishingOptions"`
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption DomainNodeToNodeEncryptionOutput `pulumi:"nodeToNodeEncryption"`
	// Snapshot related options, see below.
	SnapshotOptions DomainSnapshotOptionsPtrOutput `pulumi:"snapshotOptions"`
	// A mapping of tags to assign to the resource
	Tags pulumi.MapOutput `pulumi:"tags"`
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
	AdvancedOptions map[string]interface{} `pulumi:"advancedOptions"`
	// Amazon Resource Name (ARN) of the domain.
	Arn *string `pulumi:"arn"`
	// Cluster configuration of the domain, see below.
	ClusterConfig *DomainClusterConfig `pulumi:"clusterConfig"`
	CognitoOptions *DomainCognitoOptions `pulumi:"cognitoOptions"`
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
	// Options for publishing slow logs to CloudWatch Logs.
	LogPublishingOptions []DomainLogPublishingOption `pulumi:"logPublishingOptions"`
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption *DomainNodeToNodeEncryption `pulumi:"nodeToNodeEncryption"`
	// Snapshot related options, see below.
	SnapshotOptions *DomainSnapshotOptions `pulumi:"snapshotOptions"`
	// A mapping of tags to assign to the resource
	Tags map[string]interface{} `pulumi:"tags"`
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
	AdvancedOptions pulumi.MapInput
	// Amazon Resource Name (ARN) of the domain.
	Arn pulumi.StringPtrInput
	// Cluster configuration of the domain, see below.
	ClusterConfig DomainClusterConfigPtrInput
	CognitoOptions DomainCognitoOptionsPtrInput
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
	// Options for publishing slow logs to CloudWatch Logs.
	LogPublishingOptions DomainLogPublishingOptionArrayInput
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption DomainNodeToNodeEncryptionPtrInput
	// Snapshot related options, see below.
	SnapshotOptions DomainSnapshotOptionsPtrInput
	// A mapping of tags to assign to the resource
	Tags pulumi.MapInput
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
	AdvancedOptions map[string]interface{} `pulumi:"advancedOptions"`
	// Cluster configuration of the domain, see below.
	ClusterConfig *DomainClusterConfig `pulumi:"clusterConfig"`
	CognitoOptions *DomainCognitoOptions `pulumi:"cognitoOptions"`
	// Name of the domain.
	DomainName *string `pulumi:"domainName"`
	// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
	EbsOptions *DomainEbsOptions `pulumi:"ebsOptions"`
	// The version of Elasticsearch to deploy. Defaults to `1.5`
	ElasticsearchVersion *string `pulumi:"elasticsearchVersion"`
	// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
	EncryptAtRest *DomainEncryptAtRest `pulumi:"encryptAtRest"`
	// Options for publishing slow logs to CloudWatch Logs.
	LogPublishingOptions []DomainLogPublishingOption `pulumi:"logPublishingOptions"`
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption *DomainNodeToNodeEncryption `pulumi:"nodeToNodeEncryption"`
	// Snapshot related options, see below.
	SnapshotOptions *DomainSnapshotOptions `pulumi:"snapshotOptions"`
	// A mapping of tags to assign to the resource
	Tags map[string]interface{} `pulumi:"tags"`
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
	AdvancedOptions pulumi.MapInput
	// Cluster configuration of the domain, see below.
	ClusterConfig DomainClusterConfigPtrInput
	CognitoOptions DomainCognitoOptionsPtrInput
	// Name of the domain.
	DomainName pulumi.StringPtrInput
	// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
	EbsOptions DomainEbsOptionsPtrInput
	// The version of Elasticsearch to deploy. Defaults to `1.5`
	ElasticsearchVersion pulumi.StringPtrInput
	// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
	EncryptAtRest DomainEncryptAtRestPtrInput
	// Options for publishing slow logs to CloudWatch Logs.
	LogPublishingOptions DomainLogPublishingOptionArrayInput
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption DomainNodeToNodeEncryptionPtrInput
	// Snapshot related options, see below.
	SnapshotOptions DomainSnapshotOptionsPtrInput
	// A mapping of tags to assign to the resource
	Tags pulumi.MapInput
	// VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
	VpcOptions DomainVpcOptionsPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

