// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Generates an IAM policy document in JSON format.
 * 
 * This is a data source which can be used to construct a JSON representation of
 * an IAM policy document, for use with resources which expect policy documents,
 * such as the `aws.iam.Policy` resource.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const examplePolicyDocument = aws.iam.getPolicyDocument({
 *     statements: [
 *         {
 *             actions: [
 *                 "s3:ListAllMyBuckets",
 *                 "s3:GetBucketLocation",
 *             ],
 *             resources: ["arn:aws:s3:::*"],
 *             sid: "1",
 *         },
 *         {
 *             actions: ["s3:ListBucket"],
 *             conditions: [{
 *                 test: "StringLike",
 *                 values: [
 *                     "",
 *                     "home/",
 *                     "home/&{aws:username}/",
 *                 ],
 *                 variable: "s3:prefix",
 *             }],
 *             resources: [`arn:aws:s3:::${var_s3_bucket_name}`],
 *         },
 *         {
 *             actions: ["s3:*"],
 *             resources: [
 *                 `arn:aws:s3:::${var_s3_bucket_name}/home/&{aws:username}`,
 *                 `arn:aws:s3:::${var_s3_bucket_name}/home/&{aws:username}/*`,
 *             ],
 *         },
 *     ],
 * });
 * const examplePolicy = new aws.iam.Policy("example", {
 *     path: "/",
 *     policy: examplePolicyDocument.json,
 * });
 * ```
 * 
 * Using this data source to generate policy documents is *optional*. It is also
 * valid to use literal JSON strings within your configuration, or to use the
 * `file` interpolation function to read a raw JSON policy document from a file.
 * 
 * ## Context Variable Interpolation
 * 
 * The IAM policy document format allows context variables to be interpolated
 * into various strings within a statement. The native IAM policy document format
 * uses `${...}`-style syntax that is in conflict with interpolation
 * syntax, so this data source instead uses `&{...}` syntax for interpolations that
 * should be processed by AWS rather than by this provider.
 * 
 * ## Wildcard Principal
 * 
 * In order to define wildcard principal (a.k.a. anonymous user) use `type = "*"` and
 * `identifiers = ["*"]`. In that case the rendered json will contain `"Principal": "*"`.
 * Note, that even though the [IAM Documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html)
 * states that `"Principal": "*"` and `"Principal": {"AWS": "*"}` are equivalent,
 * those principals have different behavior for IAM Role Trust Policy. Therefore
 * this provider will normalize the principal field only in above-mentioned case and principals
 * like `type = "AWS"` and `identifiers = ["*"]` will be rendered as `"Principal": {"AWS": "*"}`.
 * 
 * ## Example with Multiple Principals
 * 
 * Showing how you can use this as an assume role policy as well as showing how you can specify multiple principal blocks with different types.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const eventStreamBucketRoleAssumeRolePolicy = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [
 *             {
 *                 identifiers: ["firehose.amazonaws.com"],
 *                 type: "Service",
 *             },
 *             {
 *                 identifiers: [varTrustedRoleArn],
 *                 type: "AWS",
 *             },
 *         ],
 *     }],
 * });
 * ```
 * 
 * ## Example with Source and Override
 * 
 * Showing how you can use `sourceJson` and `overrideJson`
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const override = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["s3:*"],
 *         resources: ["*"],
 *         sid: "SidToOverwrite",
 *     }],
 * });
 * const source = aws.iam.getPolicyDocument({
 *     statements: [
 *         {
 *             actions: ["ec2:*"],
 *             resources: ["*"],
 *         },
 *         {
 *             actions: ["s3:*"],
 *             resources: ["*"],
 *             sid: "SidToOverwrite",
 *         },
 *     ],
 * });
 * const overrideJsonExample = aws.iam.getPolicyDocument({
 *     overrideJson: override.json,
 *     statements: [
 *         {
 *             actions: ["ec2:*"],
 *             resources: ["*"],
 *         },
 *         {
 *             actions: ["s3:*"],
 *             resources: [
 *                 "arn:aws:s3:::somebucket",
 *                 "arn:aws:s3:::somebucket/*",
 *             ],
 *             sid: "SidToOverwrite",
 *         },
 *     ],
 * });
 * const sourceJsonExample = aws.iam.getPolicyDocument({
 *     sourceJson: source.json,
 *     statements: [{
 *         actions: ["s3:*"],
 *         resources: [
 *             "arn:aws:s3:::somebucket",
 *             "arn:aws:s3:::somebucket/*",
 *         ],
 *         sid: "SidToOverwrite",
 *     }],
 * });
 * ```
 * 
 * `data.aws_iam_policy_document.source_json_example.json` will evaluate to:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 * 
 * `data.aws_iam_policy_document.override_json_example.json` will evaluate to:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 * 
 * You can also combine `sourceJson` and `overrideJson` in the same document.
 * 
 * ## Example without Statement
 * 
 * Use without a `statement`:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const override = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["s3:GetObject"],
 *         resources: ["*"],
 *         sid: "OverridePlaceholder",
 *     }],
 * });
 * const source = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["ec2:DescribeAccountAttributes"],
 *         resources: ["*"],
 *         sid: "OverridePlaceholder",
 *     }],
 * });
 * const politik = aws.iam.getPolicyDocument({
 *     overrideJson: override.json,
 *     sourceJson: source.json,
 * });
 * ```
 * 
 * `data.aws_iam_policy_document.politik.json` will evaluate to:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_policy_document.html.markdown.
 */
export function getPolicyDocument(args?: GetPolicyDocumentArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyDocumentResult> & GetPolicyDocumentResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetPolicyDocumentResult> = pulumi.runtime.invoke("aws:iam/getPolicyDocument:getPolicyDocument", {
        "overrideJson": args.overrideJson,
        "policyId": args.policyId,
        "sourceJson": args.sourceJson,
        "statements": args.statements,
        "version": args.version,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getPolicyDocument.
 */
export interface GetPolicyDocumentArgs {
    /**
     * An IAM policy document to import and override the
     * current policy document.  Statements with non-blank `sid`s in the override
     * document will overwrite statements with the same `sid` in the current document.
     * Statements without an `sid` cannot be overwritten.
     */
    readonly overrideJson?: string;
    /**
     * An ID for the policy document.
     */
    readonly policyId?: string;
    /**
     * An IAM policy document to import as a base for the
     * current policy document.  Statements with non-blank `sid`s in the current
     * policy document will overwrite statements with the same `sid` in the source
     * json.  Statements without an `sid` cannot be overwritten.
     */
    readonly sourceJson?: string;
    /**
     * A nested configuration block (described below)
     * configuring one *statement* to be included in the policy document.
     */
    readonly statements?: inputApi.iam.GetPolicyDocumentStatement[];
    /**
     * IAM policy document version. Valid values: `2008-10-17`, `2012-10-17`. Defaults to `2012-10-17`. For more information, see the [AWS IAM User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html).
     */
    readonly version?: string;
}

/**
 * A collection of values returned by getPolicyDocument.
 */
export interface GetPolicyDocumentResult {
    /**
     * The above arguments serialized as a standard JSON policy document.
     */
    readonly json: string;
    readonly overrideJson?: string;
    readonly policyId?: string;
    readonly sourceJson?: string;
    readonly statements?: outputApi.iam.GetPolicyDocumentStatement[];
    readonly version?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
