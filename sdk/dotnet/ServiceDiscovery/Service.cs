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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new Aws.Ec2.VpcArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///             EnableDnsHostnames = true,
    ///             EnableDnsSupport = true,
    ///         });
    ///         var examplePrivateDnsNamespace = new Aws.ServiceDiscovery.PrivateDnsNamespace("examplePrivateDnsNamespace", new Aws.ServiceDiscovery.PrivateDnsNamespaceArgs
    ///         {
    ///             Description = "example",
    ///             Vpc = exampleVpc.Id,
    ///         });
    ///         var exampleService = new Aws.ServiceDiscovery.Service("exampleService", new Aws.ServiceDiscovery.ServiceArgs
    ///         {
    ///             DnsConfig = new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigArgs
    ///             {
    ///                 DnsRecords = 
    ///                 {
    ///                     new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigDnsRecordArgs
    ///                     {
    ///                         Ttl = 10,
    ///                         Type = "A",
    ///                     },
    ///                 },
    ///                 NamespaceId = examplePrivateDnsNamespace.Id,
    ///                 RoutingPolicy = "MULTIVALUE",
    ///             },
    ///             HealthCheckCustomConfig = new Aws.ServiceDiscovery.Inputs.ServiceHealthCheckCustomConfigArgs
    ///             {
    ///                 FailureThreshold = 1,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var examplePublicDnsNamespace = new Aws.ServiceDiscovery.PublicDnsNamespace("examplePublicDnsNamespace", new Aws.ServiceDiscovery.PublicDnsNamespaceArgs
    ///         {
    ///             Description = "example",
    ///         });
    ///         var exampleService = new Aws.ServiceDiscovery.Service("exampleService", new Aws.ServiceDiscovery.ServiceArgs
    ///         {
    ///             DnsConfig = new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigArgs
    ///             {
    ///                 DnsRecords = 
    ///                 {
    ///                     new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigDnsRecordArgs
    ///                     {
    ///                         Ttl = 10,
    ///                         Type = "A",
    ///                     },
    ///                 },
    ///                 NamespaceId = examplePublicDnsNamespace.Id,
    ///             },
    ///             HealthCheckConfig = new Aws.ServiceDiscovery.Inputs.ServiceHealthCheckConfigArgs
    ///             {
    ///                 FailureThreshold = 10,
    ///                 ResourcePath = "path",
    ///                 Type = "HTTP",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
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
            : base("aws:servicediscovery/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
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
}
