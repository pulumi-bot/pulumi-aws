// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > **Note:** `aws_alb_listener` is known as `aws_lb_listener`. The functionality is identical.
 * 
 * Provides information about a Load Balancer Listener.
 * 
 * This data source can prove useful when a module accepts an LB Listener as an
 * input variable and needs to know the LB it is attached to, or other
 * information specific to the listener in question.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const listenerArn = config.require("listenerArn");
 * 
 * const selected = pulumi.output(aws.elasticloadbalancingv2.getLoadBalancer({
 *     name: "default-public",
 * }));
 * const listener = pulumi.output(aws.elasticloadbalancingv2.getListener({
 *     arn: listenerArn,
 * }));
 * const selected443 = selected.apply(selected => aws.elasticloadbalancingv2.getListener({
 *     loadBalancerArn: selected.arn,
 *     port: 443,
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lb_listener.html.markdown.
 */
export function getListener(args?: GetListenerArgs, opts?: pulumi.InvokeOptions): Promise<GetListenerResult> & GetListenerResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetListenerResult> = pulumi.runtime.invoke("aws:elasticloadbalancingv2/getListener:getListener", {
        "arn": args.arn,
        "loadBalancerArn": args.loadBalancerArn,
        "port": args.port,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getListener.
 */
export interface GetListenerArgs {
    /**
     * The arn of the listener. Required if `load_balancer_arn` and `port` is not set.
     */
    readonly arn?: string;
    /**
     * The arn of the load balancer. Required if `arn` is not set.
     */
    readonly loadBalancerArn?: string;
    /**
     * The port of the listener. Required if `arn` is not set.
     */
    readonly port?: number;
}

/**
 * A collection of values returned by getListener.
 */
export interface GetListenerResult {
    readonly arn: string;
    readonly certificateArn: string;
    readonly defaultActions: { authenticateCognitos: { authenticationRequestExtraParams: {[key: string]: any}, onUnauthenticatedRequest: string, scope: string, sessionCookieName: string, sessionTimeout: number, userPoolArn: string, userPoolClientId: string, userPoolDomain: string }[], authenticateOidcs: { authenticationRequestExtraParams: {[key: string]: any}, authorizationEndpoint: string, clientId: string, clientSecret: string, issuer: string, onUnauthenticatedRequest: string, scope: string, sessionCookieName: string, sessionTimeout: number, tokenEndpoint: string, userInfoEndpoint: string }[], fixedResponses: { contentType: string, messageBody: string, statusCode: string }[], order: number, redirects: { host: string, path: string, port: string, protocol: string, query: string, statusCode: string }[], targetGroupArn: string, type: string }[];
    readonly loadBalancerArn: string;
    readonly port: number;
    readonly protocol: string;
    readonly sslPolicy: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
