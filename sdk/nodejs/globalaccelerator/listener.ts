// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerState, opts?: pulumi.CustomResourceOptions): Listener {
        return new Listener(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:globalaccelerator/listener:Listener';

    /**
     * Returns true if the given object is an instance of Listener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Listener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Listener.__pulumiType;
    }

    public readonly acceleratorArn!: pulumi.Output<string>;
    public readonly clientAffinity!: pulumi.Output<string | undefined>;
    public readonly portRanges!: pulumi.Output<outputs.globalaccelerator.ListenerPortRange[]>;
    public readonly protocol!: pulumi.Output<string>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerArgs | ListenerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ListenerState | undefined;
            inputs["acceleratorArn"] = state ? state.acceleratorArn : undefined;
            inputs["clientAffinity"] = state ? state.clientAffinity : undefined;
            inputs["portRanges"] = state ? state.portRanges : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
        } else {
            const args = argsOrState as ListenerArgs | undefined;
            if (!args || args.acceleratorArn === undefined) {
                throw new Error("Missing required property 'acceleratorArn'");
            }
            if (!args || args.portRanges === undefined) {
                throw new Error("Missing required property 'portRanges'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            inputs["acceleratorArn"] = args ? args.acceleratorArn : undefined;
            inputs["clientAffinity"] = args ? args.clientAffinity : undefined;
            inputs["portRanges"] = args ? args.portRanges : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Listener.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Listener resources.
 */
export interface ListenerState {
    readonly acceleratorArn?: pulumi.Input<string>;
    readonly clientAffinity?: pulumi.Input<string>;
    readonly portRanges?: pulumi.Input<pulumi.Input<inputs.globalaccelerator.ListenerPortRange>[]>;
    readonly protocol?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    readonly acceleratorArn: pulumi.Input<string>;
    readonly clientAffinity?: pulumi.Input<string>;
    readonly portRanges: pulumi.Input<pulumi.Input<inputs.globalaccelerator.ListenerPortRange>[]>;
    readonly protocol: pulumi.Input<string>;
}
