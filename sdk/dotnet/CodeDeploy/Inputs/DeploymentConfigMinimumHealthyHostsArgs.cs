// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeDeploy.Inputs
{

    public sealed class DeploymentConfigMinimumHealthyHostsArgs : Pulumi.ResourceArgs
    {
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public DeploymentConfigMinimumHealthyHostsArgs()
        {
        }
    }
}
