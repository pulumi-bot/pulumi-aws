// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./index";

/**
 * Manages an API Gateway REST Deployment. A deployment is a snapshot of the REST API configuration. The deployment can then be published to callable endpoints via the `aws.apigateway.Stage` resource and optionally managed further with the `aws.apigateway.BasePathMapping` resource, `aws.apigateway.DomainName` resource, and `awsApiMethodSettings` resource. For more information, see the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-deploy-api.html).
 *
 * To properly capture all REST API configuration in a deployment, this resource must have dependencies on all prior resources that manage resources/paths, methods, integrations, etc.
 *
 * * For REST APIs that are configured via OpenAPI specification (`aws.apigateway.RestApi` resource `body` argument), no special dependency setup is needed beyond referencing the  `id` attribute of that resource unless additional resources have further customized the REST API.
 * * When the REST API configuration involves other resources (`aws.apigateway.Integration` resource), the dependency setup can be done with implicit resource references in the `triggers` argument or explicit resource references using the [resource `dependsOn` custom option](https://www.pulumi.com/docs/intro/concepts/resources/#dependson). The `triggers` argument should be preferred over `dependsOn`, since `dependsOn` can only capture dependency ordering and will not cause the resource to recreate (redeploy the REST API) with upstream configuration changes.
 *
 * !> **WARNING:** It is recommended to use the `aws.apigateway.Stage` resource instead of managing an API Gateway Stage via the `stageName` argument of this resource. When this resource is recreated (REST API redeployment) with the `stageName` configured, the stage is deleted and recreated. This will cause a temporary service interruption, increase provide plan differences, and can require a second apply to recreate any downstream stage configuration such as associated `awsApiMethodSettings` resources.
 *
 * ## Example Usage
 */
export class Deployment extends pulumi.CustomResource {
    /**
     * Get an existing Deployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeploymentState, opts?: pulumi.CustomResourceOptions): Deployment {
        return new Deployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/deployment:Deployment';

    /**
     * Returns true if the given object is an instance of Deployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Deployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Deployment.__pulumiType;
    }

    /**
     * The creation date of the deployment
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * Description of the deployment
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The execution ARN to be used in `lambdaPermission` resource's `sourceArn`
     * when allowing API Gateway to invoke a Lambda function,
     * e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
     */
    public /*out*/ readonly executionArn!: pulumi.Output<string>;
    /**
     * The URL to invoke the API pointing to the stage,
     * e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
     */
    public /*out*/ readonly invokeUrl!: pulumi.Output<string>;
    /**
     * REST API identifier.
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * Description to set on the stage managed by the `stageName` argument.
     */
    public readonly stageDescription!: pulumi.Output<string | undefined>;
    /**
     * Name of the stage to create with this deployment. If the specified stage already exists, it will be updated to point to the new deployment. It is recommended to use the `aws.apigateway.Stage` resource instead to manage stages.
     */
    public readonly stageName!: pulumi.Output<string | undefined>;
    /**
     * Map of arbitrary keys and values that, when changed, will trigger a redeployment.
     */
    public readonly triggers!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map to set on the stage managed by the `stageName` argument.
     */
    public readonly variables!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Deployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentArgs | DeploymentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DeploymentState | undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["executionArn"] = state ? state.executionArn : undefined;
            inputs["invokeUrl"] = state ? state.invokeUrl : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["stageDescription"] = state ? state.stageDescription : undefined;
            inputs["stageName"] = state ? state.stageName : undefined;
            inputs["triggers"] = state ? state.triggers : undefined;
            inputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as DeploymentArgs | undefined;
            if ((!args || args.restApi === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'restApi'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["stageDescription"] = args ? args.stageDescription : undefined;
            inputs["stageName"] = args ? args.stageName : undefined;
            inputs["triggers"] = args ? args.triggers : undefined;
            inputs["variables"] = args ? args.variables : undefined;
            inputs["createdDate"] = undefined /*out*/;
            inputs["executionArn"] = undefined /*out*/;
            inputs["invokeUrl"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Deployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Deployment resources.
 */
export interface DeploymentState {
    /**
     * The creation date of the deployment
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * Description of the deployment
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The execution ARN to be used in `lambdaPermission` resource's `sourceArn`
     * when allowing API Gateway to invoke a Lambda function,
     * e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
     */
    readonly executionArn?: pulumi.Input<string>;
    /**
     * The URL to invoke the API pointing to the stage,
     * e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
     */
    readonly invokeUrl?: pulumi.Input<string>;
    /**
     * REST API identifier.
     */
    readonly restApi?: pulumi.Input<string | RestApi>;
    /**
     * Description to set on the stage managed by the `stageName` argument.
     */
    readonly stageDescription?: pulumi.Input<string>;
    /**
     * Name of the stage to create with this deployment. If the specified stage already exists, it will be updated to point to the new deployment. It is recommended to use the `aws.apigateway.Stage` resource instead to manage stages.
     */
    readonly stageName?: pulumi.Input<string>;
    /**
     * Map of arbitrary keys and values that, when changed, will trigger a redeployment.
     */
    readonly triggers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map to set on the stage managed by the `stageName` argument.
     */
    readonly variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Deployment resource.
 */
export interface DeploymentArgs {
    /**
     * Description of the deployment
     */
    readonly description?: pulumi.Input<string>;
    /**
     * REST API identifier.
     */
    readonly restApi: pulumi.Input<string | RestApi>;
    /**
     * Description to set on the stage managed by the `stageName` argument.
     */
    readonly stageDescription?: pulumi.Input<string>;
    /**
     * Name of the stage to create with this deployment. If the specified stage already exists, it will be updated to point to the new deployment. It is recommended to use the `aws.apigateway.Stage` resource instead to manage stages.
     */
    readonly stageName?: pulumi.Input<string>;
    /**
     * Map of arbitrary keys and values that, when changed, will trigger a redeployment.
     */
    readonly triggers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map to set on the stage managed by the `stageName` argument.
     */
    readonly variables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
