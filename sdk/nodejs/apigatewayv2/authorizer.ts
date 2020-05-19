// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Amazon API Gateway Version 2 authorizer.
 * More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
 *
 * ## Example Usage
 *
 * ### Basic WebSocket API
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apigatewayv2.Authorizer("example", {
 *     apiId: aws_apigatewayv2_api_example.id,
 *     authorizerType: "REQUEST",
 *     authorizerUri: aws_lambda_function_example.invokeArn,
 *     identitySources: ["route.request.header.Auth"],
 * });
 * ```
 *
 * ### Basic HTTP API
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apigatewayv2.Authorizer("example", {
 *     apiId: aws_apigatewayv2_api_example.id,
 *     authorizerType: "JWT",
 *     identitySources: ["$request.header.Authorization"],
 *     jwtConfiguration: {
 *         audiences: ["example"],
 *         issuer: pulumi.interpolate`https://${aws_cognito_user_pool_example.endpoint}`,
 *     },
 * });
 * ```
 */
export class Authorizer extends pulumi.CustomResource {
    /**
     * Get an existing Authorizer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthorizerState, opts?: pulumi.CustomResourceOptions): Authorizer {
        return new Authorizer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/authorizer:Authorizer';

    /**
     * Returns true if the given object is an instance of Authorizer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Authorizer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Authorizer.__pulumiType;
    }

    /**
     * The API identifier.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * The required credentials as an IAM role for API Gateway to invoke the authorizer.
     * Supported only for `REQUEST` authorizers.
     */
    public readonly authorizerCredentialsArn!: pulumi.Output<string | undefined>;
    /**
     * The authorizer type. Valid values: `JWT`, `REQUEST`.
     * For WebSocket APIs, specify `REQUEST` for a Lambda function using incoming request parameters.
     * For HTTP APIs, specify `JWT` to use JSON Web Tokens.
     */
    public readonly authorizerType!: pulumi.Output<string>;
    /**
     * The authorizer's Uniform Resource Identifier (URI).
     * For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invokeArn` attribute of the [`aws.lambda.Function`](https://www.terraform.io/docs/providers/aws/r/lambda_function.html) resource.
     * Supported only for `REQUEST` authorizers.
     */
    public readonly authorizerUri!: pulumi.Output<string | undefined>;
    /**
     * The identity sources for which authorization is requested.
     * For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
     * For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
     */
    public readonly identitySources!: pulumi.Output<string[]>;
    /**
     * The configuration of a JWT authorizer. Required for the `JWT` authorizer type.
     * Supported only for HTTP APIs.
     */
    public readonly jwtConfiguration!: pulumi.Output<outputs.apigatewayv2.AuthorizerJwtConfiguration | undefined>;
    /**
     * The name of the authorizer.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Authorizer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorizerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorizerArgs | AuthorizerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthorizerState | undefined;
            inputs["apiId"] = state ? state.apiId : undefined;
            inputs["authorizerCredentialsArn"] = state ? state.authorizerCredentialsArn : undefined;
            inputs["authorizerType"] = state ? state.authorizerType : undefined;
            inputs["authorizerUri"] = state ? state.authorizerUri : undefined;
            inputs["identitySources"] = state ? state.identitySources : undefined;
            inputs["jwtConfiguration"] = state ? state.jwtConfiguration : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as AuthorizerArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.authorizerType === undefined) {
                throw new Error("Missing required property 'authorizerType'");
            }
            if (!args || args.identitySources === undefined) {
                throw new Error("Missing required property 'identitySources'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["authorizerCredentialsArn"] = args ? args.authorizerCredentialsArn : undefined;
            inputs["authorizerType"] = args ? args.authorizerType : undefined;
            inputs["authorizerUri"] = args ? args.authorizerUri : undefined;
            inputs["identitySources"] = args ? args.identitySources : undefined;
            inputs["jwtConfiguration"] = args ? args.jwtConfiguration : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Authorizer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Authorizer resources.
 */
export interface AuthorizerState {
    /**
     * The API identifier.
     */
    readonly apiId?: pulumi.Input<string>;
    /**
     * The required credentials as an IAM role for API Gateway to invoke the authorizer.
     * Supported only for `REQUEST` authorizers.
     */
    readonly authorizerCredentialsArn?: pulumi.Input<string>;
    /**
     * The authorizer type. Valid values: `JWT`, `REQUEST`.
     * For WebSocket APIs, specify `REQUEST` for a Lambda function using incoming request parameters.
     * For HTTP APIs, specify `JWT` to use JSON Web Tokens.
     */
    readonly authorizerType?: pulumi.Input<string>;
    /**
     * The authorizer's Uniform Resource Identifier (URI).
     * For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invokeArn` attribute of the [`aws.lambda.Function`](https://www.terraform.io/docs/providers/aws/r/lambda_function.html) resource.
     * Supported only for `REQUEST` authorizers.
     */
    readonly authorizerUri?: pulumi.Input<string>;
    /**
     * The identity sources for which authorization is requested.
     * For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
     * For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
     */
    readonly identitySources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration of a JWT authorizer. Required for the `JWT` authorizer type.
     * Supported only for HTTP APIs.
     */
    readonly jwtConfiguration?: pulumi.Input<inputs.apigatewayv2.AuthorizerJwtConfiguration>;
    /**
     * The name of the authorizer.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Authorizer resource.
 */
export interface AuthorizerArgs {
    /**
     * The API identifier.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * The required credentials as an IAM role for API Gateway to invoke the authorizer.
     * Supported only for `REQUEST` authorizers.
     */
    readonly authorizerCredentialsArn?: pulumi.Input<string>;
    /**
     * The authorizer type. Valid values: `JWT`, `REQUEST`.
     * For WebSocket APIs, specify `REQUEST` for a Lambda function using incoming request parameters.
     * For HTTP APIs, specify `JWT` to use JSON Web Tokens.
     */
    readonly authorizerType: pulumi.Input<string>;
    /**
     * The authorizer's Uniform Resource Identifier (URI).
     * For `REQUEST` authorizers this must be a well-formed Lambda function URI, such as the `invokeArn` attribute of the [`aws.lambda.Function`](https://www.terraform.io/docs/providers/aws/r/lambda_function.html) resource.
     * Supported only for `REQUEST` authorizers.
     */
    readonly authorizerUri?: pulumi.Input<string>;
    /**
     * The identity sources for which authorization is requested.
     * For `REQUEST` authorizers the value is a list of one or more mapping expressions of the specified request parameters.
     * For `JWT` authorizers the single entry specifies where to extract the JSON Web Token (JWT) from inbound requests.
     */
    readonly identitySources: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration of a JWT authorizer. Required for the `JWT` authorizer type.
     * Supported only for HTTP APIs.
     */
    readonly jwtConfiguration?: pulumi.Input<inputs.apigatewayv2.AuthorizerJwtConfiguration>;
    /**
     * The name of the authorizer.
     */
    readonly name?: pulumi.Input<string>;
}
