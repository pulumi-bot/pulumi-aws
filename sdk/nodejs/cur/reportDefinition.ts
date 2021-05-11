// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages Cost and Usage Report Definitions.
 *
 * > *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.
 *
 * > *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleCurReportDefinition = new aws.cur.ReportDefinition("example_cur_report_definition", {
 *     additionalArtifacts: [
 *         "REDSHIFT",
 *         "QUICKSIGHT",
 *     ],
 *     additionalSchemaElements: ["RESOURCES"],
 *     compression: "GZIP",
 *     format: "textORcsv",
 *     reportName: "example-cur-report-definition",
 *     s3Bucket: "example-bucket-name",
 *     s3Region: "us-east-1",
 *     timeUnit: "HOURLY",
 * });
 * ```
 *
 * ## Import
 *
 * Report Definitions can be imported using the `report_name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:cur/reportDefinition:ReportDefinition example_cur_report_definition example-cur-report-definition
 * ```
 */
export class ReportDefinition extends pulumi.CustomResource {
    /**
     * Get an existing ReportDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReportDefinitionState, opts?: pulumi.CustomResourceOptions): ReportDefinition {
        return new ReportDefinition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cur/reportDefinition:ReportDefinition';

    /**
     * Returns true if the given object is an instance of ReportDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReportDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReportDefinition.__pulumiType;
    }

    /**
     * A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT, ATHENA. When ATHENA exists within additional_artifacts, no other artifact type can be declared and reportVersioning must be OVERWRITE_REPORT.
     */
    public readonly additionalArtifacts!: pulumi.Output<string[] | undefined>;
    /**
     * A list of schema elements. Valid values are: RESOURCES.
     */
    public readonly additionalSchemaElements!: pulumi.Output<string[]>;
    /**
     * Compression format for report. Valid values are: GZIP, ZIP, Parquet. If Parquet is used, then format must also be Parquet.
     */
    public readonly compression!: pulumi.Output<string>;
    /**
     * Format for report. Valid values are: textORcsv, Parquet. If Parquet is used, then Compression must also be Parquet.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
     */
    public readonly refreshClosedReports!: pulumi.Output<boolean | undefined>;
    /**
     * Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
     */
    public readonly reportName!: pulumi.Output<string>;
    /**
     * Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: CREATE_NEW_REPORT, OVERWRITE_REPORT
     */
    public readonly reportVersioning!: pulumi.Output<string | undefined>;
    /**
     * Name of the existing S3 bucket to hold generated reports.
     */
    public readonly s3Bucket!: pulumi.Output<string>;
    /**
     * Report path prefix. Limited to 256 characters.
     */
    public readonly s3Prefix!: pulumi.Output<string | undefined>;
    /**
     * Region of the existing S3 bucket to hold generated reports.
     */
    public readonly s3Region!: pulumi.Output<string>;
    /**
     * The frequency on which report data are measured and displayed.  Valid values are: HOURLY, DAILY.
     */
    public readonly timeUnit!: pulumi.Output<string>;

    /**
     * Create a ReportDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReportDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReportDefinitionArgs | ReportDefinitionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReportDefinitionState | undefined;
            inputs["additionalArtifacts"] = state ? state.additionalArtifacts : undefined;
            inputs["additionalSchemaElements"] = state ? state.additionalSchemaElements : undefined;
            inputs["compression"] = state ? state.compression : undefined;
            inputs["format"] = state ? state.format : undefined;
            inputs["refreshClosedReports"] = state ? state.refreshClosedReports : undefined;
            inputs["reportName"] = state ? state.reportName : undefined;
            inputs["reportVersioning"] = state ? state.reportVersioning : undefined;
            inputs["s3Bucket"] = state ? state.s3Bucket : undefined;
            inputs["s3Prefix"] = state ? state.s3Prefix : undefined;
            inputs["s3Region"] = state ? state.s3Region : undefined;
            inputs["timeUnit"] = state ? state.timeUnit : undefined;
        } else {
            const args = argsOrState as ReportDefinitionArgs | undefined;
            if ((!args || args.additionalSchemaElements === undefined) && !opts.urn) {
                throw new Error("Missing required property 'additionalSchemaElements'");
            }
            if ((!args || args.compression === undefined) && !opts.urn) {
                throw new Error("Missing required property 'compression'");
            }
            if ((!args || args.format === undefined) && !opts.urn) {
                throw new Error("Missing required property 'format'");
            }
            if ((!args || args.reportName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reportName'");
            }
            if ((!args || args.s3Bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 's3Bucket'");
            }
            if ((!args || args.s3Region === undefined) && !opts.urn) {
                throw new Error("Missing required property 's3Region'");
            }
            if ((!args || args.timeUnit === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeUnit'");
            }
            inputs["additionalArtifacts"] = args ? args.additionalArtifacts : undefined;
            inputs["additionalSchemaElements"] = args ? args.additionalSchemaElements : undefined;
            inputs["compression"] = args ? args.compression : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["refreshClosedReports"] = args ? args.refreshClosedReports : undefined;
            inputs["reportName"] = args ? args.reportName : undefined;
            inputs["reportVersioning"] = args ? args.reportVersioning : undefined;
            inputs["s3Bucket"] = args ? args.s3Bucket : undefined;
            inputs["s3Prefix"] = args ? args.s3Prefix : undefined;
            inputs["s3Region"] = args ? args.s3Region : undefined;
            inputs["timeUnit"] = args ? args.timeUnit : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ReportDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReportDefinition resources.
 */
export interface ReportDefinitionState {
    /**
     * A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT, ATHENA. When ATHENA exists within additional_artifacts, no other artifact type can be declared and reportVersioning must be OVERWRITE_REPORT.
     */
    additionalArtifacts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of schema elements. Valid values are: RESOURCES.
     */
    additionalSchemaElements?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Compression format for report. Valid values are: GZIP, ZIP, Parquet. If Parquet is used, then format must also be Parquet.
     */
    compression?: pulumi.Input<string>;
    /**
     * Format for report. Valid values are: textORcsv, Parquet. If Parquet is used, then Compression must also be Parquet.
     */
    format?: pulumi.Input<string>;
    /**
     * Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
     */
    refreshClosedReports?: pulumi.Input<boolean>;
    /**
     * Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
     */
    reportName?: pulumi.Input<string>;
    /**
     * Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: CREATE_NEW_REPORT, OVERWRITE_REPORT
     */
    reportVersioning?: pulumi.Input<string>;
    /**
     * Name of the existing S3 bucket to hold generated reports.
     */
    s3Bucket?: pulumi.Input<string>;
    /**
     * Report path prefix. Limited to 256 characters.
     */
    s3Prefix?: pulumi.Input<string>;
    /**
     * Region of the existing S3 bucket to hold generated reports.
     */
    s3Region?: pulumi.Input<string>;
    /**
     * The frequency on which report data are measured and displayed.  Valid values are: HOURLY, DAILY.
     */
    timeUnit?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReportDefinition resource.
 */
export interface ReportDefinitionArgs {
    /**
     * A list of additional artifacts. Valid values are: REDSHIFT, QUICKSIGHT, ATHENA. When ATHENA exists within additional_artifacts, no other artifact type can be declared and reportVersioning must be OVERWRITE_REPORT.
     */
    additionalArtifacts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of schema elements. Valid values are: RESOURCES.
     */
    additionalSchemaElements: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Compression format for report. Valid values are: GZIP, ZIP, Parquet. If Parquet is used, then format must also be Parquet.
     */
    compression: pulumi.Input<string>;
    /**
     * Format for report. Valid values are: textORcsv, Parquet. If Parquet is used, then Compression must also be Parquet.
     */
    format: pulumi.Input<string>;
    /**
     * Set to true to update your reports after they have been finalized if AWS detects charges related to previous months.
     */
    refreshClosedReports?: pulumi.Input<boolean>;
    /**
     * Unique name for the report. Must start with a number/letter and is case sensitive. Limited to 256 characters.
     */
    reportName: pulumi.Input<string>;
    /**
     * Overwrite the previous version of each report or to deliver the report in addition to the previous versions. Valid values are: CREATE_NEW_REPORT, OVERWRITE_REPORT
     */
    reportVersioning?: pulumi.Input<string>;
    /**
     * Name of the existing S3 bucket to hold generated reports.
     */
    s3Bucket: pulumi.Input<string>;
    /**
     * Report path prefix. Limited to 256 characters.
     */
    s3Prefix?: pulumi.Input<string>;
    /**
     * Region of the existing S3 bucket to hold generated reports.
     */
    s3Region: pulumi.Input<string>;
    /**
     * The frequency on which report data are measured and displayed.  Valid values are: HOURLY, DAILY.
     */
    timeUnit: pulumi.Input<string>;
}
