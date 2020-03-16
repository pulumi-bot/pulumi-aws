// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The ECR Image data source allows the details of an image with a particular tag or digest to be retrieved.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const serviceImage = aws.ecr.getImage({
 *     imageTag: "latest",
 *     repositoryName: "my/service",
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ecr_image.html.markdown.
 */
export function getImage(args: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> & GetImageResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetImageResult> = pulumi.runtime.invoke("aws:ecr/getImage:getImage", {
        "imageDigest": args.imageDigest,
        "imageTag": args.imageTag,
        "registryId": args.registryId,
        "repositoryName": args.repositoryName,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    /**
     * The sha256 digest of the image manifest. At least one of `imageDigest` or `imageTag` must be specified.
     */
    readonly imageDigest?: string;
    /**
     * The tag associated with this image. At least one of `imageDigest` or `imageTag` must be specified.
     */
    readonly imageTag?: string;
    /**
     * The ID of the Registry where the repository resides.
     */
    readonly registryId?: string;
    /**
     * The name of the ECR Repository.
     */
    readonly repositoryName: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
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
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
