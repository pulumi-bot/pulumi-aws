// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an MQ Broker Resource. This resources also manages users for the broker.
//
// For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).
//
// Changes to an MQ Broker can occur when you change a
// parameter, such as `configuration` or `user`, and are reflected in the next maintenance
// window. Because of this, this provider may report a difference in its planning
// phase because a modification has not yet taken place. You can use the
// `applyImmediately` flag to instruct the service to apply the change immediately
// (see documentation below).
//
// > **Note:** using `applyImmediately` can result in a
// brief downtime as the broker reboots.
//
// > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/mq"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = mq.NewBroker(ctx, "example", &mq.BrokerArgs{
// 			BrokerName: pulumi.String("example"),
// 			Configuration: &mq.BrokerConfigurationArgs{
// 				Id:       pulumi.String(aws_mq_configuration.Test.Id),
// 				Revision: pulumi.Int(aws_mq_configuration.Test.Latest_revision),
// 			},
// 			EngineType:       pulumi.String("ActiveMQ"),
// 			EngineVersion:    pulumi.String("5.15.0"),
// 			HostInstanceType: pulumi.String("mq.t2.micro"),
// 			SecurityGroups: pulumi.StringArray{
// 				pulumi.String(aws_security_group.Test.Id),
// 			},
// 			Users: mq.BrokerUserArray{
// 				&mq.BrokerUserArgs{
// 					Password: pulumi.String("MindTheGap"),
// 					Username: pulumi.String("ExampleUser"),
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
type Broker struct {
	pulumi.CustomResourceState

	// Specifies whether any broker modifications
	// are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately pulumi.BoolPtrOutput `pulumi:"applyImmediately"`
	// The ARN of the broker.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
	AutoMinorVersionUpgrade pulumi.BoolPtrOutput `pulumi:"autoMinorVersionUpgrade"`
	// The name of the broker.
	BrokerName pulumi.StringOutput `pulumi:"brokerName"`
	// Configuration of the broker. See below.
	Configuration BrokerConfigurationOutput `pulumi:"configuration"`
	// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
	DeploymentMode pulumi.StringPtrOutput `pulumi:"deploymentMode"`
	// Configuration block containing encryption options. See below.
	EncryptionOptions BrokerEncryptionOptionsPtrOutput `pulumi:"encryptionOptions"`
	// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
	EngineType pulumi.StringOutput `pulumi:"engineType"`
	// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
	HostInstanceType pulumi.StringOutput `pulumi:"hostInstanceType"`
	// A list of information about allocated brokers (both active & standby).
	// * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
	// * `instances.0.ip_address` - The IP Address of the broker.
	// * `instances.0.endpoints` - The broker's wire-level protocol endpoints in the following order & format referenceable e.g. as `instances.0.endpoints.0` (SSL):
	// * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
	// * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
	// * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
	// * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
	// * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
	Instances BrokerInstanceArrayOutput `pulumi:"instances"`
	// Logging configuration of the broker. See below.
	Logs BrokerLogsPtrOutput `pulumi:"logs"`
	// Maintenance window start time. See below.
	MaintenanceWindowStartTime BrokerMaintenanceWindowStartTimeOutput `pulumi:"maintenanceWindowStartTime"`
	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible pulumi.BoolPtrOutput `pulumi:"publiclyAccessible"`
	// The list of security group IDs assigned to the broker.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The list of all ActiveMQ usernames for the specified broker. See below.
	Users BrokerUserArrayOutput `pulumi:"users"`
}

// NewBroker registers a new resource with the given unique name, arguments, and options.
func NewBroker(ctx *pulumi.Context,
	name string, args *BrokerArgs, opts ...pulumi.ResourceOption) (*Broker, error) {
	if args == nil || args.BrokerName == nil {
		return nil, errors.New("missing required argument 'BrokerName'")
	}
	if args == nil || args.EngineType == nil {
		return nil, errors.New("missing required argument 'EngineType'")
	}
	if args == nil || args.EngineVersion == nil {
		return nil, errors.New("missing required argument 'EngineVersion'")
	}
	if args == nil || args.HostInstanceType == nil {
		return nil, errors.New("missing required argument 'HostInstanceType'")
	}
	if args == nil || args.SecurityGroups == nil {
		return nil, errors.New("missing required argument 'SecurityGroups'")
	}
	if args == nil || args.Users == nil {
		return nil, errors.New("missing required argument 'Users'")
	}
	if args == nil {
		args = &BrokerArgs{}
	}
	var resource Broker
	err := ctx.RegisterResource("aws:mq/broker:Broker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBroker gets an existing Broker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBroker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BrokerState, opts ...pulumi.ResourceOption) (*Broker, error) {
	var resource Broker
	err := ctx.ReadResource("aws:mq/broker:Broker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Broker resources.
type brokerState struct {
	// Specifies whether any broker modifications
	// are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// The ARN of the broker.
	Arn *string `pulumi:"arn"`
	// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The name of the broker.
	BrokerName *string `pulumi:"brokerName"`
	// Configuration of the broker. See below.
	Configuration *BrokerConfiguration `pulumi:"configuration"`
	// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// Configuration block containing encryption options. See below.
	EncryptionOptions *BrokerEncryptionOptions `pulumi:"encryptionOptions"`
	// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
	EngineType *string `pulumi:"engineType"`
	// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
	EngineVersion *string `pulumi:"engineVersion"`
	// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
	HostInstanceType *string `pulumi:"hostInstanceType"`
	// A list of information about allocated brokers (both active & standby).
	// * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
	// * `instances.0.ip_address` - The IP Address of the broker.
	// * `instances.0.endpoints` - The broker's wire-level protocol endpoints in the following order & format referenceable e.g. as `instances.0.endpoints.0` (SSL):
	// * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
	// * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
	// * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
	// * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
	// * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
	Instances []BrokerInstance `pulumi:"instances"`
	// Logging configuration of the broker. See below.
	Logs *BrokerLogs `pulumi:"logs"`
	// Maintenance window start time. See below.
	MaintenanceWindowStartTime *BrokerMaintenanceWindowStartTime `pulumi:"maintenanceWindowStartTime"`
	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// The list of security group IDs assigned to the broker.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The list of all ActiveMQ usernames for the specified broker. See below.
	Users []BrokerUser `pulumi:"users"`
}

type BrokerState struct {
	// Specifies whether any broker modifications
	// are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately pulumi.BoolPtrInput
	// The ARN of the broker.
	Arn pulumi.StringPtrInput
	// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The name of the broker.
	BrokerName pulumi.StringPtrInput
	// Configuration of the broker. See below.
	Configuration BrokerConfigurationPtrInput
	// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
	DeploymentMode pulumi.StringPtrInput
	// Configuration block containing encryption options. See below.
	EncryptionOptions BrokerEncryptionOptionsPtrInput
	// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
	EngineType pulumi.StringPtrInput
	// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
	EngineVersion pulumi.StringPtrInput
	// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
	HostInstanceType pulumi.StringPtrInput
	// A list of information about allocated brokers (both active & standby).
	// * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
	// * `instances.0.ip_address` - The IP Address of the broker.
	// * `instances.0.endpoints` - The broker's wire-level protocol endpoints in the following order & format referenceable e.g. as `instances.0.endpoints.0` (SSL):
	// * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
	// * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
	// * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
	// * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
	// * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
	Instances BrokerInstanceArrayInput
	// Logging configuration of the broker. See below.
	Logs BrokerLogsPtrInput
	// Maintenance window start time. See below.
	MaintenanceWindowStartTime BrokerMaintenanceWindowStartTimePtrInput
	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible pulumi.BoolPtrInput
	// The list of security group IDs assigned to the broker.
	SecurityGroups pulumi.StringArrayInput
	// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
	// The list of all ActiveMQ usernames for the specified broker. See below.
	Users BrokerUserArrayInput
}

func (BrokerState) ElementType() reflect.Type {
	return reflect.TypeOf((*brokerState)(nil)).Elem()
}

type brokerArgs struct {
	// Specifies whether any broker modifications
	// are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The name of the broker.
	BrokerName string `pulumi:"brokerName"`
	// Configuration of the broker. See below.
	Configuration *BrokerConfiguration `pulumi:"configuration"`
	// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// Configuration block containing encryption options. See below.
	EncryptionOptions *BrokerEncryptionOptions `pulumi:"encryptionOptions"`
	// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
	EngineType string `pulumi:"engineType"`
	// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
	EngineVersion string `pulumi:"engineVersion"`
	// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
	HostInstanceType string `pulumi:"hostInstanceType"`
	// Logging configuration of the broker. See below.
	Logs *BrokerLogs `pulumi:"logs"`
	// Maintenance window start time. See below.
	MaintenanceWindowStartTime *BrokerMaintenanceWindowStartTime `pulumi:"maintenanceWindowStartTime"`
	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// The list of security group IDs assigned to the broker.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The list of all ActiveMQ usernames for the specified broker. See below.
	Users []BrokerUser `pulumi:"users"`
}

// The set of arguments for constructing a Broker resource.
type BrokerArgs struct {
	// Specifies whether any broker modifications
	// are applied immediately, or during the next maintenance window. Default is `false`.
	ApplyImmediately pulumi.BoolPtrInput
	// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The name of the broker.
	BrokerName pulumi.StringInput
	// Configuration of the broker. See below.
	Configuration BrokerConfigurationPtrInput
	// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
	DeploymentMode pulumi.StringPtrInput
	// Configuration block containing encryption options. See below.
	EncryptionOptions BrokerEncryptionOptionsPtrInput
	// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
	EngineType pulumi.StringInput
	// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
	EngineVersion pulumi.StringInput
	// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
	HostInstanceType pulumi.StringInput
	// Logging configuration of the broker. See below.
	Logs BrokerLogsPtrInput
	// Maintenance window start time. See below.
	MaintenanceWindowStartTime BrokerMaintenanceWindowStartTimePtrInput
	// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible pulumi.BoolPtrInput
	// The list of security group IDs assigned to the broker.
	SecurityGroups pulumi.StringArrayInput
	// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
	// The list of all ActiveMQ usernames for the specified broker. See below.
	Users BrokerUserArrayInput
}

func (BrokerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*brokerArgs)(nil)).Elem()
}
