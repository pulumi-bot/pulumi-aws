// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glacier.Inputs
{

    public sealed class VaultNotificationGetArgs : Pulumi.ResourceArgs
    {
        [Input("events", required: true)]
        private InputList<string>? _events;
        public InputList<string> Events
        {
            get => _events ?? (_events = new InputList<string>());
            set => _events = value;
        }

        [Input("snsTopic", required: true)]
        public Input<string> SnsTopic { get; set; } = null!;

        public VaultNotificationGetArgs()
        {
        }
    }
}
