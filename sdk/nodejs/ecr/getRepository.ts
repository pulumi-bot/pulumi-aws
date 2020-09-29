// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * The ECR Repository data source allows the ARN, Repository URI and Registry ID to be retrieved for an ECR repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const service = pulumi.output(aws.ecr.getRepository({
 *     name: "ecr-repository",
 * }, { async: true }));
 * ```
 */
export function getRepository(args: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ecr/getRepository:getRepository", {
        "name": args.name,
        "registryId": args.registryId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryArgs {
    /**
     * The name of the ECR Repository.
     */
    readonly name: string;
    /**
     * The registry ID where the repository was created.
     */
    readonly registryId?: string;
    /**
     * A map of tags assigned to the resource.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getRepository.
 */
export interface GetRepositoryResult {
    /**
     * Full ARN of the repository.
     */
    readonly arn: string;
    /**
     * Encryption configuration for the repository. See Encryption Configuration below.
     */
    readonly encryptionConfigurations: outputs.ecr.GetRepositoryEncryptionConfiguration[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Configuration block that defines image scanning configuration for the repository. See Image Scanning Configuration below.
     */
    readonly imageScanningConfigurations: outputs.ecr.GetRepositoryImageScanningConfiguration[];
    /**
     * The tag mutability setting for the repository.
     */
    readonly imageTagMutability: string;
    readonly name: string;
    readonly registryId: string;
    /**
     * The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
     */
    readonly repositoryUrl: string;
    /**
     * A map of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}
