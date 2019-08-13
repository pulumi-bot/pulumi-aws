// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface UsagePlanApiStage {
    /**
     * API Id of the associated API stage in a usage plan.
     */
    apiId: pulumi.Input<string>;
    /**
     * API stage name of the associated API stage in a usage plan.
     */
    stage: pulumi.Input<string>;
}

export interface UsagePlanQuotaSettings {
    /**
     * The maximum number of requests that can be made in a given time period.
     */
    limit: pulumi.Input<number>;
    /**
     * The number of requests subtracted from the given limit in the initial time period.
     */
    offset?: pulumi.Input<number>;
    /**
     * The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
     */
    period: pulumi.Input<string>;
}

export interface UsagePlanThrottleSettings {
    /**
     * The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
     */
    burstLimit?: pulumi.Input<number>;
    /**
     * The API request steady-state rate limit.
     */
    rateLimit?: pulumi.Input<number>;
}
