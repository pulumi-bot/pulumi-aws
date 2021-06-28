// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class DistributionTrustedSignerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A flag that specifies whether Origin Shield is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("items")]
        private InputList<Inputs.DistributionTrustedSignerItemArgs>? _items;

        /// <summary>
        /// List of nested attributes for each trusted signer
        /// </summary>
        public InputList<Inputs.DistributionTrustedSignerItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.DistributionTrustedSignerItemArgs>());
            set => _items = value;
        }

        public DistributionTrustedSignerArgs()
        {
        }
    }
}
