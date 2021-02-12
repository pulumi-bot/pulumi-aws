// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Amazon API Gateway Version 2 deployment.
 * More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
 *
 * > **Note:** Creating a deployment for an API requires at least one `aws.apigatewayv2.Route` resource associated with that API. To avoid race conditions when all resources are being created together, you need to add implicit resource references via the `triggers` argument or explicit resource references using the [resource `dependsOn` meta-argument](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson).
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apigatewayv2.Deployment("example", {
 *     apiId: aws_apigatewayv2_route.example.api_id,
 *     description: "Example deployment",
 * });
 * ```
 *
 * ## Import
 *
 * `aws_apigatewayv2_deployment` can be imported by using the API identifier and deployment identifier, e.g.
 *
 * ```sh
 *  $ pulumi import aws:apigatewayv2/deployment:Deployment example aabbccddee/1122334
 * ```
 *
 *  The `triggers` argument cannot be imported.
 */
export class Deployment extends pulumi.CustomResource {
    /**
     * Get an existing Deployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeploymentState, opts?: pulumi.CustomResourceOptions): Deployment {
        return new Deployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/deployment:Deployment';

    /**
     * Returns true if the given object is an instance of Deployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Deployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Deployment.__pulumiType;
    }

    /**
     * The API identifier.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * Whether the deployment was automatically released.
     */
    public /*out*/ readonly autoDeployed!: pulumi.Output<boolean>;
    /**
     * The description for the deployment resource. Must be less than or equal to 1024 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A map of arbitrary keys and values that, when changed, will trigger a redeployment.
     */
    public readonly triggers!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Deployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentArgs | DeploymentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeploymentState | undefined;
            inputs["apiId"] = state ? state.apiId : undefined;
            inputs["autoDeployed"] = state ? state.autoDeployed : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["triggers"] = state ? state.triggers : undefined;
        } else {
            const args = argsOrState as DeploymentArgs | undefined;
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["triggers"] = args ? args.triggers : undefined;
            inputs["autoDeployed"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Deployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Deployment resources.
 */
export interface DeploymentState {
    /**
     * The API identifier.
     */
    readonly apiId?: pulumi.Input<string>;
    /**
     * Whether the deployment was automatically released.
     */
    readonly autoDeployed?: pulumi.Input<boolean>;
    /**
     * The description for the deployment resource. Must be less than or equal to 1024 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A map of arbitrary keys and values that, when changed, will trigger a redeployment.
     */
    readonly triggers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Deployment resource.
 */
export interface DeploymentArgs {
    /**
     * The API identifier.
     */
    readonly apiId: pulumi.Input<string>;
    /**
     * The description for the deployment resource. Must be less than or equal to 1024 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A map of arbitrary keys and values that, when changed, will trigger a redeployment.
     */
    readonly triggers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
