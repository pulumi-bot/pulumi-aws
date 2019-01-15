// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getRepository(args: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    return pulumi.runtime.invoke("aws:codecommit/getRepository:getRepository", {
        "repositoryName": args.repositoryName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryArgs {
    readonly repositoryName: string;
}

/**
 * A collection of values returned by getRepository.
 */
export interface GetRepositoryResult {
    readonly arn: string;
    readonly cloneUrlHttp: string;
    readonly cloneUrlSsh: string;
    readonly repositoryId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
