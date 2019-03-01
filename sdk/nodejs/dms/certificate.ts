// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DMS (Data Migration Service) certificate resource. DMS certificates can be created, deleted, and imported.
 * 
 * > **Note:** All arguments including the PEM encoded certificate will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateState, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, <any>state, { ...opts, id: id });
    }

    /**
     * The Amazon Resource Name (ARN) for the certificate.
     */
    public /*out*/ readonly certificateArn: pulumi.Output<string>;
    /**
     * The certificate identifier.
     */
    public readonly certificateId: pulumi.Output<string>;
    /**
     * The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
     */
    public readonly certificatePem: pulumi.Output<string | undefined>;
    /**
     * The contents of the Oracle Wallet certificate for use with SSL. Either `certificate_pem` or `certificate_wallet` must be set.
     */
    public readonly certificateWallet: pulumi.Output<string | undefined>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CertificateState = argsOrState as CertificateState | undefined;
            inputs["certificateArn"] = state ? state.certificateArn : undefined;
            inputs["certificateId"] = state ? state.certificateId : undefined;
            inputs["certificatePem"] = state ? state.certificatePem : undefined;
            inputs["certificateWallet"] = state ? state.certificateWallet : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            if (!args || args.certificateId === undefined) {
                throw new Error("Missing required property 'certificateId'");
            }
            inputs["certificateId"] = args ? args.certificateId : undefined;
            inputs["certificatePem"] = args ? args.certificatePem : undefined;
            inputs["certificateWallet"] = args ? args.certificateWallet : undefined;
            inputs["certificateArn"] = undefined /*out*/;
        }
        super("aws:dms/certificate:Certificate", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * The Amazon Resource Name (ARN) for the certificate.
     */
    readonly certificateArn?: pulumi.Input<string>;
    /**
     * The certificate identifier.
     */
    readonly certificateId?: pulumi.Input<string>;
    /**
     * The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
     */
    readonly certificatePem?: pulumi.Input<string>;
    /**
     * The contents of the Oracle Wallet certificate for use with SSL. Either `certificate_pem` or `certificate_wallet` must be set.
     */
    readonly certificateWallet?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * The certificate identifier.
     */
    readonly certificateId: pulumi.Input<string>;
    /**
     * The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
     */
    readonly certificatePem?: pulumi.Input<string>;
    /**
     * The contents of the Oracle Wallet certificate for use with SSL. Either `certificate_pem` or `certificate_wallet` must be set.
     */
    readonly certificateWallet?: pulumi.Input<string>;
}
