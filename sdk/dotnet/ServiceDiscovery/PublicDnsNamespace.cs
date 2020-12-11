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
    /// Provides a Service Discovery Public DNS Namespace resource.
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
    ///         var example = new Aws.ServiceDiscovery.PublicDnsNamespace("example", new Aws.ServiceDiscovery.PublicDnsNamespaceArgs
    ///         {
    ///             Description = "example",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Service Discovery Public DNS Namespace can be imported using the namespace ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace example 0123456789
    /// ```
    /// </summary>
    [AwsResourceType("aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace")]
    public partial class PublicDnsNamespace : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN that Amazon Route 53 assigns to the namespace when you create it.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description that you specify for the namespace when you create it.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
        /// </summary>
        [Output("hostedZone")]
        public Output<string> HostedZone { get; private set; } = null!;

        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the namespace.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a PublicDnsNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublicDnsNamespace(string name, PublicDnsNamespaceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace", name, args ?? new PublicDnsNamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublicDnsNamespace(string name, Input<string> id, PublicDnsNamespaceState? state = null, CustomResourceOptions? options = null)
            : base("aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PublicDnsNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublicDnsNamespace Get(string name, Input<string> id, PublicDnsNamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new PublicDnsNamespace(name, id, state, options);
        }
    }

    public sealed class PublicDnsNamespaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description that you specify for the namespace when you create it.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the namespace.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PublicDnsNamespaceArgs()
        {
        }
    }

    public sealed class PublicDnsNamespaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN that Amazon Route 53 assigns to the namespace when you create it.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The description that you specify for the namespace when you create it.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
        /// </summary>
        [Input("hostedZone")]
        public Input<string>? HostedZone { get; set; }

        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the namespace.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PublicDnsNamespaceState()
        {
        }
    }
}
