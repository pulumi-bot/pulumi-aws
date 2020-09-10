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
    /// Provides a load balancer policy, which can be attached to an ELB listener or backend server.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var wu_tang = new Aws.Elb.LoadBalancer("wu-tang", new Aws.Elb.LoadBalancerArgs
    ///         {
    ///             AvailabilityZones = 
    ///             {
    ///                 "us-east-1a",
    ///             },
    ///             Listeners = 
    ///             {
    ///                 new Aws.Elb.Inputs.LoadBalancerListenerArgs
    ///                 {
    ///                     InstancePort = 443,
    ///                     InstanceProtocol = "http",
    ///                     LbPort = 443,
    ///                     LbProtocol = "https",
    ///                     SslCertificateId = "arn:aws:iam::000000000000:server-certificate/wu-tang.net",
    ///                 },
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "Name", "wu-tang" },
    ///             },
    ///         });
    ///         var wu_tang_ca_pubkey_policy = new Aws.Elb.LoadBalancerPolicy("wu-tang-ca-pubkey-policy", new Aws.Elb.LoadBalancerPolicyArgs
    ///         {
    ///             LoadBalancerName = wu_tang.Name,
    ///             PolicyName = "wu-tang-ca-pubkey-policy",
    ///             PolicyTypeName = "PublicKeyPolicyType",
    ///             PolicyAttributes = 
    ///             {
    ///                 new Aws.Elb.Inputs.LoadBalancerPolicyPolicyAttributeArgs
    ///                 {
    ///                     Name = "PublicKey",
    ///                     Value = File.ReadAllText("wu-tang-pubkey"),
    ///                 },
    ///             },
    ///         });
    ///         var wu_tang_root_ca_backend_auth_policy = new Aws.Elb.LoadBalancerPolicy("wu-tang-root-ca-backend-auth-policy", new Aws.Elb.LoadBalancerPolicyArgs
    ///         {
    ///             LoadBalancerName = wu_tang.Name,
    ///             PolicyName = "wu-tang-root-ca-backend-auth-policy",
    ///             PolicyTypeName = "BackendServerAuthenticationPolicyType",
    ///             PolicyAttributes = 
    ///             {
    ///                 new Aws.Elb.Inputs.LoadBalancerPolicyPolicyAttributeArgs
    ///                 {
    ///                     Name = "PublicKeyPolicyName",
    ///                     Value = aws_load_balancer_policy.Wu_tang_root_ca_pubkey_policy.Policy_name,
    ///                 },
    ///             },
    ///         });
    ///         var wu_tang_ssl = new Aws.Elb.LoadBalancerPolicy("wu-tang-ssl", new Aws.Elb.LoadBalancerPolicyArgs
    ///         {
    ///             LoadBalancerName = wu_tang.Name,
    ///             PolicyName = "wu-tang-ssl",
    ///             PolicyTypeName = "SSLNegotiationPolicyType",
    ///             PolicyAttributes = 
    ///             {
    ///                 new Aws.Elb.Inputs.LoadBalancerPolicyPolicyAttributeArgs
    ///                 {
    ///                     Name = "ECDHE-ECDSA-AES128-GCM-SHA256",
    ///                     Value = "true",
    ///                 },
    ///                 new Aws.Elb.Inputs.LoadBalancerPolicyPolicyAttributeArgs
    ///                 {
    ///                     Name = "Protocol-TLSv1.2",
    ///                     Value = "true",
    ///                 },
    ///             },
    ///         });
    ///         var wu_tang_ssl_tls_1_1 = new Aws.Elb.LoadBalancerPolicy("wu-tang-ssl-tls-1-1", new Aws.Elb.LoadBalancerPolicyArgs
    ///         {
    ///             LoadBalancerName = wu_tang.Name,
    ///             PolicyName = "wu-tang-ssl",
    ///             PolicyTypeName = "SSLNegotiationPolicyType",
    ///             PolicyAttributes = 
    ///             {
    ///                 new Aws.Elb.Inputs.LoadBalancerPolicyPolicyAttributeArgs
    ///                 {
    ///                     Name = "Reference-Security-Policy",
    ///                     Value = "ELBSecurityPolicy-TLS-1-1-2017-01",
    ///                 },
    ///             },
    ///         });
    ///         var wu_tang_backend_auth_policies_443 = new Aws.Elb.LoadBalancerBackendServerPolicy("wu-tang-backend-auth-policies-443", new Aws.Elb.LoadBalancerBackendServerPolicyArgs
    ///         {
    ///             LoadBalancerName = wu_tang.Name,
    ///             InstancePort = 443,
    ///             PolicyNames = 
    ///             {
    ///                 wu_tang_root_ca_backend_auth_policy.PolicyName,
    ///             },
    ///         });
    ///         var wu_tang_listener_policies_443 = new Aws.Elb.ListenerPolicy("wu-tang-listener-policies-443", new Aws.Elb.ListenerPolicyArgs
    ///         {
    ///             LoadBalancerName = wu_tang.Name,
    ///             LoadBalancerPort = 443,
    ///             PolicyNames = 
    ///             {
    ///                 wu_tang_ssl.PolicyName,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Where the file `pubkey` in the current directory contains only the _public key_ of the certificate.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// This example shows how to enable backend authentication for an ELB as well as customize the TLS settings.
    /// </summary>
    public partial class LoadBalancerPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The load balancer on which the policy is defined.
        /// </summary>
        [Output("loadBalancerName")]
        public Output<string> LoadBalancerName { get; private set; } = null!;

        /// <summary>
        /// Policy attribute to apply to the policy.
        /// </summary>
        [Output("policyAttributes")]
        public Output<ImmutableArray<Outputs.LoadBalancerPolicyPolicyAttribute>> PolicyAttributes { get; private set; } = null!;

        /// <summary>
        /// The name of the load balancer policy.
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// The policy type.
        /// </summary>
        [Output("policyTypeName")]
        public Output<string> PolicyTypeName { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancerPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancerPolicy(string name, LoadBalancerPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:elb/loadBalancerPolicy:LoadBalancerPolicy", name, args ?? new LoadBalancerPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancerPolicy(string name, Input<string> id, LoadBalancerPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:elb/loadBalancerPolicy:LoadBalancerPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "aws:elasticloadbalancing/loadBalancerPolicy:LoadBalancerPolicy"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoadBalancerPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancerPolicy Get(string name, Input<string> id, LoadBalancerPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancerPolicy(name, id, state, options);
        }
    }

    public sealed class LoadBalancerPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The load balancer on which the policy is defined.
        /// </summary>
        [Input("loadBalancerName", required: true)]
        public Input<string> LoadBalancerName { get; set; } = null!;

        [Input("policyAttributes")]
        private InputList<Inputs.LoadBalancerPolicyPolicyAttributeArgs>? _policyAttributes;

        /// <summary>
        /// Policy attribute to apply to the policy.
        /// </summary>
        public InputList<Inputs.LoadBalancerPolicyPolicyAttributeArgs> PolicyAttributes
        {
            get => _policyAttributes ?? (_policyAttributes = new InputList<Inputs.LoadBalancerPolicyPolicyAttributeArgs>());
            set => _policyAttributes = value;
        }

        /// <summary>
        /// The name of the load balancer policy.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// The policy type.
        /// </summary>
        [Input("policyTypeName", required: true)]
        public Input<string> PolicyTypeName { get; set; } = null!;

        public LoadBalancerPolicyArgs()
        {
        }
    }

    public sealed class LoadBalancerPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The load balancer on which the policy is defined.
        /// </summary>
        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        [Input("policyAttributes")]
        private InputList<Inputs.LoadBalancerPolicyPolicyAttributeGetArgs>? _policyAttributes;

        /// <summary>
        /// Policy attribute to apply to the policy.
        /// </summary>
        public InputList<Inputs.LoadBalancerPolicyPolicyAttributeGetArgs> PolicyAttributes
        {
            get => _policyAttributes ?? (_policyAttributes = new InputList<Inputs.LoadBalancerPolicyPolicyAttributeGetArgs>());
            set => _policyAttributes = value;
        }

        /// <summary>
        /// The name of the load balancer policy.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// The policy type.
        /// </summary>
        [Input("policyTypeName")]
        public Input<string>? PolicyTypeName { get; set; }

        public LoadBalancerPolicyState()
        {
        }
    }
}
