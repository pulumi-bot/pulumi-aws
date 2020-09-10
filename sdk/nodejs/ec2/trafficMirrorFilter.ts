// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class TrafficMirrorFilter extends pulumi.CustomResource {
    /**
     * Get an existing TrafficMirrorFilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrafficMirrorFilterState, opts?: pulumi.CustomResourceOptions): TrafficMirrorFilter {
        return new TrafficMirrorFilter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/trafficMirrorFilter:TrafficMirrorFilter';

    /**
     * Returns true if the given object is an instance of TrafficMirrorFilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrafficMirrorFilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrafficMirrorFilter.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly networkServices!: pulumi.Output<string[] | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a TrafficMirrorFilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TrafficMirrorFilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrafficMirrorFilterArgs | TrafficMirrorFilterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TrafficMirrorFilterState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["networkServices"] = state ? state.networkServices : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as TrafficMirrorFilterArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["networkServices"] = args ? args.networkServices : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TrafficMirrorFilter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TrafficMirrorFilter resources.
 */
export interface TrafficMirrorFilterState {
    readonly description?: pulumi.Input<string>;
    readonly networkServices?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a TrafficMirrorFilter resource.
 */
export interface TrafficMirrorFilterArgs {
    readonly description?: pulumi.Input<string>;
    readonly networkServices?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
