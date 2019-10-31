// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Mq
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides information about a MQ Broker.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/mq_broker.html.markdown.
        /// </summary>
        public static Task<GetBrokerResult> GetBroker(GetBrokerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBrokerResult>("aws:mq/getBroker:getBroker", args, options.WithVersion());
    }

    public sealed class GetBrokerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique id of the mq broker.
        /// </summary>
        [Input("brokerId")]
        public Input<string>? BrokerId { get; set; }

        /// <summary>
        /// The unique name of the mq broker.
        /// </summary>
        [Input("brokerName")]
        public Input<string>? BrokerName { get; set; }

        [Input("logs")]
        public Input<Inputs.GetBrokerLogsArgs>? Logs { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetBrokerArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetBrokerResult
    {
        public readonly string Arn;
        public readonly bool AutoMinorVersionUpgrade;
        public readonly string BrokerId;
        public readonly string BrokerName;
        public readonly Outputs.GetBrokerConfigurationResult Configuration;
        public readonly string DeploymentMode;
        public readonly ImmutableArray<Outputs.GetBrokerEncryptionOptionsResult> EncryptionOptions;
        public readonly string EngineType;
        public readonly string EngineVersion;
        public readonly string HostInstanceType;
        public readonly ImmutableArray<Outputs.GetBrokerInstancesResult> Instances;
        public readonly Outputs.GetBrokerLogsResult? Logs;
        public readonly Outputs.GetBrokerMaintenanceWindowStartTimeResult MaintenanceWindowStartTime;
        public readonly bool PubliclyAccessible;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly ImmutableArray<string> SubnetIds;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly ImmutableArray<Outputs.GetBrokerUsersResult> Users;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetBrokerResult(
            string arn,
            bool autoMinorVersionUpgrade,
            string brokerId,
            string brokerName,
            Outputs.GetBrokerConfigurationResult configuration,
            string deploymentMode,
            ImmutableArray<Outputs.GetBrokerEncryptionOptionsResult> encryptionOptions,
            string engineType,
            string engineVersion,
            string hostInstanceType,
            ImmutableArray<Outputs.GetBrokerInstancesResult> instances,
            Outputs.GetBrokerLogsResult? logs,
            Outputs.GetBrokerMaintenanceWindowStartTimeResult maintenanceWindowStartTime,
            bool publiclyAccessible,
            ImmutableArray<string> securityGroups,
            ImmutableArray<string> subnetIds,
            ImmutableDictionary<string, string> tags,
            ImmutableArray<Outputs.GetBrokerUsersResult> users,
            string id)
        {
            Arn = arn;
            AutoMinorVersionUpgrade = autoMinorVersionUpgrade;
            BrokerId = brokerId;
            BrokerName = brokerName;
            Configuration = configuration;
            DeploymentMode = deploymentMode;
            EncryptionOptions = encryptionOptions;
            EngineType = engineType;
            EngineVersion = engineVersion;
            HostInstanceType = hostInstanceType;
            Instances = instances;
            Logs = logs;
            MaintenanceWindowStartTime = maintenanceWindowStartTime;
            PubliclyAccessible = publiclyAccessible;
            SecurityGroups = securityGroups;
            SubnetIds = subnetIds;
            Tags = tags;
            Users = users;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetBrokerLogsArgs : Pulumi.ResourceArgs
    {
        [Input("audit")]
        public Input<bool>? Audit { get; set; }

        [Input("general")]
        public Input<bool>? General { get; set; }

        public GetBrokerLogsArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetBrokerConfigurationResult
    {
        public readonly string Id;
        public readonly int Revision;

        [OutputConstructor]
        private GetBrokerConfigurationResult(
            string id,
            int revision)
        {
            Id = id;
            Revision = revision;
        }
    }

    [OutputType]
    public sealed class GetBrokerEncryptionOptionsResult
    {
        public readonly string KmsKeyId;
        public readonly bool UseAwsOwnedKey;

        [OutputConstructor]
        private GetBrokerEncryptionOptionsResult(
            string kmsKeyId,
            bool useAwsOwnedKey)
        {
            KmsKeyId = kmsKeyId;
            UseAwsOwnedKey = useAwsOwnedKey;
        }
    }

    [OutputType]
    public sealed class GetBrokerInstancesResult
    {
        public readonly string ConsoleUrl;
        public readonly ImmutableArray<string> Endpoints;
        public readonly string IpAddress;

        [OutputConstructor]
        private GetBrokerInstancesResult(
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
    public sealed class GetBrokerLogsResult
    {
        public readonly bool Audit;
        public readonly bool General;

        [OutputConstructor]
        private GetBrokerLogsResult(
            bool audit,
            bool general)
        {
            Audit = audit;
            General = general;
        }
    }

    [OutputType]
    public sealed class GetBrokerMaintenanceWindowStartTimeResult
    {
        public readonly string DayOfWeek;
        public readonly string TimeOfDay;
        public readonly string TimeZone;

        [OutputConstructor]
        private GetBrokerMaintenanceWindowStartTimeResult(
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
    public sealed class GetBrokerUsersResult
    {
        public readonly bool ConsoleAccess;
        public readonly ImmutableArray<string> Groups;
        public readonly string Username;

        [OutputConstructor]
        private GetBrokerUsersResult(
            bool consoleAccess,
            ImmutableArray<string> groups,
            string username)
        {
            ConsoleAccess = consoleAccess;
            Groups = groups;
            Username = username;
        }
    }
    }
}
