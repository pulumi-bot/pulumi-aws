// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ApiKey extends pulumi.CustomResource {
    /**
     * Get an existing ApiKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiKeyState, opts?: pulumi.CustomResourceOptions): ApiKey {
        return new ApiKey(name, <any>state, { ...opts, id: id });
    }

    public readonly apiId: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly expires: pulumi.Output<string | undefined>;
    public /*out*/ readonly key: pulumi.Output<string>;

    /**
     * Create a ApiKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiKeyArgs | ApiKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ApiKeyState = argsOrState as ApiKeyState | undefined;
            inputs["apiId"] = state ? state.apiId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["expires"] = state ? state.expires : undefined;
            inputs["key"] = state ? state.key : undefined;
        } else {
            const args = argsOrState as ApiKeyArgs | undefined;
            if (!args || args.apiId === undefined) {
                throw new Error("Missing required property 'apiId'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["expires"] = args ? args.expires : undefined;
            inputs["key"] = undefined /*out*/;
        }
        super("aws:appsync/apiKey:ApiKey", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiKey resources.
 */
export interface ApiKeyState {
    readonly apiId?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly expires?: pulumi.Input<string>;
    readonly key?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiKey resource.
 */
export interface ApiKeyArgs {
    readonly apiId: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly expires?: pulumi.Input<string>;
}
