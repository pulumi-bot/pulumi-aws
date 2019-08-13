// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ByteMatchSetByteMatchTuple {
    /**
     * Settings for the ByteMatchTuple. FieldToMatch documented below.
     */
    fieldToMatch: pulumi.Input<inputApi.wafregional.ByteMatchSetByteMatchTupleFieldToMatch>;
    /**
     * Within the portion of a web request that you want to search.
     */
    positionalConstraint: pulumi.Input<string>;
    /**
     * The value that you want AWS WAF to search for. The maximum length of the value is 50 bytes.
     */
    targetString?: pulumi.Input<string>;
    /**
     * The formatting way for web request.
     */
    textTransformation: pulumi.Input<string>;
}

export interface ByteMatchSetByteMatchTupleFieldToMatch {
    /**
     * When the value of Type is HEADER, enter the name of the header that you want AWS WAF to search, for example, User-Agent or Referer. If the value of Type is any other value, omit Data.
     */
    data?: pulumi.Input<string>;
    /**
     * The part of the web request that you want AWS WAF to search for a specified string.
     */
    type: pulumi.Input<string>;
}
