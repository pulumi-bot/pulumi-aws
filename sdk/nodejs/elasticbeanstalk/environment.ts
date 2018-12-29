// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {Application} from "./application";
import {ApplicationVersion} from "./applicationVersion";

export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentState, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly allSettings: pulumi.Output<{ name: string, namespace: string, resource?: string, value: string }[]>;
    public readonly application: pulumi.Output<Application>;
    public /*out*/ readonly arn: pulumi.Output<string>;
    public /*out*/ readonly autoscalingGroups: pulumi.Output<string[]>;
    public /*out*/ readonly cname: pulumi.Output<string>;
    public readonly cnamePrefix: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public /*out*/ readonly instances: pulumi.Output<string[]>;
    public /*out*/ readonly launchConfigurations: pulumi.Output<string[]>;
    public /*out*/ readonly loadBalancers: pulumi.Output<string[]>;
    public readonly name: pulumi.Output<string>;
    public readonly platformArn: pulumi.Output<string>;
    public readonly pollInterval: pulumi.Output<string | undefined>;
    public /*out*/ readonly queues: pulumi.Output<string[]>;
    public readonly settings: pulumi.Output<{ name: string, namespace: string, resource?: string, value: string }[] | undefined>;
    public readonly solutionStackName: pulumi.Output<string>;
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly templateName: pulumi.Output<string | undefined>;
    public readonly tier: pulumi.Output<string | undefined>;
    public /*out*/ readonly triggers: pulumi.Output<string[]>;
    public readonly version: pulumi.Output<ApplicationVersion>;
    public readonly waitForReadyTimeout: pulumi.Output<string | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentArgs | EnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: EnvironmentState = argsOrState as EnvironmentState | undefined;
            inputs["allSettings"] = state ? state.allSettings : undefined;
            inputs["application"] = state ? state.application : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["autoscalingGroups"] = state ? state.autoscalingGroups : undefined;
            inputs["cname"] = state ? state.cname : undefined;
            inputs["cnamePrefix"] = state ? state.cnamePrefix : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["instances"] = state ? state.instances : undefined;
            inputs["launchConfigurations"] = state ? state.launchConfigurations : undefined;
            inputs["loadBalancers"] = state ? state.loadBalancers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["platformArn"] = state ? state.platformArn : undefined;
            inputs["pollInterval"] = state ? state.pollInterval : undefined;
            inputs["queues"] = state ? state.queues : undefined;
            inputs["settings"] = state ? state.settings : undefined;
            inputs["solutionStackName"] = state ? state.solutionStackName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["templateName"] = state ? state.templateName : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["triggers"] = state ? state.triggers : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["waitForReadyTimeout"] = state ? state.waitForReadyTimeout : undefined;
        } else {
            const args = argsOrState as EnvironmentArgs | undefined;
            if (!args || args.application === undefined) {
                throw new Error("Missing required property 'application'");
            }
            inputs["application"] = args ? args.application : undefined;
            inputs["cnamePrefix"] = args ? args.cnamePrefix : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["platformArn"] = args ? args.platformArn : undefined;
            inputs["pollInterval"] = args ? args.pollInterval : undefined;
            inputs["settings"] = args ? args.settings : undefined;
            inputs["solutionStackName"] = args ? args.solutionStackName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["templateName"] = args ? args.templateName : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["waitForReadyTimeout"] = args ? args.waitForReadyTimeout : undefined;
            inputs["allSettings"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["autoscalingGroups"] = undefined /*out*/;
            inputs["cname"] = undefined /*out*/;
            inputs["instances"] = undefined /*out*/;
            inputs["launchConfigurations"] = undefined /*out*/;
            inputs["loadBalancers"] = undefined /*out*/;
            inputs["queues"] = undefined /*out*/;
            inputs["triggers"] = undefined /*out*/;
        }
        super("aws:elasticbeanstalk/environment:Environment", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Environment resources.
 */
export interface EnvironmentState {
    readonly allSettings?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, namespace: pulumi.Input<string>, resource?: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
    readonly application?: pulumi.Input<Application>;
    readonly arn?: pulumi.Input<string>;
    readonly autoscalingGroups?: pulumi.Input<pulumi.Input<string>[]>;
    readonly cname?: pulumi.Input<string>;
    readonly cnamePrefix?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly instances?: pulumi.Input<pulumi.Input<string>[]>;
    readonly launchConfigurations?: pulumi.Input<pulumi.Input<string>[]>;
    readonly loadBalancers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly platformArn?: pulumi.Input<string>;
    readonly pollInterval?: pulumi.Input<string>;
    readonly queues?: pulumi.Input<pulumi.Input<string>[]>;
    readonly settings?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, namespace: pulumi.Input<string>, resource?: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
    readonly solutionStackName?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly templateName?: pulumi.Input<string>;
    readonly tier?: pulumi.Input<string>;
    readonly triggers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly version?: pulumi.Input<ApplicationVersion>;
    readonly waitForReadyTimeout?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    readonly application: pulumi.Input<Application>;
    readonly cnamePrefix?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly platformArn?: pulumi.Input<string>;
    readonly pollInterval?: pulumi.Input<string>;
    readonly settings?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, namespace: pulumi.Input<string>, resource?: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
    readonly solutionStackName?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly templateName?: pulumi.Input<string>;
    readonly tier?: pulumi.Input<string>;
    readonly version?: pulumi.Input<ApplicationVersion>;
    readonly waitForReadyTimeout?: pulumi.Input<string>;
}
