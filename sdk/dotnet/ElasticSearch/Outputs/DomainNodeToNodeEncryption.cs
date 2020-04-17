// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticSearch.Outputs
{

    [OutputType]
    public sealed class DomainNodeToNodeEncryption
    {
        /// <summary>
        /// Whether to enable encryption at rest. If the `encrypt_at_rest` block is not provided then this defaults to `false`.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private DomainNodeToNodeEncryption(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
