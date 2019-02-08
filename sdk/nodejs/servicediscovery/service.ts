// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Service Discovery Service resource.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleVpc = new aws.ec2.Vpc("example", {
 *     cidrBlock: "10.0.0.0/16",
 * });
 * const examplePrivateDnsNamespace = new aws.servicediscovery.PrivateDnsNamespace("example", {
 *     description: "example",
 *     vpc: exampleVpc.id,
 * });
 * const exampleService = new aws.servicediscovery.Service("example", {
 *     dnsConfig: {
 *         dnsRecords: [{
 *             ttl: 10,
 *             type: "A",
 *         }],
 *         namespaceId: examplePrivateDnsNamespace.id,
 *         routingPolicy: "MULTIVALUE",
 *     },
 *     healthCheckCustomConfig: {
 *         failureThreshold: 1,
 *     },
 * });
 * ```
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const examplePublicDnsNamespace = new aws.servicediscovery.PublicDnsNamespace("example", {
 *     description: "example",
 * });
 * const exampleService = new aws.servicediscovery.Service("example", {
 *     dnsConfig: {
 *         dnsRecords: [{
 *             ttl: 10,
 *             type: "A",
 *         }],
 *         namespaceId: examplePublicDnsNamespace.id,
 *     },
 *     healthCheckConfig: {
 *         failureThreshold: 10,
 *         resourcePath: "path",
 *         type: "HTTP",
 *     },
 * });
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN of the service.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The description of the service.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
     */
    public readonly dnsConfig: pulumi.Output<{ dnsRecords: { ttl: number, type: string }[], namespaceId: string, routingPolicy?: string }>;
    /**
     * A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
     */
    public readonly healthCheckConfig: pulumi.Output<{ failureThreshold?: number, resourcePath?: string, type?: string } | undefined>;
    /**
     * A complex type that contains settings for ECS managed health checks.
     */
    public readonly healthCheckCustomConfig: pulumi.Output<{ failureThreshold?: number } | undefined>;
    /**
     * The name of the service.
     */
    public readonly name: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ServiceState = argsOrState as ServiceState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dnsConfig"] = state ? state.dnsConfig : undefined;
            inputs["healthCheckConfig"] = state ? state.healthCheckConfig : undefined;
            inputs["healthCheckCustomConfig"] = state ? state.healthCheckCustomConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if (!args || args.dnsConfig === undefined) {
                throw new Error("Missing required property 'dnsConfig'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["dnsConfig"] = args ? args.dnsConfig : undefined;
            inputs["healthCheckConfig"] = args ? args.healthCheckConfig : undefined;
            inputs["healthCheckCustomConfig"] = args ? args.healthCheckCustomConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:servicediscovery/service:Service", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * The ARN of the service.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The description of the service.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
     */
    readonly dnsConfig?: pulumi.Input<{ dnsRecords: pulumi.Input<pulumi.Input<{ ttl: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, namespaceId: pulumi.Input<string>, routingPolicy?: pulumi.Input<string> }>;
    /**
     * A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
     */
    readonly healthCheckConfig?: pulumi.Input<{ failureThreshold?: pulumi.Input<number>, resourcePath?: pulumi.Input<string>, type?: pulumi.Input<string> }>;
    /**
     * A complex type that contains settings for ECS managed health checks.
     */
    readonly healthCheckCustomConfig?: pulumi.Input<{ failureThreshold?: pulumi.Input<number> }>;
    /**
     * The name of the service.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * The description of the service.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A complex type that contains information about the resource record sets that you want Amazon Route 53 to create when you register an instance.
     */
    readonly dnsConfig: pulumi.Input<{ dnsRecords: pulumi.Input<pulumi.Input<{ ttl: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, namespaceId: pulumi.Input<string>, routingPolicy?: pulumi.Input<string> }>;
    /**
     * A complex type that contains settings for an optional health check. Only for Public DNS namespaces.
     */
    readonly healthCheckConfig?: pulumi.Input<{ failureThreshold?: pulumi.Input<number>, resourcePath?: pulumi.Input<string>, type?: pulumi.Input<string> }>;
    /**
     * A complex type that contains settings for ECS managed health checks.
     */
    readonly healthCheckCustomConfig?: pulumi.Input<{ failureThreshold?: pulumi.Input<number> }>;
    /**
     * The name of the service.
     */
    readonly name?: pulumi.Input<string>;
}
