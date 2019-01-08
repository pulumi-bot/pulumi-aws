// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CertificateAuthority extends pulumi.CustomResource {
    /**
     * Get an existing CertificateAuthority resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateAuthorityState, opts?: pulumi.CustomResourceOptions): CertificateAuthority {
        return new CertificateAuthority(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly arn: pulumi.Output<string>;
    public /*out*/ readonly certificate: pulumi.Output<string>;
    public readonly certificateAuthorityConfiguration: pulumi.Output<{ keyAlgorithm: string, signingAlgorithm: string, subject: { commonName?: string, country?: string, distinguishedNameQualifier?: string, generationQualifier?: string, givenName?: string, initials?: string, locality?: string, organization?: string, organizationalUnit?: string, pseudonym?: string, state?: string, surname?: string, title?: string } }>;
    public /*out*/ readonly certificateChain: pulumi.Output<string>;
    public /*out*/ readonly certificateSigningRequest: pulumi.Output<string>;
    public readonly enabled: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly notAfter: pulumi.Output<string>;
    public /*out*/ readonly notBefore: pulumi.Output<string>;
    public readonly revocationConfiguration: pulumi.Output<{ crlConfiguration?: { customCname?: string, enabled?: boolean, expirationInDays: number, s3BucketName?: string } } | undefined>;
    public /*out*/ readonly serial: pulumi.Output<string>;
    public /*out*/ readonly status: pulumi.Output<string>;
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly type: pulumi.Output<string | undefined>;

    /**
     * Create a CertificateAuthority resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateAuthorityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateAuthorityArgs | CertificateAuthorityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CertificateAuthorityState = argsOrState as CertificateAuthorityState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["certificateAuthorityConfiguration"] = state ? state.certificateAuthorityConfiguration : undefined;
            inputs["certificateChain"] = state ? state.certificateChain : undefined;
            inputs["certificateSigningRequest"] = state ? state.certificateSigningRequest : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["notAfter"] = state ? state.notAfter : undefined;
            inputs["notBefore"] = state ? state.notBefore : undefined;
            inputs["revocationConfiguration"] = state ? state.revocationConfiguration : undefined;
            inputs["serial"] = state ? state.serial : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as CertificateAuthorityArgs | undefined;
            if (!args || args.certificateAuthorityConfiguration === undefined) {
                throw new Error("Missing required property 'certificateAuthorityConfiguration'");
            }
            inputs["certificateAuthorityConfiguration"] = args ? args.certificateAuthorityConfiguration : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["revocationConfiguration"] = args ? args.revocationConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificate"] = undefined /*out*/;
            inputs["certificateChain"] = undefined /*out*/;
            inputs["certificateSigningRequest"] = undefined /*out*/;
            inputs["notAfter"] = undefined /*out*/;
            inputs["notBefore"] = undefined /*out*/;
            inputs["serial"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        super("aws:acmpca/certificateAuthority:CertificateAuthority", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertificateAuthority resources.
 */
export interface CertificateAuthorityState {
    readonly arn?: pulumi.Input<string>;
    readonly certificate?: pulumi.Input<string>;
    readonly certificateAuthorityConfiguration?: pulumi.Input<{ keyAlgorithm: pulumi.Input<string>, signingAlgorithm: pulumi.Input<string>, subject: pulumi.Input<{ commonName?: pulumi.Input<string>, country?: pulumi.Input<string>, distinguishedNameQualifier?: pulumi.Input<string>, generationQualifier?: pulumi.Input<string>, givenName?: pulumi.Input<string>, initials?: pulumi.Input<string>, locality?: pulumi.Input<string>, organization?: pulumi.Input<string>, organizationalUnit?: pulumi.Input<string>, pseudonym?: pulumi.Input<string>, state?: pulumi.Input<string>, surname?: pulumi.Input<string>, title?: pulumi.Input<string> }> }>;
    readonly certificateChain?: pulumi.Input<string>;
    readonly certificateSigningRequest?: pulumi.Input<string>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly notAfter?: pulumi.Input<string>;
    readonly notBefore?: pulumi.Input<string>;
    readonly revocationConfiguration?: pulumi.Input<{ crlConfiguration?: pulumi.Input<{ customCname?: pulumi.Input<string>, enabled?: pulumi.Input<boolean>, expirationInDays: pulumi.Input<number>, s3BucketName?: pulumi.Input<string> }> }>;
    readonly serial?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CertificateAuthority resource.
 */
export interface CertificateAuthorityArgs {
    readonly certificateAuthorityConfiguration: pulumi.Input<{ keyAlgorithm: pulumi.Input<string>, signingAlgorithm: pulumi.Input<string>, subject: pulumi.Input<{ commonName?: pulumi.Input<string>, country?: pulumi.Input<string>, distinguishedNameQualifier?: pulumi.Input<string>, generationQualifier?: pulumi.Input<string>, givenName?: pulumi.Input<string>, initials?: pulumi.Input<string>, locality?: pulumi.Input<string>, organization?: pulumi.Input<string>, organizationalUnit?: pulumi.Input<string>, pseudonym?: pulumi.Input<string>, state?: pulumi.Input<string>, surname?: pulumi.Input<string>, title?: pulumi.Input<string> }> }>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly revocationConfiguration?: pulumi.Input<{ crlConfiguration?: pulumi.Input<{ customCname?: pulumi.Input<string>, enabled?: pulumi.Input<boolean>, expirationInDays: pulumi.Input<number>, s3BucketName?: pulumi.Input<string> }> }>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly type?: pulumi.Input<string>;
}
