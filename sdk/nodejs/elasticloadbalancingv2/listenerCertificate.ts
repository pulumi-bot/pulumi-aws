// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * @deprecated aws.elasticloadbalancingv2.ListenerCertificate has been deprecated in favor of aws.lb.ListenerCertificate
 */
export class ListenerCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ListenerCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerCertificateState, opts?: pulumi.CustomResourceOptions): ListenerCertificate {
        pulumi.log.warn("ListenerCertificate is deprecated: aws.elasticloadbalancingv2.ListenerCertificate has been deprecated in favor of aws.lb.ListenerCertificate")
        return new ListenerCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticloadbalancingv2/listenerCertificate:ListenerCertificate';

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

    public readonly certificateArn!: pulumi.Output<string>;
    public readonly listenerArn!: pulumi.Output<string>;

    /**
     * Create a ListenerCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated aws.elasticloadbalancingv2.ListenerCertificate has been deprecated in favor of aws.lb.ListenerCertificate */
    constructor(name: string, args: ListenerCertificateArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated aws.elasticloadbalancingv2.ListenerCertificate has been deprecated in favor of aws.lb.ListenerCertificate */
    constructor(name: string, argsOrState?: ListenerCertificateArgs | ListenerCertificateState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ListenerCertificate is deprecated: aws.elasticloadbalancingv2.ListenerCertificate has been deprecated in favor of aws.lb.ListenerCertificate")
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
        super(ListenerCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListenerCertificate resources.
 */
export interface ListenerCertificateState {
    readonly certificateArn?: pulumi.Input<string>;
    readonly listenerArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ListenerCertificate resource.
 */
export interface ListenerCertificateArgs {
    readonly certificateArn: pulumi.Input<string>;
    readonly listenerArn: pulumi.Input<string>;
}
