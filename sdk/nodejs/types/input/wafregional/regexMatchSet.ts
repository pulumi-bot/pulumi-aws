// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface RegexMatchSetRegexMatchTuple {
    /**
     * The part of a web request that you want to search, such as a specified header or a query string.
     */
    fieldToMatch: pulumi.Input<inputApi.wafregional.RegexMatchSetRegexMatchTupleFieldToMatch>;
    /**
     * The ID of a [Regex Pattern Set](https://www.terraform.io/docs/providers/aws/r/waf_regex_pattern_set.html).
     */
    regexPatternSetId: pulumi.Input<string>;
    /**
     * Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
     * e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
     * See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
     * for all supported values.
     */
    textTransformation: pulumi.Input<string>;
}

export interface RegexMatchSetRegexMatchTupleFieldToMatch {
    /**
     * When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
     * If `type` is any other value, omit this field.
     */
    data?: pulumi.Input<string>;
    /**
     * The part of the web request that you want AWS WAF to search for a specified string.
     * e.g. `HEADER`, `METHOD` or `BODY`.
     * See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
     * for all supported values.
     */
    type: pulumi.Input<string>;
}
