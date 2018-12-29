// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Crawler extends pulumi.CustomResource {
    /**
     * Get an existing Crawler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CrawlerState, opts?: pulumi.CustomResourceOptions): Crawler {
        return new Crawler(name, <any>state, { ...opts, id: id });
    }

    public readonly classifiers: pulumi.Output<string[] | undefined>;
    public readonly configuration: pulumi.Output<string | undefined>;
    public readonly databaseName: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly dynamodbTargets: pulumi.Output<{ path: string }[] | undefined>;
    public readonly jdbcTargets: pulumi.Output<{ connectionName: string, exclusions?: string[], path: string }[] | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly role: pulumi.Output<string>;
    public readonly s3Targets: pulumi.Output<{ exclusions?: string[], path: string }[] | undefined>;
    public readonly schedule: pulumi.Output<string | undefined>;
    public readonly schemaChangePolicy: pulumi.Output<{ deleteBehavior?: string, updateBehavior?: string } | undefined>;
    public readonly securityConfiguration: pulumi.Output<string | undefined>;
    public readonly tablePrefix: pulumi.Output<string | undefined>;

    /**
     * Create a Crawler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CrawlerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CrawlerArgs | CrawlerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CrawlerState = argsOrState as CrawlerState | undefined;
            inputs["classifiers"] = state ? state.classifiers : undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dynamodbTargets"] = state ? state.dynamodbTargets : undefined;
            inputs["jdbcTargets"] = state ? state.jdbcTargets : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["s3Targets"] = state ? state.s3Targets : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["schemaChangePolicy"] = state ? state.schemaChangePolicy : undefined;
            inputs["securityConfiguration"] = state ? state.securityConfiguration : undefined;
            inputs["tablePrefix"] = state ? state.tablePrefix : undefined;
        } else {
            const args = argsOrState as CrawlerArgs | undefined;
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["classifiers"] = args ? args.classifiers : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dynamodbTargets"] = args ? args.dynamodbTargets : undefined;
            inputs["jdbcTargets"] = args ? args.jdbcTargets : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["s3Targets"] = args ? args.s3Targets : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["schemaChangePolicy"] = args ? args.schemaChangePolicy : undefined;
            inputs["securityConfiguration"] = args ? args.securityConfiguration : undefined;
            inputs["tablePrefix"] = args ? args.tablePrefix : undefined;
        }
        super("aws:glue/crawler:Crawler", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Crawler resources.
 */
export interface CrawlerState {
    readonly classifiers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly configuration?: pulumi.Input<string>;
    readonly databaseName?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly dynamodbTargets?: pulumi.Input<pulumi.Input<{ path: pulumi.Input<string> }>[]>;
    readonly jdbcTargets?: pulumi.Input<pulumi.Input<{ connectionName: pulumi.Input<string>, exclusions?: pulumi.Input<pulumi.Input<string>[]>, path: pulumi.Input<string> }>[]>;
    readonly name?: pulumi.Input<string>;
    readonly role?: pulumi.Input<string>;
    readonly s3Targets?: pulumi.Input<pulumi.Input<{ exclusions?: pulumi.Input<pulumi.Input<string>[]>, path: pulumi.Input<string> }>[]>;
    readonly schedule?: pulumi.Input<string>;
    readonly schemaChangePolicy?: pulumi.Input<{ deleteBehavior?: pulumi.Input<string>, updateBehavior?: pulumi.Input<string> }>;
    readonly securityConfiguration?: pulumi.Input<string>;
    readonly tablePrefix?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Crawler resource.
 */
export interface CrawlerArgs {
    readonly classifiers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly configuration?: pulumi.Input<string>;
    readonly databaseName: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly dynamodbTargets?: pulumi.Input<pulumi.Input<{ path: pulumi.Input<string> }>[]>;
    readonly jdbcTargets?: pulumi.Input<pulumi.Input<{ connectionName: pulumi.Input<string>, exclusions?: pulumi.Input<pulumi.Input<string>[]>, path: pulumi.Input<string> }>[]>;
    readonly name?: pulumi.Input<string>;
    readonly role: pulumi.Input<string>;
    readonly s3Targets?: pulumi.Input<pulumi.Input<{ exclusions?: pulumi.Input<pulumi.Input<string>[]>, path: pulumi.Input<string> }>[]>;
    readonly schedule?: pulumi.Input<string>;
    readonly schemaChangePolicy?: pulumi.Input<{ deleteBehavior?: pulumi.Input<string>, updateBehavior?: pulumi.Input<string> }>;
    readonly securityConfiguration?: pulumi.Input<string>;
    readonly tablePrefix?: pulumi.Input<string>;
}
