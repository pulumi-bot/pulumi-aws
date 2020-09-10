// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getScript(args: GetScriptArgs, opts?: pulumi.InvokeOptions): Promise<GetScriptResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:glue/getScript:getScript", {
        "dagEdges": args.dagEdges,
        "dagNodes": args.dagNodes,
        "language": args.language,
    }, opts);
}

/**
 * A collection of arguments for invoking getScript.
 */
export interface GetScriptArgs {
    readonly dagEdges: inputs.glue.GetScriptDagEdge[];
    readonly dagNodes: inputs.glue.GetScriptDagNode[];
    readonly language?: string;
}

/**
 * A collection of values returned by getScript.
 */
export interface GetScriptResult {
    readonly dagEdges: outputs.glue.GetScriptDagEdge[];
    readonly dagNodes: outputs.glue.GetScriptDagNode[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly language?: string;
    readonly pythonScript: string;
    readonly scalaCode: string;
}
