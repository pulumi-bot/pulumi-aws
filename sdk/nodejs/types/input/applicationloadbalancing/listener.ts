// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ListenerDefaultAction {
    authenticateCognito?: pulumi.Input<inputApi.applicationloadbalancing.ListenerDefaultActionAuthenticateCognito>;
    authenticateOidc?: pulumi.Input<inputApi.applicationloadbalancing.ListenerDefaultActionAuthenticateOidc>;
    /**
     * Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
     */
    fixedResponse?: pulumi.Input<inputApi.applicationloadbalancing.ListenerDefaultActionFixedResponse>;
    order?: pulumi.Input<number>;
    /**
     * Information for creating a redirect action. Required if `type` is `redirect`.
     */
    redirect?: pulumi.Input<inputApi.applicationloadbalancing.ListenerDefaultActionRedirect>;
    /**
     * The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
     */
    targetGroupArn?: pulumi.Input<string>;
    /**
     * The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
     */
    type: pulumi.Input<string>;
}

export interface ListenerDefaultActionAuthenticateCognito {
    /**
     * The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
     */
    authenticationRequestExtraParams?: pulumi.Input<{[key: string]: any}>;
    /**
     * The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
     */
    onUnauthenticatedRequest?: pulumi.Input<string>;
    /**
     * The set of user claims to be requested from the IdP.
     */
    scope?: pulumi.Input<string>;
    /**
     * The name of the cookie used to maintain session information.
     */
    sessionCookieName?: pulumi.Input<string>;
    /**
     * The maximum duration of the authentication session, in seconds.
     */
    sessionTimeout?: pulumi.Input<number>;
    /**
     * The ARN of the Cognito user pool.
     */
    userPoolArn: pulumi.Input<string>;
    /**
     * The ID of the Cognito user pool client.
     */
    userPoolClientId: pulumi.Input<string>;
    /**
     * The domain prefix or fully-qualified domain name of the Cognito user pool.
     */
    userPoolDomain: pulumi.Input<string>;
}

export interface ListenerDefaultActionAuthenticateOidc {
    /**
     * The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
     */
    authenticationRequestExtraParams?: pulumi.Input<{[key: string]: any}>;
    /**
     * The authorization endpoint of the IdP.
     */
    authorizationEndpoint: pulumi.Input<string>;
    /**
     * The OAuth 2.0 client identifier.
     */
    clientId: pulumi.Input<string>;
    /**
     * The OAuth 2.0 client secret.
     */
    clientSecret: pulumi.Input<string>;
    /**
     * The OIDC issuer identifier of the IdP.
     */
    issuer: pulumi.Input<string>;
    /**
     * The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
     */
    onUnauthenticatedRequest?: pulumi.Input<string>;
    /**
     * The set of user claims to be requested from the IdP.
     */
    scope?: pulumi.Input<string>;
    /**
     * The name of the cookie used to maintain session information.
     */
    sessionCookieName?: pulumi.Input<string>;
    /**
     * The maximum duration of the authentication session, in seconds.
     */
    sessionTimeout?: pulumi.Input<number>;
    /**
     * The token endpoint of the IdP.
     */
    tokenEndpoint: pulumi.Input<string>;
    /**
     * The user info endpoint of the IdP.
     */
    userInfoEndpoint: pulumi.Input<string>;
}

export interface ListenerDefaultActionFixedResponse {
    /**
     * The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
     */
    contentType: pulumi.Input<string>;
    /**
     * The message body.
     */
    messageBody?: pulumi.Input<string>;
    /**
     * The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
     */
    statusCode?: pulumi.Input<string>;
}

export interface ListenerDefaultActionRedirect {
    /**
     * The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
     */
    host?: pulumi.Input<string>;
    /**
     * The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
     */
    path?: pulumi.Input<string>;
    /**
     * The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
     */
    port?: pulumi.Input<string>;
    /**
     * The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
     */
    query?: pulumi.Input<string>;
    /**
     * The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.
     */
    statusCode: pulumi.Input<string>;
}
