// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The ECR Image data source allows the details of an image with a particular tag or digest to be retrieved.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const serviceImage = pulumi.output(aws.ecr.getImage({
 *     imageTag: "latest",
 *     repositoryName: "my/service",
 * }, { async: true }));
 * ```
 */
export function getImage(args: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ecr/getImage:getImage", {
        "imageDigest": args.imageDigest,
        "imageTag": args.imageTag,
        "registryId": args.registryId,
        "repositoryName": args.repositoryName,
    }, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    /**
     * The sha256 digest of the image manifest. At least one of `imageDigest` or `imageTag` must be specified.
     */
    imageDigest?: string;
    /**
     * The tag associated with this image. At least one of `imageDigest` or `imageTag` must be specified.
     */
    imageTag?: string;
    /**
     * The ID of the Registry where the repository resides.
     */
    registryId?: string;
    /**
     * The name of the ECR Repository.
     */
    repositoryName: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageDigest: string;
    /**
     * The date and time, expressed as a unix timestamp, at which the current image was pushed to the repository.
     */
    readonly imagePushedAt: number;
    /**
     * The size, in bytes, of the image in the repository.
     */
    readonly imageSizeInBytes: number;
    readonly imageTag?: string;
    /**
     * The list of tags associated with this image.
     */
    readonly imageTags: string[];
    readonly registryId: string;
    readonly repositoryName: string;
}

export function getImageApply(args: GetImageApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageResult> {
    return pulumi.output(args).apply(a => getImage(a, opts))
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageApplyArgs {
    /**
     * The sha256 digest of the image manifest. At least one of `imageDigest` or `imageTag` must be specified.
     */
    imageDigest?: pulumi.Input<string>;
    /**
     * The tag associated with this image. At least one of `imageDigest` or `imageTag` must be specified.
     */
    imageTag?: pulumi.Input<string>;
    /**
     * The ID of the Registry where the repository resides.
     */
    registryId?: pulumi.Input<string>;
    /**
     * The name of the ECR Repository.
     */
    repositoryName: pulumi.Input<string>;
}
