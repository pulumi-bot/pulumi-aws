// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeDeploy
{
    /// <summary>
    /// Provides a CodeDeploy deployment config for an application
    /// 
    /// ## Example Usage
    /// ### Server Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fooDeploymentConfig = new Aws.CodeDeploy.DeploymentConfig("fooDeploymentConfig", new Aws.CodeDeploy.DeploymentConfigArgs
    ///         {
    ///             DeploymentConfigName = "test-deployment-config",
    ///             MinimumHealthyHosts = new Aws.CodeDeploy.Inputs.DeploymentConfigMinimumHealthyHostsArgs
    ///             {
    ///                 Type = "HOST_COUNT",
    ///                 Value = 2,
    ///             },
    ///         });
    ///         var fooDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("fooDeploymentGroup", new Aws.CodeDeploy.DeploymentGroupArgs
    ///         {
    ///             AppName = aws_codedeploy_app.Foo_app.Name,
    ///             DeploymentGroupName = "bar",
    ///             ServiceRoleArn = aws_iam_role.Foo_role.Arn,
    ///             DeploymentConfigName = fooDeploymentConfig.Id,
    ///             Ec2TagFilters = 
    ///             {
    ///                 new Aws.CodeDeploy.Inputs.DeploymentGroupEc2TagFilterArgs
    ///                 {
    ///                     Key = "filterkey",
    ///                     Type = "KEY_AND_VALUE",
    ///                     Value = "filtervalue",
    ///                 },
    ///             },
    ///             TriggerConfigurations = 
    ///             {
    ///                 new Aws.CodeDeploy.Inputs.DeploymentGroupTriggerConfigurationArgs
    ///                 {
    ///                     TriggerEvents = 
    ///                     {
    ///                         "DeploymentFailure",
    ///                     },
    ///                     TriggerName = "foo-trigger",
    ///                     TriggerTargetArn = "foo-topic-arn",
    ///                 },
    ///             },
    ///             AutoRollbackConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAutoRollbackConfigurationArgs
    ///             {
    ///                 Enabled = true,
    ///                 Events = 
    ///                 {
    ///                     "DEPLOYMENT_FAILURE",
    ///                 },
    ///             },
    ///             AlarmConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAlarmConfigurationArgs
    ///             {
    ///                 Alarms = 
    ///                 {
    ///                     "my-alarm-name",
    ///                 },
    ///                 Enabled = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Lambda Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fooDeploymentConfig = new Aws.CodeDeploy.DeploymentConfig("fooDeploymentConfig", new Aws.CodeDeploy.DeploymentConfigArgs
    ///         {
    ///             DeploymentConfigName = "test-deployment-config",
    ///             ComputePlatform = "Lambda",
    ///             TrafficRoutingConfig = new Aws.CodeDeploy.Inputs.DeploymentConfigTrafficRoutingConfigArgs
    ///             {
    ///                 Type = "TimeBasedLinear",
    ///                 TimeBasedLinear = new Aws.CodeDeploy.Inputs.DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs
    ///                 {
    ///                     Interval = 10,
    ///                     Percentage = 10,
    ///                 },
    ///             },
    ///         });
    ///         var fooDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("fooDeploymentGroup", new Aws.CodeDeploy.DeploymentGroupArgs
    ///         {
    ///             AppName = aws_codedeploy_app.Foo_app.Name,
    ///             DeploymentGroupName = "bar",
    ///             ServiceRoleArn = aws_iam_role.Foo_role.Arn,
    ///             DeploymentConfigName = fooDeploymentConfig.Id,
    ///             AutoRollbackConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAutoRollbackConfigurationArgs
    ///             {
    ///                 Enabled = true,
    ///                 Events = 
    ///                 {
    ///                     "DEPLOYMENT_STOP_ON_ALARM",
    ///                 },
    ///             },
    ///             AlarmConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAlarmConfigurationArgs
    ///             {
    ///                 Alarms = 
    ///                 {
    ///                     "my-alarm-name",
    ///                 },
    ///                 Enabled = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CodeDeploy Deployment Configurations can be imported using the `deployment_config_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:codedeploy/deploymentConfig:DeploymentConfig example my-deployment-config
    /// ```
    /// </summary>
    [AwsResourceType("aws:codedeploy/deploymentConfig:DeploymentConfig")]
    public partial class DeploymentConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
        /// </summary>
        [Output("computePlatform")]
        public Output<string?> ComputePlatform { get; private set; } = null!;

        /// <summary>
        /// The AWS Assigned deployment config id
        /// </summary>
        [Output("deploymentConfigId")]
        public Output<string> DeploymentConfigId { get; private set; } = null!;

        /// <summary>
        /// The name of the deployment config.
        /// </summary>
        [Output("deploymentConfigName")]
        public Output<string> DeploymentConfigName { get; private set; } = null!;

        /// <summary>
        /// A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
        /// </summary>
        [Output("minimumHealthyHosts")]
        public Output<Outputs.DeploymentConfigMinimumHealthyHosts?> MinimumHealthyHosts { get; private set; } = null!;

        /// <summary>
        /// A traffic_routing_config block. Traffic Routing Config is documented below.
        /// </summary>
        [Output("trafficRoutingConfig")]
        public Output<Outputs.DeploymentConfigTrafficRoutingConfig?> TrafficRoutingConfig { get; private set; } = null!;


        /// <summary>
        /// Create a DeploymentConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeploymentConfig(string name, DeploymentConfigArgs args, CustomResourceOptions? options = null)
            : base("aws:codedeploy/deploymentConfig:DeploymentConfig", name, args ?? new DeploymentConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeploymentConfig(string name, Input<string> id, DeploymentConfigState? state = null, CustomResourceOptions? options = null)
            : base("aws:codedeploy/deploymentConfig:DeploymentConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeploymentConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeploymentConfig Get(string name, Input<string> id, DeploymentConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new DeploymentConfig(name, id, state, options);
        }
    }

    public sealed class DeploymentConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
        /// </summary>
        [Input("computePlatform")]
        public Input<string>? ComputePlatform { get; set; }

        /// <summary>
        /// The name of the deployment config.
        /// </summary>
        [Input("deploymentConfigName", required: true)]
        public Input<string> DeploymentConfigName { get; set; } = null!;

        /// <summary>
        /// A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
        /// </summary>
        [Input("minimumHealthyHosts")]
        public Input<Inputs.DeploymentConfigMinimumHealthyHostsArgs>? MinimumHealthyHosts { get; set; }

        /// <summary>
        /// A traffic_routing_config block. Traffic Routing Config is documented below.
        /// </summary>
        [Input("trafficRoutingConfig")]
        public Input<Inputs.DeploymentConfigTrafficRoutingConfigArgs>? TrafficRoutingConfig { get; set; }

        public DeploymentConfigArgs()
        {
        }
    }

    public sealed class DeploymentConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
        /// </summary>
        [Input("computePlatform")]
        public Input<string>? ComputePlatform { get; set; }

        /// <summary>
        /// The AWS Assigned deployment config id
        /// </summary>
        [Input("deploymentConfigId")]
        public Input<string>? DeploymentConfigId { get; set; }

        /// <summary>
        /// The name of the deployment config.
        /// </summary>
        [Input("deploymentConfigName")]
        public Input<string>? DeploymentConfigName { get; set; }

        /// <summary>
        /// A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
        /// </summary>
        [Input("minimumHealthyHosts")]
        public Input<Inputs.DeploymentConfigMinimumHealthyHostsGetArgs>? MinimumHealthyHosts { get; set; }

        /// <summary>
        /// A traffic_routing_config block. Traffic Routing Config is documented below.
        /// </summary>
        [Input("trafficRoutingConfig")]
        public Input<Inputs.DeploymentConfigTrafficRoutingConfigGetArgs>? TrafficRoutingConfig { get; set; }

        public DeploymentConfigState()
        {
        }
    }
}
