// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Inputs
{

    public sealed class UserPoolSchemaArgs : Pulumi.ResourceArgs
    {
        [Input("attributeDataType", required: true)]
        public Input<string> AttributeDataType { get; set; } = null!;

        [Input("developerOnlyAttribute")]
        public Input<bool>? DeveloperOnlyAttribute { get; set; }

        [Input("mutable")]
        public Input<bool>? Mutable { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("numberAttributeConstraints")]
        public Input<Inputs.UserPoolSchemaNumberAttributeConstraintsArgs>? NumberAttributeConstraints { get; set; }

        [Input("required")]
        public Input<bool>? Required { get; set; }

        [Input("stringAttributeConstraints")]
        public Input<Inputs.UserPoolSchemaStringAttributeConstraintsArgs>? StringAttributeConstraints { get; set; }

        public UserPoolSchemaArgs()
        {
        }
    }
}
