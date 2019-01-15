// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class NatGateway extends pulumi.CustomResource {
    /**
     * Get an existing NatGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NatGatewayState, opts?: pulumi.CustomResourceOptions): NatGateway {
        return new NatGateway(name, <any>state, { ...opts, id: id });
    }

    public readonly allocationId: pulumi.Output<string>;
    public /*out*/ readonly networkInterfaceId: pulumi.Output<string>;
    public /*out*/ readonly privateIp: pulumi.Output<string>;
    public /*out*/ readonly publicIp: pulumi.Output<string>;
    public readonly subnetId: pulumi.Output<string>;
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a NatGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NatGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NatGatewayArgs | NatGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: NatGatewayState = argsOrState as NatGatewayState | undefined;
            inputs["allocationId"] = state ? state.allocationId : undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["publicIp"] = state ? state.publicIp : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as NatGatewayArgs | undefined;
            if (!args || args.allocationId === undefined) {
                throw new Error("Missing required property 'allocationId'");
            }
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["allocationId"] = args ? args.allocationId : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["networkInterfaceId"] = undefined /*out*/;
            inputs["privateIp"] = undefined /*out*/;
            inputs["publicIp"] = undefined /*out*/;
        }
        super("aws:ec2/natGateway:NatGateway", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NatGateway resources.
 */
export interface NatGatewayState {
    readonly allocationId?: pulumi.Input<string>;
    readonly networkInterfaceId?: pulumi.Input<string>;
    readonly privateIp?: pulumi.Input<string>;
    readonly publicIp?: pulumi.Input<string>;
    readonly subnetId?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a NatGateway resource.
 */
export interface NatGatewayArgs {
    readonly allocationId: pulumi.Input<string>;
    readonly subnetId: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
