// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Load Balancer Listener Certificate resource.
 *
 * This resource is for additional certificates and does not replace the default certificate on the listener.
 *
 * > **Note:** `aws.alb.ListenerCertificate` is known as `aws.lb.ListenerCertificate`. The functionality is identical.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleCertificate = new aws.acm.Certificate("example", {});
 * const frontEndLoadBalancer = new aws.lb.LoadBalancer("frontEnd", {});
 * const frontEndListener = new aws.lb.Listener("frontEnd", {});
 * const exampleListenerCertificate = new aws.lb.ListenerCertificate("example", {
 *     certificateArn: exampleCertificate.arn,
 *     listenerArn: frontEndListener.arn,
 * });
 * ```
 */
export class ListenerCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ListenerCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerCertificateState, opts?: pulumi.CustomResourceOptions): ListenerCertificate {
        return new ListenerCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:alb/listenerCertificate:ListenerCertificate';

    /**
     * Returns true if the given object is an instance of ListenerCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ListenerCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ListenerCertificate.__pulumiType;
    }

    /**
     * The ARN of the certificate to attach to the listener.
     */
    public readonly certificateArn!: pulumi.Output<string>;
    /**
     * The ARN of the listener to which to attach the certificate.
     */
    public readonly listenerArn!: pulumi.Output<string>;

    /**
     * Create a ListenerCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerCertificateArgs | ListenerCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ListenerCertificateState | undefined;
            inputs["certificateArn"] = state ? state.certificateArn : undefined;
            inputs["listenerArn"] = state ? state.listenerArn : undefined;
        } else {
            const args = argsOrState as ListenerCertificateArgs | undefined;
            if (!args || args.certificateArn === undefined) {
                throw new Error("Missing required property 'certificateArn'");
            }
            if (!args || args.listenerArn === undefined) {
                throw new Error("Missing required property 'listenerArn'");
            }
            inputs["certificateArn"] = args ? args.certificateArn : undefined;
            inputs["listenerArn"] = args ? args.listenerArn : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "aws:applicationloadbalancing/listenerCertificate:ListenerCertificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ListenerCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListenerCertificate resources.
 */
export interface ListenerCertificateState {
    /**
     * The ARN of the certificate to attach to the listener.
     */
    readonly certificateArn?: pulumi.Input<string>;
    /**
     * The ARN of the listener to which to attach the certificate.
     */
    readonly listenerArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ListenerCertificate resource.
 */
export interface ListenerCertificateArgs {
    /**
     * The ARN of the certificate to attach to the listener.
     */
    readonly certificateArn: pulumi.Input<string>;
    /**
     * The ARN of the listener to which to attach the certificate.
     */
    readonly listenerArn: pulumi.Input<string>;
}
