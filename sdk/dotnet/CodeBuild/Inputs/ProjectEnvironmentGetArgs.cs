// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild.Inputs
{

    public sealed class ProjectEnvironmentGetArgs : Pulumi.ResourceArgs
    {
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("computeType", required: true)]
        public Input<string> ComputeType { get; set; } = null!;

        [Input("environmentVariables")]
        private InputList<Inputs.ProjectEnvironmentEnvironmentVariableGetArgs>? _environmentVariables;
        public InputList<Inputs.ProjectEnvironmentEnvironmentVariableGetArgs> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputList<Inputs.ProjectEnvironmentEnvironmentVariableGetArgs>());
            set => _environmentVariables = value;
        }

        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        [Input("imagePullCredentialsType")]
        public Input<string>? ImagePullCredentialsType { get; set; }

        [Input("privilegedMode")]
        public Input<bool>? PrivilegedMode { get; set; }

        [Input("registryCredential")]
        public Input<Inputs.ProjectEnvironmentRegistryCredentialGetArgs>? RegistryCredential { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ProjectEnvironmentGetArgs()
        {
        }
    }
}
