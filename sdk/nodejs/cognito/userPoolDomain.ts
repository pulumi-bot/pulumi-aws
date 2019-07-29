// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cognito User Pool Domain resource.
 * 
 * ## Example Usage
 * 
 * ### Amazon Cognito domain
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.cognito.UserPool("example", {});
 * const main = new aws.cognito.UserPoolDomain("main", {
 *     domain: "example-domain",
 *     userPoolId: example.id,
 * });
 * ```
 * ### Custom Cognito domain
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.cognito.UserPool("example", {});
 * const main = new aws.cognito.UserPoolDomain("main", {
 *     certificateArn: aws_acm_certificate_cert.arn,
 *     domain: "example-domain.example.com",
 *     userPoolId: example.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_user_pool_domain.html.markdown.
 */
export class UserPoolDomain extends pulumi.CustomResource {
    /**
     * Get an existing UserPoolDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPoolDomainState, opts?: pulumi.CustomResourceOptions): UserPoolDomain {
        return new UserPoolDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cognito/userPoolDomain:UserPoolDomain';

    /**
     * Returns true if the given object is an instance of UserPoolDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPoolDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPoolDomain.__pulumiType;
    }

    /**
     * The AWS account ID for the user pool owner.
     */
    public /*out*/ readonly awsAccountId!: pulumi.Output<string>;
    /**
     * The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
     */
    public readonly certificateArn!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the CloudFront distribution.
     */
    public /*out*/ readonly cloudfrontDistributionArn!: pulumi.Output<string>;
    /**
     * The domain string.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The S3 bucket where the static files for this domain are stored.
     */
    public /*out*/ readonly s3Bucket!: pulumi.Output<string>;
    /**
     * The user pool ID.
     */
    public readonly userPoolId!: pulumi.Output<string>;
    /**
     * The app version.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a UserPoolDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPoolDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPoolDomainArgs | UserPoolDomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserPoolDomainState | undefined;
            inputs["awsAccountId"] = state ? state.awsAccountId : undefined;
            inputs["certificateArn"] = state ? state.certificateArn : undefined;
            inputs["cloudfrontDistributionArn"] = state ? state.cloudfrontDistributionArn : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["s3Bucket"] = state ? state.s3Bucket : undefined;
            inputs["userPoolId"] = state ? state.userPoolId : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as UserPoolDomainArgs | undefined;
            if (!args || args.domain === undefined) {
                throw new Error("Missing required property 'domain'");
            }
            if (!args || args.userPoolId === undefined) {
                throw new Error("Missing required property 'userPoolId'");
            }
            inputs["certificateArn"] = args ? args.certificateArn : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["userPoolId"] = args ? args.userPoolId : undefined;
            inputs["awsAccountId"] = undefined /*out*/;
            inputs["cloudfrontDistributionArn"] = undefined /*out*/;
            inputs["s3Bucket"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserPoolDomain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPoolDomain resources.
 */
export interface UserPoolDomainState {
    /**
     * The AWS account ID for the user pool owner.
     */
    readonly awsAccountId?: pulumi.Input<string>;
    /**
     * The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
     */
    readonly certificateArn?: pulumi.Input<string>;
    /**
     * The ARN of the CloudFront distribution.
     */
    readonly cloudfrontDistributionArn?: pulumi.Input<string>;
    /**
     * The domain string.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * The S3 bucket where the static files for this domain are stored.
     */
    readonly s3Bucket?: pulumi.Input<string>;
    /**
     * The user pool ID.
     */
    readonly userPoolId?: pulumi.Input<string>;
    /**
     * The app version.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserPoolDomain resource.
 */
export interface UserPoolDomainArgs {
    /**
     * The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
     */
    readonly certificateArn?: pulumi.Input<string>;
    /**
     * The domain string.
     */
    readonly domain: pulumi.Input<string>;
    /**
     * The user pool ID.
     */
    readonly userPoolId: pulumi.Input<string>;
}
