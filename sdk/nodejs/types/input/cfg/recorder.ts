// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface RecorderRecordingGroup {
    /**
     * Specifies whether AWS Config records configuration changes
     * for every supported type of regional resource (which includes any new type that will become supported in the future).
     * Conflicts with `resourceTypes`. Defaults to `true`.
     */
    allSupported?: pulumi.Input<boolean>;
    /**
     * Specifies whether AWS Config includes all supported types of *global resources*
     * with the resources that it records. Requires `allSupported = true`. Conflicts with `resourceTypes`.
     */
    includeGlobalResourceTypes?: pulumi.Input<boolean>;
    /**
     * A list that specifies the types of AWS resources for which
     * AWS Config records configuration changes (for example, `AWS::EC2::Instance` or `AWS::CloudTrail::Trail`).
     * See [relevant part of AWS Docs](http://docs.aws.amazon.com/config/latest/APIReference/API_ResourceIdentifier.html#config-Type-ResourceIdentifier-resourceType) for available types.
     */
    resourceTypes?: pulumi.Input<pulumi.Input<string>[]>;
}
