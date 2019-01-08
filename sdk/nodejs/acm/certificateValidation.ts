// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CertificateValidation extends pulumi.CustomResource {
    /**
     * Get an existing CertificateValidation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateValidationState, opts?: pulumi.CustomResourceOptions): CertificateValidation {
        return new CertificateValidation(name, <any>state, { ...opts, id: id });
    }

    public readonly certificateArn: pulumi.Output<string>;
    public readonly validationRecordFqdns: pulumi.Output<string[] | undefined>;

    /**
     * Create a CertificateValidation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateValidationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateValidationArgs | CertificateValidationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CertificateValidationState = argsOrState as CertificateValidationState | undefined;
            inputs["certificateArn"] = state ? state.certificateArn : undefined;
            inputs["validationRecordFqdns"] = state ? state.validationRecordFqdns : undefined;
        } else {
            const args = argsOrState as CertificateValidationArgs | undefined;
            if (!args || args.certificateArn === undefined) {
                throw new Error("Missing required property 'certificateArn'");
            }
            inputs["certificateArn"] = args ? args.certificateArn : undefined;
            inputs["validationRecordFqdns"] = args ? args.validationRecordFqdns : undefined;
        }
        super("aws:acm/certificateValidation:CertificateValidation", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertificateValidation resources.
 */
export interface CertificateValidationState {
    readonly certificateArn?: pulumi.Input<string>;
    readonly validationRecordFqdns?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a CertificateValidation resource.
 */
export interface CertificateValidationArgs {
    readonly certificateArn: pulumi.Input<string>;
    readonly validationRecordFqdns?: pulumi.Input<pulumi.Input<string>[]>;
}
