// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an RDS DB proxy resource. For additional information, see the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html).
//
// ## Import
//
// DB proxies can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:rds/proxy:Proxy example example
// ```
type Proxy struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) for the proxy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths ProxyAuthArrayOutput `pulumi:"auths"`
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging pulumi.BoolPtrOutput `pulumi:"debugLogging"`
	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily pulumi.StringOutput `pulumi:"engineFamily"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout pulumi.IntOutput `pulumi:"idleClientTimeout"`
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name pulumi.StringOutput `pulumi:"name"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls pulumi.BoolPtrOutput `pulumi:"requireTls"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds pulumi.StringArrayOutput `pulumi:"vpcSecurityGroupIds"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds pulumi.StringArrayOutput `pulumi:"vpcSubnetIds"`
}

// NewProxy registers a new resource with the given unique name, arguments, and options.
func NewProxy(ctx *pulumi.Context,
	name string, args *ProxyArgs, opts ...pulumi.ResourceOption) (*Proxy, error) {
	if args == nil || args.Auths == nil {
		return nil, errors.New("missing required argument 'Auths'")
	}
	if args == nil || args.EngineFamily == nil {
		return nil, errors.New("missing required argument 'EngineFamily'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil || args.VpcSubnetIds == nil {
		return nil, errors.New("missing required argument 'VpcSubnetIds'")
	}
	if args == nil {
		args = &ProxyArgs{}
	}
	var resource Proxy
	err := ctx.RegisterResource("aws:rds/proxy:Proxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProxy gets an existing Proxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProxyState, opts ...pulumi.ResourceOption) (*Proxy, error) {
	var resource Proxy
	err := ctx.ReadResource("aws:rds/proxy:Proxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Proxy resources.
type proxyState struct {
	// The Amazon Resource Name (ARN) for the proxy.
	Arn *string `pulumi:"arn"`
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths []ProxyAuth `pulumi:"auths"`
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging *bool `pulumi:"debugLogging"`
	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint *string `pulumi:"endpoint"`
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily *string `pulumi:"engineFamily"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout *int `pulumi:"idleClientTimeout"`
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name *string `pulumi:"name"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls *bool `pulumi:"requireTls"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn *string `pulumi:"roleArn"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds []string `pulumi:"vpcSubnetIds"`
}

type ProxyState struct {
	// The Amazon Resource Name (ARN) for the proxy.
	Arn pulumi.StringPtrInput
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths ProxyAuthArrayInput
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging pulumi.BoolPtrInput
	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint pulumi.StringPtrInput
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily pulumi.StringPtrInput
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout pulumi.IntPtrInput
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name pulumi.StringPtrInput
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds pulumi.StringArrayInput
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds pulumi.StringArrayInput
}

func (ProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyState)(nil)).Elem()
}

type proxyArgs struct {
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths []ProxyAuth `pulumi:"auths"`
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging *bool `pulumi:"debugLogging"`
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily string `pulumi:"engineFamily"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout *int `pulumi:"idleClientTimeout"`
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name *string `pulumi:"name"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls *bool `pulumi:"requireTls"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn string `pulumi:"roleArn"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds []string `pulumi:"vpcSubnetIds"`
}

// The set of arguments for constructing a Proxy resource.
type ProxyArgs struct {
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths ProxyAuthArrayInput
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging pulumi.BoolPtrInput
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily pulumi.StringInput
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout pulumi.IntPtrInput
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name pulumi.StringPtrInput
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds pulumi.StringArrayInput
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds pulumi.StringArrayInput
}

func (ProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyArgs)(nil)).Elem()
}

type ProxyInput interface {
	pulumi.Input

	ToProxyOutput() ProxyOutput
	ToProxyOutputWithContext(ctx context.Context) ProxyOutput
}

func (Proxy) ElementType() reflect.Type {
	return reflect.TypeOf((*Proxy)(nil)).Elem()
}

func (i Proxy) ToProxyOutput() ProxyOutput {
	return i.ToProxyOutputWithContext(context.Background())
}

func (i Proxy) ToProxyOutputWithContext(ctx context.Context) ProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyOutput)
}

type ProxyOutput struct {
	*pulumi.OutputState
}

func (ProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyOutput)(nil)).Elem()
}

func (o ProxyOutput) ToProxyOutput() ProxyOutput {
	return o
}

func (o ProxyOutput) ToProxyOutputWithContext(ctx context.Context) ProxyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProxyOutput{})
}
