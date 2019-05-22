// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./restApi";

/**
 * Connects a custom domain name registered via `aws_api_gateway_domain_name`
 * with a deployed API so that its methods can be called via the
 * custom domain name.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 * 
 * const exampleDeployment = new aws.apigateway.Deployment("example", {
 *     // See aws_api_gateway_rest_api docs for how to create this
 *     restApi: aws_api_gateway_rest_api_MyDemoAPI.id,
 *     stageName: "live",
 * });
 * const exampleDomainName = new aws.apigateway.DomainName("example", {
 *     certificateBody: fs.readFileSync(`./example.com/example.crt`, "utf-8"),
 *     certificateChain: fs.readFileSync(`./example.com/ca.crt`, "utf-8"),
 *     certificateName: "example-api",
 *     certificatePrivateKey: fs.readFileSync(`./example.com/example.key`, "utf-8"),
 *     domainName: "example.com",
 * });
 * const test = new aws.apigateway.BasePathMapping("test", {
 *     restApi: aws_api_gateway_rest_api_MyDemoAPI.id,
 *     domainName: exampleDomainName.domainName,
 *     stageName: exampleDeployment.stageName,
 * });
 * ```
 */
export class BasePathMapping extends pulumi.CustomResource {
    /**
     * Get an existing BasePathMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BasePathMappingState, opts?: pulumi.CustomResourceOptions): BasePathMapping {
        return new BasePathMapping(name, <any>state, { ...opts, id: id });
    }

    /**
     * The id of the API to connect.
     */
    public readonly restApi!: pulumi.Output<RestApi>;
    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     */
    public readonly basePath!: pulumi.Output<string | undefined>;
    /**
     * The already-registered domain name to connect the API to.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     */
    public readonly stageName!: pulumi.Output<string | undefined>;

    /**
     * Create a BasePathMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BasePathMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BasePathMappingArgs | BasePathMappingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BasePathMappingState | undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["basePath"] = state ? state.basePath : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["stageName"] = state ? state.stageName : undefined;
        } else {
            const args = argsOrState as BasePathMappingArgs | undefined;
            if (!args || args.restApi === undefined) {
                throw new Error("Missing required property 'restApi'");
            }
            if (!args || args.domainName === undefined) {
                throw new Error("Missing required property 'domainName'");
            }
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["basePath"] = args ? args.basePath : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["stageName"] = args ? args.stageName : undefined;
        }
        super("aws:apigateway/basePathMapping:BasePathMapping", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BasePathMapping resources.
 */
export interface BasePathMappingState {
    /**
     * The id of the API to connect.
     */
    readonly restApi?: pulumi.Input<RestApi>;
    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     */
    readonly basePath?: pulumi.Input<string>;
    /**
     * The already-registered domain name to connect the API to.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     */
    readonly stageName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BasePathMapping resource.
 */
export interface BasePathMappingArgs {
    /**
     * The id of the API to connect.
     */
    readonly restApi: pulumi.Input<RestApi>;
    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     */
    readonly basePath?: pulumi.Input<string>;
    /**
     * The already-registered domain name to connect the API to.
     */
    readonly domainName: pulumi.Input<string>;
    /**
     * The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     */
    readonly stageName?: pulumi.Input<string>;
}
