// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ClusterParameterGroupParameter {
    /**
     * Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
     */
    applyMethod?: string;
    /**
     * The name of the documentDB parameter.
     */
    name: string;
    /**
     * The value of the documentDB parameter.
     */
    value: string;
}
