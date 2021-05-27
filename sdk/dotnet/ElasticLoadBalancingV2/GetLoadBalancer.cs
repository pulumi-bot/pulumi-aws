// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2
{
    [Obsolete(@"aws.elasticloadbalancingv2.getLoadBalancer has been deprecated in favor of aws.lb.getLoadBalancer")]
    public static class GetLoadBalancer
    {
        /// <summary>
        /// &gt; **Note:** `aws.alb.LoadBalancer` is known as `aws.lb.LoadBalancer`. The functionality is identical.
        /// 
        /// Provides information about a Load Balancer.
        /// 
        /// This data source can prove useful when a module accepts an LB as an input
        /// variable and needs to, for example, determine the security groups associated
        /// with it, etc.
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
        ///         var config = new Config();
        ///         var lbArn = config.Get("lbArn") ?? "";
        ///         var lbName = config.Get("lbName") ?? "";
        ///         var test = Output.Create(Aws.LB.GetLoadBalancer.InvokeAsync(new Aws.LB.GetLoadBalancerArgs
        ///         {
        ///             Arn = lbArn,
        ///             Name = lbName,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLoadBalancerResult> InvokeAsync(GetLoadBalancerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancerResult>("aws:elasticloadbalancingv2/getLoadBalancer:getLoadBalancer", args ?? new GetLoadBalancerArgs(), options.WithVersion());

        public static Output<GetLoadBalancerResult> Invoke(GetLoadBalancerOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetLoadBalancerOutputArgs();
            return Pulumi.Output.All(
                args.Arn.Box(),
                args.Name.Box(),
                args.Tags.ToDict().Box()
            ).Apply(a => {
                    var args = new GetLoadBalancerArgs();
                    a[0].Set(args, nameof(args.Arn));
                    a[1].Set(args, nameof(args.Name));
                    a[2].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetLoadBalancerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The full ARN of the load balancer.
        /// </summary>
        [Input("arn")]
        public string? Arn { get; set; }

        /// <summary>
        /// The unique name of the load balancer.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetLoadBalancerArgs()
        {
        }
    }

    public sealed class GetLoadBalancerOutputArgs
    {
        /// <summary>
        /// The full ARN of the load balancer.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The unique name of the load balancer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetLoadBalancerOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLoadBalancerResult
    {
        public readonly Outputs.GetLoadBalancerAccessLogsResult AccessLogs;
        public readonly string Arn;
        public readonly string ArnSuffix;
        public readonly string CustomerOwnedIpv4Pool;
        public readonly string DnsName;
        public readonly bool DropInvalidHeaderFields;
        public readonly bool EnableDeletionProtection;
        public readonly bool EnableHttp2;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int IdleTimeout;
        public readonly bool Internal;
        public readonly string IpAddressType;
        public readonly string LoadBalancerType;
        public readonly string Name;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly ImmutableArray<Outputs.GetLoadBalancerSubnetMappingResult> SubnetMappings;
        public readonly ImmutableArray<string> Subnets;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string VpcId;
        public readonly string ZoneId;

        [OutputConstructor]
        private GetLoadBalancerResult(
            Outputs.GetLoadBalancerAccessLogsResult accessLogs,

            string arn,

            string arnSuffix,

            string customerOwnedIpv4Pool,

            string dnsName,

            bool dropInvalidHeaderFields,

            bool enableDeletionProtection,

            bool enableHttp2,

            string id,

            int idleTimeout,

            bool @internal,

            string ipAddressType,

            string loadBalancerType,

            string name,

            ImmutableArray<string> securityGroups,

            ImmutableArray<Outputs.GetLoadBalancerSubnetMappingResult> subnetMappings,

            ImmutableArray<string> subnets,

            ImmutableDictionary<string, string> tags,

            string vpcId,

            string zoneId)
        {
            AccessLogs = accessLogs;
            Arn = arn;
            ArnSuffix = arnSuffix;
            CustomerOwnedIpv4Pool = customerOwnedIpv4Pool;
            DnsName = dnsName;
            DropInvalidHeaderFields = dropInvalidHeaderFields;
            EnableDeletionProtection = enableDeletionProtection;
            EnableHttp2 = enableHttp2;
            Id = id;
            IdleTimeout = idleTimeout;
            Internal = @internal;
            IpAddressType = ipAddressType;
            LoadBalancerType = loadBalancerType;
            Name = name;
            SecurityGroups = securityGroups;
            SubnetMappings = subnetMappings;
            Subnets = subnets;
            Tags = tags;
            VpcId = vpcId;
            ZoneId = zoneId;
        }
    }
}
