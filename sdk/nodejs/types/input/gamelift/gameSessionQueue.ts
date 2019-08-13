// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface GameSessionQueuePlayerLatencyPolicy {
    /**
     * Maximum latency value that is allowed for any player.
     */
    maximumIndividualPlayerLatencyMilliseconds: pulumi.Input<number>;
    /**
     * Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
     */
    policyDurationSeconds?: pulumi.Input<number>;
}
