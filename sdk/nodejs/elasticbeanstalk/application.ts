// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an Elastic Beanstalk Application Resource. Elastic Beanstalk allows
 * you to deploy and manage applications in the AWS cloud without worrying about
 * the infrastructure that runs those applications.
 *
 * This resource creates an application that has one configuration template named
 * `default`, and no application versions
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const tftest = new aws.elasticbeanstalk.Application("tftest", {
 *     description: "tf-test-desc",
 *     appversionLifecycle: {
 *         serviceRole: aws_iam_role.beanstalk_service.arn,
 *         maxCount: 128,
 *         deleteSourceFromS3: true,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Elastic Beanstalk Applications can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:elasticbeanstalk/application:Application tf_test tf-test-name
 * ```
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticbeanstalk/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    public readonly appversionLifecycle!: pulumi.Output<outputs.elasticbeanstalk.ApplicationAppversionLifecycle | undefined>;
    /**
     * The ARN assigned by AWS for this Elastic Beanstalk Application.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Short description of the application
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the application, must be unique within your account
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value map of tags for the Elastic Beanstalk Application.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            inputs["appversionLifecycle"] = state ? state.appversionLifecycle : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            inputs["appversionLifecycle"] = args ? args.appversionLifecycle : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    readonly appversionLifecycle?: pulumi.Input<inputs.elasticbeanstalk.ApplicationAppversionLifecycle>;
    /**
     * The ARN assigned by AWS for this Elastic Beanstalk Application.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Short description of the application
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the application, must be unique within your account
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value map of tags for the Elastic Beanstalk Application.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    readonly appversionLifecycle?: pulumi.Input<inputs.elasticbeanstalk.ApplicationAppversionLifecycle>;
    /**
     * Short description of the application
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the application, must be unique within your account
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value map of tags for the Elastic Beanstalk Application.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
