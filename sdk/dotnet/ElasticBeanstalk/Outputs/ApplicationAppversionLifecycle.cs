// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticBeanstalk.Outputs
{

    [OutputType]
    public sealed class ApplicationAppversionLifecycle
    {
        public readonly bool? DeleteSourceFromS3;
        public readonly int? MaxAgeInDays;
        public readonly int? MaxCount;
        public readonly string ServiceRole;

        [OutputConstructor]
        private ApplicationAppversionLifecycle(
            bool? deleteSourceFromS3,

            int? maxAgeInDays,

            int? maxCount,

            string serviceRole)
        {
            DeleteSourceFromS3 = deleteSourceFromS3;
            MaxAgeInDays = maxAgeInDays;
            MaxCount = maxCount;
            ServiceRole = serviceRole;
        }
    }
}
