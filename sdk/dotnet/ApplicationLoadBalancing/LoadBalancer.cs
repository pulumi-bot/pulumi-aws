// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApplicationLoadBalancing
{
    /// <summary>
    /// Provides a Load Balancer resource.
    /// 
    /// &gt; **Note:** `aws.alb.LoadBalancer` is known as `aws.lb.LoadBalancer`. The functionality is identical.
    /// 
    /// 
    /// 
    /// Deprecated: aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer
    /// </summary>
    [Obsolete(@"aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer")]
    public partial class LoadBalancer : Pulumi.CustomResource
    {
        /// <summary>
        /// An Access Logs block. Access Logs documented below.
        /// </summary>
        [Output("accessLogs")]
        public Output<Outputs.LoadBalancerAccessLogs?> AccessLogs { get; private set; } = null!;

        /// <summary>
        /// The ARN of the load balancer (matches `id`).
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ARN suffix for use with CloudWatch Metrics.
        /// </summary>
        [Output("arnSuffix")]
        public Output<string> ArnSuffix { get; private set; } = null!;

        /// <summary>
        /// The DNS name of the load balancer.
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
        /// </summary>
        [Output("dropInvalidHeaderFields")]
        public Output<bool?> DropInvalidHeaderFields { get; private set; } = null!;

        /// <summary>
        /// If true, cross-zone load balancing of the load balancer will be enabled.
        /// This is a `network` load balancer feature. Defaults to `false`.
        /// </summary>
        [Output("enableCrossZoneLoadBalancing")]
        public Output<bool?> EnableCrossZoneLoadBalancing { get; private set; } = null!;

        /// <summary>
        /// If true, deletion of the load balancer will be disabled via
        /// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
        /// </summary>
        [Output("enableDeletionProtection")]
        public Output<bool?> EnableDeletionProtection { get; private set; } = null!;

        /// <summary>
        /// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
        /// </summary>
        [Output("enableHttp2")]
        public Output<bool?> EnableHttp2 { get; private set; } = null!;

        /// <summary>
        /// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
        /// </summary>
        [Output("idleTimeout")]
        public Output<int?> IdleTimeout { get; private set; } = null!;

        /// <summary>
        /// If true, the LB will be internal.
        /// </summary>
        [Output("internal")]
        public Output<bool> Internal { get; private set; } = null!;

        /// <summary>
        /// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
        /// </summary>
        [Output("ipAddressType")]
        public Output<string> IpAddressType { get; private set; } = null!;

        /// <summary>
        /// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
        /// </summary>
        [Output("loadBalancerType")]
        public Output<string?> LoadBalancerType { get; private set; } = null!;

        /// <summary>
        /// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
        /// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
        /// this provider will autogenerate a name beginning with `tf-lb`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// A subnet mapping block as documented below.
        /// </summary>
        [Output("subnetMappings")]
        public Output<ImmutableArray<Outputs.LoadBalancerSubnetMapping>> SubnetMappings { get; private set; } = null!;

        /// <summary>
        /// A list of subnet IDs to attach to the LB. Subnets
        /// cannot be updated for Load Balancers of type `network`. Changing this value
        /// for load balancers of type `network` will force a recreation of the resource.
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<string>> Subnets { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancer(string name, LoadBalancerArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:applicationloadbalancing/loadBalancer:LoadBalancer", name, args ?? new LoadBalancerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancer(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
            : base("aws:applicationloadbalancing/loadBalancer:LoadBalancer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancer Get(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancer(name, id, state, options);
        }
    }

    public sealed class LoadBalancerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An Access Logs block. Access Logs documented below.
        /// </summary>
        [Input("accessLogs")]
        public Input<Inputs.LoadBalancerAccessLogsArgs>? AccessLogs { get; set; }

        /// <summary>
        /// Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
        /// </summary>
        [Input("dropInvalidHeaderFields")]
        public Input<bool>? DropInvalidHeaderFields { get; set; }

        /// <summary>
        /// If true, cross-zone load balancing of the load balancer will be enabled.
        /// This is a `network` load balancer feature. Defaults to `false`.
        /// </summary>
        [Input("enableCrossZoneLoadBalancing")]
        public Input<bool>? EnableCrossZoneLoadBalancing { get; set; }

        /// <summary>
        /// If true, deletion of the load balancer will be disabled via
        /// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
        /// </summary>
        [Input("enableDeletionProtection")]
        public Input<bool>? EnableDeletionProtection { get; set; }

        /// <summary>
        /// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
        /// </summary>
        [Input("enableHttp2")]
        public Input<bool>? EnableHttp2 { get; set; }

        /// <summary>
        /// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// If true, the LB will be internal.
        /// </summary>
        [Input("internal")]
        public Input<bool>? Internal { get; set; }

        /// <summary>
        /// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        /// <summary>
        /// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
        /// </summary>
        [Input("loadBalancerType")]
        public Input<string>? LoadBalancerType { get; set; }

        /// <summary>
        /// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
        /// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
        /// this provider will autogenerate a name beginning with `tf-lb`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnetMappings")]
        private InputList<Inputs.LoadBalancerSubnetMappingArgs>? _subnetMappings;

        /// <summary>
        /// A subnet mapping block as documented below.
        /// </summary>
        public InputList<Inputs.LoadBalancerSubnetMappingArgs> SubnetMappings
        {
            get => _subnetMappings ?? (_subnetMappings = new InputList<Inputs.LoadBalancerSubnetMappingArgs>());
            set => _subnetMappings = value;
        }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// A list of subnet IDs to attach to the LB. Subnets
        /// cannot be updated for Load Balancers of type `network`. Changing this value
        /// for load balancers of type `network` will force a recreation of the resource.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public LoadBalancerArgs()
        {
        }
    }

    public sealed class LoadBalancerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An Access Logs block. Access Logs documented below.
        /// </summary>
        [Input("accessLogs")]
        public Input<Inputs.LoadBalancerAccessLogsGetArgs>? AccessLogs { get; set; }

        /// <summary>
        /// The ARN of the load balancer (matches `id`).
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ARN suffix for use with CloudWatch Metrics.
        /// </summary>
        [Input("arnSuffix")]
        public Input<string>? ArnSuffix { get; set; }

        /// <summary>
        /// The DNS name of the load balancer.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
        /// </summary>
        [Input("dropInvalidHeaderFields")]
        public Input<bool>? DropInvalidHeaderFields { get; set; }

        /// <summary>
        /// If true, cross-zone load balancing of the load balancer will be enabled.
        /// This is a `network` load balancer feature. Defaults to `false`.
        /// </summary>
        [Input("enableCrossZoneLoadBalancing")]
        public Input<bool>? EnableCrossZoneLoadBalancing { get; set; }

        /// <summary>
        /// If true, deletion of the load balancer will be disabled via
        /// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
        /// </summary>
        [Input("enableDeletionProtection")]
        public Input<bool>? EnableDeletionProtection { get; set; }

        /// <summary>
        /// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
        /// </summary>
        [Input("enableHttp2")]
        public Input<bool>? EnableHttp2 { get; set; }

        /// <summary>
        /// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// If true, the LB will be internal.
        /// </summary>
        [Input("internal")]
        public Input<bool>? Internal { get; set; }

        /// <summary>
        /// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        /// <summary>
        /// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
        /// </summary>
        [Input("loadBalancerType")]
        public Input<string>? LoadBalancerType { get; set; }

        /// <summary>
        /// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
        /// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
        /// this provider will autogenerate a name beginning with `tf-lb`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnetMappings")]
        private InputList<Inputs.LoadBalancerSubnetMappingGetArgs>? _subnetMappings;

        /// <summary>
        /// A subnet mapping block as documented below.
        /// </summary>
        public InputList<Inputs.LoadBalancerSubnetMappingGetArgs> SubnetMappings
        {
            get => _subnetMappings ?? (_subnetMappings = new InputList<Inputs.LoadBalancerSubnetMappingGetArgs>());
            set => _subnetMappings = value;
        }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// A list of subnet IDs to attach to the LB. Subnets
        /// cannot be updated for Load Balancers of type `network`. Changing this value
        /// for load balancers of type `network` will force a recreation of the resource.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public LoadBalancerState()
        {
        }
    }
}
