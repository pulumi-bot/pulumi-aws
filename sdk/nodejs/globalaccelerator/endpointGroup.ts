// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator endpoint group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.globalaccelerator.EndpointGroup("example", {
 *     listenerArn: aws_globalaccelerator_listener.example.id,
 *     endpointConfigurations: [{
 *         endpointId: aws_lb.example.arn,
 *         weight: 100,
 *     }],
 * });
 * ```
 */
export class EndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing EndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointGroupState, opts?: pulumi.CustomResourceOptions): EndpointGroup {
        return new EndpointGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:globalaccelerator/endpointGroup:EndpointGroup';

    /**
     * Returns true if the given object is an instance of EndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointGroup.__pulumiType;
    }

    /**
     * The list of endpoint objects. Fields documented below.
     */
    public readonly endpointConfigurations!: pulumi.Output<outputs.globalaccelerator.EndpointGroupEndpointConfiguration[] | undefined>;
    /**
     * The name of the AWS Region where the endpoint group is located.
     */
    public readonly endpointGroupRegion!: pulumi.Output<string>;
    /**
     * The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
     */
    public readonly healthCheckIntervalSeconds!: pulumi.Output<number | undefined>;
    public readonly healthCheckPath!: pulumi.Output<string>;
    public readonly healthCheckPort!: pulumi.Output<number>;
    /**
     * The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
     */
    public readonly healthCheckProtocol!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the listener.
     */
    public readonly listenerArn!: pulumi.Output<string>;
    /**
     * The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
     */
    public readonly thresholdCount!: pulumi.Output<number | undefined>;
    /**
     * The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
     */
    public readonly trafficDialPercentage!: pulumi.Output<number | undefined>;

    /**
     * Create a EndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointGroupArgs | EndpointGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EndpointGroupState | undefined;
            inputs["endpointConfigurations"] = state ? state.endpointConfigurations : undefined;
            inputs["endpointGroupRegion"] = state ? state.endpointGroupRegion : undefined;
            inputs["healthCheckIntervalSeconds"] = state ? state.healthCheckIntervalSeconds : undefined;
            inputs["healthCheckPath"] = state ? state.healthCheckPath : undefined;
            inputs["healthCheckPort"] = state ? state.healthCheckPort : undefined;
            inputs["healthCheckProtocol"] = state ? state.healthCheckProtocol : undefined;
            inputs["listenerArn"] = state ? state.listenerArn : undefined;
            inputs["thresholdCount"] = state ? state.thresholdCount : undefined;
            inputs["trafficDialPercentage"] = state ? state.trafficDialPercentage : undefined;
        } else {
            const args = argsOrState as EndpointGroupArgs | undefined;
            if (!args || args.listenerArn === undefined) {
                throw new Error("Missing required property 'listenerArn'");
            }
            inputs["endpointConfigurations"] = args ? args.endpointConfigurations : undefined;
            inputs["endpointGroupRegion"] = args ? args.endpointGroupRegion : undefined;
            inputs["healthCheckIntervalSeconds"] = args ? args.healthCheckIntervalSeconds : undefined;
            inputs["healthCheckPath"] = args ? args.healthCheckPath : undefined;
            inputs["healthCheckPort"] = args ? args.healthCheckPort : undefined;
            inputs["healthCheckProtocol"] = args ? args.healthCheckProtocol : undefined;
            inputs["listenerArn"] = args ? args.listenerArn : undefined;
            inputs["thresholdCount"] = args ? args.thresholdCount : undefined;
            inputs["trafficDialPercentage"] = args ? args.trafficDialPercentage : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EndpointGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointGroup resources.
 */
export interface EndpointGroupState {
    /**
     * The list of endpoint objects. Fields documented below.
     */
    readonly endpointConfigurations?: pulumi.Input<pulumi.Input<inputs.globalaccelerator.EndpointGroupEndpointConfiguration>[]>;
    /**
     * The name of the AWS Region where the endpoint group is located.
     */
    readonly endpointGroupRegion?: pulumi.Input<string>;
    /**
     * The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
     */
    readonly healthCheckIntervalSeconds?: pulumi.Input<number>;
    readonly healthCheckPath?: pulumi.Input<string>;
    readonly healthCheckPort?: pulumi.Input<number>;
    /**
     * The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
     */
    readonly healthCheckProtocol?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the listener.
     */
    readonly listenerArn?: pulumi.Input<string>;
    /**
     * The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
     */
    readonly thresholdCount?: pulumi.Input<number>;
    /**
     * The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
     */
    readonly trafficDialPercentage?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a EndpointGroup resource.
 */
export interface EndpointGroupArgs {
    /**
     * The list of endpoint objects. Fields documented below.
     */
    readonly endpointConfigurations?: pulumi.Input<pulumi.Input<inputs.globalaccelerator.EndpointGroupEndpointConfiguration>[]>;
    /**
     * The name of the AWS Region where the endpoint group is located.
     */
    readonly endpointGroupRegion?: pulumi.Input<string>;
    /**
     * The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
     */
    readonly healthCheckIntervalSeconds?: pulumi.Input<number>;
    readonly healthCheckPath?: pulumi.Input<string>;
    readonly healthCheckPort?: pulumi.Input<number>;
    /**
     * The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
     */
    readonly healthCheckProtocol?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the listener.
     */
    readonly listenerArn: pulumi.Input<string>;
    /**
     * The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
     */
    readonly thresholdCount?: pulumi.Input<number>;
    /**
     * The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
     */
    readonly trafficDialPercentage?: pulumi.Input<number>;
}
