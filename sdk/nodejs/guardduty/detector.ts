// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage a GuardDuty detector.
 * 
 * > **NOTE:** Deleting this resource is equivalent to "disabling" GuardDuty for an AWS region, which removes all existing findings. You can set the `enable` attribute to `false` to instead "suspend" monitoring and feedback reporting while keeping existing data. See the [Suspending or Disabling Amazon GuardDuty documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_suspend-disable.html) for more information.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const myDetector = new aws.guardduty.Detector("MyDetector", {
 *     enable: true,
 * });
 * ```
 */
export class Detector extends pulumi.CustomResource {
    /**
     * Get an existing Detector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DetectorState, opts?: pulumi.CustomResourceOptions): Detector {
        return new Detector(name, <any>state, { ...opts, id: id });
    }

    /**
     * The AWS account ID of the GuardDuty detector
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty master account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty master accounts, it must be configured in Terraform to enable drift detection. Valid values for standalone and master accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
     */
    public readonly findingPublishingFrequency!: pulumi.Output<string>;

    /**
     * Create a Detector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DetectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DetectorArgs | DetectorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DetectorState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["enable"] = state ? state.enable : undefined;
            inputs["findingPublishingFrequency"] = state ? state.findingPublishingFrequency : undefined;
        } else {
            const args = argsOrState as DetectorArgs | undefined;
            inputs["enable"] = args ? args.enable : undefined;
            inputs["findingPublishingFrequency"] = args ? args.findingPublishingFrequency : undefined;
            inputs["accountId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("aws:guardduty/detector:Detector", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Detector resources.
 */
export interface DetectorState {
    /**
     * The AWS account ID of the GuardDuty detector
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
     */
    readonly enable?: pulumi.Input<boolean>;
    /**
     * Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty master account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty master accounts, it must be configured in Terraform to enable drift detection. Valid values for standalone and master accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
     */
    readonly findingPublishingFrequency?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Detector resource.
 */
export interface DetectorArgs {
    /**
     * Enable monitoring and feedback reporting. Setting to `false` is equivalent to "suspending" GuardDuty. Defaults to `true`.
     */
    readonly enable?: pulumi.Input<boolean>;
    /**
     * Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty master account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty master accounts, it must be configured in Terraform to enable drift detection. Valid values for standalone and master accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
     */
    readonly findingPublishingFrequency?: pulumi.Input<string>;
}
