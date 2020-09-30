// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS Inspector Rules Packages data source allows access to the list of AWS
 * Inspector Rules Packages which can be used by AWS Inspector within the region
 * configured in the provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const rules = aws.inspector.getRulesPackages({});
 * // e.g. Use in aws_inspector_assessment_template
 * const group = new aws.inspector.ResourceGroup("group", {tags: {
 *     test: "test",
 * }});
 * const assessmentAssessmentTarget = new aws.inspector.AssessmentTarget("assessmentAssessmentTarget", {resourceGroupArn: group.arn});
 * const assessmentAssessmentTemplate = new aws.inspector.AssessmentTemplate("assessmentAssessmentTemplate", {
 *     targetArn: assessmentAssessmentTarget.arn,
 *     duration: "60",
 *     rulesPackageArns: rules.then(rules => rules.arns),
 * });
 * ```
 */
export function getRulesPackages(opts?: pulumi.InvokeOptions): Promise<GetRulesPackagesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:inspector/getRulesPackages:getRulesPackages", {
    }, opts);
}

/**
 * A collection of values returned by getRulesPackages.
 */
export interface GetRulesPackagesResult {
    /**
     * A list of the AWS Inspector Rules Packages arns available in the AWS region.
     */
    readonly arns: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
