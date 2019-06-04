// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage AWS Secrets Manager secret metadata. To manage a secret value, see the [`aws_secretsmanager_secret_version` resource](https://www.terraform.io/docs/providers/aws/r/secretsmanager_secret_version.html).
 * 
 * ## Example Usage
 * 
 * ### Basic
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.secretsmanager.Secret("example", {});
 * ```
 * 
 * ### Rotation Configuration
 * 
 * To enable automatic secret rotation, the Secrets Manager service requires usage of a Lambda function. The [Rotate Secrets section in the Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html) provides additional information about deploying a prebuilt Lambda functions for supported credential rotation (e.g. RDS) or deploying a custom Lambda function.
 * 
 * > **NOTE:** Configuring rotation causes the secret to rotate once as soon as you store the secret. Before you do this, you must ensure that all of your applications that use the credentials stored in the secret are updated to retrieve the secret from AWS Secrets Manager. The old credentials might no longer be usable after the initial rotation and any applications that you fail to update will break as soon as the old credentials are no longer valid.
 * 
 * > **NOTE:** If you cancel a rotation that is in progress (by removing the `rotation` configuration), it can leave the VersionStage labels in an unexpected state. Depending on what step of the rotation was in progress, you might need to remove the staging label AWSPENDING from the partially created version, specified by the SecretVersionId response value. You should also evaluate the partially rotated new version to see if it should be deleted, which you can do by removing all staging labels from the new version's VersionStage field.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const rotation_example = new aws.secretsmanager.Secret("rotation-example", {
 *     rotationLambdaArn: aws_lambda_function_example.arn,
 *     rotationRules: {
 *         automaticallyAfterDays: 7,
 *     },
 * });
 * ```
 */
export class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState, opts?: pulumi.CustomResourceOptions): Secret {
        return new Secret(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:secretsmanager/secret:Secret';

    /**
     * Returns true if the given object is an instance of Secret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Secret {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === Secret.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the secret.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A description of the secret.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ARN or alias of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
     */
    public readonly policy!: pulumi.Output<string | undefined>;
    /**
     * Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
     */
    public readonly recoveryWindowInDays!: pulumi.Output<number | undefined>;
    /**
     * Specifies whether automatic rotation is enabled for this secret.
     */
    public /*out*/ readonly rotationEnabled!: pulumi.Output<boolean>;
    /**
     * Specifies the ARN of the Lambda function that can rotate the secret.
     */
    public readonly rotationLambdaArn!: pulumi.Output<string | undefined>;
    /**
     * A structure that defines the rotation configuration for this secret. Defined below.
     */
    public readonly rotationRules!: pulumi.Output<{ automaticallyAfterDays: number } | undefined>;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the secret.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Secret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretArgs | SecretState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["recoveryWindowInDays"] = state ? state.recoveryWindowInDays : undefined;
            inputs["rotationEnabled"] = state ? state.rotationEnabled : undefined;
            inputs["rotationLambdaArn"] = state ? state.rotationLambdaArn : undefined;
            inputs["rotationRules"] = state ? state.rotationRules : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SecretArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["recoveryWindowInDays"] = args ? args.recoveryWindowInDays : undefined;
            inputs["rotationLambdaArn"] = args ? args.rotationLambdaArn : undefined;
            inputs["rotationRules"] = args ? args.rotationRules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["rotationEnabled"] = undefined /*out*/;
        }
        super(Secret.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Secret resources.
 */
export interface SecretState {
    /**
     * Amazon Resource Name (ARN) of the secret.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A description of the secret.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the ARN or alias of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
     */
    readonly recoveryWindowInDays?: pulumi.Input<number>;
    /**
     * Specifies whether automatic rotation is enabled for this secret.
     */
    readonly rotationEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the ARN of the Lambda function that can rotate the secret.
     */
    readonly rotationLambdaArn?: pulumi.Input<string>;
    /**
     * A structure that defines the rotation configuration for this secret. Defined below.
     */
    readonly rotationRules?: pulumi.Input<{ automaticallyAfterDays: pulumi.Input<number> }>;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the secret.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Secret resource.
 */
export interface SecretArgs {
    /**
     * A description of the secret.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the ARN or alias of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
     */
    readonly recoveryWindowInDays?: pulumi.Input<number>;
    /**
     * Specifies the ARN of the Lambda function that can rotate the secret.
     */
    readonly rotationLambdaArn?: pulumi.Input<string>;
    /**
     * A structure that defines the rotation configuration for this secret. Defined below.
     */
    readonly rotationRules?: pulumi.Input<{ automaticallyAfterDays: pulumi.Input<number> }>;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the secret.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
