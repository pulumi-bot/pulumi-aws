// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Load Balancer Listener resource.
 *
 * > **Note:** `aws.alb.Listener` is known as `aws.lb.Listener`. The functionality is identical.
 *
 * ## Example Usage
 * ### Forward Action
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const frontEndLoadBalancer = new aws.lb.LoadBalancer("frontEndLoadBalancer", {});
 * // ...
 * const frontEndTargetGroup = new aws.lb.TargetGroup("frontEndTargetGroup", {});
 * // ...
 * const frontEndListener = new aws.lb.Listener("frontEndListener", {
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: "443",
 *     protocol: "HTTPS",
 *     sslPolicy: "ELBSecurityPolicy-2016-08",
 *     certificateArn: "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4",
 *     defaultActions: [{
 *         type: "forward",
 *         targetGroupArn: frontEndTargetGroup.arn,
 *     }],
 * });
 * ```
 *
 * To a NLB:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const frontEnd = new aws.lb.Listener("frontEnd", {
 *     loadBalancerArn: aws_lb.front_end.arn,
 *     port: "443",
 *     protocol: "TLS",
 *     certificateArn: "arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4",
 *     alpnPolicy: "HTTP2Preferred",
 *     defaultActions: [{
 *         type: "forward",
 *         targetGroupArn: aws_lb_target_group.front_end.arn,
 *     }],
 * });
 * ```
 * ### Redirect Action
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const frontEndLoadBalancer = new aws.lb.LoadBalancer("frontEndLoadBalancer", {});
 * // ...
 * const frontEndListener = new aws.lb.Listener("frontEndListener", {
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: "80",
 *     protocol: "HTTP",
 *     defaultActions: [{
 *         type: "redirect",
 *         redirect: {
 *             port: "443",
 *             protocol: "HTTPS",
 *             statusCode: "HTTP_301",
 *         },
 *     }],
 * });
 * ```
 * ### Fixed-response Action
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const frontEndLoadBalancer = new aws.lb.LoadBalancer("frontEndLoadBalancer", {});
 * // ...
 * const frontEndListener = new aws.lb.Listener("frontEndListener", {
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: "80",
 *     protocol: "HTTP",
 *     defaultActions: [{
 *         type: "fixed-response",
 *         fixedResponse: {
 *             contentType: "text/plain",
 *             messageBody: "Fixed response content",
 *             statusCode: "200",
 *         },
 *     }],
 * });
 * ```
 * ### Authenticate-cognito Action
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const frontEndLoadBalancer = new aws.lb.LoadBalancer("frontEndLoadBalancer", {});
 * // ...
 * const frontEndTargetGroup = new aws.lb.TargetGroup("frontEndTargetGroup", {});
 * // ...
 * const pool = new aws.cognito.UserPool("pool", {});
 * // ...
 * const client = new aws.cognito.UserPoolClient("client", {});
 * // ...
 * const domain = new aws.cognito.UserPoolDomain("domain", {});
 * // ...
 * const frontEndListener = new aws.lb.Listener("frontEndListener", {
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: "80",
 *     protocol: "HTTP",
 *     defaultActions: [
 *         {
 *             type: "authenticate-cognito",
 *             authenticateCognito: {
 *                 userPoolArn: pool.arn,
 *                 userPoolClientId: client.id,
 *                 userPoolDomain: domain.domain,
 *             },
 *         },
 *         {
 *             type: "forward",
 *             targetGroupArn: frontEndTargetGroup.arn,
 *         },
 *     ],
 * });
 * ```
 * ### Authenticate-OIDC Action
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const frontEndLoadBalancer = new aws.lb.LoadBalancer("frontEndLoadBalancer", {});
 * // ...
 * const frontEndTargetGroup = new aws.lb.TargetGroup("frontEndTargetGroup", {});
 * // ...
 * const frontEndListener = new aws.lb.Listener("frontEndListener", {
 *     loadBalancerArn: frontEndLoadBalancer.arn,
 *     port: "80",
 *     protocol: "HTTP",
 *     defaultActions: [
 *         {
 *             type: "authenticate-oidc",
 *             authenticateOidc: {
 *                 authorizationEndpoint: "https://example.com/authorization_endpoint",
 *                 clientId: "client_id",
 *                 clientSecret: "client_secret",
 *                 issuer: "https://example.com",
 *                 tokenEndpoint: "https://example.com/token_endpoint",
 *                 userInfoEndpoint: "https://example.com/user_info_endpoint",
 *             },
 *         },
 *         {
 *             type: "forward",
 *             targetGroupArn: frontEndTargetGroup.arn,
 *         },
 *     ],
 * });
 * ```
 * ### Gateway Load Balancer Listener
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleLoadBalancer = new aws.lb.LoadBalancer("exampleLoadBalancer", {
 *     loadBalancerType: "gateway",
 *     subnetMappings: [{
 *         subnetId: aws_subnet.example.id,
 *     }],
 * });
 * const exampleTargetGroup = new aws.lb.TargetGroup("exampleTargetGroup", {
 *     port: 6081,
 *     protocol: "GENEVE",
 *     vpcId: aws_vpc.example.id,
 *     healthCheck: {
 *         port: 80,
 *         protocol: "HTTP",
 *     },
 * });
 * const exampleListener = new aws.lb.Listener("exampleListener", {
 *     loadBalancerArn: exampleLoadBalancer.id,
 *     defaultActions: [{
 *         targetGroupArn: exampleTargetGroup.id,
 *         type: "forward",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Listeners can be imported using their ARN, e.g.
 *
 * ```sh
 *  $ pulumi import aws:alb/listener:Listener front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:listener/app/front-end-alb/8e4497da625e2d8a/9ab28ade35828f96
 * ```
 */
export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerState, opts?: pulumi.CustomResourceOptions): Listener {
        return new Listener(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:alb/listener:Listener';

    /**
     * Returns true if the given object is an instance of Listener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Listener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Listener.__pulumiType;
    }

    /**
     * Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
     */
    public readonly alpnPolicy!: pulumi.Output<string | undefined>;
    /**
     * ARN of the target group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `aws.lb.ListenerCertificate` resource.
     */
    public readonly certificateArn!: pulumi.Output<string | undefined>;
    /**
     * Configuration block for default actions. Detailed below.
     */
    public readonly defaultActions!: pulumi.Output<outputs.alb.ListenerDefaultAction[]>;
    /**
     * ARN of the load balancer.
     */
    public readonly loadBalancerArn!: pulumi.Output<string>;
    /**
     * Port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * Protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
     */
    public readonly sslPolicy!: pulumi.Output<string>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerArgs | ListenerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ListenerState | undefined;
            inputs["alpnPolicy"] = state ? state.alpnPolicy : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificateArn"] = state ? state.certificateArn : undefined;
            inputs["defaultActions"] = state ? state.defaultActions : undefined;
            inputs["loadBalancerArn"] = state ? state.loadBalancerArn : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["sslPolicy"] = state ? state.sslPolicy : undefined;
        } else {
            const args = argsOrState as ListenerArgs | undefined;
            if ((!args || args.defaultActions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultActions'");
            }
            if ((!args || args.loadBalancerArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerArn'");
            }
            inputs["alpnPolicy"] = args ? args.alpnPolicy : undefined;
            inputs["certificateArn"] = args ? args.certificateArn : undefined;
            inputs["defaultActions"] = args ? args.defaultActions : undefined;
            inputs["loadBalancerArn"] = args ? args.loadBalancerArn : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "aws:applicationloadbalancing/listener:Listener" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Listener.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Listener resources.
 */
export interface ListenerState {
    /**
     * Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
     */
    alpnPolicy?: pulumi.Input<string>;
    /**
     * ARN of the target group.
     */
    arn?: pulumi.Input<string>;
    /**
     * ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `aws.lb.ListenerCertificate` resource.
     */
    certificateArn?: pulumi.Input<string>;
    /**
     * Configuration block for default actions. Detailed below.
     */
    defaultActions?: pulumi.Input<pulumi.Input<inputs.alb.ListenerDefaultAction>[]>;
    /**
     * ARN of the load balancer.
     */
    loadBalancerArn?: pulumi.Input<string>;
    /**
     * Port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
     */
    port?: pulumi.Input<number>;
    /**
     * Protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
     */
    sslPolicy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
     */
    alpnPolicy?: pulumi.Input<string>;
    /**
     * ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `aws.lb.ListenerCertificate` resource.
     */
    certificateArn?: pulumi.Input<string>;
    /**
     * Configuration block for default actions. Detailed below.
     */
    defaultActions: pulumi.Input<pulumi.Input<inputs.alb.ListenerDefaultAction>[]>;
    /**
     * ARN of the load balancer.
     */
    loadBalancerArn: pulumi.Input<string>;
    /**
     * Port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
     */
    port?: pulumi.Input<number>;
    /**
     * Protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
     */
    sslPolicy?: pulumi.Input<string>;
}
