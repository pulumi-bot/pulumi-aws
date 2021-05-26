// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApplicationLoadBalancing
{
    [Obsolete(@"aws.applicationloadbalancing.getListener has been deprecated in favor of aws.alb.getListener")]
    public static class GetListener
    {
        /// <summary>
        /// &gt; **Note:** `aws.alb.Listener` is known as `aws.lb.Listener`. The functionality is identical.
        /// 
        /// Provides information about a Load Balancer Listener.
        /// 
        /// This data source can prove useful when a module accepts an LB Listener as an input variable and needs to know the LB it is attached to, or other information specific to the listener in question.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var config = new Config();
        ///         var listenerArn = config.Require("listenerArn");
        ///         var listener = Output.Create(Aws.LB.GetListener.InvokeAsync(new Aws.LB.GetListenerArgs
        ///         {
        ///             Arn = listenerArn,
        ///         }));
        ///         var selected = Output.Create(Aws.LB.GetLoadBalancer.InvokeAsync(new Aws.LB.GetLoadBalancerArgs
        ///         {
        ///             Name = "default-public",
        ///         }));
        ///         var selected443 = selected.Apply(selected =&gt; Output.Create(Aws.LB.GetListener.InvokeAsync(new Aws.LB.GetListenerArgs
        ///         {
        ///             LoadBalancerArn = selected.Arn,
        ///             Port = 443,
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetListenerResult> InvokeAsync(GetListenerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetListenerResult>("aws:applicationloadbalancing/getListener:getListener", args ?? new GetListenerArgs(), options.WithVersion());

        public static Output<GetListenerResult> Apply(GetListenerApplyArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetListenerApplyArgs();
            return Pulumi.Output.All(
                args.Arn.Box(),
                args.LoadBalancerArn.Box(),
                args.Port.Box(),
                args.Tags.Box()
            ).Apply(a => {
                    var args = new GetListenerArgs();
                    a[0].Set(args, nameof(args.Arn));
                    a[1].Set(args, nameof(args.LoadBalancerArn));
                    a[2].Set(args, nameof(args.Port));
                    a[3].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetListenerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ARN of the listener. Required if `load_balancer_arn` and `port` is not set.
        /// </summary>
        [Input("arn")]
        public string? Arn { get; set; }

        /// <summary>
        /// ARN of the load balancer. Required if `arn` is not set.
        /// </summary>
        [Input("loadBalancerArn")]
        public string? LoadBalancerArn { get; set; }

        /// <summary>
        /// Port of the listener. Required if `arn` is not set.
        /// </summary>
        [Input("port")]
        public int? Port { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetListenerArgs()
        {
        }
    }

    public sealed class GetListenerApplyArgs
    {
        /// <summary>
        /// ARN of the listener. Required if `load_balancer_arn` and `port` is not set.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// ARN of the load balancer. Required if `arn` is not set.
        /// </summary>
        [Input("loadBalancerArn")]
        public Input<string>? LoadBalancerArn { get; set; }

        /// <summary>
        /// Port of the listener. Required if `arn` is not set.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetListenerApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetListenerResult
    {
        public readonly string AlpnPolicy;
        public readonly string Arn;
        public readonly string CertificateArn;
        public readonly ImmutableArray<Outputs.GetListenerDefaultActionResult> DefaultActions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LoadBalancerArn;
        public readonly int Port;
        public readonly string Protocol;
        public readonly string SslPolicy;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetListenerResult(
            string alpnPolicy,

            string arn,

            string certificateArn,

            ImmutableArray<Outputs.GetListenerDefaultActionResult> defaultActions,

            string id,

            string loadBalancerArn,

            int port,

            string protocol,

            string sslPolicy,

            ImmutableDictionary<string, string> tags)
        {
            AlpnPolicy = alpnPolicy;
            Arn = arn;
            CertificateArn = certificateArn;
            DefaultActions = defaultActions;
            Id = id;
            LoadBalancerArn = loadBalancerArn;
            Port = port;
            Protocol = protocol;
            SslPolicy = sslPolicy;
            Tags = tags;
        }
    }
}
