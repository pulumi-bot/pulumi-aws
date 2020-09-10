// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Workspaces.Inputs
{

    public sealed class DirectorySelfServicePermissionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("changeComputeType")]
        public Input<bool>? ChangeComputeType { get; set; }

        [Input("increaseVolumeSize")]
        public Input<bool>? IncreaseVolumeSize { get; set; }

        [Input("rebuildWorkspace")]
        public Input<bool>? RebuildWorkspace { get; set; }

        [Input("restartWorkspace")]
        public Input<bool>? RestartWorkspace { get; set; }

        [Input("switchRunningMode")]
        public Input<bool>? SwitchRunningMode { get; set; }

        public DirectorySelfServicePermissionsGetArgs()
        {
        }
    }
}
