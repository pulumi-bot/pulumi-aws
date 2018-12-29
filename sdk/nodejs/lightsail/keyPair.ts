// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class KeyPair extends pulumi.CustomResource {
    /**
     * Get an existing KeyPair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyPairState, opts?: pulumi.CustomResourceOptions): KeyPair {
        return new KeyPair(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly arn: pulumi.Output<string>;
    public /*out*/ readonly encryptedFingerprint: pulumi.Output<string>;
    public /*out*/ readonly encryptedPrivateKey: pulumi.Output<string>;
    public /*out*/ readonly fingerprint: pulumi.Output<string>;
    public readonly name: pulumi.Output<string>;
    public readonly namePrefix: pulumi.Output<string | undefined>;
    public readonly pgpKey: pulumi.Output<string | undefined>;
    public /*out*/ readonly privateKey: pulumi.Output<string>;
    public readonly publicKey: pulumi.Output<string>;

    /**
     * Create a KeyPair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KeyPairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyPairArgs | KeyPairState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: KeyPairState = argsOrState as KeyPairState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["encryptedFingerprint"] = state ? state.encryptedFingerprint : undefined;
            inputs["encryptedPrivateKey"] = state ? state.encryptedPrivateKey : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["pgpKey"] = state ? state.pgpKey : undefined;
            inputs["privateKey"] = state ? state.privateKey : undefined;
            inputs["publicKey"] = state ? state.publicKey : undefined;
        } else {
            const args = argsOrState as KeyPairArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["pgpKey"] = args ? args.pgpKey : undefined;
            inputs["publicKey"] = args ? args.publicKey : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["encryptedFingerprint"] = undefined /*out*/;
            inputs["encryptedPrivateKey"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["privateKey"] = undefined /*out*/;
        }
        super("aws:lightsail/keyPair:KeyPair", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyPair resources.
 */
export interface KeyPairState {
    readonly arn?: pulumi.Input<string>;
    readonly encryptedFingerprint?: pulumi.Input<string>;
    readonly encryptedPrivateKey?: pulumi.Input<string>;
    readonly fingerprint?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly pgpKey?: pulumi.Input<string>;
    readonly privateKey?: pulumi.Input<string>;
    readonly publicKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyPair resource.
 */
export interface KeyPairArgs {
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly pgpKey?: pulumi.Input<string>;
    readonly publicKey?: pulumi.Input<string>;
}
