// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lb
{
    public static partial class Invokes
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
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lb_target_group.html.markdown.
        /// </summary>
        public static Task<GetTargetGroupResult> GetTargetGroup(GetTargetGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetGroupResult>("aws:lb/getTargetGroup:getTargetGroup", args, options.WithVersion());
    }

    public sealed class GetTargetGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full ARN of the target group.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The unique name of the target group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
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
        public readonly bool LambdaMultiValueHeadersEnabled;
        public readonly string Name;
        public readonly int Port;
        public readonly string Protocol;
        public readonly bool ProxyProtocolV2;
        public readonly int SlowStart;
        public readonly Outputs.GetTargetGroupStickinessResult Stickiness;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string TargetType;
        public readonly string VpcId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetTargetGroupResult(
            string arn,
            string arnSuffix,
            int deregistrationDelay,
            Outputs.GetTargetGroupHealthCheckResult healthCheck,
            bool lambdaMultiValueHeadersEnabled,
            string name,
            int port,
            string protocol,
            bool proxyProtocolV2,
            int slowStart,
            Outputs.GetTargetGroupStickinessResult stickiness,
            ImmutableDictionary<string, string> tags,
            string targetType,
            string vpcId,
            string id)
        {
            Arn = arn;
            ArnSuffix = arnSuffix;
            DeregistrationDelay = deregistrationDelay;
            HealthCheck = healthCheck;
            LambdaMultiValueHeadersEnabled = lambdaMultiValueHeadersEnabled;
            Name = name;
            Port = port;
            Protocol = protocol;
            ProxyProtocolV2 = proxyProtocolV2;
            SlowStart = slowStart;
            Stickiness = stickiness;
            Tags = tags;
            TargetType = targetType;
            VpcId = vpcId;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetTargetGroupHealthCheckResult
    {
        public readonly bool Enabled;
        public readonly int HealthyThreshold;
        public readonly int Interval;
        public readonly string Matcher;
        public readonly string Path;
        public readonly string Port;
        public readonly string Protocol;
        public readonly int Timeout;
        public readonly int UnhealthyThreshold;

        [OutputConstructor]
        private GetTargetGroupHealthCheckResult(
            bool enabled,
            int healthyThreshold,
            int interval,
            string matcher,
            string path,
            string port,
            string protocol,
            int timeout,
            int unhealthyThreshold)
        {
            Enabled = enabled;
            HealthyThreshold = healthyThreshold;
            Interval = interval;
            Matcher = matcher;
            Path = path;
            Port = port;
            Protocol = protocol;
            Timeout = timeout;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }

    [OutputType]
    public sealed class GetTargetGroupStickinessResult
    {
        public readonly int CookieDuration;
        public readonly bool Enabled;
        public readonly string Type;

        [OutputConstructor]
        private GetTargetGroupStickinessResult(
            int cookieDuration,
            bool enabled,
            string type)
        {
            CookieDuration = cookieDuration;
            Enabled = enabled;
            Type = type;
        }
    }
    }
}
