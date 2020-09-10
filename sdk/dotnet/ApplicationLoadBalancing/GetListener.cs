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
        public static Task<GetListenerResult> InvokeAsync(GetListenerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetListenerResult>("aws:applicationloadbalancing/getListener:getListener", args ?? new GetListenerArgs(), options.WithVersion());
    }


    public sealed class GetListenerArgs : Pulumi.InvokeArgs
    {
        [Input("arn")]
        public string? Arn { get; set; }

        [Input("loadBalancerArn")]
        public string? LoadBalancerArn { get; set; }

        [Input("port")]
        public int? Port { get; set; }

        public GetListenerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetListenerResult
    {
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

        [OutputConstructor]
        private GetListenerResult(
            string arn,

            string certificateArn,

            ImmutableArray<Outputs.GetListenerDefaultActionResult> defaultActions,

            string id,

            string loadBalancerArn,

            int port,

            string protocol,

            string sslPolicy)
        {
            Arn = arn;
            CertificateArn = certificateArn;
            DefaultActions = defaultActions;
            Id = id;
            LoadBalancerArn = loadBalancerArn;
            Port = port;
            Protocol = protocol;
            SslPolicy = sslPolicy;
        }
    }
}
