// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class RouteTable extends pulumi.CustomResource {
    /**
     * Get an existing RouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteTableState, opts?: pulumi.CustomResourceOptions): RouteTable {
        return new RouteTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/routeTable:RouteTable';

    /**
     * Returns true if the given object is an instance of RouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteTable.__pulumiType;
    }

    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    public readonly propagatingVgws!: pulumi.Output<string[]>;
    public readonly routes!: pulumi.Output<outputs.ec2.RouteTableRoute[]>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a RouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteTableArgs | RouteTableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouteTableState | undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["propagatingVgws"] = state ? state.propagatingVgws : undefined;
            inputs["routes"] = state ? state.routes : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as RouteTableArgs | undefined;
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["propagatingVgws"] = args ? args.propagatingVgws : undefined;
            inputs["routes"] = args ? args.routes : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["ownerId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RouteTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteTable resources.
 */
export interface RouteTableState {
    readonly ownerId?: pulumi.Input<string>;
    readonly propagatingVgws?: pulumi.Input<pulumi.Input<string>[]>;
    readonly routes?: pulumi.Input<pulumi.Input<inputs.ec2.RouteTableRoute>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteTable resource.
 */
export interface RouteTableArgs {
    readonly propagatingVgws?: pulumi.Input<pulumi.Input<string>[]>;
    readonly routes?: pulumi.Input<pulumi.Input<inputs.ec2.RouteTableRoute>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly vpcId: pulumi.Input<string>;
}
