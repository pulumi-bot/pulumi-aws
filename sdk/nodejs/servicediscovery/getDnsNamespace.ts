// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Retrieves information about a Service Discovery private or public DNS namespace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.servicediscovery.getDnsNamespace({
 *     name: "example.service.local",
 *     type: "DNS_PRIVATE",
 * }));
 * ```
 */
export function getDnsNamespace(args: GetDnsNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetDnsNamespaceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:servicediscovery/getDnsNamespace:getDnsNamespace", {
        "name": args.name,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getDnsNamespace.
 */
export interface GetDnsNamespaceArgs {
    /**
     * The name of the namespace.
     */
    readonly name: string;
    /**
     * The type of the namespace. Allowed values are `DNS_PUBLIC` or `DNS_PRIVATE`.
     */
    readonly type: string;
}

/**
 * A collection of values returned by getDnsNamespace.
 */
export interface GetDnsNamespaceResult {
    /**
     * The Amazon Resource Name (ARN) of the namespace.
     */
    readonly arn: string;
    /**
     * A description of the namespace.
     */
    readonly description: string;
    /**
     * The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
     */
    readonly hostedZone: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly type: string;
}
