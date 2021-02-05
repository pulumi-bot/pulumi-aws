// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2
{
    [Obsolete(@"aws.elasticloadbalancingv2.getTargetGroup has been deprecated in favor of aws.lb.getTargetGroup")]
    public static class GetTargetGroup
    {
        /// <summary>
        /// &gt; **Note:** `aws.alb.TargetGroup` is known as `aws.lb.TargetGroup`. The functionality is identical.
        /// 
        /// Provides information about a Load Balancer Target Group.
        /// 
        /// This data source can prove useful when a module accepts an LB Target Group as an
        /// input variable and needs to know its attributes. It can also be used to get the ARN of
        /// an LB Target Group for use in other resources, given LB Target Group name.
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
        ///         var lbTgArn = config.Get("lbTgArn") ?? "";
        ///         var lbTgName = config.Get("lbTgName") ?? "";
        ///         var test = Output.Create(Aws.LB.GetTargetGroup.InvokeAsync(new Aws.LB.GetTargetGroupArgs
        ///         {
        ///             Arn = lbTgArn,
        ///             Name = lbTgName,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTargetGroupResult> InvokeAsync(GetTargetGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetGroupResult>("aws:elasticloadbalancingv2/getTargetGroup:getTargetGroup", args ?? new GetTargetGroupArgs(), options.WithVersion());
    }


    public sealed class GetTargetGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The full ARN of the target group.
        /// </summary>
        [Input("arn")]
        public string? Arn { get; set; }

        /// <summary>
        /// The unique name of the target group.
        /// </summary>
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
        public readonly string ProtocolVersion;
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

            string protocolVersion,

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
            ProtocolVersion = protocolVersion;
            ProxyProtocolV2 = proxyProtocolV2;
            SlowStart = slowStart;
            Stickiness = stickiness;
            Tags = tags;
            TargetType = targetType;
            VpcId = vpcId;
        }
    }
}
