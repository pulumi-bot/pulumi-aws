// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild.Inputs
{

    public sealed class ProjectLogsConfigArgs : Pulumi.ResourceArgs
    {
        [Input("cloudwatchLogs")]
        public Input<Inputs.ProjectLogsConfigCloudwatchLogsArgs>? CloudwatchLogs { get; set; }

        [Input("s3Logs")]
        public Input<Inputs.ProjectLogsConfigS3LogsArgs>? S3Logs { get; set; }

        public ProjectLogsConfigArgs()
        {
        }
    }
}
