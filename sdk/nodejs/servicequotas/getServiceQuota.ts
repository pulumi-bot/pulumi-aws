// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Retrieve information about a Service Quota.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const byQuotaCode = pulumi.output(aws.servicequotas.getServiceQuota({
 *     quotaCode: "L-F678F1CE",
 *     serviceCode: "vpc",
 * }));
 * const byQuotaName = pulumi.output(aws.servicequotas.getServiceQuota({
 *     quotaName: "VPCs per Region",
 *     serviceCode: "vpc",
 * }));
 * ```
 */
export function getServiceQuota(args: GetServiceQuotaArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceQuotaResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:servicequotas/getServiceQuota:getServiceQuota", {
        "quotaCode": args.quotaCode,
        "quotaName": args.quotaName,
        "serviceCode": args.serviceCode,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceQuota.
 */
export interface GetServiceQuotaArgs {
    /**
     * Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
     */
    readonly quotaCode?: string;
    /**
     * Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
     */
    readonly quotaName?: string;
    /**
     * Service code for the quota. Available values can be found with the `aws.servicequotas.getService` data source or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
     */
    readonly serviceCode: string;
}

/**
 * A collection of values returned by getServiceQuota.
 */
export interface GetServiceQuotaResult {
    /**
     * Whether the service quota is adjustable.
     */
    readonly adjustable: boolean;
    /**
     * Amazon Resource Name (ARN) of the service quota.
     */
    readonly arn: string;
    /**
     * Default value of the service quota.
     */
    readonly defaultValue: number;
    /**
     * Whether the service quota is global for the AWS account.
     */
    readonly globalQuota: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly quotaCode: string;
    readonly quotaName: string;
    readonly serviceCode: string;
    /**
     * Name of the service.
     */
    readonly serviceName: string;
    /**
     * Current value of the service quota.
     */
    readonly value: number;
}
