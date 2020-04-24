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
    /// Provides a load balancer cookie stickiness policy, which allows an ELB to control the sticky session lifetime of the browser.
    /// 
    /// 
    /// 
    /// Deprecated: aws.LoadBalancerCookieStickinessPolicy has been deprecated in favour of aws.LoadBalancerCookieStickinessPolicy
    /// </summary>
    public partial class LoadBalancerCookieStickinessPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The time period after which
        /// the session cookie should be considered stale, expressed in seconds.
        /// </summary>
        [Output("cookieExpirationPeriod")]
        public Output<int?> CookieExpirationPeriod { get; private set; } = null!;

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
        /// The name of the stickiness policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancerCookieStickinessPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancerCookieStickinessPolicy(string name, LoadBalancerCookieStickinessPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticloadbalancing/loadBalancerCookieStickinessPolicy:LoadBalancerCookieStickinessPolicy", name, args ?? new LoadBalancerCookieStickinessPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancerCookieStickinessPolicy(string name, Input<string> id, LoadBalancerCookieStickinessPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticloadbalancing/loadBalancerCookieStickinessPolicy:LoadBalancerCookieStickinessPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancerCookieStickinessPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancerCookieStickinessPolicy Get(string name, Input<string> id, LoadBalancerCookieStickinessPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancerCookieStickinessPolicy(name, id, state, options);
        }
    }

    public sealed class LoadBalancerCookieStickinessPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time period after which
        /// the session cookie should be considered stale, expressed in seconds.
        /// </summary>
        [Input("cookieExpirationPeriod")]
        public Input<int>? CookieExpirationPeriod { get; set; }

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
        /// The name of the stickiness policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public LoadBalancerCookieStickinessPolicyArgs()
        {
        }
    }

    public sealed class LoadBalancerCookieStickinessPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time period after which
        /// the session cookie should be considered stale, expressed in seconds.
        /// </summary>
        [Input("cookieExpirationPeriod")]
        public Input<int>? CookieExpirationPeriod { get; set; }

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
        /// The name of the stickiness policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public LoadBalancerCookieStickinessPolicyState()
        {
        }
    }
}
