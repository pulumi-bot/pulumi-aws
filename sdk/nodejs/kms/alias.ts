// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Alias extends pulumi.CustomResource {
    /**
     * Get an existing Alias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AliasState, opts?: pulumi.CustomResourceOptions): Alias {
        return new Alias(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:kms/alias:Alias';

    /**
     * Returns true if the given object is an instance of Alias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alias.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    public /*out*/ readonly targetKeyArn!: pulumi.Output<string>;
    public readonly targetKeyId!: pulumi.Output<string>;

    /**
     * Create a Alias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AliasArgs | AliasState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AliasState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["targetKeyArn"] = state ? state.targetKeyArn : undefined;
            inputs["targetKeyId"] = state ? state.targetKeyId : undefined;
        } else {
            const args = argsOrState as AliasArgs | undefined;
            if (!args || args.targetKeyId === undefined) {
                throw new Error("Missing required property 'targetKeyId'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["targetKeyId"] = args ? args.targetKeyId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["targetKeyArn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Alias.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alias resources.
 */
export interface AliasState {
    readonly arn?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly targetKeyArn?: pulumi.Input<string>;
    readonly targetKeyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Alias resource.
 */
export interface AliasArgs {
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly targetKeyId: pulumi.Input<string>;
}
