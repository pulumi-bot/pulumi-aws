// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage AWS Device Farm Projects.
 * Please keep in mind that this feature is only supported on the "us-west-2" region.
 * This resource will error if you try to create a project in another region.
 *
 * For more information about Device Farm Projects, see the AWS Documentation on
 * [Device Farm Projects][aws-get-project].
 *
 * ## Basic Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const awesomeDevices = new aws.devicefarm.Project("awesome_devices", {});
 * ```
 *
 * ## Import
 *
 * DeviceFarm Projects can be imported by their arn
 *
 * ```sh
 *  $ pulumi import aws:devicefarm/project:Project example arn:aws:devicefarm:us-west-2:123456789012:project:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:devicefarm/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * The Amazon Resource Name of this project
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the project
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * The Amazon Resource Name of this project
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name of the project
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The name of the project
     */
    readonly name?: pulumi.Input<string>;
}
