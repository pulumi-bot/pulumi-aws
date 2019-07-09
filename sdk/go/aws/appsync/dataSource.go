// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AppSync DataSource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_datasource.html.markdown.
type DataSource struct {
	s *pulumi.ResourceState
}

// NewDataSource registers a new resource with the given unique name, arguments, and options.
func NewDataSource(ctx *pulumi.Context,
	name string, args *DataSourceArgs, opts ...pulumi.ResourceOpt) (*DataSource, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiId"] = nil
		inputs["description"] = nil
		inputs["dynamodbConfig"] = nil
		inputs["elasticsearchConfig"] = nil
		inputs["httpConfig"] = nil
		inputs["lambdaConfig"] = nil
		inputs["name"] = nil
		inputs["serviceRoleArn"] = nil
		inputs["type"] = nil
	} else {
		inputs["apiId"] = args.ApiId
		inputs["description"] = args.Description
		inputs["dynamodbConfig"] = args.DynamodbConfig
		inputs["elasticsearchConfig"] = args.ElasticsearchConfig
		inputs["httpConfig"] = args.HttpConfig
		inputs["lambdaConfig"] = args.LambdaConfig
		inputs["name"] = args.Name
		inputs["serviceRoleArn"] = args.ServiceRoleArn
		inputs["type"] = args.Type
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:appsync/dataSource:DataSource", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DataSource{s: s}, nil
}

// GetDataSource gets an existing DataSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSource(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DataSourceState, opts ...pulumi.ResourceOpt) (*DataSource, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiId"] = state.ApiId
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["dynamodbConfig"] = state.DynamodbConfig
		inputs["elasticsearchConfig"] = state.ElasticsearchConfig
		inputs["httpConfig"] = state.HttpConfig
		inputs["lambdaConfig"] = state.LambdaConfig
		inputs["name"] = state.Name
		inputs["serviceRoleArn"] = state.ServiceRoleArn
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("aws:appsync/dataSource:DataSource", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DataSource{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DataSource) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DataSource) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The API ID for the GraphQL API for the DataSource.
func (r *DataSource) ApiId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["apiId"])
}

// The ARN
func (r *DataSource) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// A description of the DataSource.
func (r *DataSource) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// DynamoDB settings. See below
func (r *DataSource) DynamodbConfig() *pulumi.Output {
	return r.s.State["dynamodbConfig"]
}

// Amazon Elasticsearch settings. See below
func (r *DataSource) ElasticsearchConfig() *pulumi.Output {
	return r.s.State["elasticsearchConfig"]
}

// HTTP settings. See below
func (r *DataSource) HttpConfig() *pulumi.Output {
	return r.s.State["httpConfig"]
}

// AWS Lambda settings. See below
func (r *DataSource) LambdaConfig() *pulumi.Output {
	return r.s.State["lambdaConfig"]
}

// A user-supplied name for the DataSource.
func (r *DataSource) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The IAM service role ARN for the data source.
func (r *DataSource) ServiceRoleArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceRoleArn"])
}

// The type of the DataSource. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`.
func (r *DataSource) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering DataSource resources.
type DataSourceState struct {
	// The API ID for the GraphQL API for the DataSource.
	ApiId interface{}
	// The ARN
	Arn interface{}
	// A description of the DataSource.
	Description interface{}
	// DynamoDB settings. See below
	DynamodbConfig interface{}
	// Amazon Elasticsearch settings. See below
	ElasticsearchConfig interface{}
	// HTTP settings. See below
	HttpConfig interface{}
	// AWS Lambda settings. See below
	LambdaConfig interface{}
	// A user-supplied name for the DataSource.
	Name interface{}
	// The IAM service role ARN for the data source.
	ServiceRoleArn interface{}
	// The type of the DataSource. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`.
	Type interface{}
}

// The set of arguments for constructing a DataSource resource.
type DataSourceArgs struct {
	// The API ID for the GraphQL API for the DataSource.
	ApiId interface{}
	// A description of the DataSource.
	Description interface{}
	// DynamoDB settings. See below
	DynamodbConfig interface{}
	// Amazon Elasticsearch settings. See below
	ElasticsearchConfig interface{}
	// HTTP settings. See below
	HttpConfig interface{}
	// AWS Lambda settings. See below
	LambdaConfig interface{}
	// A user-supplied name for the DataSource.
	Name interface{}
	// The IAM service role ARN for the data source.
	ServiceRoleArn interface{}
	// The type of the DataSource. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`.
	Type interface{}
}
