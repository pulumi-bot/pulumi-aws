// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementGeoMatchStatementArgs : Pulumi.ResourceArgs
    {
        [Input("countryCodes", required: true)]
        private InputList<string>? _countryCodes;
        public InputList<string> CountryCodes
        {
            get => _countryCodes ?? (_countryCodes = new InputList<string>());
            set => _countryCodes = value;
        }

        public WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementGeoMatchStatementArgs()
        {
        }
    }
}
