// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface GraphQLApiLogConfig {
    /**
     * Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
     */
    cloudwatchLogsRoleArn: pulumi.Input<string>;
    /**
     * Field logging level. Valid values: `ALL`, `ERROR`, `NONE`.
     */
    fieldLogLevel: pulumi.Input<string>;
}

export interface GraphQLApiOpenidConnectConfig {
    /**
     * Number of milliseconds a token is valid after being authenticated.
     */
    authTtl?: pulumi.Input<number>;
    /**
     * Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
     */
    clientId?: pulumi.Input<string>;
    /**
     * Number of milliseconds a token is valid after being issued to a user.
     */
    iatTtl?: pulumi.Input<number>;
    /**
     * Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
     */
    issuer: pulumi.Input<string>;
}

export interface GraphQLApiUserPoolConfig {
    /**
     * A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
     */
    appIdClientRegex?: pulumi.Input<string>;
    /**
     * The AWS region in which the user pool was created.
     */
    awsRegion?: pulumi.Input<string>;
    /**
     * The action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: `ALLOW` and `DENY`
     */
    defaultAction: pulumi.Input<string>;
    /**
     * The user pool ID.
     */
    userPoolId: pulumi.Input<string>;
}
