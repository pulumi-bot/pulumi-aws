// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class DomainIdentity extends pulumi.CustomResource {
    /**
     * Get an existing DomainIdentity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainIdentityState, opts?: pulumi.CustomResourceOptions): DomainIdentity {
        return new DomainIdentity(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly arn: pulumi.Output<string>;
    public readonly domain: pulumi.Output<string>;
    public /*out*/ readonly verificationToken: pulumi.Output<string>;

    /**
     * Create a DomainIdentity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainIdentityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainIdentityArgs | DomainIdentityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DomainIdentityState = argsOrState as DomainIdentityState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["verificationToken"] = state ? state.verificationToken : undefined;
        } else {
            const args = argsOrState as DomainIdentityArgs | undefined;
            if (!args || args.domain === undefined) {
                throw new Error("Missing required property 'domain'");
            }
            inputs["domain"] = args ? args.domain : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["verificationToken"] = undefined /*out*/;
        }
        super("aws:ses/domainIdentity:DomainIdentity", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainIdentity resources.
 */
export interface DomainIdentityState {
    readonly arn?: pulumi.Input<string>;
    readonly domain?: pulumi.Input<string>;
    readonly verificationToken?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainIdentity resource.
 */
export interface DomainIdentityArgs {
    readonly domain: pulumi.Input<string>;
}
