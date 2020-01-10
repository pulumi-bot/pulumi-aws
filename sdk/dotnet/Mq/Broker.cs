// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Mq
{
    /// <summary>
    /// Provides an MQ Broker Resource. This resources also manages users for the broker.
    /// 
    /// For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).
    /// 
    /// Changes to an MQ Broker can occur when you change a
    /// parameter, such as `configuration` or `user`, and are reflected in the next maintenance
    /// window. Because of this, this provider may report a difference in its planning
    /// phase because a modification has not yet taken place. You can use the
    /// `apply_immediately` flag to instruct the service to apply the change immediately
    /// (see documentation below).
    /// 
    /// &gt; **Note:** using `apply_immediately` can result in a
    /// brief downtime as the broker reboots.
    /// 
    /// &gt; **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
    /// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/mq_broker.html.markdown.
    /// </summary>
    public partial class Broker : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether any broker modifications
        /// are applied immediately, or during the next maintenance window. Default is `false`.
        /// </summary>
        [Output("applyImmediately")]
        public Output<bool?> ApplyImmediately { get; private set; } = null!;

        /// <summary>
        /// The ARN of the broker.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
        /// </summary>
        [Output("autoMinorVersionUpgrade")]
        public Output<bool?> AutoMinorVersionUpgrade { get; private set; } = null!;

        /// <summary>
        /// The name of the broker.
        /// </summary>
        [Output("brokerName")]
        public Output<string> BrokerName { get; private set; } = null!;

        /// <summary>
        /// Configuration of the broker. See below.
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.BrokerConfiguration> Configuration { get; private set; } = null!;

        /// <summary>
        /// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
        /// </summary>
        [Output("deploymentMode")]
        public Output<string?> DeploymentMode { get; private set; } = null!;

        /// <summary>
        /// Configuration block containing encryption options. See below.
        /// </summary>
        [Output("encryptionOptions")]
        public Output<Outputs.BrokerEncryptionOptions?> EncryptionOptions { get; private set; } = null!;

        /// <summary>
        /// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
        /// </summary>
        [Output("engineType")]
        public Output<string> EngineType { get; private set; } = null!;

        /// <summary>
        /// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
        /// </summary>
        [Output("hostInstanceType")]
        public Output<string> HostInstanceType { get; private set; } = null!;

        /// <summary>
        /// A list of information about allocated brokers (both active &amp; standby).
        /// * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
        /// * `instances.0.ip_address` - The IP Address of the broker.
        /// * `instances.0.endpoints` - The broker's wire-level protocol endpoints in the following order &amp; format referenceable e.g. as `instances.0.endpoints.0` (SSL):
        /// * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
        /// * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
        /// * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
        /// * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
        /// * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<Outputs.BrokerInstances>> Instances { get; private set; } = null!;

        /// <summary>
        /// Logging configuration of the broker. See below.
        /// </summary>
        [Output("logs")]
        public Output<Outputs.BrokerLogs?> Logs { get; private set; } = null!;

        /// <summary>
        /// Maintenance window start time. See below.
        /// </summary>
        [Output("maintenanceWindowStartTime")]
        public Output<Outputs.BrokerMaintenanceWindowStartTime> MaintenanceWindowStartTime { get; private set; } = null!;

        /// <summary>
        /// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
        /// </summary>
        [Output("publiclyAccessible")]
        public Output<bool?> PubliclyAccessible { get; private set; } = null!;

        /// <summary>
        /// The list of security group IDs assigned to the broker.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The list of all ActiveMQ usernames for the specified broker. See below.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.BrokerUsers>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a Broker resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Broker(string name, BrokerArgs args, CustomResourceOptions? options = null)
            : base("aws:mq/broker:Broker", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Broker(string name, Input<string> id, BrokerState? state = null, CustomResourceOptions? options = null)
            : base("aws:mq/broker:Broker", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Broker resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Broker Get(string name, Input<string> id, BrokerState? state = null, CustomResourceOptions? options = null)
        {
            return new Broker(name, id, state, options);
        }
    }

    public sealed class BrokerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether any broker modifications
        /// are applied immediately, or during the next maintenance window. Default is `false`.
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        /// <summary>
        /// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
        /// </summary>
        [Input("autoMinorVersionUpgrade")]
        public Input<bool>? AutoMinorVersionUpgrade { get; set; }

        /// <summary>
        /// The name of the broker.
        /// </summary>
        [Input("brokerName", required: true)]
        public Input<string> BrokerName { get; set; } = null!;

        /// <summary>
        /// Configuration of the broker. See below.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.BrokerConfigurationArgs>? Configuration { get; set; }

        /// <summary>
        /// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
        /// </summary>
        [Input("deploymentMode")]
        public Input<string>? DeploymentMode { get; set; }

        /// <summary>
        /// Configuration block containing encryption options. See below.
        /// </summary>
        [Input("encryptionOptions")]
        public Input<Inputs.BrokerEncryptionOptionsArgs>? EncryptionOptions { get; set; }

        /// <summary>
        /// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
        /// </summary>
        [Input("engineType", required: true)]
        public Input<string> EngineType { get; set; } = null!;

        /// <summary>
        /// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
        /// </summary>
        [Input("engineVersion", required: true)]
        public Input<string> EngineVersion { get; set; } = null!;

        /// <summary>
        /// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
        /// </summary>
        [Input("hostInstanceType", required: true)]
        public Input<string> HostInstanceType { get; set; } = null!;

        /// <summary>
        /// Logging configuration of the broker. See below.
        /// </summary>
        [Input("logs")]
        public Input<Inputs.BrokerLogsArgs>? Logs { get; set; }

        /// <summary>
        /// Maintenance window start time. See below.
        /// </summary>
        [Input("maintenanceWindowStartTime")]
        public Input<Inputs.BrokerMaintenanceWindowStartTimeArgs>? MaintenanceWindowStartTime { get; set; }

        /// <summary>
        /// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
        /// </summary>
        [Input("publiclyAccessible")]
        public Input<bool>? PubliclyAccessible { get; set; }

        [Input("securityGroups", required: true)]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// The list of security group IDs assigned to the broker.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("users", required: true)]
        private InputList<Inputs.BrokerUsersArgs>? _users;

        /// <summary>
        /// The list of all ActiveMQ usernames for the specified broker. See below.
        /// </summary>
        public InputList<Inputs.BrokerUsersArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.BrokerUsersArgs>());
            set => _users = value;
        }

        public BrokerArgs()
        {
        }
    }

    public sealed class BrokerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether any broker modifications
        /// are applied immediately, or during the next maintenance window. Default is `false`.
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        /// <summary>
        /// The ARN of the broker.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
        /// </summary>
        [Input("autoMinorVersionUpgrade")]
        public Input<bool>? AutoMinorVersionUpgrade { get; set; }

        /// <summary>
        /// The name of the broker.
        /// </summary>
        [Input("brokerName")]
        public Input<string>? BrokerName { get; set; }

        /// <summary>
        /// Configuration of the broker. See below.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.BrokerConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
        /// </summary>
        [Input("deploymentMode")]
        public Input<string>? DeploymentMode { get; set; }

        /// <summary>
        /// Configuration block containing encryption options. See below.
        /// </summary>
        [Input("encryptionOptions")]
        public Input<Inputs.BrokerEncryptionOptionsGetArgs>? EncryptionOptions { get; set; }

        /// <summary>
        /// The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
        /// </summary>
        [Input("engineType")]
        public Input<string>? EngineType { get; set; }

        /// <summary>
        /// The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
        /// </summary>
        [Input("hostInstanceType")]
        public Input<string>? HostInstanceType { get; set; }

        [Input("instances")]
        private InputList<Inputs.BrokerInstancesGetArgs>? _instances;

        /// <summary>
        /// A list of information about allocated brokers (both active &amp; standby).
        /// * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
        /// * `instances.0.ip_address` - The IP Address of the broker.
        /// * `instances.0.endpoints` - The broker's wire-level protocol endpoints in the following order &amp; format referenceable e.g. as `instances.0.endpoints.0` (SSL):
        /// * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
        /// * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
        /// * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
        /// * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
        /// * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
        /// </summary>
        public InputList<Inputs.BrokerInstancesGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.BrokerInstancesGetArgs>());
            set => _instances = value;
        }

        /// <summary>
        /// Logging configuration of the broker. See below.
        /// </summary>
        [Input("logs")]
        public Input<Inputs.BrokerLogsGetArgs>? Logs { get; set; }

        /// <summary>
        /// Maintenance window start time. See below.
        /// </summary>
        [Input("maintenanceWindowStartTime")]
        public Input<Inputs.BrokerMaintenanceWindowStartTimeGetArgs>? MaintenanceWindowStartTime { get; set; }

        /// <summary>
        /// Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
        /// </summary>
        [Input("publiclyAccessible")]
        public Input<bool>? PubliclyAccessible { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// The list of security group IDs assigned to the broker.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("users")]
        private InputList<Inputs.BrokerUsersGetArgs>? _users;

        /// <summary>
        /// The list of all ActiveMQ usernames for the specified broker. See below.
        /// </summary>
        public InputList<Inputs.BrokerUsersGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.BrokerUsersGetArgs>());
            set => _users = value;
        }

        public BrokerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class BrokerConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Configuration ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Revision of the Configuration.
        /// </summary>
        [Input("revision")]
        public Input<int>? Revision { get; set; }

        public BrokerConfigurationArgs()
        {
        }
    }

    public sealed class BrokerConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Configuration ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Revision of the Configuration.
        /// </summary>
        [Input("revision")]
        public Input<int>? Revision { get; set; }

        public BrokerConfigurationGetArgs()
        {
        }
    }

    public sealed class BrokerEncryptionOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting `use_aws_owned_key` to `false`. To perform drift detection when AWS managed CMKs or customer managed CMKs are in use, this value must be configured.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Boolean to enable an AWS owned Key Management Service (KMS) Customer Master Key (CMK) that is not in your account. Defaults to `true`. Setting to `false` without configuring `kms_key_id` will create an AWS managed Customer Master Key (CMK) aliased to `aws/mq` in your account.
        /// </summary>
        [Input("useAwsOwnedKey")]
        public Input<bool>? UseAwsOwnedKey { get; set; }

        public BrokerEncryptionOptionsArgs()
        {
        }
    }

    public sealed class BrokerEncryptionOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting `use_aws_owned_key` to `false`. To perform drift detection when AWS managed CMKs or customer managed CMKs are in use, this value must be configured.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Boolean to enable an AWS owned Key Management Service (KMS) Customer Master Key (CMK) that is not in your account. Defaults to `true`. Setting to `false` without configuring `kms_key_id` will create an AWS managed Customer Master Key (CMK) aliased to `aws/mq` in your account.
        /// </summary>
        [Input("useAwsOwnedKey")]
        public Input<bool>? UseAwsOwnedKey { get; set; }

        public BrokerEncryptionOptionsGetArgs()
        {
        }
    }

    public sealed class BrokerInstancesGetArgs : Pulumi.ResourceArgs
    {
        [Input("consoleUrl")]
        public Input<string>? ConsoleUrl { get; set; }

        [Input("endpoints")]
        private InputList<string>? _endpoints;
        public InputList<string> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<string>());
            set => _endpoints = value;
        }

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        public BrokerInstancesGetArgs()
        {
        }
    }

    public sealed class BrokerLogsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables audit logging. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
        /// </summary>
        [Input("audit")]
        public Input<bool>? Audit { get; set; }

        /// <summary>
        /// Enables general logging via CloudWatch. Defaults to `false`.
        /// </summary>
        [Input("general")]
        public Input<bool>? General { get; set; }

        public BrokerLogsArgs()
        {
        }
    }

    public sealed class BrokerLogsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables audit logging. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
        /// </summary>
        [Input("audit")]
        public Input<bool>? Audit { get; set; }

        /// <summary>
        /// Enables general logging via CloudWatch. Defaults to `false`.
        /// </summary>
        [Input("general")]
        public Input<bool>? General { get; set; }

        public BrokerLogsGetArgs()
        {
        }
    }

    public sealed class BrokerMaintenanceWindowStartTimeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The day of the week. e.g. `MONDAY`, `TUESDAY`, or `WEDNESDAY`
        /// </summary>
        [Input("dayOfWeek", required: true)]
        public Input<string> DayOfWeek { get; set; } = null!;

        /// <summary>
        /// The time, in 24-hour format. e.g. `02:00`
        /// </summary>
        [Input("timeOfDay", required: true)]
        public Input<string> TimeOfDay { get; set; } = null!;

        /// <summary>
        /// The time zone, UTC by default, in either the Country/City format, or the UTC offset format. e.g. `CET`
        /// </summary>
        [Input("timeZone", required: true)]
        public Input<string> TimeZone { get; set; } = null!;

        public BrokerMaintenanceWindowStartTimeArgs()
        {
        }
    }

    public sealed class BrokerMaintenanceWindowStartTimeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The day of the week. e.g. `MONDAY`, `TUESDAY`, or `WEDNESDAY`
        /// </summary>
        [Input("dayOfWeek", required: true)]
        public Input<string> DayOfWeek { get; set; } = null!;

        /// <summary>
        /// The time, in 24-hour format. e.g. `02:00`
        /// </summary>
        [Input("timeOfDay", required: true)]
        public Input<string> TimeOfDay { get; set; } = null!;

        /// <summary>
        /// The time zone, UTC by default, in either the Country/City format, or the UTC offset format. e.g. `CET`
        /// </summary>
        [Input("timeZone", required: true)]
        public Input<string> TimeZone { get; set; } = null!;

        public BrokerMaintenanceWindowStartTimeGetArgs()
        {
        }
    }

    public sealed class BrokerUsersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user.
        /// </summary>
        [Input("consoleAccess")]
        public Input<bool>? ConsoleAccess { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// The list of groups (20 maximum) to which the ActiveMQ user belongs.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// The password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The username of the user.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public BrokerUsersArgs()
        {
        }
    }

    public sealed class BrokerUsersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user.
        /// </summary>
        [Input("consoleAccess")]
        public Input<bool>? ConsoleAccess { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// The list of groups (20 maximum) to which the ActiveMQ user belongs.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// The password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The username of the user.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public BrokerUsersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class BrokerConfiguration
    {
        /// <summary>
        /// The Configuration ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Revision of the Configuration.
        /// </summary>
        public readonly int Revision;

        [OutputConstructor]
        private BrokerConfiguration(
            string id,
            int revision)
        {
            Id = id;
            Revision = revision;
        }
    }

    [OutputType]
    public sealed class BrokerEncryptionOptions
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of Key Management Service (KMS) Customer Master Key (CMK) to use for encryption at rest. Requires setting `use_aws_owned_key` to `false`. To perform drift detection when AWS managed CMKs or customer managed CMKs are in use, this value must be configured.
        /// </summary>
        public readonly string KmsKeyId;
        /// <summary>
        /// Boolean to enable an AWS owned Key Management Service (KMS) Customer Master Key (CMK) that is not in your account. Defaults to `true`. Setting to `false` without configuring `kms_key_id` will create an AWS managed Customer Master Key (CMK) aliased to `aws/mq` in your account.
        /// </summary>
        public readonly bool? UseAwsOwnedKey;

        [OutputConstructor]
        private BrokerEncryptionOptions(
            string kmsKeyId,
            bool? useAwsOwnedKey)
        {
            KmsKeyId = kmsKeyId;
            UseAwsOwnedKey = useAwsOwnedKey;
        }
    }

    [OutputType]
    public sealed class BrokerInstances
    {
        public readonly string ConsoleUrl;
        public readonly ImmutableArray<string> Endpoints;
        public readonly string IpAddress;

        [OutputConstructor]
        private BrokerInstances(
            string consoleUrl,
            ImmutableArray<string> endpoints,
            string ipAddress)
        {
            ConsoleUrl = consoleUrl;
            Endpoints = endpoints;
            IpAddress = ipAddress;
        }
    }

    [OutputType]
    public sealed class BrokerLogs
    {
        /// <summary>
        /// Enables audit logging. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
        /// </summary>
        public readonly bool? Audit;
        /// <summary>
        /// Enables general logging via CloudWatch. Defaults to `false`.
        /// </summary>
        public readonly bool? General;

        [OutputConstructor]
        private BrokerLogs(
            bool? audit,
            bool? general)
        {
            Audit = audit;
            General = general;
        }
    }

    [OutputType]
    public sealed class BrokerMaintenanceWindowStartTime
    {
        /// <summary>
        /// The day of the week. e.g. `MONDAY`, `TUESDAY`, or `WEDNESDAY`
        /// </summary>
        public readonly string DayOfWeek;
        /// <summary>
        /// The time, in 24-hour format. e.g. `02:00`
        /// </summary>
        public readonly string TimeOfDay;
        /// <summary>
        /// The time zone, UTC by default, in either the Country/City format, or the UTC offset format. e.g. `CET`
        /// </summary>
        public readonly string TimeZone;

        [OutputConstructor]
        private BrokerMaintenanceWindowStartTime(
            string dayOfWeek,
            string timeOfDay,
            string timeZone)
        {
            DayOfWeek = dayOfWeek;
            TimeOfDay = timeOfDay;
            TimeZone = timeZone;
        }
    }

    [OutputType]
    public sealed class BrokerUsers
    {
        /// <summary>
        /// Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user.
        /// </summary>
        public readonly bool? ConsoleAccess;
        /// <summary>
        /// The list of groups (20 maximum) to which the ActiveMQ user belongs.
        /// </summary>
        public readonly ImmutableArray<string> Groups;
        /// <summary>
        /// The password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The username of the user.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private BrokerUsers(
            bool? consoleAccess,
            ImmutableArray<string> groups,
            string password,
            string username)
        {
            ConsoleAccess = consoleAccess;
            Groups = groups;
            Password = password;
            Username = username;
        }
    }
    }
}
