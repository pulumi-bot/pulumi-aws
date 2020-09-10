// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Organizations.Inputs
{

    public sealed class OrganizationRootArgs : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyTypes")]
        private InputList<Inputs.OrganizationRootPolicyTypeArgs>? _policyTypes;
        public InputList<Inputs.OrganizationRootPolicyTypeArgs> PolicyTypes
        {
            get => _policyTypes ?? (_policyTypes = new InputList<Inputs.OrganizationRootPolicyTypeArgs>());
            set => _policyTypes = value;
        }

        public OrganizationRootArgs()
        {
        }
    }
}
