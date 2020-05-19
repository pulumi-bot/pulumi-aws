// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The CloudFormation Stack data source allows access to stack
 * outputs and other useful data including the template body.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const network = pulumi.output(aws.cloudformation.getStack({
 *     name: "my-network-stack",
 * }, { async: true }));
 * const web = new aws.ec2.Instance("web", {
 *     ami: "ami-abb07bcb",
 *     instanceType: "t1.micro",
 *     subnetId: network.apply(network => network.outputs["SubnetId"]),
 *     tags: {
 *         Name: "HelloWorld",
 *     },
 * });
 * ```
 */
export function getStack(args: GetStackArgs, opts?: pulumi.InvokeOptions): Promise<GetStackResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:cloudformation/getStack:getStack", {
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getStack.
 */
export interface GetStackArgs {
    /**
     * The name of the stack
     */
    readonly name: string;
    /**
     * A map of tags associated with this stack.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getStack.
 */
export interface GetStackResult {
    /**
     * A list of capabilities
     */
    readonly capabilities: string[];
    /**
     * Description of the stack
     */
    readonly description: string;
    /**
     * Whether the rollback of the stack is disabled when stack creation fails
     */
    readonly disableRollback: boolean;
    /**
     * The ARN of the IAM role used to create the stack.
     */
    readonly iamRoleArn: string;
    readonly name: string;
    /**
     * A list of SNS topic ARNs to publish stack related events
     */
    readonly notificationArns: string[];
    /**
     * A map of outputs from the stack.
     */
    readonly outputs: {[key: string]: any};
    /**
     * A map of parameters that specify input parameters for the stack.
     */
    readonly parameters: {[key: string]: any};
    /**
     * A map of tags associated with this stack.
     */
    readonly tags: {[key: string]: any};
    /**
     * Structure containing the template body.
     */
    readonly templateBody: string;
    /**
     * The amount of time that can pass before the stack status becomes `CREATE_FAILED`
     */
    readonly timeoutInMinutes: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
