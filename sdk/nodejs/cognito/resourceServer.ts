// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cognito Resource Server.
 *
 * ## Example Usage
 * ### Create a basic resource server
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const pool = new aws.cognito.UserPool("pool", {});
 * const resource = new aws.cognito.ResourceServer("resource", {
 *     identifier: "https://example.com",
 *     userPoolId: pool.id,
 * });
 * ```
 * ### Create a resource server with sample-scope
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const pool = new aws.cognito.UserPool("pool", {});
 * const resource = new aws.cognito.ResourceServer("resource", {
 *     identifier: "https://example.com",
 *     scopes: [{
 *         scopeDescription: "a Sample Scope Description",
 *         scopeName: "sample-scope",
 *     }],
 *     userPoolId: pool.id,
 * });
 * ```
 */
export class ResourceServer extends pulumi.CustomResource {
    /**
     * Get an existing ResourceServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceServerState, opts?: pulumi.CustomResourceOptions): ResourceServer {
        return new ResourceServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cognito/resourceServer:ResourceServer';

    /**
     * Returns true if the given object is an instance of ResourceServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceServer.__pulumiType;
    }

    /**
     * An identifier for the resource server.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * A name for the resource server.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of all scopes configured for this resource server in the format identifier/scope_name.
     */
    public /*out*/ readonly scopeIdentifiers!: pulumi.Output<string[]>;
    /**
     * A list of Authorization Scope.
     */
    public readonly scopes!: pulumi.Output<outputs.cognito.ResourceServerScope[] | undefined>;
    public readonly userPoolId!: pulumi.Output<string>;

    /**
     * Create a ResourceServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceServerArgs | ResourceServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResourceServerState | undefined;
            inputs["identifier"] = state ? state.identifier : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["scopeIdentifiers"] = state ? state.scopeIdentifiers : undefined;
            inputs["scopes"] = state ? state.scopes : undefined;
            inputs["userPoolId"] = state ? state.userPoolId : undefined;
        } else {
            const args = argsOrState as ResourceServerArgs | undefined;
            if (!args || args.identifier === undefined) {
                throw new Error("Missing required property 'identifier'");
            }
            if (!args || args.userPoolId === undefined) {
                throw new Error("Missing required property 'userPoolId'");
            }
            inputs["identifier"] = args ? args.identifier : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scopes"] = args ? args.scopes : undefined;
            inputs["userPoolId"] = args ? args.userPoolId : undefined;
            inputs["scopeIdentifiers"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResourceServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceServer resources.
 */
export interface ResourceServerState {
    /**
     * An identifier for the resource server.
     */
    readonly identifier?: pulumi.Input<string>;
    /**
     * A name for the resource server.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of all scopes configured for this resource server in the format identifier/scope_name.
     */
    readonly scopeIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of Authorization Scope.
     */
    readonly scopes?: pulumi.Input<pulumi.Input<inputs.cognito.ResourceServerScope>[]>;
    readonly userPoolId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceServer resource.
 */
export interface ResourceServerArgs {
    /**
     * An identifier for the resource server.
     */
    readonly identifier: pulumi.Input<string>;
    /**
     * A name for the resource server.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of Authorization Scope.
     */
    readonly scopes?: pulumi.Input<pulumi.Input<inputs.cognito.ResourceServerScope>[]>;
    readonly userPoolId: pulumi.Input<string>;
}
