// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Classifier extends pulumi.CustomResource {
    /**
     * Get an existing Classifier resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClassifierState, opts?: pulumi.CustomResourceOptions): Classifier {
        return new Classifier(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/classifier:Classifier';

    /**
     * Returns true if the given object is an instance of Classifier.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Classifier {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Classifier.__pulumiType;
    }

    public readonly csvClassifier!: pulumi.Output<outputs.glue.ClassifierCsvClassifier | undefined>;
    public readonly grokClassifier!: pulumi.Output<outputs.glue.ClassifierGrokClassifier | undefined>;
    public readonly jsonClassifier!: pulumi.Output<outputs.glue.ClassifierJsonClassifier | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly xmlClassifier!: pulumi.Output<outputs.glue.ClassifierXmlClassifier | undefined>;

    /**
     * Create a Classifier resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClassifierArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClassifierArgs | ClassifierState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClassifierState | undefined;
            inputs["csvClassifier"] = state ? state.csvClassifier : undefined;
            inputs["grokClassifier"] = state ? state.grokClassifier : undefined;
            inputs["jsonClassifier"] = state ? state.jsonClassifier : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["xmlClassifier"] = state ? state.xmlClassifier : undefined;
        } else {
            const args = argsOrState as ClassifierArgs | undefined;
            inputs["csvClassifier"] = args ? args.csvClassifier : undefined;
            inputs["grokClassifier"] = args ? args.grokClassifier : undefined;
            inputs["jsonClassifier"] = args ? args.jsonClassifier : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["xmlClassifier"] = args ? args.xmlClassifier : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Classifier.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Classifier resources.
 */
export interface ClassifierState {
    readonly csvClassifier?: pulumi.Input<inputs.glue.ClassifierCsvClassifier>;
    readonly grokClassifier?: pulumi.Input<inputs.glue.ClassifierGrokClassifier>;
    readonly jsonClassifier?: pulumi.Input<inputs.glue.ClassifierJsonClassifier>;
    readonly name?: pulumi.Input<string>;
    readonly xmlClassifier?: pulumi.Input<inputs.glue.ClassifierXmlClassifier>;
}

/**
 * The set of arguments for constructing a Classifier resource.
 */
export interface ClassifierArgs {
    readonly csvClassifier?: pulumi.Input<inputs.glue.ClassifierCsvClassifier>;
    readonly grokClassifier?: pulumi.Input<inputs.glue.ClassifierGrokClassifier>;
    readonly jsonClassifier?: pulumi.Input<inputs.glue.ClassifierJsonClassifier>;
    readonly name?: pulumi.Input<string>;
    readonly xmlClassifier?: pulumi.Input<inputs.glue.ClassifierXmlClassifier>;
}
