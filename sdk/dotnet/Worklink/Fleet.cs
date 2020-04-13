// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WorkLink
{
    /// <summary>
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/worklink_fleet.html.markdown.
    /// </summary>
    public partial class Fleet : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the created WorkLink Fleet.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the Amazon Kinesis data stream that receives the audit events.
        /// </summary>
        [Output("auditStreamArn")]
        public Output<string?> AuditStreamArn { get; private set; } = null!;

        /// <summary>
        /// The identifier used by users to sign in to the Amazon WorkLink app.
        /// </summary>
        [Output("companyCode")]
        public Output<string> CompanyCode { get; private set; } = null!;

        /// <summary>
        /// The time that the fleet was created.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        /// </summary>
        [Output("deviceCaCertificate")]
        public Output<string?> DeviceCaCertificate { get; private set; } = null!;

        /// <summary>
        /// The name of the fleet.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// is cannot removed without force recreating.
        /// </summary>
        [Output("identityProvider")]
        public Output<Outputs.FleetIdentityProvider?> IdentityProvider { get; private set; } = null!;

        /// <summary>
        /// The time that the fleet was last updated.
        /// </summary>
        [Output("lastUpdatedTime")]
        public Output<string> LastUpdatedTime { get; private set; } = null!;

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// is cannot removed without force recreating.
        /// </summary>
        [Output("network")]
        public Output<Outputs.FleetNetwork?> Network { get; private set; } = null!;

        /// <summary>
        /// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
        /// </summary>
        [Output("optimizeForEndUserLocation")]
        public Output<bool?> OptimizeForEndUserLocation { get; private set; } = null!;


        /// <summary>
        /// Create a Fleet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fleet(string name, FleetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:worklink/fleet:Fleet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Fleet(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
            : base("aws:worklink/fleet:Fleet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fleet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fleet Get(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
        {
            return new Fleet(name, id, state, options);
        }
    }

    public sealed class FleetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Amazon Kinesis data stream that receives the audit events.
        /// </summary>
        [Input("auditStreamArn")]
        public Input<string>? AuditStreamArn { get; set; }

        /// <summary>
        /// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        /// </summary>
        [Input("deviceCaCertificate")]
        public Input<string>? DeviceCaCertificate { get; set; }

        /// <summary>
        /// The name of the fleet.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// is cannot removed without force recreating.
        /// </summary>
        [Input("identityProvider")]
        public Input<Inputs.FleetIdentityProviderArgs>? IdentityProvider { get; set; }

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// is cannot removed without force recreating.
        /// </summary>
        [Input("network")]
        public Input<Inputs.FleetNetworkArgs>? Network { get; set; }

        /// <summary>
        /// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
        /// </summary>
        [Input("optimizeForEndUserLocation")]
        public Input<bool>? OptimizeForEndUserLocation { get; set; }

        public FleetArgs()
        {
        }
    }

    public sealed class FleetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the created WorkLink Fleet.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ARN of the Amazon Kinesis data stream that receives the audit events.
        /// </summary>
        [Input("auditStreamArn")]
        public Input<string>? AuditStreamArn { get; set; }

        /// <summary>
        /// The identifier used by users to sign in to the Amazon WorkLink app.
        /// </summary>
        [Input("companyCode")]
        public Input<string>? CompanyCode { get; set; }

        /// <summary>
        /// The time that the fleet was created.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        /// </summary>
        [Input("deviceCaCertificate")]
        public Input<string>? DeviceCaCertificate { get; set; }

        /// <summary>
        /// The name of the fleet.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// is cannot removed without force recreating.
        /// </summary>
        [Input("identityProvider")]
        public Input<Inputs.FleetIdentityProviderGetArgs>? IdentityProvider { get; set; }

        /// <summary>
        /// The time that the fleet was last updated.
        /// </summary>
        [Input("lastUpdatedTime")]
        public Input<string>? LastUpdatedTime { get; set; }

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// is cannot removed without force recreating.
        /// </summary>
        [Input("network")]
        public Input<Inputs.FleetNetworkGetArgs>? Network { get; set; }

        /// <summary>
        /// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
        /// </summary>
        [Input("optimizeForEndUserLocation")]
        public Input<bool>? OptimizeForEndUserLocation { get; set; }

        public FleetState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class FleetIdentityProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SAML metadata document provided by the customer’s identity provider.
        /// </summary>
        [Input("samlMetadata", required: true)]
        public Input<string> SamlMetadata { get; set; } = null!;

        /// <summary>
        /// The type of identity provider.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FleetIdentityProviderArgs()
        {
        }
    }

    public sealed class FleetIdentityProviderGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SAML metadata document provided by the customer’s identity provider.
        /// </summary>
        [Input("samlMetadata", required: true)]
        public Input<string> SamlMetadata { get; set; } = null!;

        /// <summary>
        /// The type of identity provider.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FleetIdentityProviderGetArgs()
        {
        }
    }

    public sealed class FleetNetworkArgs : Pulumi.ResourceArgs
    {
        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs associated with access to the provided subnets.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of subnet IDs used for X-ENI connections from Amazon WorkLink rendering containers.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The VPC ID with connectivity to associated websites.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public FleetNetworkArgs()
        {
        }
    }

    public sealed class FleetNetworkGetArgs : Pulumi.ResourceArgs
    {
        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs associated with access to the provided subnets.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of subnet IDs used for X-ENI connections from Amazon WorkLink rendering containers.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The VPC ID with connectivity to associated websites.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public FleetNetworkGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class FleetIdentityProvider
    {
        /// <summary>
        /// The SAML metadata document provided by the customer’s identity provider.
        /// </summary>
        public readonly string SamlMetadata;
        /// <summary>
        /// The type of identity provider.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private FleetIdentityProvider(
            string samlMetadata,
            string type)
        {
            SamlMetadata = samlMetadata;
            Type = type;
        }
    }

    [OutputType]
    public sealed class FleetNetwork
    {
        /// <summary>
        /// A list of security group IDs associated with access to the provided subnets.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// A list of subnet IDs used for X-ENI connections from Amazon WorkLink rendering containers.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The VPC ID with connectivity to associated websites.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private FleetNetwork(
            ImmutableArray<string> securityGroupIds,
            ImmutableArray<string> subnetIds,
            string vpcId)
        {
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
    }
}
