// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Elb
{
    /// <summary>
    /// Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.
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
    ///         var lb = new Aws.Elb.LoadBalancer("lb", new Aws.Elb.LoadBalancerArgs
    ///         {
    ///             AvailabilityZones = 
    ///             {
    ///                 "us-east-1a",
    ///             },
    ///             Listeners = 
    ///             {
    ///                 new Aws.Elb.Inputs.LoadBalancerListenerArgs
    ///                 {
    ///                     InstancePort = 8000,
    ///                     InstanceProtocol = "https",
    ///                     LbPort = 443,
    ///                     LbProtocol = "https",
    ///                     SslCertificateId = "arn:aws:iam::123456789012:server-certificate/certName",
    ///                 },
    ///             },
    ///         });
    ///         var foo = new Aws.Elb.SslNegotiationPolicy("foo", new Aws.Elb.SslNegotiationPolicyArgs
    ///         {
    ///             Attributes = 
    ///             {
    ///                 new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
    ///                 {
    ///                     Name = "Protocol-TLSv1",
    ///                     Value = "false",
    ///                 },
    ///                 new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
    ///                 {
    ///                     Name = "Protocol-TLSv1.1",
    ///                     Value = "false",
    ///                 },
    ///                 new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
    ///                 {
    ///                     Name = "Protocol-TLSv1.2",
    ///                     Value = "true",
    ///                 },
    ///                 new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
    ///                 {
    ///                     Name = "Server-Defined-Cipher-Order",
    ///                     Value = "true",
    ///                 },
    ///                 new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
    ///                 {
    ///                     Name = "ECDHE-RSA-AES128-GCM-SHA256",
    ///                     Value = "true",
    ///                 },
    ///                 new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
    ///                 {
    ///                     Name = "AES128-GCM-SHA256",
    ///                     Value = "true",
    ///                 },
    ///                 new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
    ///                 {
    ///                     Name = "EDH-RSA-DES-CBC3-SHA",
    ///                     Value = "false",
    ///                 },
    ///             },
    ///             LbPort = 443,
    ///             LoadBalancer = lb.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SslNegotiationPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// An SSL Negotiation policy attribute. Each has two properties:
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableArray<Outputs.SslNegotiationPolicyAttribute>> Attributes { get; private set; } = null!;

        /// <summary>
        /// The load balancer port to which the policy
        /// should be applied. This must be an active listener on the load
        /// balancer.
        /// </summary>
        [Output("lbPort")]
        public Output<int> LbPort { get; private set; } = null!;

        /// <summary>
        /// The load balancer to which the policy
        /// should be attached.
        /// </summary>
        [Output("loadBalancer")]
        public Output<string> LoadBalancer { get; private set; } = null!;

        /// <summary>
        /// The name of the attribute
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a SslNegotiationPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SslNegotiationPolicy(string name, SslNegotiationPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:elb/sslNegotiationPolicy:SslNegotiationPolicy", name, args ?? new SslNegotiationPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SslNegotiationPolicy(string name, Input<string> id, SslNegotiationPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:elb/sslNegotiationPolicy:SslNegotiationPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Alias { Type = "aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SslNegotiationPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SslNegotiationPolicy Get(string name, Input<string> id, SslNegotiationPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SslNegotiationPolicy(name, id, state, options);
        }
    }

    public sealed class SslNegotiationPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputList<Inputs.SslNegotiationPolicyAttributeArgs>? _attributes;

        /// <summary>
        /// An SSL Negotiation policy attribute. Each has two properties:
        /// </summary>
        public InputList<Inputs.SslNegotiationPolicyAttributeArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.SslNegotiationPolicyAttributeArgs>());
            set => _attributes = value;
        }

        /// <summary>
        /// The load balancer port to which the policy
        /// should be applied. This must be an active listener on the load
        /// balancer.
        /// </summary>
        [Input("lbPort", required: true)]
        public Input<int> LbPort { get; set; } = null!;

        /// <summary>
        /// The load balancer to which the policy
        /// should be attached.
        /// </summary>
        [Input("loadBalancer", required: true)]
        public Input<string> LoadBalancer { get; set; } = null!;

        /// <summary>
        /// The name of the attribute
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SslNegotiationPolicyArgs()
        {
        }
    }

    public sealed class SslNegotiationPolicyState : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputList<Inputs.SslNegotiationPolicyAttributeGetArgs>? _attributes;

        /// <summary>
        /// An SSL Negotiation policy attribute. Each has two properties:
        /// </summary>
        public InputList<Inputs.SslNegotiationPolicyAttributeGetArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.SslNegotiationPolicyAttributeGetArgs>());
            set => _attributes = value;
        }

        /// <summary>
        /// The load balancer port to which the policy
        /// should be applied. This must be an active listener on the load
        /// balancer.
        /// </summary>
        [Input("lbPort")]
        public Input<int>? LbPort { get; set; }

        /// <summary>
        /// The load balancer to which the policy
        /// should be attached.
        /// </summary>
        [Input("loadBalancer")]
        public Input<string>? LoadBalancer { get; set; }

        /// <summary>
        /// The name of the attribute
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SslNegotiationPolicyState()
        {
        }
    }
}
