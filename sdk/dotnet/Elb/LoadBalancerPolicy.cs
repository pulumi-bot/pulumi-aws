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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/load_balancer_policy.html.markdown.
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
        public Output<ImmutableArray<Outputs.LoadBalancerPolicyPolicyAttributes>> PolicyAttributes { get; private set; } = null!;

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
            : base("aws:elb/loadBalancerPolicy:LoadBalancerPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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
                Version = Utilities.Version,                Aliases = { new Alias { Type = "aws:elasticloadbalancing/loadBalancerPolicy:LoadBalancerPolicy" } },
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
        private InputList<Inputs.LoadBalancerPolicyPolicyAttributesArgs>? _policyAttributes;

        /// <summary>
        /// Policy attribute to apply to the policy.
        /// </summary>
        public InputList<Inputs.LoadBalancerPolicyPolicyAttributesArgs> PolicyAttributes
        {
            get => _policyAttributes ?? (_policyAttributes = new InputList<Inputs.LoadBalancerPolicyPolicyAttributesArgs>());
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
        private InputList<Inputs.LoadBalancerPolicyPolicyAttributesGetArgs>? _policyAttributes;

        /// <summary>
        /// Policy attribute to apply to the policy.
        /// </summary>
        public InputList<Inputs.LoadBalancerPolicyPolicyAttributesGetArgs> PolicyAttributes
        {
            get => _policyAttributes ?? (_policyAttributes = new InputList<Inputs.LoadBalancerPolicyPolicyAttributesGetArgs>());
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

    namespace Inputs
    {

    public sealed class LoadBalancerPolicyPolicyAttributesArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public LoadBalancerPolicyPolicyAttributesArgs()
        {
        }
    }

    public sealed class LoadBalancerPolicyPolicyAttributesGetArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public LoadBalancerPolicyPolicyAttributesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class LoadBalancerPolicyPolicyAttributes
    {
        public readonly string? Name;
        public readonly string? Value;

        [OutputConstructor]
        private LoadBalancerPolicyPolicyAttributes(
            string? name,
            string? value)
        {
            Name = name;
            Value = value;
        }
    }
    }
}
