// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticBeanstalk.Inputs
{

    public sealed class ConfigurationTemplateSettingArgs : Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        [Input("resource")]
        public Input<string>? Resource { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ConfigurationTemplateSettingArgs()
        {
        }
    }
}
