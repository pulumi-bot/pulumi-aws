// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3.Outputs
{

    [OutputType]
    public sealed class AnalyticsConfigurationStorageClassAnalysisDataExportDestination
    {
        public readonly Outputs.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestination S3BucketDestination;

        [OutputConstructor]
        private AnalyticsConfigurationStorageClassAnalysisDataExportDestination(Outputs.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestination s3BucketDestination)
        {
            S3BucketDestination = s3BucketDestination;
        }
    }
}
