// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Inspector assessment template
 */
export class AssessmentTemplate extends pulumi.CustomResource {
    /**
     * Get an existing AssessmentTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssessmentTemplateState, opts?: pulumi.CustomResourceOptions): AssessmentTemplate {
        return new AssessmentTemplate(name, <any>state, { ...opts, id: id });
    }

    /**
     * The template assessment ARN.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The duration of the inspector run.
     */
    public readonly duration: pulumi.Output<number>;
    /**
     * The name of the assessment template.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The rules to be used during the run.
     */
    public readonly rulesPackageArns: pulumi.Output<string[]>;
    /**
     * The assessment target ARN to attach the template to.
     */
    public readonly targetArn: pulumi.Output<string>;

    /**
     * Create a AssessmentTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssessmentTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssessmentTemplateArgs | AssessmentTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AssessmentTemplateState = argsOrState as AssessmentTemplateState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["duration"] = state ? state.duration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rulesPackageArns"] = state ? state.rulesPackageArns : undefined;
            inputs["targetArn"] = state ? state.targetArn : undefined;
        } else {
            const args = argsOrState as AssessmentTemplateArgs | undefined;
            if (!args || args.duration === undefined) {
                throw new Error("Missing required property 'duration'");
            }
            if (!args || args.rulesPackageArns === undefined) {
                throw new Error("Missing required property 'rulesPackageArns'");
            }
            if (!args || args.targetArn === undefined) {
                throw new Error("Missing required property 'targetArn'");
            }
            inputs["duration"] = args ? args.duration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rulesPackageArns"] = args ? args.rulesPackageArns : undefined;
            inputs["targetArn"] = args ? args.targetArn : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:inspector/assessmentTemplate:AssessmentTemplate", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AssessmentTemplate resources.
 */
export interface AssessmentTemplateState {
    /**
     * The template assessment ARN.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The duration of the inspector run.
     */
    readonly duration?: pulumi.Input<number>;
    /**
     * The name of the assessment template.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The rules to be used during the run.
     */
    readonly rulesPackageArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The assessment target ARN to attach the template to.
     */
    readonly targetArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AssessmentTemplate resource.
 */
export interface AssessmentTemplateArgs {
    /**
     * The duration of the inspector run.
     */
    readonly duration: pulumi.Input<number>;
    /**
     * The name of the assessment template.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The rules to be used during the run.
     */
    readonly rulesPackageArns: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The assessment target ARN to attach the template to.
     */
    readonly targetArn: pulumi.Input<string>;
}
