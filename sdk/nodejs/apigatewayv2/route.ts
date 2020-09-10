// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteState, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/route:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    public readonly apiId!: pulumi.Output<string>;
    public readonly apiKeyRequired!: pulumi.Output<boolean | undefined>;
    public readonly authorizationScopes!: pulumi.Output<string[] | undefined>;
    public readonly authorizationType!: pulumi.Output<string | undefined>;
    public readonly authorizerId!: pulumi.Output<string | undefined>;
    public readonly modelSelectionExpression!: pulumi.Output<string | undefined>;
    public readonly operationName!: pulumi.Output<string | undefined>;
    public readonly requestModels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly routeKey!: pulumi.Output<string>;
    public readonly routeResponseSelectionExpression!: pulumi.Output<string | undefined>;
    public readonly target!: pulumi.Output<string | undefined>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteArgs | RouteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouteState | undefined;
            inputs["apiId"] = state ? state.apiId : undefined;
            inputs["apiKeyRequired"] = state ? state.apiKeyRequired : undefined;
            inputs["authorizationScopes"] = state ? state.authorizationScopes : undefined;
            inputs["authorizationType"] = state ? state.authorizationType : undefined;
            inputs["authorizerId"] = state ? state.authorizerId : undefined;
            inputs["modelSelectionExpression"] = state ? state.modelSelectionExpression : undefined;
            inputs["operationName"] = state ? state.operationName : undefined;
            inputs["requestModels"] = state ? state.requestModels : undefined;
            inputs["routeKey"] = state ? state.routeKey : undefined;
            inputs["routeResponseSelectionExpression"] = state ? state.routeResponseSelectionExpression : undefined;
            inputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as RouteArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            if (!args || args.routeKey === undefined) {
                throw new Error("Missing required property 'routeKey'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["apiKeyRequired"] = args ? args.apiKeyRequired : undefined;
            inputs["authorizationScopes"] = args ? args.authorizationScopes : undefined;
            inputs["authorizationType"] = args ? args.authorizationType : undefined;
            inputs["authorizerId"] = args ? args.authorizerId : undefined;
            inputs["modelSelectionExpression"] = args ? args.modelSelectionExpression : undefined;
            inputs["operationName"] = args ? args.operationName : undefined;
            inputs["requestModels"] = args ? args.requestModels : undefined;
            inputs["routeKey"] = args ? args.routeKey : undefined;
            inputs["routeResponseSelectionExpression"] = args ? args.routeResponseSelectionExpression : undefined;
            inputs["target"] = args ? args.target : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Route.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Route resources.
 */
export interface RouteState {
    readonly apiId?: pulumi.Input<string>;
    readonly apiKeyRequired?: pulumi.Input<boolean>;
    readonly authorizationScopes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly authorizationType?: pulumi.Input<string>;
    readonly authorizerId?: pulumi.Input<string>;
    readonly modelSelectionExpression?: pulumi.Input<string>;
    readonly operationName?: pulumi.Input<string>;
    readonly requestModels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly routeKey?: pulumi.Input<string>;
    readonly routeResponseSelectionExpression?: pulumi.Input<string>;
    readonly target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    readonly apiId: pulumi.Input<string>;
    readonly apiKeyRequired?: pulumi.Input<boolean>;
    readonly authorizationScopes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly authorizationType?: pulumi.Input<string>;
    readonly authorizerId?: pulumi.Input<string>;
    readonly modelSelectionExpression?: pulumi.Input<string>;
    readonly operationName?: pulumi.Input<string>;
    readonly requestModels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly routeKey: pulumi.Input<string>;
    readonly routeResponseSelectionExpression?: pulumi.Input<string>;
    readonly target?: pulumi.Input<string>;
}
