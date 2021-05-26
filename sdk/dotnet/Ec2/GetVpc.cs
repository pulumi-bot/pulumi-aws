// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetVpc
    {
        /// <summary>
        /// `aws.ec2.Vpc` provides details about a specific VPC.
        /// 
        /// This resource can prove useful when a module accepts a vpc id as
        /// an input variable and needs to, for example, determine the CIDR block of that
        /// VPC.
        /// </summary>
        public static Task<GetVpcResult> InvokeAsync(GetVpcArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcResult>("aws:ec2/getVpc:getVpc", args ?? new GetVpcArgs(), options.WithVersion());

        public static Output<GetVpcResult> Apply(GetVpcApplyArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetVpcApplyArgs();
            return Pulumi.Output.All(
                args.CidrBlock.Box(),
                args.Default.Box(),
                args.DhcpOptionsId.Box(),
                args.Filters.Box(),
                args.Id.Box(),
                args.State.Box(),
                args.Tags.Box()
            ).Apply(a => {
                    var args = new GetVpcArgs();
                    a[0].Set(args, nameof(args.CidrBlock));
                    a[1].Set(args, nameof(args.Default));
                    a[2].Set(args, nameof(args.DhcpOptionsId));
                    a[3].Set(args, nameof(args.Filters));
                    a[4].Set(args, nameof(args.Id));
                    a[5].Set(args, nameof(args.State));
                    a[6].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetVpcArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cidr block of the desired VPC.
        /// </summary>
        [Input("cidrBlock")]
        public string? CidrBlock { get; set; }

        /// <summary>
        /// Boolean constraint on whether the desired VPC is
        /// the default VPC for the region.
        /// </summary>
        [Input("default")]
        public bool? Default { get; set; }

        /// <summary>
        /// The DHCP options id of the desired VPC.
        /// </summary>
        [Input("dhcpOptionsId")]
        public string? DhcpOptionsId { get; set; }

        [Input("filters")]
        private List<Inputs.GetVpcFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetVpcFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The id of the specific VPC to retrieve.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The current state of the desired VPC.
        /// Can be either `"pending"` or `"available"`.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired VPC.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetVpcArgs()
        {
        }
    }

    public sealed class GetVpcApplyArgs
    {
        /// <summary>
        /// The cidr block of the desired VPC.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Boolean constraint on whether the desired VPC is
        /// the default VPC for the region.
        /// </summary>
        [Input("default")]
        public Input<bool>? Default { get; set; }

        /// <summary>
        /// The DHCP options id of the desired VPC.
        /// </summary>
        [Input("dhcpOptionsId")]
        public Input<string>? DhcpOptionsId { get; set; }

        [Input("filters")]
        private InputList<Inputs.GetVpcFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public InputList<Inputs.GetVpcFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetVpcFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The id of the specific VPC to retrieve.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The current state of the desired VPC.
        /// Can be either `"pending"` or `"available"`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired VPC.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetVpcApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of VPC
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The CIDR block for the association.
        /// </summary>
        public readonly string CidrBlock;
        public readonly ImmutableArray<Outputs.GetVpcCidrBlockAssociationResult> CidrBlockAssociations;
        public readonly bool Default;
        public readonly string DhcpOptionsId;
        /// <summary>
        /// Whether or not the VPC has DNS hostname support
        /// </summary>
        public readonly bool EnableDnsHostnames;
        /// <summary>
        /// Whether or not the VPC has DNS support
        /// </summary>
        public readonly bool EnableDnsSupport;
        public readonly ImmutableArray<Outputs.GetVpcFilterResult> Filters;
        public readonly string Id;
        /// <summary>
        /// The allowed tenancy of instances launched into the
        /// selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
        /// </summary>
        public readonly string InstanceTenancy;
        /// <summary>
        /// The association ID for the IPv6 CIDR block.
        /// </summary>
        public readonly string Ipv6AssociationId;
        /// <summary>
        /// The IPv6 CIDR block.
        /// </summary>
        public readonly string Ipv6CidrBlock;
        /// <summary>
        /// The ID of the main route table associated with this VPC.
        /// </summary>
        public readonly string MainRouteTableId;
        /// <summary>
        /// The ID of the AWS account that owns the VPC.
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// The State of the association.
        /// </summary>
        public readonly string State;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetVpcResult(
            string arn,

            string cidrBlock,

            ImmutableArray<Outputs.GetVpcCidrBlockAssociationResult> cidrBlockAssociations,

            bool @default,

            string dhcpOptionsId,

            bool enableDnsHostnames,

            bool enableDnsSupport,

            ImmutableArray<Outputs.GetVpcFilterResult> filters,

            string id,

            string instanceTenancy,

            string ipv6AssociationId,

            string ipv6CidrBlock,

            string mainRouteTableId,

            string ownerId,

            string state,

            ImmutableDictionary<string, string> tags)
        {
            Arn = arn;
            CidrBlock = cidrBlock;
            CidrBlockAssociations = cidrBlockAssociations;
            Default = @default;
            DhcpOptionsId = dhcpOptionsId;
            EnableDnsHostnames = enableDnsHostnames;
            EnableDnsSupport = enableDnsSupport;
            Filters = filters;
            Id = id;
            InstanceTenancy = instanceTenancy;
            Ipv6AssociationId = ipv6AssociationId;
            Ipv6CidrBlock = ipv6CidrBlock;
            MainRouteTableId = mainRouteTableId;
            OwnerId = ownerId;
            State = state;
            Tags = tags;
        }
    }
}
