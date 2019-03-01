// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Inspector assessment target
 */
export class AssessmentTarget extends pulumi.CustomResource {
    /**
     * Get an existing AssessmentTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssessmentTargetState, opts?: pulumi.CustomResourceOptions): AssessmentTarget {
        return new AssessmentTarget(name, <any>state, { ...opts, id: id });
    }

    /**
     * The target assessment ARN.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The name of the assessment target.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
     */
    public readonly resourceGroupArn: pulumi.Output<string | undefined>;

    /**
     * Create a AssessmentTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AssessmentTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssessmentTargetArgs | AssessmentTargetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AssessmentTargetState = argsOrState as AssessmentTargetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupArn"] = state ? state.resourceGroupArn : undefined;
        } else {
            const args = argsOrState as AssessmentTargetArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupArn"] = args ? args.resourceGroupArn : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:inspector/assessmentTarget:AssessmentTarget", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AssessmentTarget resources.
 */
export interface AssessmentTargetState {
    /**
     * The target assessment ARN.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name of the assessment target.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
     */
    readonly resourceGroupArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AssessmentTarget resource.
 */
export interface AssessmentTargetArgs {
    /**
     * The name of the assessment target.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
     */
    readonly resourceGroupArn?: pulumi.Input<string>;
}
