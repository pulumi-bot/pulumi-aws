// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface IdentityPoolCognitoIdentityProvider {
    /**
     * The client ID for the Amazon Cognito Identity User Pool.
     */
    clientId?: string;
    /**
     * The provider name for an Amazon Cognito Identity User Pool.
     */
    providerName?: string;
    /**
     * Whether server-side token validation is enabled for the identity provider’s token or not.
     */
    serverSideTokenCheck?: boolean;
}
