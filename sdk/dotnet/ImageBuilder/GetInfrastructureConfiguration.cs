// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ImageBuilder
{
    public static class GetInfrastructureConfiguration
    {
        /// <summary>
        /// Provides details about an Image Builder Infrastructure Configuration.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.ImageBuilder.GetInfrastructureConfiguration.InvokeAsync(new Aws.ImageBuilder.GetInfrastructureConfigurationArgs
        ///         {
        ///             Arn = "arn:aws:imagebuilder:us-west-2:aws:infrastructure-configuration/example",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInfrastructureConfigurationResult> InvokeAsync(GetInfrastructureConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInfrastructureConfigurationResult>("aws:imagebuilder/getInfrastructureConfiguration:getInfrastructureConfiguration", args ?? new GetInfrastructureConfigurationArgs(), options.WithVersion());

        public static Output<GetInfrastructureConfigurationResult> Invoke(GetInfrastructureConfigurationOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Arn.Box(),
                args.ResourceTags.ToDict().Box(),
                args.Tags.ToDict().Box()
            ).Apply(a => {
                    var args = new GetInfrastructureConfigurationArgs();
                    a[0].Set(args, nameof(args.Arn));
                    a[1].Set(args, nameof(args.ResourceTags));
                    a[2].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetInfrastructureConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the infrastructure configuration.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        [Input("resourceTags")]
        private Dictionary<string, string>? _resourceTags;

        /// <summary>
        /// Key-value map of resource tags for the infrastructure created by the infrastructure configuration.
        /// </summary>
        public Dictionary<string, string> ResourceTags
        {
            get => _resourceTags ?? (_resourceTags = new Dictionary<string, string>());
            set => _resourceTags = value;
        }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Key-value map of resource tags for the infrastructure configuration.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetInfrastructureConfigurationArgs()
        {
        }
    }

    public sealed class GetInfrastructureConfigurationOutputArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the infrastructure configuration.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("resourceTags")]
        private InputMap<string>? _resourceTags;

        /// <summary>
        /// Key-value map of resource tags for the infrastructure created by the infrastructure configuration.
        /// </summary>
        public InputMap<string> ResourceTags
        {
            get => _resourceTags ?? (_resourceTags = new InputMap<string>());
            set => _resourceTags = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags for the infrastructure configuration.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetInfrastructureConfigurationOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInfrastructureConfigurationResult
    {
        public readonly string Arn;
        /// <summary>
        /// Date the infrastructure configuration was updated.
        /// </summary>
        public readonly string DateCreated;
        public readonly string DateUpdated;
        /// <summary>
        /// Description of the infrastructure configuration.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the IAM Instance Profile associated with the configuration.
        /// </summary>
        public readonly string InstanceProfileName;
        /// <summary>
        /// Set of EC2 Instance Types associated with the configuration.
        /// </summary>
        public readonly ImmutableArray<string> InstanceTypes;
        /// <summary>
        /// Name of the EC2 Key Pair associated with the configuration.
        /// </summary>
        public readonly string KeyPair;
        /// <summary>
        /// Nested list of logging settings.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInfrastructureConfigurationLoggingResult> Loggings;
        /// <summary>
        /// Name of the infrastructure configuration.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Key-value map of resource tags for the infrastructure created by the infrastructure configuration.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ResourceTags;
        /// <summary>
        /// Set of EC2 Security Group identifiers associated with the configuration.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// Amazon Resource Name (ARN) of the SNS Topic associated with the configuration.
        /// </summary>
        public readonly string SnsTopicArn;
        /// <summary>
        /// Identifier of the EC2 Subnet associated with the configuration.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// Key-value map of resource tags for the infrastructure configuration.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Whether instances are terminated on failure.
        /// </summary>
        public readonly bool TerminateInstanceOnFailure;

        [OutputConstructor]
        private GetInfrastructureConfigurationResult(
            string arn,

            string dateCreated,

            string dateUpdated,

            string description,

            string id,

            string instanceProfileName,

            ImmutableArray<string> instanceTypes,

            string keyPair,

            ImmutableArray<Outputs.GetInfrastructureConfigurationLoggingResult> loggings,

            string name,

            ImmutableDictionary<string, string> resourceTags,

            ImmutableArray<string> securityGroupIds,

            string snsTopicArn,

            string subnetId,

            ImmutableDictionary<string, string> tags,

            bool terminateInstanceOnFailure)
        {
            Arn = arn;
            DateCreated = dateCreated;
            DateUpdated = dateUpdated;
            Description = description;
            Id = id;
            InstanceProfileName = instanceProfileName;
            InstanceTypes = instanceTypes;
            KeyPair = keyPair;
            Loggings = loggings;
            Name = name;
            ResourceTags = resourceTags;
            SecurityGroupIds = securityGroupIds;
            SnsTopicArn = snsTopicArn;
            SubnetId = subnetId;
            Tags = tags;
            TerminateInstanceOnFailure = terminateInstanceOnFailure;
        }
    }
}
