// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages an AWS IoT certificate.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 * 
 * const cert = new aws.iot.Certificate("cert", {
 *     active: true,
 *     csr: fs.readFileSync("/my/csr.pem", "utf-8"),
 * });
 * ```
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
     * Boolean flag to indicate if the certificate should be active
     */
    public readonly active!: pulumi.Output<boolean>;
    /**
     * The ARN of the created AWS IoT certificate
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The certificate signing request. Review the
     * [IoT API Reference Guide] (http://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
     * for more information on creating a certificate from a certificate signing request (CSR).
     */
    public readonly csr!: pulumi.Output<string>;

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
            const state = argsOrState as CertificateState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["csr"] = state ? state.csr : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            if (!args || args.active === undefined) {
                throw new Error("Missing required property 'active'");
            }
            if (!args || args.csr === undefined) {
                throw new Error("Missing required property 'csr'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["csr"] = args ? args.csr : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:iot/certificate:Certificate", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * Boolean flag to indicate if the certificate should be active
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * The ARN of the created AWS IoT certificate
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The certificate signing request. Review the
     * [IoT API Reference Guide] (http://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
     * for more information on creating a certificate from a certificate signing request (CSR).
     */
    readonly csr?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Boolean flag to indicate if the certificate should be active
     */
    readonly active: pulumi.Input<boolean>;
    /**
     * The certificate signing request. Review the
     * [IoT API Reference Guide] (http://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
     * for more information on creating a certificate from a certificate signing request (CSR).
     */
    readonly csr: pulumi.Input<string>;
}
