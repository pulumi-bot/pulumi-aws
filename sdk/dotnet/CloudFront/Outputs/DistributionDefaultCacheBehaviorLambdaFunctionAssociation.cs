// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Outputs
{

    [OutputType]
    public sealed class DistributionDefaultCacheBehaviorLambdaFunctionAssociation
    {
        public readonly string EventType;
        public readonly bool? IncludeBody;
        public readonly string LambdaArn;

        [OutputConstructor]
        private DistributionDefaultCacheBehaviorLambdaFunctionAssociation(
            string eventType,

            bool? includeBody,

            string lambdaArn)
        {
            EventType = eventType;
            IncludeBody = includeBody;
            LambdaArn = lambdaArn;
        }
    }
}
