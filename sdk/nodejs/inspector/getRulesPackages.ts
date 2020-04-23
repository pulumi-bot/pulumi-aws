// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS Inspector Rules Packages data source allows access to the list of AWS
 * Inspector Rules Packages which can be used by AWS Inspector within the region
 * configured in the provider.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * // Declare the data source
 * const rules = aws.inspector.getRulesPackages();
 * // e.g. Use in aws.inspector.AssessmentTemplate
 * const group = new aws.inspector.ResourceGroup("group", {
 *     tags: {
 *         test: "test",
 *     },
 * });
 * const assessmentAssessmentTarget = new aws.inspector.AssessmentTarget("assessment", {
 *     resourceGroupArn: group.arn,
 * });
 * const assessmentAssessmentTemplate = new aws.inspector.AssessmentTemplate("assessment", {
 *     duration: 60,
 *     rulesPackageArns: rules.arns,
 *     targetArn: assessmentAssessmentTarget.arn,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/inspector_rules_packages.html.markdown.
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
