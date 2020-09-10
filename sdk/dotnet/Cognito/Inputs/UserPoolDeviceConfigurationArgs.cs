// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Inputs
{

    public sealed class UserPoolDeviceConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("challengeRequiredOnNewDevice")]
        public Input<bool>? ChallengeRequiredOnNewDevice { get; set; }

        [Input("deviceOnlyRememberedOnUserPrompt")]
        public Input<bool>? DeviceOnlyRememberedOnUserPrompt { get; set; }

        public UserPoolDeviceConfigurationArgs()
        {
        }
    }
}
