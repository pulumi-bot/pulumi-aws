// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GlobalAccelerator
{
    /// <summary>
    /// Creates a Global Accelerator accelerator.
    /// </summary>
    public partial class Accelerator : Pulumi.CustomResource
    {
        /// <summary>
        /// The attributes of the accelerator. Fields documented below.
        /// </summary>
        [Output("attributes")]
        public Output<Outputs.AcceleratorAttributes?> Attributes { get; private set; } = null!;

        /// <summary>
        /// The DNS name of the accelerator. For example, `a5d53ff5ee6bca4ce.awsglobalaccelerator.com`.
        /// * `hosted_zone_id` --  The Global Accelerator Route 53 zone ID that can be used to
        /// route an [Alias Resource Record Set](https://docs.aws.amazon.com/Route53/latest/APIReference/API_AliasTarget.html) to the Global Accelerator. This attribute
        /// is simply an alias for the zone ID `Z2BJ6XQ5FK7U4H`.
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the accelerator is enabled. The value is true or false. The default value is true.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("hostedZoneId")]
        public Output<string> HostedZoneId { get; private set; } = null!;

        /// <summary>
        /// The value for the address type must be `IPV4`.
        /// </summary>
        [Output("ipAddressType")]
        public Output<string?> IpAddressType { get; private set; } = null!;

        /// <summary>
        /// IP address set associated with the accelerator.
        /// </summary>
        [Output("ipSets")]
        public Output<ImmutableArray<Outputs.AcceleratorIpSet>> IpSets { get; private set; } = null!;

        /// <summary>
        /// The name of the accelerator.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Accelerator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Accelerator(string name, AcceleratorArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:globalaccelerator/accelerator:Accelerator", name, args ?? new AcceleratorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Accelerator(string name, Input<string> id, AcceleratorState? state = null, CustomResourceOptions? options = null)
            : base("aws:globalaccelerator/accelerator:Accelerator", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Accelerator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Accelerator Get(string name, Input<string> id, AcceleratorState? state = null, CustomResourceOptions? options = null)
        {
            return new Accelerator(name, id, state, options);
        }
    }

    public sealed class AcceleratorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The attributes of the accelerator. Fields documented below.
        /// </summary>
        [Input("attributes")]
        public Input<Inputs.AcceleratorAttributesArgs>? Attributes { get; set; }

        /// <summary>
        /// Indicates whether the accelerator is enabled. The value is true or false. The default value is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The value for the address type must be `IPV4`.
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        /// <summary>
        /// The name of the accelerator.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AcceleratorArgs()
        {
        }
    }

    public sealed class AcceleratorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The attributes of the accelerator. Fields documented below.
        /// </summary>
        [Input("attributes")]
        public Input<Inputs.AcceleratorAttributesGetArgs>? Attributes { get; set; }

        /// <summary>
        /// The DNS name of the accelerator. For example, `a5d53ff5ee6bca4ce.awsglobalaccelerator.com`.
        /// * `hosted_zone_id` --  The Global Accelerator Route 53 zone ID that can be used to
        /// route an [Alias Resource Record Set](https://docs.aws.amazon.com/Route53/latest/APIReference/API_AliasTarget.html) to the Global Accelerator. This attribute
        /// is simply an alias for the zone ID `Z2BJ6XQ5FK7U4H`.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// Indicates whether the accelerator is enabled. The value is true or false. The default value is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("hostedZoneId")]
        public Input<string>? HostedZoneId { get; set; }

        /// <summary>
        /// The value for the address type must be `IPV4`.
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        [Input("ipSets")]
        private InputList<Inputs.AcceleratorIpSetGetArgs>? _ipSets;

        /// <summary>
        /// IP address set associated with the accelerator.
        /// </summary>
        public InputList<Inputs.AcceleratorIpSetGetArgs> IpSets
        {
            get => _ipSets ?? (_ipSets = new InputList<Inputs.AcceleratorIpSetGetArgs>());
            set => _ipSets = value;
        }

        /// <summary>
        /// The name of the accelerator.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AcceleratorState()
        {
        }
    }
}
