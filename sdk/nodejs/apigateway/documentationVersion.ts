// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an API Gateway Documentation Version.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_api_gateway_rest_api_example = new aws.apigateway.RestApi("example", {
 *     name: "example_api",
 * });
 * const aws_api_gateway_documentation_part_example = new aws.apigateway.DocumentationPart("example", {
 *     location: {
 *         type: "API",
 *     },
 *     properties: "{\"description\":\"Example\"}",
 *     restApiId: aws_api_gateway_rest_api_example.id,
 * });
 * const aws_api_gateway_documentation_version_example = new aws.apigateway.DocumentationVersion("example", {
 *     description: "Example description",
 *     restApiId: aws_api_gateway_rest_api_example.id,
 *     version: "example_version",
 * }, {dependsOn: [aws_api_gateway_documentation_part_example]});
 * ```
 */
export class DocumentationVersion extends pulumi.CustomResource {
    /**
     * Get an existing DocumentationVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentationVersionState, opts?: pulumi.CustomResourceOptions): DocumentationVersion {
        return new DocumentationVersion(name, <any>state, { ...opts, id: id });
    }

    /**
     * The description of the API documentation version.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The ID of the associated Rest API
     */
    public readonly restApiId: pulumi.Output<string>;
    /**
     * The version identifier of the API documentation snapshot.
     */
    public readonly version: pulumi.Output<string>;

    /**
     * Create a DocumentationVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentationVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DocumentationVersionArgs | DocumentationVersionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DocumentationVersionState = argsOrState as DocumentationVersionState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["restApiId"] = state ? state.restApiId : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as DocumentationVersionArgs | undefined;
            if (!args || args.restApiId === undefined) {
                throw new Error("Missing required property 'restApiId'");
            }
            if (!args || args.version === undefined) {
                throw new Error("Missing required property 'version'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["restApiId"] = args ? args.restApiId : undefined;
            inputs["version"] = args ? args.version : undefined;
        }
        super("aws:apigateway/documentationVersion:DocumentationVersion", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DocumentationVersion resources.
 */
export interface DocumentationVersionState {
    /**
     * The description of the API documentation version.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the associated Rest API
     */
    readonly restApiId?: pulumi.Input<string>;
    /**
     * The version identifier of the API documentation snapshot.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DocumentationVersion resource.
 */
export interface DocumentationVersionArgs {
    /**
     * The description of the API documentation version.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the associated Rest API
     */
    readonly restApiId: pulumi.Input<string>;
    /**
     * The version identifier of the API documentation snapshot.
     */
    readonly version: pulumi.Input<string>;
}
