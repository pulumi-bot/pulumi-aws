// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides details about a specific Amazon API Gateway Version 2 API.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.apigatewayv2.getApi({
 *     apiId: "aabbccddee",
 * }));
 * ```
 */
export function getApi(args: GetApiArgs, opts?: pulumi.InvokeOptions): Promise<GetApiResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:apigatewayv2/getApi:getApi", {
        "apiId": args.apiId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getApi.
 */
export interface GetApiArgs {
    /**
     * The API identifier.
     */
    apiId: string;
    /**
     * A map of resource tags.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getApi.
 */
export interface GetApiResult {
    /**
     * The URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
     */
    readonly apiEndpoint: string;
    readonly apiId: string;
    /**
     * An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
     * Applicable for WebSocket APIs.
     */
    readonly apiKeySelectionExpression: string;
    /**
     * The ARN of the API.
     */
    readonly arn: string;
    /**
     * The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html).
     * Applicable for HTTP APIs.
     */
    readonly corsConfigurations: outputs.apigatewayv2.GetApiCorsConfiguration[];
    /**
     * The description of the API.
     */
    readonly description: string;
    /**
     * Whether clients can invoke the API by using the default `execute-api` endpoint.
     */
    readonly disableExecuteApiEndpoint: boolean;
    /**
     * The ARN prefix to be used in an `aws.lambda.Permission`'s `sourceArn` attribute
     * or in an `aws.iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
     * See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
     */
    readonly executionArn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the API.
     */
    readonly name: string;
    /**
     * The API protocol.
     */
    readonly protocolType: string;
    /**
     * The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
     */
    readonly routeSelectionExpression: string;
    /**
     * A map of resource tags.
     */
    readonly tags: {[key: string]: string};
    /**
     * A version identifier for the API.
     */
    readonly version: string;
}

export function getApiApply(args: GetApiApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApiResult> {
    return pulumi.output(args).apply(a => getApi(a, opts))
}

/**
 * A collection of arguments for invoking getApi.
 */
export interface GetApiApplyArgs {
    /**
     * The API identifier.
     */
    apiId: pulumi.Input<string>;
    /**
     * A map of resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
