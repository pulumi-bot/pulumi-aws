// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Get information on an Amazon MSK Configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.msk.getConfiguration({
 *     name: "example",
 * }, { async: true }));
 * ```
 */
export function getConfiguration(args: GetConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:msk/getConfiguration:getConfiguration", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getConfiguration.
 */
export interface GetConfigurationArgs {
    /**
     * Name of the configuration.
     */
    name: string;
}

/**
 * A collection of values returned by getConfiguration.
 */
export interface GetConfigurationResult {
    /**
     * Amazon Resource Name (ARN) of the configuration.
     */
    readonly arn: string;
    /**
     * Description of the configuration.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of Apache Kafka versions which can use this configuration.
     */
    readonly kafkaVersions: string[];
    /**
     * Latest revision of the configuration.
     */
    readonly latestRevision: number;
    readonly name: string;
    /**
     * Contents of the server.properties file.
     */
    readonly serverProperties: string;
}
