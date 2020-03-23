// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    /// <summary>
    /// Provides a CloudWatch Event Target resource.
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_target.html.markdown.
    /// </summary>
    public partial class EventTarget : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) associated of the target.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("batchTarget")]
        public Output<Outputs.EventTargetBatchTarget?> BatchTarget { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("ecsTarget")]
        public Output<Outputs.EventTargetEcsTarget?> EcsTarget { get; private set; } = null!;

        /// <summary>
        /// Valid JSON text passed to the target.
        /// </summary>
        [Output("input")]
        public Output<string?> Input { get; private set; } = null!;

        /// <summary>
        /// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
        /// that is used for extracting part of the matched event when passing it to the target.
        /// </summary>
        [Output("inputPath")]
        public Output<string?> InputPath { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are providing a custom input to a target based on certain event data.
        /// </summary>
        [Output("inputTransformer")]
        public Output<Outputs.EventTargetInputTransformer?> InputTransformer { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("kinesisTarget")]
        public Output<Outputs.EventTargetKinesisTarget?> KinesisTarget { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The name of the rule you want to add targets to.
        /// </summary>
        [Output("rule")]
        public Output<string> Rule { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        /// </summary>
        [Output("runCommandTargets")]
        public Output<ImmutableArray<Outputs.EventTargetRunCommandTargets>> RunCommandTargets { get; private set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Output("sqsTarget")]
        public Output<Outputs.EventTargetSqsTarget?> SqsTarget { get; private set; } = null!;

        /// <summary>
        /// The unique target assignment ID.  If missing, will generate a random, unique id.
        /// </summary>
        [Output("targetId")]
        public Output<string> TargetId { get; private set; } = null!;


        /// <summary>
        /// Create a EventTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventTarget(string name, EventTargetArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventTarget:EventTarget", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private EventTarget(string name, Input<string> id, EventTargetState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventTarget:EventTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventTarget Get(string name, Input<string> id, EventTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new EventTarget(name, id, state, options);
        }
    }

    public sealed class EventTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) associated of the target.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("batchTarget")]
        public Input<Inputs.EventTargetBatchTargetArgs>? BatchTarget { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("ecsTarget")]
        public Input<Inputs.EventTargetEcsTargetArgs>? EcsTarget { get; set; }

        /// <summary>
        /// Valid JSON text passed to the target.
        /// </summary>
        [Input("input")]
        public Input<string>? Input { get; set; }

        /// <summary>
        /// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
        /// that is used for extracting part of the matched event when passing it to the target.
        /// </summary>
        [Input("inputPath")]
        public Input<string>? InputPath { get; set; }

        /// <summary>
        /// Parameters used when you are providing a custom input to a target based on certain event data.
        /// </summary>
        [Input("inputTransformer")]
        public Input<Inputs.EventTargetInputTransformerArgs>? InputTransformer { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("kinesisTarget")]
        public Input<Inputs.EventTargetKinesisTargetArgs>? KinesisTarget { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The name of the rule you want to add targets to.
        /// </summary>
        [Input("rule", required: true)]
        public Input<string> Rule { get; set; } = null!;

        [Input("runCommandTargets")]
        private InputList<Inputs.EventTargetRunCommandTargetsArgs>? _runCommandTargets;

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        /// </summary>
        public InputList<Inputs.EventTargetRunCommandTargetsArgs> RunCommandTargets
        {
            get => _runCommandTargets ?? (_runCommandTargets = new InputList<Inputs.EventTargetRunCommandTargetsArgs>());
            set => _runCommandTargets = value;
        }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("sqsTarget")]
        public Input<Inputs.EventTargetSqsTargetArgs>? SqsTarget { get; set; }

        /// <summary>
        /// The unique target assignment ID.  If missing, will generate a random, unique id.
        /// </summary>
        [Input("targetId")]
        public Input<string>? TargetId { get; set; }

        public EventTargetArgs()
        {
        }
    }

    public sealed class EventTargetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) associated of the target.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("batchTarget")]
        public Input<Inputs.EventTargetBatchTargetGetArgs>? BatchTarget { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("ecsTarget")]
        public Input<Inputs.EventTargetEcsTargetGetArgs>? EcsTarget { get; set; }

        /// <summary>
        /// Valid JSON text passed to the target.
        /// </summary>
        [Input("input")]
        public Input<string>? Input { get; set; }

        /// <summary>
        /// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
        /// that is used for extracting part of the matched event when passing it to the target.
        /// </summary>
        [Input("inputPath")]
        public Input<string>? InputPath { get; set; }

        /// <summary>
        /// Parameters used when you are providing a custom input to a target based on certain event data.
        /// </summary>
        [Input("inputTransformer")]
        public Input<Inputs.EventTargetInputTransformerGetArgs>? InputTransformer { get; set; }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("kinesisTarget")]
        public Input<Inputs.EventTargetKinesisTargetGetArgs>? KinesisTarget { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The name of the rule you want to add targets to.
        /// </summary>
        [Input("rule")]
        public Input<string>? Rule { get; set; }

        [Input("runCommandTargets")]
        private InputList<Inputs.EventTargetRunCommandTargetsGetArgs>? _runCommandTargets;

        /// <summary>
        /// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        /// </summary>
        public InputList<Inputs.EventTargetRunCommandTargetsGetArgs> RunCommandTargets
        {
            get => _runCommandTargets ?? (_runCommandTargets = new InputList<Inputs.EventTargetRunCommandTargetsGetArgs>());
            set => _runCommandTargets = value;
        }

        /// <summary>
        /// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        /// </summary>
        [Input("sqsTarget")]
        public Input<Inputs.EventTargetSqsTargetGetArgs>? SqsTarget { get; set; }

        /// <summary>
        /// The unique target assignment ID.  If missing, will generate a random, unique id.
        /// </summary>
        [Input("targetId")]
        public Input<string>? TargetId { get; set; }

        public EventTargetState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class EventTargetBatchTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
        /// </summary>
        [Input("arraySize")]
        public Input<int>? ArraySize { get; set; }

        /// <summary>
        /// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
        /// </summary>
        [Input("jobAttempts")]
        public Input<int>? JobAttempts { get; set; }

        /// <summary>
        /// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
        /// </summary>
        [Input("jobDefinition", required: true)]
        public Input<string> JobDefinition { get; set; } = null!;

        /// <summary>
        /// The name to use for this execution of the job, if the target is an AWS Batch job.
        /// </summary>
        [Input("jobName", required: true)]
        public Input<string> JobName { get; set; } = null!;

        public EventTargetBatchTargetArgs()
        {
        }
    }

    public sealed class EventTargetBatchTargetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
        /// </summary>
        [Input("arraySize")]
        public Input<int>? ArraySize { get; set; }

        /// <summary>
        /// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
        /// </summary>
        [Input("jobAttempts")]
        public Input<int>? JobAttempts { get; set; }

        /// <summary>
        /// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
        /// </summary>
        [Input("jobDefinition", required: true)]
        public Input<string> JobDefinition { get; set; } = null!;

        /// <summary>
        /// The name to use for this execution of the job, if the target is an AWS Batch job.
        /// </summary>
        [Input("jobName", required: true)]
        public Input<string> JobName { get; set; } = null!;

        public EventTargetBatchTargetGetArgs()
        {
        }
    }

    public sealed class EventTargetEcsTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies an ECS task group for the task. The maximum length is 255 characters.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. Valid values are EC2 or FARGATE.
        /// </summary>
        [Input("launchType")]
        public Input<string>? LaunchType { get; set; }

        /// <summary>
        /// Use this if the ECS task uses the awsvpc network mode. This specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. Required if launch_type is FARGATE because the awsvpc mode is required for Fargate tasks.
        /// </summary>
        [Input("networkConfiguration")]
        public Input<EventTargetEcsTargetNetworkConfigurationArgs>? NetworkConfiguration { get; set; }

        /// <summary>
        /// Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as 1.1.0. This is used only if LaunchType is FARGATE. For more information about valid platform versions, see [AWS Fargate Platform Versions](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
        /// </summary>
        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        /// <summary>
        /// The number of tasks to create based on the TaskDefinition. The default is 1.
        /// </summary>
        [Input("taskCount")]
        public Input<int>? TaskCount { get; set; }

        /// <summary>
        /// The ARN of the task definition to use if the event target is an Amazon ECS cluster.
        /// </summary>
        [Input("taskDefinitionArn", required: true)]
        public Input<string> TaskDefinitionArn { get; set; } = null!;

        public EventTargetEcsTargetArgs()
        {
        }
    }

    public sealed class EventTargetEcsTargetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies an ECS task group for the task. The maximum length is 255 characters.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. Valid values are EC2 or FARGATE.
        /// </summary>
        [Input("launchType")]
        public Input<string>? LaunchType { get; set; }

        /// <summary>
        /// Use this if the ECS task uses the awsvpc network mode. This specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. Required if launch_type is FARGATE because the awsvpc mode is required for Fargate tasks.
        /// </summary>
        [Input("networkConfiguration")]
        public Input<EventTargetEcsTargetNetworkConfigurationGetArgs>? NetworkConfiguration { get; set; }

        /// <summary>
        /// Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as 1.1.0. This is used only if LaunchType is FARGATE. For more information about valid platform versions, see [AWS Fargate Platform Versions](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
        /// </summary>
        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        /// <summary>
        /// The number of tasks to create based on the TaskDefinition. The default is 1.
        /// </summary>
        [Input("taskCount")]
        public Input<int>? TaskCount { get; set; }

        /// <summary>
        /// The ARN of the task definition to use if the event target is an Amazon ECS cluster.
        /// </summary>
        [Input("taskDefinitionArn", required: true)]
        public Input<string> TaskDefinitionArn { get; set; } = null!;

        public EventTargetEcsTargetGetArgs()
        {
        }
    }

    public sealed class EventTargetEcsTargetNetworkConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Assign a public IP address to the ENI (Fargate launch type only). Valid values are `true` or `false`. Default `false`.
        /// </summary>
        [Input("assignPublicIp")]
        public Input<bool>? AssignPublicIp { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnets", required: true)]
        private InputList<string>? _subnets;

        /// <summary>
        /// The subnets associated with the task or service.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        public EventTargetEcsTargetNetworkConfigurationArgs()
        {
        }
    }

    public sealed class EventTargetEcsTargetNetworkConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Assign a public IP address to the ENI (Fargate launch type only). Valid values are `true` or `false`. Default `false`.
        /// </summary>
        [Input("assignPublicIp")]
        public Input<bool>? AssignPublicIp { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnets", required: true)]
        private InputList<string>? _subnets;

        /// <summary>
        /// The subnets associated with the task or service.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        public EventTargetEcsTargetNetworkConfigurationGetArgs()
        {
        }
    }

    public sealed class EventTargetInputTransformerArgs : Pulumi.ResourceArgs
    {
        [Input("inputPaths")]
        private InputMap<object>? _inputPaths;

        /// <summary>
        /// Key value pairs specified in the form of JSONPath (for example, time = $.time)
        /// </summary>
        public InputMap<object> InputPaths
        {
            get => _inputPaths ?? (_inputPaths = new InputMap<object>());
            set => _inputPaths = value;
        }

        /// <summary>
        /// Structure containing the template body.
        /// </summary>
        [Input("inputTemplate", required: true)]
        public Input<string> InputTemplate { get; set; } = null!;

        public EventTargetInputTransformerArgs()
        {
        }
    }

    public sealed class EventTargetInputTransformerGetArgs : Pulumi.ResourceArgs
    {
        [Input("inputPaths")]
        private InputMap<object>? _inputPaths;

        /// <summary>
        /// Key value pairs specified in the form of JSONPath (for example, time = $.time)
        /// </summary>
        public InputMap<object> InputPaths
        {
            get => _inputPaths ?? (_inputPaths = new InputMap<object>());
            set => _inputPaths = value;
        }

        /// <summary>
        /// Structure containing the template body.
        /// </summary>
        [Input("inputTemplate", required: true)]
        public Input<string> InputTemplate { get; set; } = null!;

        public EventTargetInputTransformerGetArgs()
        {
        }
    }

    public sealed class EventTargetKinesisTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The JSON path to be extracted from the event and used as the partition key.
        /// </summary>
        [Input("partitionKeyPath")]
        public Input<string>? PartitionKeyPath { get; set; }

        public EventTargetKinesisTargetArgs()
        {
        }
    }

    public sealed class EventTargetKinesisTargetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The JSON path to be extracted from the event and used as the partition key.
        /// </summary>
        [Input("partitionKeyPath")]
        public Input<string>? PartitionKeyPath { get; set; }

        public EventTargetKinesisTargetGetArgs()
        {
        }
    }

    public sealed class EventTargetRunCommandTargetsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can be either `tag:tag-key` or `InstanceIds`.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// If Key is `tag:tag-key`, Values is a list of tag values. If Key is `InstanceIds`, Values is a list of Amazon EC2 instance IDs.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public EventTargetRunCommandTargetsArgs()
        {
        }
    }

    public sealed class EventTargetRunCommandTargetsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can be either `tag:tag-key` or `InstanceIds`.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// If Key is `tag:tag-key`, Values is a list of tag values. If Key is `InstanceIds`, Values is a list of Amazon EC2 instance IDs.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public EventTargetRunCommandTargetsGetArgs()
        {
        }
    }

    public sealed class EventTargetSqsTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The FIFO message group ID to use as the target.
        /// </summary>
        [Input("messageGroupId")]
        public Input<string>? MessageGroupId { get; set; }

        public EventTargetSqsTargetArgs()
        {
        }
    }

    public sealed class EventTargetSqsTargetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The FIFO message group ID to use as the target.
        /// </summary>
        [Input("messageGroupId")]
        public Input<string>? MessageGroupId { get; set; }

        public EventTargetSqsTargetGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EventTargetBatchTarget
    {
        /// <summary>
        /// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
        /// </summary>
        public readonly int? ArraySize;
        /// <summary>
        /// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
        /// </summary>
        public readonly int? JobAttempts;
        /// <summary>
        /// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
        /// </summary>
        public readonly string JobDefinition;
        /// <summary>
        /// The name to use for this execution of the job, if the target is an AWS Batch job.
        /// </summary>
        public readonly string JobName;

        [OutputConstructor]
        private EventTargetBatchTarget(
            int? arraySize,
            int? jobAttempts,
            string jobDefinition,
            string jobName)
        {
            ArraySize = arraySize;
            JobAttempts = jobAttempts;
            JobDefinition = jobDefinition;
            JobName = jobName;
        }
    }

    [OutputType]
    public sealed class EventTargetEcsTarget
    {
        /// <summary>
        /// Specifies an ECS task group for the task. The maximum length is 255 characters.
        /// </summary>
        public readonly string? Group;
        /// <summary>
        /// Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. Valid values are EC2 or FARGATE.
        /// </summary>
        public readonly string? LaunchType;
        /// <summary>
        /// Use this if the ECS task uses the awsvpc network mode. This specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. Required if launch_type is FARGATE because the awsvpc mode is required for Fargate tasks.
        /// </summary>
        public readonly EventTargetEcsTargetNetworkConfiguration? NetworkConfiguration;
        /// <summary>
        /// Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as 1.1.0. This is used only if LaunchType is FARGATE. For more information about valid platform versions, see [AWS Fargate Platform Versions](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
        /// </summary>
        public readonly string? PlatformVersion;
        /// <summary>
        /// The number of tasks to create based on the TaskDefinition. The default is 1.
        /// </summary>
        public readonly int? TaskCount;
        /// <summary>
        /// The ARN of the task definition to use if the event target is an Amazon ECS cluster.
        /// </summary>
        public readonly string TaskDefinitionArn;

        [OutputConstructor]
        private EventTargetEcsTarget(
            string? group,
            string? launchType,
            EventTargetEcsTargetNetworkConfiguration? networkConfiguration,
            string? platformVersion,
            int? taskCount,
            string taskDefinitionArn)
        {
            Group = group;
            LaunchType = launchType;
            NetworkConfiguration = networkConfiguration;
            PlatformVersion = platformVersion;
            TaskCount = taskCount;
            TaskDefinitionArn = taskDefinitionArn;
        }
    }

    [OutputType]
    public sealed class EventTargetEcsTargetNetworkConfiguration
    {
        /// <summary>
        /// Assign a public IP address to the ENI (Fargate launch type only). Valid values are `true` or `false`. Default `false`.
        /// </summary>
        public readonly bool? AssignPublicIp;
        /// <summary>
        /// The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// The subnets associated with the task or service.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;

        [OutputConstructor]
        private EventTargetEcsTargetNetworkConfiguration(
            bool? assignPublicIp,
            ImmutableArray<string> securityGroups,
            ImmutableArray<string> subnets)
        {
            AssignPublicIp = assignPublicIp;
            SecurityGroups = securityGroups;
            Subnets = subnets;
        }
    }

    [OutputType]
    public sealed class EventTargetInputTransformer
    {
        /// <summary>
        /// Key value pairs specified in the form of JSONPath (for example, time = $.time)
        /// </summary>
        public readonly ImmutableDictionary<string, object>? InputPaths;
        /// <summary>
        /// Structure containing the template body.
        /// </summary>
        public readonly string InputTemplate;

        [OutputConstructor]
        private EventTargetInputTransformer(
            ImmutableDictionary<string, object>? inputPaths,
            string inputTemplate)
        {
            InputPaths = inputPaths;
            InputTemplate = inputTemplate;
        }
    }

    [OutputType]
    public sealed class EventTargetKinesisTarget
    {
        /// <summary>
        /// The JSON path to be extracted from the event and used as the partition key.
        /// </summary>
        public readonly string? PartitionKeyPath;

        [OutputConstructor]
        private EventTargetKinesisTarget(string? partitionKeyPath)
        {
            PartitionKeyPath = partitionKeyPath;
        }
    }

    [OutputType]
    public sealed class EventTargetRunCommandTargets
    {
        /// <summary>
        /// Can be either `tag:tag-key` or `InstanceIds`.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// If Key is `tag:tag-key`, Values is a list of tag values. If Key is `InstanceIds`, Values is a list of Amazon EC2 instance IDs.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private EventTargetRunCommandTargets(
            string key,
            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }

    [OutputType]
    public sealed class EventTargetSqsTarget
    {
        /// <summary>
        /// The FIFO message group ID to use as the target.
        /// </summary>
        public readonly string? MessageGroupId;

        [OutputConstructor]
        private EventTargetSqsTarget(string? messageGroupId)
        {
            MessageGroupId = messageGroupId;
        }
    }
    }
}
