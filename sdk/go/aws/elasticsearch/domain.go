// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Domain struct {
	s *pulumi.ResourceState
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOpt) (*Domain, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessPolicies"] = nil
		inputs["advancedOptions"] = nil
		inputs["clusterConfig"] = nil
		inputs["cognitoOptions"] = nil
		inputs["domainName"] = nil
		inputs["ebsOptions"] = nil
		inputs["elasticsearchVersion"] = nil
		inputs["encryptAtRest"] = nil
		inputs["logPublishingOptions"] = nil
		inputs["nodeToNodeEncryption"] = nil
		inputs["snapshotOptions"] = nil
		inputs["tags"] = nil
		inputs["vpcOptions"] = nil
	} else {
		inputs["accessPolicies"] = args.AccessPolicies
		inputs["advancedOptions"] = args.AdvancedOptions
		inputs["clusterConfig"] = args.ClusterConfig
		inputs["cognitoOptions"] = args.CognitoOptions
		inputs["domainName"] = args.DomainName
		inputs["ebsOptions"] = args.EbsOptions
		inputs["elasticsearchVersion"] = args.ElasticsearchVersion
		inputs["encryptAtRest"] = args.EncryptAtRest
		inputs["logPublishingOptions"] = args.LogPublishingOptions
		inputs["nodeToNodeEncryption"] = args.NodeToNodeEncryption
		inputs["snapshotOptions"] = args.SnapshotOptions
		inputs["tags"] = args.Tags
		inputs["vpcOptions"] = args.VpcOptions
	}
	inputs["arn"] = nil
	inputs["domainId"] = nil
	inputs["endpoint"] = nil
	inputs["kibanaEndpoint"] = nil
	s, err := ctx.RegisterResource("aws:elasticsearch/domain:Domain", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DomainState, opts ...pulumi.ResourceOpt) (*Domain, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessPolicies"] = state.AccessPolicies
		inputs["advancedOptions"] = state.AdvancedOptions
		inputs["arn"] = state.Arn
		inputs["clusterConfig"] = state.ClusterConfig
		inputs["cognitoOptions"] = state.CognitoOptions
		inputs["domainId"] = state.DomainId
		inputs["domainName"] = state.DomainName
		inputs["ebsOptions"] = state.EbsOptions
		inputs["elasticsearchVersion"] = state.ElasticsearchVersion
		inputs["encryptAtRest"] = state.EncryptAtRest
		inputs["endpoint"] = state.Endpoint
		inputs["kibanaEndpoint"] = state.KibanaEndpoint
		inputs["logPublishingOptions"] = state.LogPublishingOptions
		inputs["nodeToNodeEncryption"] = state.NodeToNodeEncryption
		inputs["snapshotOptions"] = state.SnapshotOptions
		inputs["tags"] = state.Tags
		inputs["vpcOptions"] = state.VpcOptions
	}
	s, err := ctx.ReadResource("aws:elasticsearch/domain:Domain", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Domain) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Domain) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// IAM policy document specifying the access policies for the domain
func (r *Domain) AccessPolicies() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessPolicies"])
}

// Key-value string pairs to specify advanced configuration options.
// Note that the values for these configuration options must be strings (wrapped in quotes) or they
// may be wrong and cause a perpetual diff, causing Terraform to want to recreate your Elasticsearch
// domain on every apply.
func (r *Domain) AdvancedOptions() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["advancedOptions"])
}

// Amazon Resource Name (ARN) of the domain.
func (r *Domain) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Cluster configuration of the domain, see below.
func (r *Domain) ClusterConfig() *pulumi.Output {
	return r.s.State["clusterConfig"]
}

func (r *Domain) CognitoOptions() *pulumi.Output {
	return r.s.State["cognitoOptions"]
}

// Unique identifier for the domain.
func (r *Domain) DomainId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["domainId"])
}

// Name of the domain.
func (r *Domain) DomainName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["domainName"])
}

// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
func (r *Domain) EbsOptions() *pulumi.Output {
	return r.s.State["ebsOptions"]
}

// The version of Elasticsearch to deploy. Defaults to `1.5`
func (r *Domain) ElasticsearchVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["elasticsearchVersion"])
}

// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
func (r *Domain) EncryptAtRest() *pulumi.Output {
	return r.s.State["encryptAtRest"]
}

// Domain-specific endpoint used to submit index, search, and data upload requests.
func (r *Domain) Endpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endpoint"])
}

// Domain-specific endpoint for kibana without https scheme.
// * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnet_ids` were created inside.
// * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
func (r *Domain) KibanaEndpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["kibanaEndpoint"])
}

// Options for publishing slow logs to CloudWatch Logs.
func (r *Domain) LogPublishingOptions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["logPublishingOptions"])
}

// Node-to-node encryption options. See below.
func (r *Domain) NodeToNodeEncryption() *pulumi.Output {
	return r.s.State["nodeToNodeEncryption"]
}

// Snapshot related options, see below.
func (r *Domain) SnapshotOptions() *pulumi.Output {
	return r.s.State["snapshotOptions"]
}

// A mapping of tags to assign to the resource
func (r *Domain) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
func (r *Domain) VpcOptions() *pulumi.Output {
	return r.s.State["vpcOptions"]
}

// Input properties used for looking up and filtering Domain resources.
type DomainState struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies interface{}
	// Key-value string pairs to specify advanced configuration options.
	// Note that the values for these configuration options must be strings (wrapped in quotes) or they
	// may be wrong and cause a perpetual diff, causing Terraform to want to recreate your Elasticsearch
	// domain on every apply.
	AdvancedOptions interface{}
	// Amazon Resource Name (ARN) of the domain.
	Arn interface{}
	// Cluster configuration of the domain, see below.
	ClusterConfig interface{}
	CognitoOptions interface{}
	// Unique identifier for the domain.
	DomainId interface{}
	// Name of the domain.
	DomainName interface{}
	// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
	EbsOptions interface{}
	// The version of Elasticsearch to deploy. Defaults to `1.5`
	ElasticsearchVersion interface{}
	// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
	EncryptAtRest interface{}
	// Domain-specific endpoint used to submit index, search, and data upload requests.
	Endpoint interface{}
	// Domain-specific endpoint for kibana without https scheme.
	// * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnet_ids` were created inside.
	// * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
	KibanaEndpoint interface{}
	// Options for publishing slow logs to CloudWatch Logs.
	LogPublishingOptions interface{}
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption interface{}
	// Snapshot related options, see below.
	SnapshotOptions interface{}
	// A mapping of tags to assign to the resource
	Tags interface{}
	// VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
	VpcOptions interface{}
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies interface{}
	// Key-value string pairs to specify advanced configuration options.
	// Note that the values for these configuration options must be strings (wrapped in quotes) or they
	// may be wrong and cause a perpetual diff, causing Terraform to want to recreate your Elasticsearch
	// domain on every apply.
	AdvancedOptions interface{}
	// Cluster configuration of the domain, see below.
	ClusterConfig interface{}
	CognitoOptions interface{}
	// Name of the domain.
	DomainName interface{}
	// EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
	EbsOptions interface{}
	// The version of Elasticsearch to deploy. Defaults to `1.5`
	ElasticsearchVersion interface{}
	// Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
	EncryptAtRest interface{}
	// Options for publishing slow logs to CloudWatch Logs.
	LogPublishingOptions interface{}
	// Node-to-node encryption options. See below.
	NodeToNodeEncryption interface{}
	// Snapshot related options, see below.
	SnapshotOptions interface{}
	// A mapping of tags to assign to the resource
	Tags interface{}
	// VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
	VpcOptions interface{}
}
