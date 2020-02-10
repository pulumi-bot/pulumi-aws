// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ServiceDiscovery
{
    /// <summary>
    /// Provides a Service Discovery Service resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/service_discovery_service.html.markdown.
    /// </summary>
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the service.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the service.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
        /// </summary>
        [Output("dnsConfig")]
        public Output<Outputs.ServiceDnsConfig?> DnsConfig { get; private set; } = null!;

        /// <summary>
        /// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
        /// </summary>
        [Output("healthCheckConfig")]
        public Output<Outputs.ServiceHealthCheckConfig?> HealthCheckConfig { get; private set; } = null!;

        /// <summary>
        /// A complex type that contains settings for ECS managed health checks.
        /// </summary>
        [Output("healthCheckCustomConfig")]
        public Output<Outputs.ServiceHealthCheckCustomConfig?> HealthCheckCustomConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the service.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the namespace to use for DNS configuration.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:servicediscovery/service:Service", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("aws:servicediscovery/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the service.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
        /// </summary>
        [Input("dnsConfig")]
        public Input<Inputs.ServiceDnsConfigArgs>? DnsConfig { get; set; }

        /// <summary>
        /// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
        /// </summary>
        [Input("healthCheckConfig")]
        public Input<Inputs.ServiceHealthCheckConfigArgs>? HealthCheckConfig { get; set; }

        /// <summary>
        /// A complex type that contains settings for ECS managed health checks.
        /// </summary>
        [Input("healthCheckCustomConfig")]
        public Input<Inputs.ServiceHealthCheckCustomConfigArgs>? HealthCheckCustomConfig { get; set; }

        /// <summary>
        /// The name of the service.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the namespace to use for DNS configuration.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the service.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The description of the service.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
        /// </summary>
        [Input("dnsConfig")]
        public Input<Inputs.ServiceDnsConfigGetArgs>? DnsConfig { get; set; }

        /// <summary>
        /// A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
        /// </summary>
        [Input("healthCheckConfig")]
        public Input<Inputs.ServiceHealthCheckConfigGetArgs>? HealthCheckConfig { get; set; }

        /// <summary>
        /// A complex type that contains settings for ECS managed health checks.
        /// </summary>
        [Input("healthCheckCustomConfig")]
        public Input<Inputs.ServiceHealthCheckCustomConfigGetArgs>? HealthCheckCustomConfig { get; set; }

        /// <summary>
        /// The name of the service.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the namespace to use for DNS configuration.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        public ServiceState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ServiceDnsConfigArgs : Pulumi.ResourceArgs
    {
        [Input("dnsRecords", required: true)]
        private InputList<ServiceDnsConfigDnsRecordsArgs>? _dnsRecords;

        /// <summary>
        /// An array that contains one DnsRecord object for each resource record set.
        /// </summary>
        public InputList<ServiceDnsConfigDnsRecordsArgs> DnsRecords
        {
            get => _dnsRecords ?? (_dnsRecords = new InputList<ServiceDnsConfigDnsRecordsArgs>());
            set => _dnsRecords = value;
        }

        /// <summary>
        /// The ID of the namespace to use for DNS configuration.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        /// <summary>
        /// The routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED
        /// </summary>
        [Input("routingPolicy")]
        public Input<string>? RoutingPolicy { get; set; }

        public ServiceDnsConfigArgs()
        {
        }
    }

    public sealed class ServiceDnsConfigDnsRecordsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set.
        /// </summary>
        [Input("ttl", required: true)]
        public Input<int> Ttl { get; set; } = null!;

        /// <summary>
        /// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ServiceDnsConfigDnsRecordsArgs()
        {
        }
    }

    public sealed class ServiceDnsConfigDnsRecordsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set.
        /// </summary>
        [Input("ttl", required: true)]
        public Input<int> Ttl { get; set; } = null!;

        /// <summary>
        /// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ServiceDnsConfigDnsRecordsGetArgs()
        {
        }
    }

    public sealed class ServiceDnsConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("dnsRecords", required: true)]
        private InputList<ServiceDnsConfigDnsRecordsGetArgs>? _dnsRecords;

        /// <summary>
        /// An array that contains one DnsRecord object for each resource record set.
        /// </summary>
        public InputList<ServiceDnsConfigDnsRecordsGetArgs> DnsRecords
        {
            get => _dnsRecords ?? (_dnsRecords = new InputList<ServiceDnsConfigDnsRecordsGetArgs>());
            set => _dnsRecords = value;
        }

        /// <summary>
        /// The ID of the namespace to use for DNS configuration.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        /// <summary>
        /// The routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED
        /// </summary>
        [Input("routingPolicy")]
        public Input<string>? RoutingPolicy { get; set; }

        public ServiceDnsConfigGetArgs()
        {
        }
    }

    public sealed class ServiceHealthCheckConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance.  Maximum value of 10.
        /// </summary>
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        /// <summary>
        /// The path that you want Route 53 to request when performing health checks. Route 53 automatically adds the DNS name for the service. If you don't specify a value, the default value is /.
        /// </summary>
        [Input("resourcePath")]
        public Input<string>? ResourcePath { get; set; }

        /// <summary>
        /// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ServiceHealthCheckConfigArgs()
        {
        }
    }

    public sealed class ServiceHealthCheckConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance.  Maximum value of 10.
        /// </summary>
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        /// <summary>
        /// The path that you want Route 53 to request when performing health checks. Route 53 automatically adds the DNS name for the service. If you don't specify a value, the default value is /.
        /// </summary>
        [Input("resourcePath")]
        public Input<string>? ResourcePath { get; set; }

        /// <summary>
        /// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ServiceHealthCheckConfigGetArgs()
        {
        }
    }

    public sealed class ServiceHealthCheckCustomConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance.  Maximum value of 10.
        /// </summary>
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        public ServiceHealthCheckCustomConfigArgs()
        {
        }
    }

    public sealed class ServiceHealthCheckCustomConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance.  Maximum value of 10.
        /// </summary>
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        public ServiceHealthCheckCustomConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ServiceDnsConfig
    {
        /// <summary>
        /// An array that contains one DnsRecord object for each resource record set.
        /// </summary>
        public readonly ImmutableArray<ServiceDnsConfigDnsRecords> DnsRecords;
        /// <summary>
        /// The ID of the namespace to use for DNS configuration.
        /// </summary>
        public readonly string NamespaceId;
        /// <summary>
        /// The routing policy that you want to apply to all records that Route 53 creates when you register an instance and specify the service. Valid Values: MULTIVALUE, WEIGHTED
        /// </summary>
        public readonly string? RoutingPolicy;

        [OutputConstructor]
        private ServiceDnsConfig(
            ImmutableArray<ServiceDnsConfigDnsRecords> dnsRecords,
            string namespaceId,
            string? routingPolicy)
        {
            DnsRecords = dnsRecords;
            NamespaceId = namespaceId;
            RoutingPolicy = routingPolicy;
        }
    }

    [OutputType]
    public sealed class ServiceDnsConfigDnsRecords
    {
        /// <summary>
        /// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this resource record set.
        /// </summary>
        public readonly int Ttl;
        /// <summary>
        /// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ServiceDnsConfigDnsRecords(
            int ttl,
            string type)
        {
            Ttl = ttl;
            Type = type;
        }
    }

    [OutputType]
    public sealed class ServiceHealthCheckConfig
    {
        /// <summary>
        /// The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance.  Maximum value of 10.
        /// </summary>
        public readonly int? FailureThreshold;
        /// <summary>
        /// The path that you want Route 53 to request when performing health checks. Route 53 automatically adds the DNS name for the service. If you don't specify a value, the default value is /.
        /// </summary>
        public readonly string? ResourcePath;
        /// <summary>
        /// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ServiceHealthCheckConfig(
            int? failureThreshold,
            string? resourcePath,
            string? type)
        {
            FailureThreshold = failureThreshold;
            ResourcePath = resourcePath;
            Type = type;
        }
    }

    [OutputType]
    public sealed class ServiceHealthCheckCustomConfig
    {
        /// <summary>
        /// The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance.  Maximum value of 10.
        /// </summary>
        public readonly int? FailureThreshold;

        [OutputConstructor]
        private ServiceHealthCheckCustomConfig(int? failureThreshold)
        {
            FailureThreshold = failureThreshold;
        }
    }
    }
}
