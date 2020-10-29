// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2
{
    /// <summary>
    /// Manages an Amazon API Gateway Version 2 VPC Link.
    /// 
    /// &gt; **Note:** Amazon API Gateway Version 2 VPC Links enable private integrations that connect HTTP APIs to private resources in a VPC.
    /// To enable private integration for REST APIs, use the `Amazon API Gateway Version 1 VPC Link` resource.
    /// </summary>
    public partial class VpcLink : Pulumi.CustomResource
    {
        /// <summary>
        /// The VPC Link ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the VPC Link.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Security group IDs for the VPC Link.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Subnet IDs for the VPC Link.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the VPC Link.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a VpcLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcLink(string name, VpcLinkArgs args, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/vpcLink:VpcLink", name, args ?? new VpcLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcLink(string name, Input<string> id, VpcLinkState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/vpcLink:VpcLink", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcLink Get(string name, Input<string> id, VpcLinkState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcLink(name, id, state, options);
        }
    }

    public sealed class VpcLinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the VPC Link.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group IDs for the VPC Link.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// Subnet IDs for the VPC Link.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the VPC Link.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public VpcLinkArgs()
        {
        }
    }

    public sealed class VpcLinkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The VPC Link ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the VPC Link.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group IDs for the VPC Link.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// Subnet IDs for the VPC Link.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the VPC Link.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public VpcLinkState()
        {
        }
    }
}
