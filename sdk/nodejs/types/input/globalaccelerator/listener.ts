// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ListenerPortRange {
    /**
     * The first port in the range of ports, inclusive.
     */
    fromPort?: pulumi.Input<number>;
    /**
     * The last port in the range of ports, inclusive.
     */
    toPort?: pulumi.Input<number>;
}
