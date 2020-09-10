// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AccessKey extends pulumi.CustomResource {
    /**
     * Get an existing AccessKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessKeyState, opts?: pulumi.CustomResourceOptions): AccessKey {
        return new AccessKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/accessKey:AccessKey';

    /**
     * Returns true if the given object is an instance of AccessKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessKey.__pulumiType;
    }

    public /*out*/ readonly encryptedSecret!: pulumi.Output<string>;
    public /*out*/ readonly keyFingerprint!: pulumi.Output<string>;
    public readonly pgpKey!: pulumi.Output<string | undefined>;
    public /*out*/ readonly secret!: pulumi.Output<string>;
    public /*out*/ readonly sesSmtpPasswordV4!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly user!: pulumi.Output<string>;

    /**
     * Create a AccessKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessKeyArgs | AccessKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccessKeyState | undefined;
            inputs["encryptedSecret"] = state ? state.encryptedSecret : undefined;
            inputs["keyFingerprint"] = state ? state.keyFingerprint : undefined;
            inputs["pgpKey"] = state ? state.pgpKey : undefined;
            inputs["secret"] = state ? state.secret : undefined;
            inputs["sesSmtpPasswordV4"] = state ? state.sesSmtpPasswordV4 : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as AccessKeyArgs | undefined;
            if (!args || args.user === undefined) {
                throw new Error("Missing required property 'user'");
            }
            inputs["pgpKey"] = args ? args.pgpKey : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["user"] = args ? args.user : undefined;
            inputs["encryptedSecret"] = undefined /*out*/;
            inputs["keyFingerprint"] = undefined /*out*/;
            inputs["secret"] = undefined /*out*/;
            inputs["sesSmtpPasswordV4"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccessKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessKey resources.
 */
export interface AccessKeyState {
    readonly encryptedSecret?: pulumi.Input<string>;
    readonly keyFingerprint?: pulumi.Input<string>;
    readonly pgpKey?: pulumi.Input<string>;
    readonly secret?: pulumi.Input<string>;
    readonly sesSmtpPasswordV4?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessKey resource.
 */
export interface AccessKeyArgs {
    readonly pgpKey?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly user: pulumi.Input<string>;
}
