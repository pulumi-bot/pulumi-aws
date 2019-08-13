// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface IdentityPoolRoleAttachmentRoleMapping {
    /**
     * Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
     */
    ambiguousRoleResolution?: string;
    /**
     * A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
     */
    identityProvider: string;
    /**
     * The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
     */
    mappingRules?: outputApi.cognito.IdentityPoolRoleAttachmentRoleMappingMappingRule[];
    /**
     * The role mapping type.
     */
    type: string;
}

export interface IdentityPoolRoleAttachmentRoleMappingMappingRule {
    /**
     * The claim name that must be present in the token, for example, "isAdmin" or "paid".
     */
    claim: string;
    /**
     * The match condition that specifies how closely the claim value in the IdP token must match Value.
     */
    matchType: string;
    /**
     * The role ARN.
     */
    roleArn: string;
    /**
     * A brief string that the claim must match, for example, "paid" or "yes".
     */
    value: string;
}

export interface IdentityPoolRoleAttachmentRoles {
    authenticated?: string;
    unauthenticated?: string;
}
