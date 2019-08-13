// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface SelectionSelectionTag {
    /**
     * The key in a key-value pair.
     */
    key: string;
    /**
     * An operation, such as `StringEquals`, that is applied to a key-value pair used to filter resources in a selection.
     */
    type: string;
    /**
     * The value in a key-value pair.
     */
    value: string;
}
