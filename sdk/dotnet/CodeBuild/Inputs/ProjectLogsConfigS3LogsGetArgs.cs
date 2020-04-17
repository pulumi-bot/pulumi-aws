// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild.Inputs
{

    public sealed class ProjectLogsConfigS3LogsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, output artifacts will not be encrypted. If `type` is set to `NO_ARTIFACTS` then this value will be ignored. Defaults to `false`.
        /// * `override_artifact_name` (Optional) If set to true, a name specified in the build spec file overrides the artifact name.
        /// * `override_artifact_name` (Optional) If set to true, a name specified in the build spec file overrides the artifact name.
        /// </summary>
        [Input("encryptionDisabled")]
        public Input<bool>? EncryptionDisabled { get; set; }

        /// <summary>
        /// Information about the build output artifact location. If `type` is set to `CODEPIPELINE` or `NO_ARTIFACTS` then this value will be ignored. If `type` is set to `S3`, this is the name of the output bucket.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Current status of logs in CloudWatch Logs for a build project. Valid values: `ENABLED`, `DISABLED`. Defaults to `ENABLED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ProjectLogsConfigS3LogsGetArgs()
        {
        }
    }
}
