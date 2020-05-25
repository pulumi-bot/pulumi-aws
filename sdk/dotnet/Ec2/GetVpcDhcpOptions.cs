// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetVpcDhcpOptions
    {
        /// <summary>
        /// Retrieve information about an EC2 DHCP Options configuration.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// 
        /// {{% example %}}
        /// ### Lookup by DHCP Options ID
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Ec2.GetVpcDhcpOptions.InvokeAsync(new Aws.Ec2.GetVpcDhcpOptionsArgs
        ///         {
        ///             DhcpOptionsId = "dopts-12345678",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Lookup by Filter
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Ec2.GetVpcDhcpOptions.InvokeAsync(new Aws.Ec2.GetVpcDhcpOptionsArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Ec2.Inputs.GetVpcDhcpOptionsFilterArgs
        ///                 {
        ///                     Name = "key",
        ///                     Values = 
        ///                     {
        ///                         "domain-name",
        ///                     },
        ///                 },
        ///                 new Aws.Ec2.Inputs.GetVpcDhcpOptionsFilterArgs
        ///                 {
        ///                     Name = "value",
        ///                     Values = 
        ///                     {
        ///                         "example.com",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcDhcpOptionsResult> InvokeAsync(GetVpcDhcpOptionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcDhcpOptionsResult>("aws:ec2/getVpcDhcpOptions:getVpcDhcpOptions", args ?? new GetVpcDhcpOptionsArgs(), options.WithVersion());
    }


    public sealed class GetVpcDhcpOptionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The EC2 DHCP Options ID.
        /// </summary>
        [Input("dhcpOptionsId")]
        public string? DhcpOptionsId { get; set; }

        [Input("filters")]
        private List<Inputs.GetVpcDhcpOptionsFilterArgs>? _filters;

        /// <summary>
        /// List of custom filters as described below.
        /// </summary>
        public List<Inputs.GetVpcDhcpOptionsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcDhcpOptionsFilterArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags assigned to the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetVpcDhcpOptionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcDhcpOptionsResult
    {
        /// <summary>
        /// EC2 DHCP Options ID
        /// </summary>
        public readonly string DhcpOptionsId;
        /// <summary>
        /// The suffix domain name to used when resolving non Fully Qualified Domain Names. e.g. the `search` value in the `/etc/resolv.conf` file.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// List of name servers.
        /// </summary>
        public readonly ImmutableArray<string> DomainNameServers;
        public readonly ImmutableArray<Outputs.GetVpcDhcpOptionsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of NETBIOS name servers.
        /// </summary>
        public readonly ImmutableArray<string> NetbiosNameServers;
        /// <summary>
        /// The NetBIOS node type (1, 2, 4, or 8). For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
        /// </summary>
        public readonly string NetbiosNodeType;
        /// <summary>
        /// List of NTP servers.
        /// </summary>
        public readonly ImmutableArray<string> NtpServers;
        /// <summary>
        /// The ID of the AWS account that owns the DHCP options set.
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// A map of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;

        [OutputConstructor]
        private GetVpcDhcpOptionsResult(
            string dhcpOptionsId,

            string domainName,

            ImmutableArray<string> domainNameServers,

            ImmutableArray<Outputs.GetVpcDhcpOptionsFilterResult> filters,

            string id,

            ImmutableArray<string> netbiosNameServers,

            string netbiosNodeType,

            ImmutableArray<string> ntpServers,

            string ownerId,

            ImmutableDictionary<string, object> tags)
        {
            DhcpOptionsId = dhcpOptionsId;
            DomainName = domainName;
            DomainNameServers = domainNameServers;
            Filters = filters;
            Id = id;
            NetbiosNameServers = netbiosNameServers;
            NetbiosNodeType = netbiosNodeType;
            NtpServers = ntpServers;
            OwnerId = ownerId;
            Tags = tags;
        }
    }
}
