// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an AppSync GraphQL API.
 * 
 * {{% examples %}}
 * ## Example Usage
 * 
 * {{% example %}}
 * ### API Key Authentication
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.appsync.GraphQLApi("example", {
 *     authenticationType: "API_KEY",
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### AWS Cognito User Pool Authentication
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.appsync.GraphQLApi("example", {
 *     authenticationType: "AMAZON_COGNITO_USER_POOLS",
 *     userPoolConfig: {
 *         awsRegion: aws_region_current.name,
 *         defaultAction: "DENY",
 *         userPoolId: aws_cognito_user_pool_example.id,
 *     },
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### AWS IAM Authentication
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.appsync.GraphQLApi("example", {
 *     authenticationType: "AWS_IAM",
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### With Schema
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.appsync.GraphQLApi("example", {
 *     authenticationType: "AWS_IAM",
 *     schema: `schema {
 * 	query: Query
 * }
 * type Query {
 *   test: Int
 * }
 * `,
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### OpenID Connect Authentication
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.appsync.GraphQLApi("example", {
 *     authenticationType: "OPENID_CONNECT",
 *     openidConnectConfig: {
 *         issuer: "https://example.com",
 *     },
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### With Multiple Authentication Providers
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.appsync.GraphQLApi("example", {
 *     additionalAuthenticationProviders: [{
 *         authenticationType: "AWS_IAM",
 *     }],
 *     authenticationType: "API_KEY",
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### Enabling Logging
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleRole = new aws.iam.Role("example", {
 *     assumeRolePolicy: `{
 *     "Version": "2012-10-17",
 *     "Statement": [
 *         {
 *         "Effect": "Allow",
 *         "Principal": {
 *             "Service": "appsync.amazonaws.com"
 *         },
 *         "Action": "sts:AssumeRole"
 *         }
 *     ]
 * }
 * `,
 * });
 * const exampleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("example", {
 *     policyArn: "arn:aws:iam::aws:policy/service-role/AWSAppSyncPushToCloudWatchLogs",
 *     role: exampleRole.name,
 * });
 * const exampleGraphQLApi = new aws.appsync.GraphQLApi("example", {
 *     logConfig: {
 *         cloudwatchLogsRoleArn: exampleRole.arn,
 *         fieldLogLevel: "ERROR",
 *     },
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_graphql_api.html.markdown.
 */
export class GraphQLApi extends pulumi.CustomResource {
    /**
     * Get an existing GraphQLApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GraphQLApiState, opts?: pulumi.CustomResourceOptions): GraphQLApi {
        return new GraphQLApi(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appsync/graphQLApi:GraphQLApi';

    /**
     * Returns true if the given object is an instance of GraphQLApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GraphQLApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GraphQLApi.__pulumiType;
    }

    /**
     * One or more additional authentication providers for the GraphqlApi. Defined below.
     */
    public readonly additionalAuthenticationProviders!: pulumi.Output<outputs.appsync.GraphQLApiAdditionalAuthenticationProvider[] | undefined>;
    /**
     * The ARN
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
     */
    public readonly authenticationType!: pulumi.Output<string>;
    /**
     * Nested argument containing logging configuration. Defined below.
     */
    public readonly logConfig!: pulumi.Output<outputs.appsync.GraphQLApiLogConfig | undefined>;
    /**
     * A user-supplied name for the GraphqlApi.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Nested argument containing OpenID Connect configuration. Defined below.
     */
    public readonly openidConnectConfig!: pulumi.Output<outputs.appsync.GraphQLApiOpenidConnectConfig | undefined>;
    /**
     * The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
     */
    public readonly schema!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
     */
    public /*out*/ readonly uris!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Amazon Cognito User Pool configuration. Defined below.
     */
    public readonly userPoolConfig!: pulumi.Output<outputs.appsync.GraphQLApiUserPoolConfig | undefined>;

    /**
     * Create a GraphQLApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GraphQLApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GraphQLApiArgs | GraphQLApiState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GraphQLApiState | undefined;
            inputs["additionalAuthenticationProviders"] = state ? state.additionalAuthenticationProviders : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["authenticationType"] = state ? state.authenticationType : undefined;
            inputs["logConfig"] = state ? state.logConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["openidConnectConfig"] = state ? state.openidConnectConfig : undefined;
            inputs["schema"] = state ? state.schema : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["uris"] = state ? state.uris : undefined;
            inputs["userPoolConfig"] = state ? state.userPoolConfig : undefined;
        } else {
            const args = argsOrState as GraphQLApiArgs | undefined;
            if (!args || args.authenticationType === undefined) {
                throw new Error("Missing required property 'authenticationType'");
            }
            inputs["additionalAuthenticationProviders"] = args ? args.additionalAuthenticationProviders : undefined;
            inputs["authenticationType"] = args ? args.authenticationType : undefined;
            inputs["logConfig"] = args ? args.logConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["openidConnectConfig"] = args ? args.openidConnectConfig : undefined;
            inputs["schema"] = args ? args.schema : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userPoolConfig"] = args ? args.userPoolConfig : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["uris"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GraphQLApi.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GraphQLApi resources.
 */
export interface GraphQLApiState {
    /**
     * One or more additional authentication providers for the GraphqlApi. Defined below.
     */
    readonly additionalAuthenticationProviders?: pulumi.Input<pulumi.Input<inputs.appsync.GraphQLApiAdditionalAuthenticationProvider>[]>;
    /**
     * The ARN
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
     */
    readonly authenticationType?: pulumi.Input<string>;
    /**
     * Nested argument containing logging configuration. Defined below.
     */
    readonly logConfig?: pulumi.Input<inputs.appsync.GraphQLApiLogConfig>;
    /**
     * A user-supplied name for the GraphqlApi.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Nested argument containing OpenID Connect configuration. Defined below.
     */
    readonly openidConnectConfig?: pulumi.Input<inputs.appsync.GraphQLApiOpenidConnectConfig>;
    /**
     * The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
     */
    readonly uris?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Amazon Cognito User Pool configuration. Defined below.
     */
    readonly userPoolConfig?: pulumi.Input<inputs.appsync.GraphQLApiUserPoolConfig>;
}

/**
 * The set of arguments for constructing a GraphQLApi resource.
 */
export interface GraphQLApiArgs {
    /**
     * One or more additional authentication providers for the GraphqlApi. Defined below.
     */
    readonly additionalAuthenticationProviders?: pulumi.Input<pulumi.Input<inputs.appsync.GraphQLApiAdditionalAuthenticationProvider>[]>;
    /**
     * The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
     */
    readonly authenticationType: pulumi.Input<string>;
    /**
     * Nested argument containing logging configuration. Defined below.
     */
    readonly logConfig?: pulumi.Input<inputs.appsync.GraphQLApiLogConfig>;
    /**
     * A user-supplied name for the GraphqlApi.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Nested argument containing OpenID Connect configuration. Defined below.
     */
    readonly openidConnectConfig?: pulumi.Input<inputs.appsync.GraphQLApiOpenidConnectConfig>;
    /**
     * The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The Amazon Cognito User Pool configuration. Defined below.
     */
    readonly userPoolConfig?: pulumi.Input<inputs.appsync.GraphQLApiUserPoolConfig>;
}
