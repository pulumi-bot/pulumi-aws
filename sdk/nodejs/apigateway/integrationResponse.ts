// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./restApi";

export class IntegrationResponse extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationResponse resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationResponseState, opts?: pulumi.CustomResourceOptions): IntegrationResponse {
        return new IntegrationResponse(name, <any>state, { ...opts, id: id });
    }

    public readonly contentHandling: pulumi.Output<string | undefined>;
    public readonly httpMethod: pulumi.Output<string>;
    public readonly resourceId: pulumi.Output<string>;
    public readonly responseParameters: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly responseParametersInJson: pulumi.Output<string | undefined>;
    public readonly responseTemplates: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly restApi: pulumi.Output<RestApi>;
    public readonly selectionPattern: pulumi.Output<string | undefined>;
    public readonly statusCode: pulumi.Output<string>;

    /**
     * Create a IntegrationResponse resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationResponseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationResponseArgs | IntegrationResponseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: IntegrationResponseState = argsOrState as IntegrationResponseState | undefined;
            inputs["contentHandling"] = state ? state.contentHandling : undefined;
            inputs["httpMethod"] = state ? state.httpMethod : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["responseParameters"] = state ? state.responseParameters : undefined;
            inputs["responseParametersInJson"] = state ? state.responseParametersInJson : undefined;
            inputs["responseTemplates"] = state ? state.responseTemplates : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["selectionPattern"] = state ? state.selectionPattern : undefined;
            inputs["statusCode"] = state ? state.statusCode : undefined;
        } else {
            const args = argsOrState as IntegrationResponseArgs | undefined;
            if (!args || args.httpMethod === undefined) {
                throw new Error("Missing required property 'httpMethod'");
            }
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.restApi === undefined) {
                throw new Error("Missing required property 'restApi'");
            }
            if (!args || args.statusCode === undefined) {
                throw new Error("Missing required property 'statusCode'");
            }
            inputs["contentHandling"] = args ? args.contentHandling : undefined;
            inputs["httpMethod"] = args ? args.httpMethod : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["responseParameters"] = args ? args.responseParameters : undefined;
            inputs["responseParametersInJson"] = args ? args.responseParametersInJson : undefined;
            inputs["responseTemplates"] = args ? args.responseTemplates : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["selectionPattern"] = args ? args.selectionPattern : undefined;
            inputs["statusCode"] = args ? args.statusCode : undefined;
        }
        super("aws:apigateway/integrationResponse:IntegrationResponse", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationResponse resources.
 */
export interface IntegrationResponseState {
    readonly contentHandling?: pulumi.Input<string>;
    readonly httpMethod?: pulumi.Input<string>;
    readonly resourceId?: pulumi.Input<string>;
    readonly responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly responseParametersInJson?: pulumi.Input<string>;
    readonly responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly restApi?: pulumi.Input<RestApi>;
    readonly selectionPattern?: pulumi.Input<string>;
    readonly statusCode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationResponse resource.
 */
export interface IntegrationResponseArgs {
    readonly contentHandling?: pulumi.Input<string>;
    readonly httpMethod: pulumi.Input<string>;
    readonly resourceId: pulumi.Input<string>;
    readonly responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly responseParametersInJson?: pulumi.Input<string>;
    readonly responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly restApi: pulumi.Input<RestApi>;
    readonly selectionPattern?: pulumi.Input<string>;
    readonly statusCode: pulumi.Input<string>;
}
