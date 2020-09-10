// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CachesIscsiVolume extends pulumi.CustomResource {
    /**
     * Get an existing CachesIscsiVolume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CachesIscsiVolumeState, opts?: pulumi.CustomResourceOptions): CachesIscsiVolume {
        return new CachesIscsiVolume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume';

    /**
     * Returns true if the given object is an instance of CachesIscsiVolume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CachesIscsiVolume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CachesIscsiVolume.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly chapEnabled!: pulumi.Output<boolean>;
    public readonly gatewayArn!: pulumi.Output<string>;
    public /*out*/ readonly lunNumber!: pulumi.Output<number>;
    public readonly networkInterfaceId!: pulumi.Output<string>;
    public /*out*/ readonly networkInterfacePort!: pulumi.Output<number>;
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    public readonly sourceVolumeArn!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly targetArn!: pulumi.Output<string>;
    public readonly targetName!: pulumi.Output<string>;
    public /*out*/ readonly volumeArn!: pulumi.Output<string>;
    public /*out*/ readonly volumeId!: pulumi.Output<string>;
    public readonly volumeSizeInBytes!: pulumi.Output<number>;

    /**
     * Create a CachesIscsiVolume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CachesIscsiVolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CachesIscsiVolumeArgs | CachesIscsiVolumeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CachesIscsiVolumeState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["chapEnabled"] = state ? state.chapEnabled : undefined;
            inputs["gatewayArn"] = state ? state.gatewayArn : undefined;
            inputs["lunNumber"] = state ? state.lunNumber : undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            inputs["networkInterfacePort"] = state ? state.networkInterfacePort : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
            inputs["sourceVolumeArn"] = state ? state.sourceVolumeArn : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetArn"] = state ? state.targetArn : undefined;
            inputs["targetName"] = state ? state.targetName : undefined;
            inputs["volumeArn"] = state ? state.volumeArn : undefined;
            inputs["volumeId"] = state ? state.volumeId : undefined;
            inputs["volumeSizeInBytes"] = state ? state.volumeSizeInBytes : undefined;
        } else {
            const args = argsOrState as CachesIscsiVolumeArgs | undefined;
            if (!args || args.gatewayArn === undefined) {
                throw new Error("Missing required property 'gatewayArn'");
            }
            if (!args || args.networkInterfaceId === undefined) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            if (!args || args.targetName === undefined) {
                throw new Error("Missing required property 'targetName'");
            }
            if (!args || args.volumeSizeInBytes === undefined) {
                throw new Error("Missing required property 'volumeSizeInBytes'");
            }
            inputs["gatewayArn"] = args ? args.gatewayArn : undefined;
            inputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            inputs["snapshotId"] = args ? args.snapshotId : undefined;
            inputs["sourceVolumeArn"] = args ? args.sourceVolumeArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetName"] = args ? args.targetName : undefined;
            inputs["volumeSizeInBytes"] = args ? args.volumeSizeInBytes : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["chapEnabled"] = undefined /*out*/;
            inputs["lunNumber"] = undefined /*out*/;
            inputs["networkInterfacePort"] = undefined /*out*/;
            inputs["targetArn"] = undefined /*out*/;
            inputs["volumeArn"] = undefined /*out*/;
            inputs["volumeId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CachesIscsiVolume.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CachesIscsiVolume resources.
 */
export interface CachesIscsiVolumeState {
    readonly arn?: pulumi.Input<string>;
    readonly chapEnabled?: pulumi.Input<boolean>;
    readonly gatewayArn?: pulumi.Input<string>;
    readonly lunNumber?: pulumi.Input<number>;
    readonly networkInterfaceId?: pulumi.Input<string>;
    readonly networkInterfacePort?: pulumi.Input<number>;
    readonly snapshotId?: pulumi.Input<string>;
    readonly sourceVolumeArn?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly targetArn?: pulumi.Input<string>;
    readonly targetName?: pulumi.Input<string>;
    readonly volumeArn?: pulumi.Input<string>;
    readonly volumeId?: pulumi.Input<string>;
    readonly volumeSizeInBytes?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CachesIscsiVolume resource.
 */
export interface CachesIscsiVolumeArgs {
    readonly gatewayArn: pulumi.Input<string>;
    readonly networkInterfaceId: pulumi.Input<string>;
    readonly snapshotId?: pulumi.Input<string>;
    readonly sourceVolumeArn?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly targetName: pulumi.Input<string>;
    readonly volumeSizeInBytes: pulumi.Input<number>;
}
