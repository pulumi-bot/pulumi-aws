// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an Amazon Lex Slot Type resource. For more information see
 * [Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const flowerTypes = new aws.lex.SlotType("flower_types", {
 *     createVersion: true,
 *     description: "Types of flowers to order",
 *     enumerationValues: [
 *         {
 *             synonyms: [
 *                 "Lirium",
 *                 "Martagon",
 *             ],
 *             value: "lilies",
 *         },
 *         {
 *             synonyms: [
 *                 "Eduardoregelia",
 *                 "Podonix",
 *             ],
 *             value: "tulips",
 *         },
 *     ],
 *     name: "FlowerTypes",
 *     valueSelectionStrategy: "ORIGINAL_VALUE",
 * });
 * ```
 */
export class SlotType extends pulumi.CustomResource {
    /**
     * Get an existing SlotType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SlotTypeState, opts?: pulumi.CustomResourceOptions): SlotType {
        return new SlotType(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lex/slotType:SlotType';

    /**
     * Returns true if the given object is an instance of SlotType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SlotType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SlotType.__pulumiType;
    }

    /**
     * Checksum identifying the version of the slot type that was created. The checksum is
     * not included as an argument because the resource will add it automatically when updating the slot type.
     */
    public /*out*/ readonly checksum!: pulumi.Output<string>;
    /**
     * Determines if a new slot type version is created when the initial resource is created and on each
     * update. Defaults to `false`.
     */
    public readonly createVersion!: pulumi.Output<boolean | undefined>;
    /**
     * The date when the slot type version was created.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * A description of the slot type. Must be less than or equal to 200 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of EnumerationValue objects that defines the values that
     * the slot type can take. Each value can have a list of synonyms, which are additional values that help
     * train the machine learning model about the values that it resolves for a slot. Attributes are
     * documented under enumeration_value.
     */
    public readonly enumerationValues!: pulumi.Output<outputs.lex.SlotTypeEnumerationValue[]>;
    /**
     * The date when the `$LATEST` version of this slot type was updated.
     */
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    /**
     * The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Determines the slot resolution strategy that Amazon Lex
     * uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
     * value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
     * if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
     */
    public readonly valueSelectionStrategy!: pulumi.Output<string | undefined>;
    /**
     * The version of the slot type.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a SlotType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SlotTypeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SlotTypeArgs | SlotTypeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SlotTypeState | undefined;
            inputs["checksum"] = state ? state.checksum : undefined;
            inputs["createVersion"] = state ? state.createVersion : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enumerationValues"] = state ? state.enumerationValues : undefined;
            inputs["lastUpdatedDate"] = state ? state.lastUpdatedDate : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["valueSelectionStrategy"] = state ? state.valueSelectionStrategy : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as SlotTypeArgs | undefined;
            if (!args || args.enumerationValues === undefined) {
                throw new Error("Missing required property 'enumerationValues'");
            }
            inputs["createVersion"] = args ? args.createVersion : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enumerationValues"] = args ? args.enumerationValues : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["valueSelectionStrategy"] = args ? args.valueSelectionStrategy : undefined;
            inputs["checksum"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["lastUpdatedDate"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SlotType.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SlotType resources.
 */
export interface SlotTypeState {
    /**
     * Checksum identifying the version of the slot type that was created. The checksum is
     * not included as an argument because the resource will add it automatically when updating the slot type.
     */
    readonly checksum?: pulumi.Input<string>;
    /**
     * Determines if a new slot type version is created when the initial resource is created and on each
     * update. Defaults to `false`.
     */
    readonly createVersion?: pulumi.Input<boolean>;
    /**
     * The date when the slot type version was created.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * A description of the slot type. Must be less than or equal to 200 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of EnumerationValue objects that defines the values that
     * the slot type can take. Each value can have a list of synonyms, which are additional values that help
     * train the machine learning model about the values that it resolves for a slot. Attributes are
     * documented under enumeration_value.
     */
    readonly enumerationValues?: pulumi.Input<pulumi.Input<inputs.lex.SlotTypeEnumerationValue>[]>;
    /**
     * The date when the `$LATEST` version of this slot type was updated.
     */
    readonly lastUpdatedDate?: pulumi.Input<string>;
    /**
     * The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Determines the slot resolution strategy that Amazon Lex
     * uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
     * value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
     * if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
     */
    readonly valueSelectionStrategy?: pulumi.Input<string>;
    /**
     * The version of the slot type.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SlotType resource.
 */
export interface SlotTypeArgs {
    /**
     * Determines if a new slot type version is created when the initial resource is created and on each
     * update. Defaults to `false`.
     */
    readonly createVersion?: pulumi.Input<boolean>;
    /**
     * A description of the slot type. Must be less than or equal to 200 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of EnumerationValue objects that defines the values that
     * the slot type can take. Each value can have a list of synonyms, which are additional values that help
     * train the machine learning model about the values that it resolves for a slot. Attributes are
     * documented under enumeration_value.
     */
    readonly enumerationValues: pulumi.Input<pulumi.Input<inputs.lex.SlotTypeEnumerationValue>[]>;
    /**
     * The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Determines the slot resolution strategy that Amazon Lex
     * uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
     * value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
     * if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
     */
    readonly valueSelectionStrategy?: pulumi.Input<string>;
}
