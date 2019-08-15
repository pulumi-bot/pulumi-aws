// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cognito User Pool Client resource.
 * 
 * ## Example Usage
 * 
 * ### Create a basic user pool client
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const pool = new aws.cognito.UserPool("pool", {});
 * const client = new aws.cognito.UserPoolClient("client", {
 *     userPoolId: pool.id,
 * });
 * ```
 * 
 * ### Create a user pool client with no SRP authentication
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const pool = new aws.cognito.UserPool("pool", {});
 * const client = new aws.cognito.UserPoolClient("client", {
 *     explicitAuthFlows: ["ADMIN_NO_SRP_AUTH"],
 *     generateSecret: true,
 *     userPoolId: pool.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_user_pool_client.html.markdown.
 */
export class UserPoolClient extends pulumi.CustomResource {
    /**
     * Get an existing UserPoolClient resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPoolClientState, opts?: pulumi.CustomResourceOptions): UserPoolClient {
        return new UserPoolClient(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cognito/userPoolClient:UserPoolClient';

    /**
     * Returns true if the given object is an instance of UserPoolClient.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPoolClient {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPoolClient.__pulumiType;
    }

    /**
     * List of allowed OAuth flows (code, implicit, client_credentials).
     */
    public readonly allowedOauthFlows!: pulumi.Output<string[] | undefined>;
    /**
     * Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
     */
    public readonly allowedOauthFlowsUserPoolClient!: pulumi.Output<boolean | undefined>;
    /**
     * List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
     */
    public readonly allowedOauthScopes!: pulumi.Output<string[] | undefined>;
    /**
     * List of allowed callback URLs for the identity providers.
     */
    public readonly callbackUrls!: pulumi.Output<string[] | undefined>;
    /**
     * The client secret of the user pool client.
     */
    public /*out*/ readonly clientSecret!: pulumi.Output<string>;
    /**
     * The default redirect URI. Must be in the list of callback URLs.
     */
    public readonly defaultRedirectUri!: pulumi.Output<string | undefined>;
    /**
     * List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH).
     */
    public readonly explicitAuthFlows!: pulumi.Output<string[] | undefined>;
    /**
     * Should an application secret be generated.
     */
    public readonly generateSecret!: pulumi.Output<boolean | undefined>;
    /**
     * List of allowed logout URLs for the identity providers.
     */
    public readonly logoutUrls!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the application client.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of user pool attributes the application client can read from.
     */
    public readonly readAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * The time limit in days refresh tokens are valid for.
     */
    public readonly refreshTokenValidity!: pulumi.Output<number | undefined>;
    /**
     * List of provider names for the identity providers that are supported on this client.
     */
    public readonly supportedIdentityProviders!: pulumi.Output<string[] | undefined>;
    /**
     * The user pool the client belongs to.
     */
    public readonly userPoolId!: pulumi.Output<string>;
    /**
     * List of user pool attributes the application client can write to.
     */
    public readonly writeAttributes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a UserPoolClient resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPoolClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPoolClientArgs | UserPoolClientState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserPoolClientState | undefined;
            inputs["allowedOauthFlows"] = state ? state.allowedOauthFlows : undefined;
            inputs["allowedOauthFlowsUserPoolClient"] = state ? state.allowedOauthFlowsUserPoolClient : undefined;
            inputs["allowedOauthScopes"] = state ? state.allowedOauthScopes : undefined;
            inputs["callbackUrls"] = state ? state.callbackUrls : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["defaultRedirectUri"] = state ? state.defaultRedirectUri : undefined;
            inputs["explicitAuthFlows"] = state ? state.explicitAuthFlows : undefined;
            inputs["generateSecret"] = state ? state.generateSecret : undefined;
            inputs["logoutUrls"] = state ? state.logoutUrls : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["readAttributes"] = state ? state.readAttributes : undefined;
            inputs["refreshTokenValidity"] = state ? state.refreshTokenValidity : undefined;
            inputs["supportedIdentityProviders"] = state ? state.supportedIdentityProviders : undefined;
            inputs["userPoolId"] = state ? state.userPoolId : undefined;
            inputs["writeAttributes"] = state ? state.writeAttributes : undefined;
        } else {
            const args = argsOrState as UserPoolClientArgs | undefined;
            if (!args || args.userPoolId === undefined) {
                throw new Error("Missing required property 'userPoolId'");
            }
            inputs["allowedOauthFlows"] = args ? args.allowedOauthFlows : undefined;
            inputs["allowedOauthFlowsUserPoolClient"] = args ? args.allowedOauthFlowsUserPoolClient : undefined;
            inputs["allowedOauthScopes"] = args ? args.allowedOauthScopes : undefined;
            inputs["callbackUrls"] = args ? args.callbackUrls : undefined;
            inputs["defaultRedirectUri"] = args ? args.defaultRedirectUri : undefined;
            inputs["explicitAuthFlows"] = args ? args.explicitAuthFlows : undefined;
            inputs["generateSecret"] = args ? args.generateSecret : undefined;
            inputs["logoutUrls"] = args ? args.logoutUrls : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["readAttributes"] = args ? args.readAttributes : undefined;
            inputs["refreshTokenValidity"] = args ? args.refreshTokenValidity : undefined;
            inputs["supportedIdentityProviders"] = args ? args.supportedIdentityProviders : undefined;
            inputs["userPoolId"] = args ? args.userPoolId : undefined;
            inputs["writeAttributes"] = args ? args.writeAttributes : undefined;
            inputs["clientSecret"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserPoolClient.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPoolClient resources.
 */
export interface UserPoolClientState {
    /**
     * List of allowed OAuth flows (code, implicit, client_credentials).
     */
    readonly allowedOauthFlows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
     */
    readonly allowedOauthFlowsUserPoolClient?: pulumi.Input<boolean>;
    /**
     * List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
     */
    readonly allowedOauthScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of allowed callback URLs for the identity providers.
     */
    readonly callbackUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The client secret of the user pool client.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * The default redirect URI. Must be in the list of callback URLs.
     */
    readonly defaultRedirectUri?: pulumi.Input<string>;
    /**
     * List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH).
     */
    readonly explicitAuthFlows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Should an application secret be generated.
     */
    readonly generateSecret?: pulumi.Input<boolean>;
    /**
     * List of allowed logout URLs for the identity providers.
     */
    readonly logoutUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the application client.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of user pool attributes the application client can read from.
     */
    readonly readAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time limit in days refresh tokens are valid for.
     */
    readonly refreshTokenValidity?: pulumi.Input<number>;
    /**
     * List of provider names for the identity providers that are supported on this client.
     */
    readonly supportedIdentityProviders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user pool the client belongs to.
     */
    readonly userPoolId?: pulumi.Input<string>;
    /**
     * List of user pool attributes the application client can write to.
     */
    readonly writeAttributes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a UserPoolClient resource.
 */
export interface UserPoolClientArgs {
    /**
     * List of allowed OAuth flows (code, implicit, client_credentials).
     */
    readonly allowedOauthFlows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the client is allowed to follow the OAuth protocol when interacting with Cognito user pools.
     */
    readonly allowedOauthFlowsUserPoolClient?: pulumi.Input<boolean>;
    /**
     * List of allowed OAuth scopes (phone, email, openid, profile, and aws.cognito.signin.user.admin).
     */
    readonly allowedOauthScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of allowed callback URLs for the identity providers.
     */
    readonly callbackUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default redirect URI. Must be in the list of callback URLs.
     */
    readonly defaultRedirectUri?: pulumi.Input<string>;
    /**
     * List of authentication flows (ADMIN_NO_SRP_AUTH, CUSTOM_AUTH_FLOW_ONLY, USER_PASSWORD_AUTH).
     */
    readonly explicitAuthFlows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Should an application secret be generated.
     */
    readonly generateSecret?: pulumi.Input<boolean>;
    /**
     * List of allowed logout URLs for the identity providers.
     */
    readonly logoutUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the application client.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of user pool attributes the application client can read from.
     */
    readonly readAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time limit in days refresh tokens are valid for.
     */
    readonly refreshTokenValidity?: pulumi.Input<number>;
    /**
     * List of provider names for the identity providers that are supported on this client.
     */
    readonly supportedIdentityProviders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user pool the client belongs to.
     */
    readonly userPoolId: pulumi.Input<string>;
    /**
     * List of user pool attributes the application client can write to.
     */
    readonly writeAttributes?: pulumi.Input<pulumi.Input<string>[]>;
}
