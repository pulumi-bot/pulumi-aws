// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides information about a Lambda Layer Version.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const layerName = config.require("layerName");
 * 
 * const existing = aws.lambda.getLayerVersion({
 *     layerName: layerName,
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lambda_layer_version.html.markdown.
 */
export function getLayerVersion(args: GetLayerVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetLayerVersionResult> & GetLayerVersionResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetLayerVersionResult> = pulumi.runtime.invoke("aws:lambda/getLayerVersion:getLayerVersion", {
        "compatibleRuntime": args.compatibleRuntime,
        "layerName": args.layerName,
        "version": args.version,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getLayerVersion.
 */
export interface GetLayerVersionArgs {
    /**
     * Specific runtime the layer version must support. Conflicts with `version`. If specified, the latest available layer version supporting the provided runtime will be used.
     */
    readonly compatibleRuntime?: string;
    /**
     * Name of the lambda layer.
     */
    readonly layerName: string;
    /**
     * Specific layer version. Conflicts with `compatibleRuntime`. If omitted, the latest available layer version will be used.
     */
    readonly version?: number;
}

/**
 * A collection of values returned by getLayerVersion.
 */
export interface GetLayerVersionResult {
    /**
     * The Amazon Resource Name (ARN) of the Lambda Layer with version.
     */
    readonly arn: string;
    readonly compatibleRuntime?: string;
    /**
     * A list of [Runtimes][1] the specific Lambda Layer version is compatible with.
     */
    readonly compatibleRuntimes: string[];
    /**
     * The date this resource was created.
     */
    readonly createdDate: string;
    /**
     * Description of the specific Lambda Layer version.
     */
    readonly description: string;
    /**
     * The Amazon Resource Name (ARN) of the Lambda Layer without version.
     */
    readonly layerArn: string;
    readonly layerName: string;
    /**
     * License info associated with the specific Lambda Layer version.
     */
    readonly licenseInfo: string;
    /**
     * Base64-encoded representation of raw SHA-256 sum of the zip file.
     */
    readonly sourceCodeHash: string;
    /**
     * The size in bytes of the function .zip file.
     */
    readonly sourceCodeSize: number;
    /**
     * This Lamba Layer version.
     */
    readonly version: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
