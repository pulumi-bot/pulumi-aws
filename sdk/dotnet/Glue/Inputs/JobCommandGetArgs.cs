// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Inputs
{

    public sealed class JobCommandGetArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pythonVersion")]
        public Input<string>? PythonVersion { get; set; }

        [Input("scriptLocation", required: true)]
        public Input<string> ScriptLocation { get; set; } = null!;

        public JobCommandGetArgs()
        {
        }
    }
}
