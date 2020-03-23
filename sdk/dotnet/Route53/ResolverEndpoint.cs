// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    /// <summary>
    /// Provides a Route 53 Resolver endpoint resource.
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_resolver_endpoint.html.markdown.
    /// </summary>
    public partial class ResolverEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Route 53 Resolver endpoint.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The direction of DNS queries to or from the Route 53 Resolver endpoint.
        /// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
        /// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC that you want to create the resolver endpoint in.
        /// </summary>
        [Output("hostVpcId")]
        public Output<string> HostVpcId { get; private set; } = null!;

        /// <summary>
        /// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
        /// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
        /// </summary>
        [Output("ipAddresses")]
        public Output<ImmutableArray<Outputs.ResolverEndpointIpAddresses>> IpAddresses { get; private set; } = null!;

        /// <summary>
        /// The friendly name of the Route 53 Resolver endpoint.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of one or more security groups that you want to use to control access to this VPC.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ResolverEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResolverEndpoint(string name, ResolverEndpointArgs args, CustomResourceOptions? options = null)
            : base("aws:route53/resolverEndpoint:ResolverEndpoint", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ResolverEndpoint(string name, Input<string> id, ResolverEndpointState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/resolverEndpoint:ResolverEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResolverEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResolverEndpoint Get(string name, Input<string> id, ResolverEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new ResolverEndpoint(name, id, state, options);
        }
    }

    public sealed class ResolverEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The direction of DNS queries to or from the Route 53 Resolver endpoint.
        /// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
        /// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
        /// </summary>
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        [Input("ipAddresses", required: true)]
        private InputList<Inputs.ResolverEndpointIpAddressesArgs>? _ipAddresses;

        /// <summary>
        /// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
        /// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
        /// </summary>
        public InputList<Inputs.ResolverEndpointIpAddressesArgs> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<Inputs.ResolverEndpointIpAddressesArgs>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// The friendly name of the Route 53 Resolver endpoint.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The ID of one or more security groups that you want to use to control access to this VPC.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
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

        public ResolverEndpointArgs()
        {
        }
    }

    public sealed class ResolverEndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Route 53 Resolver endpoint.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The direction of DNS queries to or from the Route 53 Resolver endpoint.
        /// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
        /// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The ID of the VPC that you want to create the resolver endpoint in.
        /// </summary>
        [Input("hostVpcId")]
        public Input<string>? HostVpcId { get; set; }

        [Input("ipAddresses")]
        private InputList<Inputs.ResolverEndpointIpAddressesGetArgs>? _ipAddresses;

        /// <summary>
        /// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
        /// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
        /// </summary>
        public InputList<Inputs.ResolverEndpointIpAddressesGetArgs> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<Inputs.ResolverEndpointIpAddressesGetArgs>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// The friendly name of the Route 53 Resolver endpoint.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The ID of one or more security groups that you want to use to control access to this VPC.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
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

        public ResolverEndpointState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ResolverEndpointIpAddressesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address in the subnet that you want to use for DNS queries.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        /// <summary>
        /// The ID of the subnet that contains the IP address.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public ResolverEndpointIpAddressesArgs()
        {
        }
    }

    public sealed class ResolverEndpointIpAddressesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address in the subnet that you want to use for DNS queries.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        /// <summary>
        /// The ID of the subnet that contains the IP address.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public ResolverEndpointIpAddressesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ResolverEndpointIpAddresses
    {
        /// <summary>
        /// The IP address in the subnet that you want to use for DNS queries.
        /// </summary>
        public readonly string Ip;
        public readonly string IpId;
        /// <summary>
        /// The ID of the subnet that contains the IP address.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private ResolverEndpointIpAddresses(
            string ip,
            string ipId,
            string subnetId)
        {
            Ip = ip;
            IpId = ipId;
            SubnetId = subnetId;
        }
    }
    }
}
