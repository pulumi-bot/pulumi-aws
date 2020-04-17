// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds.Inputs
{

    public sealed class OptionGroupOptionOptionSettingGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the option group. If omitted, this provider will assign a random, unique name. Must be lowercase, to match as it is stored in AWS.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Value of the setting.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public OptionGroupOptionOptionSettingGetArgs()
        {
        }
    }
}
