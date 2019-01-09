// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class SnapshotCopy extends pulumi.CustomResource {
    /**
     * Get an existing SnapshotCopy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotCopyState, opts?: pulumi.CustomResourceOptions): SnapshotCopy {
        return new SnapshotCopy(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly dataEncryptionKeyId: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly encrypted: pulumi.Output<boolean | undefined>;
    public readonly kmsKeyId: pulumi.Output<string | undefined>;
    public /*out*/ readonly ownerAlias: pulumi.Output<string>;
    public /*out*/ readonly ownerId: pulumi.Output<string>;
    public readonly sourceRegion: pulumi.Output<string>;
    public readonly sourceSnapshotId: pulumi.Output<string>;
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    public /*out*/ readonly volumeId: pulumi.Output<string>;
    public /*out*/ readonly volumeSize: pulumi.Output<number>;

    /**
     * Create a SnapshotCopy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotCopyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotCopyArgs | SnapshotCopyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SnapshotCopyState = argsOrState as SnapshotCopyState | undefined;
            inputs["dataEncryptionKeyId"] = state ? state.dataEncryptionKeyId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["ownerAlias"] = state ? state.ownerAlias : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["sourceRegion"] = state ? state.sourceRegion : undefined;
            inputs["sourceSnapshotId"] = state ? state.sourceSnapshotId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["volumeId"] = state ? state.volumeId : undefined;
            inputs["volumeSize"] = state ? state.volumeSize : undefined;
        } else {
            const args = argsOrState as SnapshotCopyArgs | undefined;
            if (!args || args.sourceRegion === undefined) {
                throw new Error("Missing required property 'sourceRegion'");
            }
            if (!args || args.sourceSnapshotId === undefined) {
                throw new Error("Missing required property 'sourceSnapshotId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["sourceRegion"] = args ? args.sourceRegion : undefined;
            inputs["sourceSnapshotId"] = args ? args.sourceSnapshotId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["dataEncryptionKeyId"] = undefined /*out*/;
            inputs["ownerAlias"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["volumeId"] = undefined /*out*/;
            inputs["volumeSize"] = undefined /*out*/;
        }
        super("aws:ebs/snapshotCopy:SnapshotCopy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnapshotCopy resources.
 */
export interface SnapshotCopyState {
    readonly dataEncryptionKeyId?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly encrypted?: pulumi.Input<boolean>;
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly ownerAlias?: pulumi.Input<string>;
    readonly ownerId?: pulumi.Input<string>;
    readonly sourceRegion?: pulumi.Input<string>;
    readonly sourceSnapshotId?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly volumeId?: pulumi.Input<string>;
    readonly volumeSize?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SnapshotCopy resource.
 */
export interface SnapshotCopyArgs {
    readonly description?: pulumi.Input<string>;
    readonly encrypted?: pulumi.Input<boolean>;
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly sourceRegion: pulumi.Input<string>;
    readonly sourceSnapshotId: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
