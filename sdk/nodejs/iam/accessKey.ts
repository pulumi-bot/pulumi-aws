// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an IAM access key. This is a set of credentials that allow API requests to be made as an IAM user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const lbUser = new aws.iam.User("lb", {
 *     path: "/system/",
 * });
 * const lbAccessKey = new aws.iam.AccessKey("lb", {
 *     pgpKey: "keybase:some_person_that_exists",
 *     user: lbUser.name,
 * });
 * const lbRo = new aws.iam.UserPolicy("lb_ro", {
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "ec2:Describe*"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": "*"
 *     }
 *   ]
 * }
 * `,
 *     user: lbUser.name,
 * });
 *
 * export const secret = lbAccessKey.encryptedSecret;
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testUser = new aws.iam.User("testUser", {path: "/test/"});
 * const testAccessKey = new aws.iam.AccessKey("testAccessKey", {user: testUser.name});
 * export const awsIamSmtpPasswordV4 = testAccessKey.sesSmtpPasswordV4;
 * ```
 */
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

    /**
     * The encrypted secret, base64 encoded, if `pgpKey` was specified.
     * > **NOTE:** The encrypted secret may be decrypted using the command line,
     */
    public /*out*/ readonly encryptedSecret!: pulumi.Output<string>;
    /**
     * The fingerprint of the PGP key used to encrypt
     * the secret
     */
    public /*out*/ readonly keyFingerprint!: pulumi.Output<string>;
    /**
     * Either a base-64 encoded PGP public key, or a
     * keybase username in the form `keybase:some_person_that_exists`, for use
     * in the `encryptedSecret` output attribute.
     */
    public readonly pgpKey!: pulumi.Output<string | undefined>;
    /**
     * The secret access key. Note that this will be written
     * to the state file. If you use this, please protect your backend state file
     * judiciously. Alternatively, you may supply a `pgpKey` instead, which will
     * prevent the secret from being stored in plaintext, at the cost of preventing
     * the use of the secret key in automation.
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;
    /**
     * **DEPRECATED** The secret access key converted into an SES SMTP
     * password by applying [AWS's documented conversion
     *
     * @deprecated AWS SigV2 for SES SMTP passwords isy deprecated.
Use 'ses_smtp_password_v4' for region-specific AWS SigV4 signed SES SMTP password instead.
     */
    public /*out*/ readonly sesSmtpPassword!: pulumi.Output<string>;
    /**
     * The secret access key converted into an SES SMTP
     * password by applying [AWS's documented Sigv4 conversion
     * algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert).
     * As SigV4 is region specific, valid Provider regions are `ap-south-1`, `ap-southeast-2`, `eu-central-1`, `eu-west-1`, `us-east-1` and `us-west-2`. See current [AWS SES regions](https://docs.aws.amazon.com/general/latest/gr/rande.html#ses_region)
     */
    public /*out*/ readonly sesSmtpPasswordV4!: pulumi.Output<string>;
    /**
     * The access key status to apply. Defaults to `Active`.
     * Valid values are `Active` and `Inactive`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The IAM user to associate with this access key.
     */
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
            inputs["sesSmtpPassword"] = state ? state.sesSmtpPassword : undefined;
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
            inputs["sesSmtpPassword"] = undefined /*out*/;
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
    /**
     * The encrypted secret, base64 encoded, if `pgpKey` was specified.
     * > **NOTE:** The encrypted secret may be decrypted using the command line,
     */
    readonly encryptedSecret?: pulumi.Input<string>;
    /**
     * The fingerprint of the PGP key used to encrypt
     * the secret
     */
    readonly keyFingerprint?: pulumi.Input<string>;
    /**
     * Either a base-64 encoded PGP public key, or a
     * keybase username in the form `keybase:some_person_that_exists`, for use
     * in the `encryptedSecret` output attribute.
     */
    readonly pgpKey?: pulumi.Input<string>;
    /**
     * The secret access key. Note that this will be written
     * to the state file. If you use this, please protect your backend state file
     * judiciously. Alternatively, you may supply a `pgpKey` instead, which will
     * prevent the secret from being stored in plaintext, at the cost of preventing
     * the use of the secret key in automation.
     */
    readonly secret?: pulumi.Input<string>;
    /**
     * **DEPRECATED** The secret access key converted into an SES SMTP
     * password by applying [AWS's documented conversion
     *
     * @deprecated AWS SigV2 for SES SMTP passwords isy deprecated.
Use 'ses_smtp_password_v4' for region-specific AWS SigV4 signed SES SMTP password instead.
     */
    readonly sesSmtpPassword?: pulumi.Input<string>;
    /**
     * The secret access key converted into an SES SMTP
     * password by applying [AWS's documented Sigv4 conversion
     * algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert).
     * As SigV4 is region specific, valid Provider regions are `ap-south-1`, `ap-southeast-2`, `eu-central-1`, `eu-west-1`, `us-east-1` and `us-west-2`. See current [AWS SES regions](https://docs.aws.amazon.com/general/latest/gr/rande.html#ses_region)
     */
    readonly sesSmtpPasswordV4?: pulumi.Input<string>;
    /**
     * The access key status to apply. Defaults to `Active`.
     * Valid values are `Active` and `Inactive`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The IAM user to associate with this access key.
     */
    readonly user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessKey resource.
 */
export interface AccessKeyArgs {
    /**
     * Either a base-64 encoded PGP public key, or a
     * keybase username in the form `keybase:some_person_that_exists`, for use
     * in the `encryptedSecret` output attribute.
     */
    readonly pgpKey?: pulumi.Input<string>;
    /**
     * The access key status to apply. Defaults to `Active`.
     * Valid values are `Active` and `Inactive`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The IAM user to associate with this access key.
     */
    readonly user: pulumi.Input<string>;
}
