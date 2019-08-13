// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

import {RestApi} from "../../../apigateway/restApi";

export interface MethodSettingsSettings {
    /**
     * Specifies whether the cached responses are encrypted.
     */
    cacheDataEncrypted?: pulumi.Input<boolean>;
    /**
     * Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
     */
    cacheTtlInSeconds?: pulumi.Input<number>;
    /**
     * Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached. 
     */
    cachingEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies whether data trace logging is enabled for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
     */
    dataTraceEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs. The available levels are `OFF`, `ERROR`, and `INFO`.
     */
    loggingLevel?: pulumi.Input<string>;
    /**
     * Specifies whether Amazon CloudWatch metrics are enabled for this method.
     */
    metricsEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies whether authorization is required for a cache invalidation request.
     */
    requireAuthorizationForCacheControl?: pulumi.Input<boolean>;
    /**
     * Specifies the throttling burst limit.
     */
    throttlingBurstLimit?: pulumi.Input<number>;
    /**
     * Specifies the throttling rate limit.
     */
    throttlingRateLimit?: pulumi.Input<number>;
    /**
     * Specifies how to handle unauthorized requests for cache invalidation. The available values are `FAIL_WITH_403`, `SUCCEED_WITH_RESPONSE_HEADER`, `SUCCEED_WITHOUT_RESPONSE_HEADER`.
     */
    unauthorizedCacheControlHeaderStrategy?: pulumi.Input<string>;
}
