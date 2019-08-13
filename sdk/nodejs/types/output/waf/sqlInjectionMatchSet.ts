// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface SqlInjectionMatchSetSqlInjectionMatchTuple {
    /**
     * Specifies where in a web request to look for snippets of malicious SQL code.
     */
    fieldToMatch: outputApi.waf.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch;
    /**
     * Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
     * If you specify a transformation, AWS WAF performs the transformation on `fieldToMatch` before inspecting a request for a match.
     * e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
     * See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SqlInjectionMatchTuple.html#WAF-Type-SqlInjectionMatchTuple-TextTransformation)
     * for all supported values.
     */
    textTransformation: string;
}

export interface SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch {
    /**
     * When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
     * If `type` is any other value, omit this field.
     */
    data?: string;
    /**
     * The part of the web request that you want AWS WAF to search for a specified string.
     * e.g. `HEADER`, `METHOD` or `BODY`.
     * See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
     * for all supported values.
     */
    type: string;
}
