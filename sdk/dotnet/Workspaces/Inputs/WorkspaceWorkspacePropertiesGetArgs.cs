// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Workspaces.Inputs
{

    public sealed class WorkspaceWorkspacePropertiesGetArgs : Pulumi.ResourceArgs
    {
        [Input("computeTypeName")]
        public Input<string>? ComputeTypeName { get; set; }

        [Input("rootVolumeSizeGib")]
        public Input<int>? RootVolumeSizeGib { get; set; }

        [Input("runningMode")]
        public Input<string>? RunningMode { get; set; }

        [Input("runningModeAutoStopTimeoutInMinutes")]
        public Input<int>? RunningModeAutoStopTimeoutInMinutes { get; set; }

        [Input("userVolumeSizeGib")]
        public Input<int>? UserVolumeSizeGib { get; set; }

        public WorkspaceWorkspacePropertiesGetArgs()
        {
        }
    }
}
