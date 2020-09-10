// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LB
{
    public static class GetTargetGroup
    {
        public static Task<GetTargetGroupResult> InvokeAsync(GetTargetGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetGroupResult>("aws:lb/getTargetGroup:getTargetGroup", args ?? new GetTargetGroupArgs(), options.WithVersion());
    }


    public sealed class GetTargetGroupArgs : Pulumi.InvokeArgs
    {
        [Input("arn")]
        public string? Arn { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetTargetGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetGroupResult
    {
        public readonly string Arn;
        public readonly string ArnSuffix;
        public readonly int DeregistrationDelay;
        public readonly Outputs.GetTargetGroupHealthCheckResult HealthCheck;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool LambdaMultiValueHeadersEnabled;
        public readonly string LoadBalancingAlgorithmType;
        public readonly string Name;
        public readonly int Port;
        public readonly string Protocol;
        public readonly bool ProxyProtocolV2;
        public readonly int SlowStart;
        public readonly Outputs.GetTargetGroupStickinessResult Stickiness;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string TargetType;
        public readonly string VpcId;

        [OutputConstructor]
        private GetTargetGroupResult(
            string arn,

            string arnSuffix,

            int deregistrationDelay,

            Outputs.GetTargetGroupHealthCheckResult healthCheck,

            string id,

            bool lambdaMultiValueHeadersEnabled,

            string loadBalancingAlgorithmType,

            string name,

            int port,

            string protocol,

            bool proxyProtocolV2,

            int slowStart,

            Outputs.GetTargetGroupStickinessResult stickiness,

            ImmutableDictionary<string, string> tags,

            string targetType,

            string vpcId)
        {
            Arn = arn;
            ArnSuffix = arnSuffix;
            DeregistrationDelay = deregistrationDelay;
            HealthCheck = healthCheck;
            Id = id;
            LambdaMultiValueHeadersEnabled = lambdaMultiValueHeadersEnabled;
            LoadBalancingAlgorithmType = loadBalancingAlgorithmType;
            Name = name;
            Port = port;
            Protocol = protocol;
            ProxyProtocolV2 = proxyProtocolV2;
            SlowStart = slowStart;
            Stickiness = stickiness;
            Tags = tags;
            TargetType = targetType;
            VpcId = vpcId;
        }
    }
}
