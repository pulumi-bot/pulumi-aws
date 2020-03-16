// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a SMB Location within AWS DataSync.
 * 
 * > **NOTE:** The DataSync Agents must be available before creating this resource.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.datasync.LocationSmb("example", {
 *     agentArns: [aws_datasync_agent_example.arn],
 *     password: "ANotGreatPassword",
 *     serverHostname: "smb.example.com",
 *     subdirectory: "/exported/path",
 *     user: "Guest",
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/datasync_location_smb.html.markdown.
 */
export class LocationSmb extends pulumi.CustomResource {
    /**
     * Get an existing LocationSmb resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocationSmbState, opts?: pulumi.CustomResourceOptions): LocationSmb {
        return new LocationSmb(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:datasync/locationSmb:LocationSmb';

    /**
     * Returns true if the given object is an instance of LocationSmb.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocationSmb {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocationSmb.__pulumiType;
    }

    /**
     * A list of DataSync Agent ARNs with which this location will be associated.
     */
    public readonly agentArns!: pulumi.Output<string[]>;
    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the Windows domain the SMB server belongs to.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
     */
    public readonly mountOptions!: pulumi.Output<outputs.datasync.LocationSmbMountOptions | undefined>;
    /**
     * The password of the user who can mount the share and has file permissions in the SMB.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
     */
    public readonly serverHostname!: pulumi.Output<string>;
    /**
     * Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
     */
    public readonly subdirectory!: pulumi.Output<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    public /*out*/ readonly uri!: pulumi.Output<string>;
    /**
     * The user who can mount the share and has file and folder permissions in the SMB share.
     */
    public readonly user!: pulumi.Output<string>;

    /**
     * Create a LocationSmb resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocationSmbArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocationSmbArgs | LocationSmbState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LocationSmbState | undefined;
            inputs["agentArns"] = state ? state.agentArns : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["mountOptions"] = state ? state.mountOptions : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["serverHostname"] = state ? state.serverHostname : undefined;
            inputs["subdirectory"] = state ? state.subdirectory : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["uri"] = state ? state.uri : undefined;
            inputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as LocationSmbArgs | undefined;
            if (!args || args.agentArns === undefined) {
                throw new Error("Missing required property 'agentArns'");
            }
            if (!args || args.password === undefined) {
                throw new Error("Missing required property 'password'");
            }
            if (!args || args.serverHostname === undefined) {
                throw new Error("Missing required property 'serverHostname'");
            }
            if (!args || args.subdirectory === undefined) {
                throw new Error("Missing required property 'subdirectory'");
            }
            if (!args || args.user === undefined) {
                throw new Error("Missing required property 'user'");
            }
            inputs["agentArns"] = args ? args.agentArns : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["mountOptions"] = args ? args.mountOptions : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["serverHostname"] = args ? args.serverHostname : undefined;
            inputs["subdirectory"] = args ? args.subdirectory : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["user"] = args ? args.user : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LocationSmb.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocationSmb resources.
 */
export interface LocationSmbState {
    /**
     * A list of DataSync Agent ARNs with which this location will be associated.
     */
    readonly agentArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name of the Windows domain the SMB server belongs to.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
     */
    readonly mountOptions?: pulumi.Input<inputs.datasync.LocationSmbMountOptions>;
    /**
     * The password of the user who can mount the share and has file permissions in the SMB.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
     */
    readonly serverHostname?: pulumi.Input<string>;
    /**
     * Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
     */
    readonly subdirectory?: pulumi.Input<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly uri?: pulumi.Input<string>;
    /**
     * The user who can mount the share and has file and folder permissions in the SMB share.
     */
    readonly user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LocationSmb resource.
 */
export interface LocationSmbArgs {
    /**
     * A list of DataSync Agent ARNs with which this location will be associated.
     */
    readonly agentArns: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Windows domain the SMB server belongs to.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
     */
    readonly mountOptions?: pulumi.Input<inputs.datasync.LocationSmbMountOptions>;
    /**
     * The password of the user who can mount the share and has file permissions in the SMB.
     */
    readonly password: pulumi.Input<string>;
    /**
     * Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
     */
    readonly serverHostname: pulumi.Input<string>;
    /**
     * Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
     */
    readonly subdirectory: pulumi.Input<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The user who can mount the share and has file and folder permissions in the SMB share.
     */
    readonly user: pulumi.Input<string>;
}
