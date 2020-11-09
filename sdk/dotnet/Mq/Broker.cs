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
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Mq.Broker("example", new Aws.Mq.BrokerArgs
    ///         {
    ///             BrokerName = "example",
    ///             Configuration = new Aws.Mq.Inputs.BrokerConfigurationArgs
    ///             {
    ///                 Id = aws_mq_configuration.Test.Id,
    ///                 Revision = aws_mq_configuration.Test.Latest_revision,
    ///             },
    ///             EngineType = "ActiveMQ",
    ///             EngineVersion = "5.15.0",
    ///             HostInstanceType = "mq.t2.micro",
    ///             SecurityGroups = 
    ///             {
    ///                 aws_security_group.Test.Id,
    ///             },
    ///             Users = 
    ///             {
    ///                 new Aws.Mq.Inputs.BrokerUserArgs
    ///                 {
    ///                     Username = "ExampleUser",
    ///                     Password = "MindTheGap",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// MQ Brokers can be imported using their broker id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:mq/broker:Broker example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
    /// ```
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
        public Output<ImmutableArray<Outputs.BrokerInstance>> Instances { get; private set; } = null!;

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
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The list of all ActiveMQ usernames for the specified broker. See below.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.BrokerUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a Broker resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Broker(string name, BrokerArgs args, CustomResourceOptions? options = null)
            : base("aws:mq/broker:Broker", name, args ?? new BrokerArgs(), MakeResourceOptions(options, ""))
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
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("users", required: true)]
        private InputList<Inputs.BrokerUserArgs>? _users;

        /// <summary>
        /// The list of all ActiveMQ usernames for the specified broker. See below.
        /// </summary>
        public InputList<Inputs.BrokerUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.BrokerUserArgs>());
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
        private InputList<Inputs.BrokerInstanceGetArgs>? _instances;

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
        public InputList<Inputs.BrokerInstanceGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.BrokerInstanceGetArgs>());
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
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("users")]
        private InputList<Inputs.BrokerUserGetArgs>? _users;

        /// <summary>
        /// The list of all ActiveMQ usernames for the specified broker. See below.
        /// </summary>
        public InputList<Inputs.BrokerUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.BrokerUserGetArgs>());
            set => _users = value;
        }

        public BrokerState()
        {
        }
    }
}
