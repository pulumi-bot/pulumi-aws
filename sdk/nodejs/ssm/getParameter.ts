// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SSM Parameter data source.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const foo = pulumi.output(aws.ssm.getParameter({
 *     name: "foo",
 * }));
 * ```
 * 
 * > **Note:** The unencrypted value of a SecureString will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 * 
 * 
 * > **Note:** The data source is currently following the behavior of the [SSM API](https://docs.aws.amazon.com/sdk-for-go/api/service/ssm/#Parameter) to return a string value, regardless of parameter type. For type `StringList`, we can use the built-in [split()](https://www.terraform.io/docs/configuration/functions/split.html) function to get values in a list. Example: `split(",", data.aws_ssm_parameter.subnets.value)`
 */
export function getParameter(args: GetParameterArgs, opts?: pulumi.InvokeOptions): Promise<GetParameterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ssm/getParameter:getParameter", {
        "name": args.name,
        "withDecryption": args.withDecryption,
    }, opts);
}

/**
 * A collection of arguments for invoking getParameter.
 */
export interface GetParameterArgs {
    /**
     * The name of the parameter.
     */
    readonly name: string;
    /**
     * Whether to return decrypted `SecureString` value. Defaults to `true`.
     */
    readonly withDecryption?: boolean;
}

/**
 * A collection of values returned by getParameter.
 */
export interface GetParameterResult {
    readonly arn: string;
    readonly name: string;
    readonly type: string;
    readonly value: string;
    readonly withDecryption?: boolean;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
