// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an Elastic Load Balancer resource, also known as a "Classic
 * Load Balancer" after the release of
 * `Application/Network Load Balancers`.
 *
 * > **NOTE on ELB Instances and ELB Attachments:** This provider currently
 * provides both a standalone ELB Attachment resource
 * (describing an instance attached to an ELB), and an ELB resource with
 * `instances` defined in-line. At this time you cannot use an ELB with in-line
 * instances in conjunction with a ELB Attachment resources. Doing so will cause a
 * conflict and will overwrite attachments.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // Create a new load balancer
 * const bar = new aws.elb.LoadBalancer("bar", {
 *     availabilityZones: [
 *         "us-west-2a",
 *         "us-west-2b",
 *         "us-west-2c",
 *     ],
 *     accessLogs: {
 *         bucket: "foo",
 *         bucketPrefix: "bar",
 *         interval: 60,
 *     },
 *     listeners: [
 *         {
 *             instancePort: 8000,
 *             instanceProtocol: "http",
 *             lbPort: 80,
 *             lbProtocol: "http",
 *         },
 *         {
 *             instancePort: 8000,
 *             instanceProtocol: "http",
 *             lbPort: 443,
 *             lbProtocol: "https",
 *             sslCertificateId: "arn:aws:iam::123456789012:server-certificate/certName",
 *         },
 *     ],
 *     healthCheck: {
 *         healthyThreshold: 2,
 *         unhealthyThreshold: 2,
 *         timeout: 3,
 *         target: "HTTP:8000/",
 *         interval: 30,
 *     },
 *     instances: [aws_instance.foo.id],
 *     crossZoneLoadBalancing: true,
 *     idleTimeout: 400,
 *     connectionDraining: true,
 *     connectionDrainingTimeout: 400,
 *     tags: {
 *         Name: "foobar-elb",
 *     },
 * });
 * ```
 * ## Note on ECDSA Key Algorithm
 *
 * If the ARN of the `sslCertificateId` that is pointed to references a
 * certificate that was signed by an ECDSA key, note that ELB only supports the
 * P256 and P384 curves.  Using a certificate signed by a key using a different
 * curve could produce the error `ERR_SSL_VERSION_OR_CIPHER_MISMATCH` in your
 * browser.
 *
 * ## Import
 *
 * ELBs can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:elb/loadBalancer:LoadBalancer bar elb-production-12345
 * ```
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elb/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * An Access Logs block. Access Logs documented below.
     */
    public readonly accessLogs!: pulumi.Output<outputs.elb.LoadBalancerAccessLogs | undefined>;
    /**
     * The ARN of the ELB
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The AZ's to serve traffic in.
     */
    public readonly availabilityZones!: pulumi.Output<string[]>;
    /**
     * Boolean to enable connection draining. Default: `false`
     */
    public readonly connectionDraining!: pulumi.Output<boolean | undefined>;
    /**
     * The time in seconds to allow for connections to drain. Default: `300`
     */
    public readonly connectionDrainingTimeout!: pulumi.Output<number | undefined>;
    /**
     * Enable cross-zone load balancing. Default: `true`
     */
    public readonly crossZoneLoadBalancing!: pulumi.Output<boolean | undefined>;
    /**
     * The DNS name of the ELB
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * A healthCheck block. Health Check documented below.
     */
    public readonly healthCheck!: pulumi.Output<outputs.elb.LoadBalancerHealthCheck>;
    /**
     * The time in seconds that the connection is allowed to be idle. Default: `60`
     */
    public readonly idleTimeout!: pulumi.Output<number | undefined>;
    /**
     * A list of instance ids to place in the ELB pool.
     */
    public readonly instances!: pulumi.Output<string[]>;
    /**
     * If true, ELB will be an internal ELB.
     */
    public readonly internal!: pulumi.Output<boolean>;
    /**
     * A list of listener blocks. Listeners documented below.
     */
    public readonly listeners!: pulumi.Output<outputs.elb.LoadBalancerListener[]>;
    /**
     * The name of the ELB. By default generated by this provider.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * A list of security group IDs to assign to the ELB.
     * Only valid if creating an ELB within a VPC
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * The name of the security group that you can use as
     * part of your inbound rules for your load balancer's back-end application
     * instances. Use this for Classic or Default VPC only.
     */
    public readonly sourceSecurityGroup!: pulumi.Output<string>;
    /**
     * The ID of the security group that you can use as
     * part of your inbound rules for your load balancer's back-end application
     * instances. Only available on ELBs launched in a VPC.
     */
    public /*out*/ readonly sourceSecurityGroupId!: pulumi.Output<string>;
    /**
     * A list of subnet IDs to attach to the ELB.
     */
    public readonly subnets!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The canonical hosted zone ID of the ELB (to be used in a Route 53 Alias record)
     */
    public /*out*/ readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            inputs["accessLogs"] = state ? state.accessLogs : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["availabilityZones"] = state ? state.availabilityZones : undefined;
            inputs["connectionDraining"] = state ? state.connectionDraining : undefined;
            inputs["connectionDrainingTimeout"] = state ? state.connectionDrainingTimeout : undefined;
            inputs["crossZoneLoadBalancing"] = state ? state.crossZoneLoadBalancing : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["healthCheck"] = state ? state.healthCheck : undefined;
            inputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            inputs["instances"] = state ? state.instances : undefined;
            inputs["internal"] = state ? state.internal : undefined;
            inputs["listeners"] = state ? state.listeners : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["sourceSecurityGroup"] = state ? state.sourceSecurityGroup : undefined;
            inputs["sourceSecurityGroupId"] = state ? state.sourceSecurityGroupId : undefined;
            inputs["subnets"] = state ? state.subnets : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            if ((!args || args.listeners === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listeners'");
            }
            inputs["accessLogs"] = args ? args.accessLogs : undefined;
            inputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            inputs["connectionDraining"] = args ? args.connectionDraining : undefined;
            inputs["connectionDrainingTimeout"] = args ? args.connectionDrainingTimeout : undefined;
            inputs["crossZoneLoadBalancing"] = args ? args.crossZoneLoadBalancing : undefined;
            inputs["healthCheck"] = args ? args.healthCheck : undefined;
            inputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            inputs["instances"] = args ? args.instances : undefined;
            inputs["internal"] = args ? args.internal : undefined;
            inputs["listeners"] = args ? args.listeners : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["sourceSecurityGroup"] = args ? args.sourceSecurityGroup : undefined;
            inputs["subnets"] = args ? args.subnets : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["dnsName"] = undefined /*out*/;
            inputs["sourceSecurityGroupId"] = undefined /*out*/;
            inputs["zoneId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "aws:elasticloadbalancing/loadBalancer:LoadBalancer" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(LoadBalancer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * An Access Logs block. Access Logs documented below.
     */
    readonly accessLogs?: pulumi.Input<inputs.elb.LoadBalancerAccessLogs>;
    /**
     * The ARN of the ELB
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The AZ's to serve traffic in.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean to enable connection draining. Default: `false`
     */
    readonly connectionDraining?: pulumi.Input<boolean>;
    /**
     * The time in seconds to allow for connections to drain. Default: `300`
     */
    readonly connectionDrainingTimeout?: pulumi.Input<number>;
    /**
     * Enable cross-zone load balancing. Default: `true`
     */
    readonly crossZoneLoadBalancing?: pulumi.Input<boolean>;
    /**
     * The DNS name of the ELB
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * A healthCheck block. Health Check documented below.
     */
    readonly healthCheck?: pulumi.Input<inputs.elb.LoadBalancerHealthCheck>;
    /**
     * The time in seconds that the connection is allowed to be idle. Default: `60`
     */
    readonly idleTimeout?: pulumi.Input<number>;
    /**
     * A list of instance ids to place in the ELB pool.
     */
    readonly instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, ELB will be an internal ELB.
     */
    readonly internal?: pulumi.Input<boolean>;
    /**
     * A list of listener blocks. Listeners documented below.
     */
    readonly listeners?: pulumi.Input<pulumi.Input<inputs.elb.LoadBalancerListener>[]>;
    /**
     * The name of the ELB. By default generated by this provider.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * A list of security group IDs to assign to the ELB.
     * Only valid if creating an ELB within a VPC
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the security group that you can use as
     * part of your inbound rules for your load balancer's back-end application
     * instances. Use this for Classic or Default VPC only.
     */
    readonly sourceSecurityGroup?: pulumi.Input<string>;
    /**
     * The ID of the security group that you can use as
     * part of your inbound rules for your load balancer's back-end application
     * instances. Only available on ELBs launched in a VPC.
     */
    readonly sourceSecurityGroupId?: pulumi.Input<string>;
    /**
     * A list of subnet IDs to attach to the ELB.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The canonical hosted zone ID of the ELB (to be used in a Route 53 Alias record)
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * An Access Logs block. Access Logs documented below.
     */
    readonly accessLogs?: pulumi.Input<inputs.elb.LoadBalancerAccessLogs>;
    /**
     * The AZ's to serve traffic in.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean to enable connection draining. Default: `false`
     */
    readonly connectionDraining?: pulumi.Input<boolean>;
    /**
     * The time in seconds to allow for connections to drain. Default: `300`
     */
    readonly connectionDrainingTimeout?: pulumi.Input<number>;
    /**
     * Enable cross-zone load balancing. Default: `true`
     */
    readonly crossZoneLoadBalancing?: pulumi.Input<boolean>;
    /**
     * A healthCheck block. Health Check documented below.
     */
    readonly healthCheck?: pulumi.Input<inputs.elb.LoadBalancerHealthCheck>;
    /**
     * The time in seconds that the connection is allowed to be idle. Default: `60`
     */
    readonly idleTimeout?: pulumi.Input<number>;
    /**
     * A list of instance ids to place in the ELB pool.
     */
    readonly instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, ELB will be an internal ELB.
     */
    readonly internal?: pulumi.Input<boolean>;
    /**
     * A list of listener blocks. Listeners documented below.
     */
    readonly listeners: pulumi.Input<pulumi.Input<inputs.elb.LoadBalancerListener>[]>;
    /**
     * The name of the ELB. By default generated by this provider.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * A list of security group IDs to assign to the ELB.
     * Only valid if creating an ELB within a VPC
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the security group that you can use as
     * part of your inbound rules for your load balancer's back-end application
     * instances. Use this for Classic or Default VPC only.
     */
    readonly sourceSecurityGroup?: pulumi.Input<string>;
    /**
     * A list of subnet IDs to attach to the ELB.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
