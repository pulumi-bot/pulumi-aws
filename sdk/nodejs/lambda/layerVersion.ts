// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Lambda Layer Version resource. Lambda Layers allow you to reuse shared bits of code across multiple lambda functions.
 * 
 * For information about Lambda Layers and how to use them, see [AWS Lambda Layers][1]
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const lambdaLayer = new aws.lambda.LayerVersion("lambdaLayer", {
 *     compatibleRuntimes: ["nodejs8.10"],
 *     code: new pulumi.asset.FileArchive("lambda_layer_payload.zip"),
 *     layerName: "lambdaLayerName",
 * });
 * ```
 * 
 * ## Specifying the Deployment Package
 * 
 * AWS Lambda Layers expect source code to be provided as a deployment package whose structure varies depending on which `compatibleRuntimes` this layer specifies.
 * See [Runtimes][2] for the valid values of `compatibleRuntimes`.
 * 
 * Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or
 * indirectly via Amazon S3 (using the `s3Bucket`, `s3Key` and `s3ObjectVersion` arguments). When providing the deployment
 * package via S3 it may be useful to use the `aws.s3.BucketObject` resource to upload it.
 * 
 * For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading
 * large files efficiently.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lambda_layer_version.html.markdown.
 */
export class LayerVersion extends pulumi.CustomResource {
    /**
     * Get an existing LayerVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LayerVersionState, opts?: pulumi.CustomResourceOptions): LayerVersion {
        return new LayerVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lambda/layerVersion:LayerVersion';

    /**
     * Returns true if the given object is an instance of LayerVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LayerVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LayerVersion.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the Lambda Layer with version.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of [Runtimes][2] this layer is compatible with. Up to 5 runtimes can be specified.
     */
    public readonly compatibleRuntimes!: pulumi.Output<string[] | undefined>;
    /**
     * The date this resource was created.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * Description of what your Lambda Layer does.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
     */
    public readonly code!: pulumi.Output<pulumi.asset.Archive | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Lambda Layer without version.
     */
    public /*out*/ readonly layerArn!: pulumi.Output<string>;
    /**
     * A unique name for your Lambda Layer
     */
    public readonly layerName!: pulumi.Output<string>;
    /**
     * License info for your Lambda Layer. See [License Info][3].
     */
    public readonly licenseInfo!: pulumi.Output<string | undefined>;
    /**
     * The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
     */
    public readonly s3Bucket!: pulumi.Output<string | undefined>;
    /**
     * The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    public readonly s3Key!: pulumi.Output<string | undefined>;
    /**
     * The object version containing the function's deployment package. Conflicts with `filename`.
     */
    public readonly s3ObjectVersion!: pulumi.Output<string | undefined>;
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
     */
    public readonly sourceCodeHash!: pulumi.Output<string>;
    /**
     * The size in bytes of the function .zip file.
     */
    public /*out*/ readonly sourceCodeSize!: pulumi.Output<number>;
    /**
     * This Lamba Layer version.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a LayerVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LayerVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LayerVersionArgs | LayerVersionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LayerVersionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["compatibleRuntimes"] = state ? state.compatibleRuntimes : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["code"] = state ? state.code : undefined;
            inputs["layerArn"] = state ? state.layerArn : undefined;
            inputs["layerName"] = state ? state.layerName : undefined;
            inputs["licenseInfo"] = state ? state.licenseInfo : undefined;
            inputs["s3Bucket"] = state ? state.s3Bucket : undefined;
            inputs["s3Key"] = state ? state.s3Key : undefined;
            inputs["s3ObjectVersion"] = state ? state.s3ObjectVersion : undefined;
            inputs["sourceCodeHash"] = state ? state.sourceCodeHash : undefined;
            inputs["sourceCodeSize"] = state ? state.sourceCodeSize : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as LayerVersionArgs | undefined;
            if (!args || args.layerName === undefined) {
                throw new Error("Missing required property 'layerName'");
            }
            inputs["compatibleRuntimes"] = args ? args.compatibleRuntimes : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["code"] = args ? args.code : undefined;
            inputs["layerName"] = args ? args.layerName : undefined;
            inputs["licenseInfo"] = args ? args.licenseInfo : undefined;
            inputs["s3Bucket"] = args ? args.s3Bucket : undefined;
            inputs["s3Key"] = args ? args.s3Key : undefined;
            inputs["s3ObjectVersion"] = args ? args.s3ObjectVersion : undefined;
            inputs["sourceCodeHash"] = args ? args.sourceCodeHash : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["layerArn"] = undefined /*out*/;
            inputs["sourceCodeSize"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LayerVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LayerVersion resources.
 */
export interface LayerVersionState {
    /**
     * The Amazon Resource Name (ARN) of the Lambda Layer with version.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A list of [Runtimes][2] this layer is compatible with. Up to 5 runtimes can be specified.
     */
    readonly compatibleRuntimes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date this resource was created.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * Description of what your Lambda Layer does.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
     */
    readonly code?: pulumi.Input<pulumi.asset.Archive>;
    /**
     * The Amazon Resource Name (ARN) of the Lambda Layer without version.
     */
    readonly layerArn?: pulumi.Input<string>;
    /**
     * A unique name for your Lambda Layer
     */
    readonly layerName?: pulumi.Input<string>;
    /**
     * License info for your Lambda Layer. See [License Info][3].
     */
    readonly licenseInfo?: pulumi.Input<string>;
    /**
     * The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
     */
    readonly s3Bucket?: pulumi.Input<string>;
    /**
     * The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    readonly s3Key?: pulumi.Input<string>;
    /**
     * The object version containing the function's deployment package. Conflicts with `filename`.
     */
    readonly s3ObjectVersion?: pulumi.Input<string>;
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
     */
    readonly sourceCodeHash?: pulumi.Input<string>;
    /**
     * The size in bytes of the function .zip file.
     */
    readonly sourceCodeSize?: pulumi.Input<number>;
    /**
     * This Lamba Layer version.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LayerVersion resource.
 */
export interface LayerVersionArgs {
    /**
     * A list of [Runtimes][2] this layer is compatible with. Up to 5 runtimes can be specified.
     */
    readonly compatibleRuntimes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of what your Lambda Layer does.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
     */
    readonly code?: pulumi.Input<pulumi.asset.Archive>;
    /**
     * A unique name for your Lambda Layer
     */
    readonly layerName: pulumi.Input<string>;
    /**
     * License info for your Lambda Layer. See [License Info][3].
     */
    readonly licenseInfo?: pulumi.Input<string>;
    /**
     * The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
     */
    readonly s3Bucket?: pulumi.Input<string>;
    /**
     * The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    readonly s3Key?: pulumi.Input<string>;
    /**
     * The object version containing the function's deployment package. Conflicts with `filename`.
     */
    readonly s3ObjectVersion?: pulumi.Input<string>;
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3Key`. The usual way to set this is `${filebase64sha256("file.zip")}` (this provider 0.11.12 or later) or `${base64sha256(file("file.zip"))}` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda layer source archive.
     */
    readonly sourceCodeHash?: pulumi.Input<string>;
}
