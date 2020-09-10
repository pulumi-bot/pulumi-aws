// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class FileSystem extends pulumi.CustomResource {
    /**
     * Get an existing FileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FileSystemState, opts?: pulumi.CustomResourceOptions): FileSystem {
        return new FileSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:efs/fileSystem:FileSystem';

    /**
     * Returns true if the given object is an instance of FileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FileSystem.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly creationToken!: pulumi.Output<string>;
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    public readonly encrypted!: pulumi.Output<boolean>;
    public readonly kmsKeyId!: pulumi.Output<string>;
    public readonly lifecyclePolicy!: pulumi.Output<outputs.efs.FileSystemLifecyclePolicy | undefined>;
    public readonly performanceMode!: pulumi.Output<string>;
    public readonly provisionedThroughputInMibps!: pulumi.Output<number | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly throughputMode!: pulumi.Output<string | undefined>;

    /**
     * Create a FileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileSystemArgs | FileSystemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FileSystemState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["creationToken"] = state ? state.creationToken : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["lifecyclePolicy"] = state ? state.lifecyclePolicy : undefined;
            inputs["performanceMode"] = state ? state.performanceMode : undefined;
            inputs["provisionedThroughputInMibps"] = state ? state.provisionedThroughputInMibps : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["throughputMode"] = state ? state.throughputMode : undefined;
        } else {
            const args = argsOrState as FileSystemArgs | undefined;
            inputs["creationToken"] = args ? args.creationToken : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["lifecyclePolicy"] = args ? args.lifecyclePolicy : undefined;
            inputs["performanceMode"] = args ? args.performanceMode : undefined;
            inputs["provisionedThroughputInMibps"] = args ? args.provisionedThroughputInMibps : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["throughputMode"] = args ? args.throughputMode : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["dnsName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FileSystem.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FileSystem resources.
 */
export interface FileSystemState {
    readonly arn?: pulumi.Input<string>;
    readonly creationToken?: pulumi.Input<string>;
    readonly dnsName?: pulumi.Input<string>;
    readonly encrypted?: pulumi.Input<boolean>;
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly lifecyclePolicy?: pulumi.Input<inputs.efs.FileSystemLifecyclePolicy>;
    readonly performanceMode?: pulumi.Input<string>;
    readonly provisionedThroughputInMibps?: pulumi.Input<number>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly throughputMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FileSystem resource.
 */
export interface FileSystemArgs {
    readonly creationToken?: pulumi.Input<string>;
    readonly encrypted?: pulumi.Input<boolean>;
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly lifecyclePolicy?: pulumi.Input<inputs.efs.FileSystemLifecyclePolicy>;
    readonly performanceMode?: pulumi.Input<string>;
    readonly provisionedThroughputInMibps?: pulumi.Input<number>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly throughputMode?: pulumi.Input<string>;
}
