// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancing
{
    /// <summary>
    /// Attaches a load balancer policy to an ELB backend server.
    /// 
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/load_balancer_backend_server_policy.html.markdown.
    /// </summary>
    public partial class LoadBalancerBackendServerPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The instance port to apply the policy to.
        /// </summary>
        [Output("instancePort")]
        public Output<int> InstancePort { get; private set; } = null!;

        /// <summary>
        /// The load balancer to attach the policy to.
        /// </summary>
        [Output("loadBalancerName")]
        public Output<string> LoadBalancerName { get; private set; } = null!;

        /// <summary>
        /// List of Policy Names to apply to the backend server.
        /// </summary>
        [Output("policyNames")]
        public Output<ImmutableArray<string>> PolicyNames { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancerBackendServerPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancerBackendServerPolicy(string name, LoadBalancerBackendServerPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancerBackendServerPolicy(string name, Input<string> id, LoadBalancerBackendServerPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancerBackendServerPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancerBackendServerPolicy Get(string name, Input<string> id, LoadBalancerBackendServerPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancerBackendServerPolicy(name, id, state, options);
        }
    }

    public sealed class LoadBalancerBackendServerPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance port to apply the policy to.
        /// </summary>
        [Input("instancePort", required: true)]
        public Input<int> InstancePort { get; set; } = null!;

        /// <summary>
        /// The load balancer to attach the policy to.
        /// </summary>
        [Input("loadBalancerName", required: true)]
        public Input<string> LoadBalancerName { get; set; } = null!;

        [Input("policyNames")]
        private InputList<string>? _policyNames;

        /// <summary>
        /// List of Policy Names to apply to the backend server.
        /// </summary>
        public InputList<string> PolicyNames
        {
            get => _policyNames ?? (_policyNames = new InputList<string>());
            set => _policyNames = value;
        }

        public LoadBalancerBackendServerPolicyArgs()
        {
        }
    }

    public sealed class LoadBalancerBackendServerPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance port to apply the policy to.
        /// </summary>
        [Input("instancePort")]
        public Input<int>? InstancePort { get; set; }

        /// <summary>
        /// The load balancer to attach the policy to.
        /// </summary>
        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        [Input("policyNames")]
        private InputList<string>? _policyNames;

        /// <summary>
        /// List of Policy Names to apply to the backend server.
        /// </summary>
        public InputList<string> PolicyNames
        {
            get => _policyNames ?? (_policyNames = new InputList<string>());
            set => _policyNames = value;
        }

        public LoadBalancerBackendServerPolicyState()
        {
        }
    }
}
