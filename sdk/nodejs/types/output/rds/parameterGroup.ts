// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ParameterGroupParameter {
    /**
     * "immediate" (default), or "pending-reboot". Some
     * engines can't apply some parameters without a reboot, and you will need to
     * specify "pending-reboot" here.
     */
    applyMethod?: string;
    /**
     * The name of the DB parameter.
     */
    name: string;
    /**
     * The value of the DB parameter.
     */
    value: string;
}
