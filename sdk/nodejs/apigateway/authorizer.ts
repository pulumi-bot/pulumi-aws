// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

import {RestApi} from "./restApi";

/**
 * Provides an API Gateway Authorizer.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const demoRestApi = new aws.apigateway.RestApi("demo", {});
 * const invocationRole = new aws.iam.Role("invocationRole", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "apigateway.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 *     path: "/",
 * });
 * const lambda = new aws.iam.Role("lambda", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "lambda.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 * });
 * const authorizer = new aws.lambda.Function("authorizer", {
 *     code: new pulumi.asset.FileArchive("lambda-function.zip"),
 *     handler: "exports.example",
 *     role: lambda.arn,
 * });
 * const demoAuthorizer = new aws.apigateway.Authorizer("demo", {
 *     authorizerCredentials: invocationRole.arn,
 *     authorizerUri: authorizer.invokeArn,
 *     restApi: demoRestApi.id,
 * });
 * const invocationPolicy = new aws.iam.RolePolicy("invocationPolicy", {
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "lambda:InvokeFunction",
 *       "Effect": "Allow",
 *       "Resource": "${authorizer.arn}"
 *     }
 *   ]
 * }
 * `,
 *     role: invocationRole.id,
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_authorizer.html.markdown.
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
    public static readonly __pulumiType = 'aws:apigateway/authorizer:Authorizer';

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
     * The credentials required for the authorizer.
     * To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
     */
    public readonly authorizerCredentials!: pulumi.Output<string | undefined>;
    /**
     * The TTL of cached authorizer results in seconds.
     * Defaults to `300`.
     */
    public readonly authorizerResultTtlInSeconds!: pulumi.Output<number | undefined>;
    /**
     * The authorizer's Uniform Resource Identifier (URI).
     * This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
     * e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
     */
    public readonly authorizerUri!: pulumi.Output<string | undefined>;
    /**
     * The source of the identity in an incoming request.
     * Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
     */
    public readonly identitySource!: pulumi.Output<string | undefined>;
    /**
     * A validation expression for the incoming identity.
     * For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
     * against this expression, and will proceed if the token matches. If the token doesn't match,
     * the client receives a 401 Unauthorized response.
     */
    public readonly identityValidationExpression!: pulumi.Output<string | undefined>;
    /**
     * The name of the authorizer
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of the Amazon Cognito user pool ARNs.
     * Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
     */
    public readonly providerArns!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the associated REST API
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
     * Defaults to `TOKEN`.
     */
    public readonly type!: pulumi.Output<string | undefined>;

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
            inputs["authorizerCredentials"] = state ? state.authorizerCredentials : undefined;
            inputs["authorizerResultTtlInSeconds"] = state ? state.authorizerResultTtlInSeconds : undefined;
            inputs["authorizerUri"] = state ? state.authorizerUri : undefined;
            inputs["identitySource"] = state ? state.identitySource : undefined;
            inputs["identityValidationExpression"] = state ? state.identityValidationExpression : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["providerArns"] = state ? state.providerArns : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AuthorizerArgs | undefined;
            if (!args || args.restApi === undefined) {
                throw new Error("Missing required property 'restApi'");
            }
            inputs["authorizerCredentials"] = args ? args.authorizerCredentials : undefined;
            inputs["authorizerResultTtlInSeconds"] = args ? args.authorizerResultTtlInSeconds : undefined;
            inputs["authorizerUri"] = args ? args.authorizerUri : undefined;
            inputs["identitySource"] = args ? args.identitySource : undefined;
            inputs["identityValidationExpression"] = args ? args.identityValidationExpression : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["providerArns"] = args ? args.providerArns : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["type"] = args ? args.type : undefined;
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
     * The credentials required for the authorizer.
     * To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
     */
    readonly authorizerCredentials?: pulumi.Input<string>;
    /**
     * The TTL of cached authorizer results in seconds.
     * Defaults to `300`.
     */
    readonly authorizerResultTtlInSeconds?: pulumi.Input<number>;
    /**
     * The authorizer's Uniform Resource Identifier (URI).
     * This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
     * e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
     */
    readonly authorizerUri?: pulumi.Input<string>;
    /**
     * The source of the identity in an incoming request.
     * Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
     */
    readonly identitySource?: pulumi.Input<string>;
    /**
     * A validation expression for the incoming identity.
     * For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
     * against this expression, and will proceed if the token matches. If the token doesn't match,
     * the client receives a 401 Unauthorized response.
     */
    readonly identityValidationExpression?: pulumi.Input<string>;
    /**
     * The name of the authorizer
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of the Amazon Cognito user pool ARNs.
     * Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
     */
    readonly providerArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi?: pulumi.Input<string | RestApi>;
    /**
     * The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
     * Defaults to `TOKEN`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Authorizer resource.
 */
export interface AuthorizerArgs {
    /**
     * The credentials required for the authorizer.
     * To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
     */
    readonly authorizerCredentials?: pulumi.Input<string>;
    /**
     * The TTL of cached authorizer results in seconds.
     * Defaults to `300`.
     */
    readonly authorizerResultTtlInSeconds?: pulumi.Input<number>;
    /**
     * The authorizer's Uniform Resource Identifier (URI).
     * This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
     * e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
     */
    readonly authorizerUri?: pulumi.Input<string>;
    /**
     * The source of the identity in an incoming request.
     * Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
     */
    readonly identitySource?: pulumi.Input<string>;
    /**
     * A validation expression for the incoming identity.
     * For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
     * against this expression, and will proceed if the token matches. If the token doesn't match,
     * the client receives a 401 Unauthorized response.
     */
    readonly identityValidationExpression?: pulumi.Input<string>;
    /**
     * The name of the authorizer
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of the Amazon Cognito user pool ARNs.
     * Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
     */
    readonly providerArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi: pulumi.Input<string | RestApi>;
    /**
     * The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
     * Defaults to `TOKEN`.
     */
    readonly type?: pulumi.Input<string>;
}
