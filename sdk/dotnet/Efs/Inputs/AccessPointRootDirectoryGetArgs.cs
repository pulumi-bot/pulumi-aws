// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Efs.Inputs
{

    public sealed class AccessPointRootDirectoryGetArgs : Pulumi.ResourceArgs
    {
        [Input("creationInfo")]
        public Input<Inputs.AccessPointRootDirectoryCreationInfoGetArgs>? CreationInfo { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        public AccessPointRootDirectoryGetArgs()
        {
        }
    }
}
