// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild.Inputs
{

    public sealed class ProjectSourceGetArgs : Pulumi.ResourceArgs
    {
        [Input("auths")]
        private InputList<Inputs.ProjectSourceAuthGetArgs>? _auths;
        public InputList<Inputs.ProjectSourceAuthGetArgs> Auths
        {
            get => _auths ?? (_auths = new InputList<Inputs.ProjectSourceAuthGetArgs>());
            set => _auths = value;
        }

        [Input("buildspec")]
        public Input<string>? Buildspec { get; set; }

        [Input("gitCloneDepth")]
        public Input<int>? GitCloneDepth { get; set; }

        [Input("gitSubmodulesConfig")]
        public Input<Inputs.ProjectSourceGitSubmodulesConfigGetArgs>? GitSubmodulesConfig { get; set; }

        [Input("insecureSsl")]
        public Input<bool>? InsecureSsl { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("reportBuildStatus")]
        public Input<bool>? ReportBuildStatus { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ProjectSourceGetArgs()
        {
        }
    }
}
