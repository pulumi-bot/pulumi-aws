// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an OpsWorks application resource.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 * 
 * const foo_app = new aws.opsworks.Application("foo-app", {
 *     appSources: [{
 *         revision: "master",
 *         type: "git",
 *         url: "https://github.com/example.git",
 *     }],
 *     autoBundleOnDeploy: "true",
 *     description: "This is a Rails application",
 *     documentRoot: "public",
 *     domains: [
 *         "example.com",
 *         "sub.example.com",
 *     ],
 *     enableSsl: true,
 *     environments: [{
 *         key: "key",
 *         secure: false,
 *         value: "value",
 *     }],
 *     railsEnv: "staging",
 *     shortName: "foobar",
 *     sslConfigurations: [{
 *         certificate: fs.readFileSync("./foobar.crt", "utf-8"),
 *         privateKey: fs.readFileSync("./foobar.key", "utf-8"),
 *     }],
 *     stackId: aws_opsworks_stack_main.id,
 *     type: "rails",
 * });
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
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /**
     * SCM configuration of the app as described below.
     */
    public readonly appSources!: pulumi.Output<{ password?: string, revision?: string, sshKey?: string, type: string, url?: string, username?: string }[]>;
    /**
     * Run bundle install when deploying for application of type `rails`.
     */
    public readonly autoBundleOnDeploy!: pulumi.Output<string | undefined>;
    /**
     * Specify activity and workflow workers for your app using the aws-flow gem.
     */
    public readonly awsFlowRubySettings!: pulumi.Output<string | undefined>;
    /**
     * The data source's ARN.
     */
    public readonly dataSourceArn!: pulumi.Output<string | undefined>;
    /**
     * The database name.
     */
    public readonly dataSourceDatabaseName!: pulumi.Output<string | undefined>;
    /**
     * The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
     */
    public readonly dataSourceType!: pulumi.Output<string | undefined>;
    /**
     * A description of the app.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Subfolder for the document root for application of type `rails`.
     */
    public readonly documentRoot!: pulumi.Output<string | undefined>;
    /**
     * A list of virtual host alias.
     */
    public readonly domains!: pulumi.Output<string[] | undefined>;
    /**
     * Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
     */
    public readonly enableSsl!: pulumi.Output<boolean | undefined>;
    /**
     * Object to define environment variables.  Object is described below.
     */
    public readonly environments!: pulumi.Output<{ key: string, secure?: boolean, value: string }[] | undefined>;
    /**
     * A human-readable name for the application.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Rails environment for application of type `rails`.
     */
    public readonly railsEnv!: pulumi.Output<string | undefined>;
    /**
     * A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
     */
    public readonly shortName!: pulumi.Output<string>;
    /**
     * The SSL configuration of the app. Object is described below.
     */
    public readonly sslConfigurations!: pulumi.Output<{ certificate: string, chain?: string, privateKey: string }[] | undefined>;
    /**
     * The id of the stack the application will belong to.
     */
    public readonly stackId!: pulumi.Output<string>;
    /**
     * The type of source to use. For example, "archive".
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            inputs["appSources"] = state ? state.appSources : undefined;
            inputs["autoBundleOnDeploy"] = state ? state.autoBundleOnDeploy : undefined;
            inputs["awsFlowRubySettings"] = state ? state.awsFlowRubySettings : undefined;
            inputs["dataSourceArn"] = state ? state.dataSourceArn : undefined;
            inputs["dataSourceDatabaseName"] = state ? state.dataSourceDatabaseName : undefined;
            inputs["dataSourceType"] = state ? state.dataSourceType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["documentRoot"] = state ? state.documentRoot : undefined;
            inputs["domains"] = state ? state.domains : undefined;
            inputs["enableSsl"] = state ? state.enableSsl : undefined;
            inputs["environments"] = state ? state.environments : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["railsEnv"] = state ? state.railsEnv : undefined;
            inputs["shortName"] = state ? state.shortName : undefined;
            inputs["sslConfigurations"] = state ? state.sslConfigurations : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if (!args || args.stackId === undefined) {
                throw new Error("Missing required property 'stackId'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["appSources"] = args ? args.appSources : undefined;
            inputs["autoBundleOnDeploy"] = args ? args.autoBundleOnDeploy : undefined;
            inputs["awsFlowRubySettings"] = args ? args.awsFlowRubySettings : undefined;
            inputs["dataSourceArn"] = args ? args.dataSourceArn : undefined;
            inputs["dataSourceDatabaseName"] = args ? args.dataSourceDatabaseName : undefined;
            inputs["dataSourceType"] = args ? args.dataSourceType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["documentRoot"] = args ? args.documentRoot : undefined;
            inputs["domains"] = args ? args.domains : undefined;
            inputs["enableSsl"] = args ? args.enableSsl : undefined;
            inputs["environments"] = args ? args.environments : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["railsEnv"] = args ? args.railsEnv : undefined;
            inputs["shortName"] = args ? args.shortName : undefined;
            inputs["sslConfigurations"] = args ? args.sslConfigurations : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        super("aws:opsworks/application:Application", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * SCM configuration of the app as described below.
     */
    readonly appSources?: pulumi.Input<pulumi.Input<{ password?: pulumi.Input<string>, revision?: pulumi.Input<string>, sshKey?: pulumi.Input<string>, type: pulumi.Input<string>, url?: pulumi.Input<string>, username?: pulumi.Input<string> }>[]>;
    /**
     * Run bundle install when deploying for application of type `rails`.
     */
    readonly autoBundleOnDeploy?: pulumi.Input<string>;
    /**
     * Specify activity and workflow workers for your app using the aws-flow gem.
     */
    readonly awsFlowRubySettings?: pulumi.Input<string>;
    /**
     * The data source's ARN.
     */
    readonly dataSourceArn?: pulumi.Input<string>;
    /**
     * The database name.
     */
    readonly dataSourceDatabaseName?: pulumi.Input<string>;
    /**
     * The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
     */
    readonly dataSourceType?: pulumi.Input<string>;
    /**
     * A description of the app.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Subfolder for the document root for application of type `rails`.
     */
    readonly documentRoot?: pulumi.Input<string>;
    /**
     * A list of virtual host alias.
     */
    readonly domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
     */
    readonly enableSsl?: pulumi.Input<boolean>;
    /**
     * Object to define environment variables.  Object is described below.
     */
    readonly environments?: pulumi.Input<pulumi.Input<{ key: pulumi.Input<string>, secure?: pulumi.Input<boolean>, value: pulumi.Input<string> }>[]>;
    /**
     * A human-readable name for the application.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Rails environment for application of type `rails`.
     */
    readonly railsEnv?: pulumi.Input<string>;
    /**
     * A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
     */
    readonly shortName?: pulumi.Input<string>;
    /**
     * The SSL configuration of the app. Object is described below.
     */
    readonly sslConfigurations?: pulumi.Input<pulumi.Input<{ certificate: pulumi.Input<string>, chain?: pulumi.Input<string>, privateKey: pulumi.Input<string> }>[]>;
    /**
     * The id of the stack the application will belong to.
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * The type of source to use. For example, "archive".
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * SCM configuration of the app as described below.
     */
    readonly appSources?: pulumi.Input<pulumi.Input<{ password?: pulumi.Input<string>, revision?: pulumi.Input<string>, sshKey?: pulumi.Input<string>, type: pulumi.Input<string>, url?: pulumi.Input<string>, username?: pulumi.Input<string> }>[]>;
    /**
     * Run bundle install when deploying for application of type `rails`.
     */
    readonly autoBundleOnDeploy?: pulumi.Input<string>;
    /**
     * Specify activity and workflow workers for your app using the aws-flow gem.
     */
    readonly awsFlowRubySettings?: pulumi.Input<string>;
    /**
     * The data source's ARN.
     */
    readonly dataSourceArn?: pulumi.Input<string>;
    /**
     * The database name.
     */
    readonly dataSourceDatabaseName?: pulumi.Input<string>;
    /**
     * The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
     */
    readonly dataSourceType?: pulumi.Input<string>;
    /**
     * A description of the app.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Subfolder for the document root for application of type `rails`.
     */
    readonly documentRoot?: pulumi.Input<string>;
    /**
     * A list of virtual host alias.
     */
    readonly domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
     */
    readonly enableSsl?: pulumi.Input<boolean>;
    /**
     * Object to define environment variables.  Object is described below.
     */
    readonly environments?: pulumi.Input<pulumi.Input<{ key: pulumi.Input<string>, secure?: pulumi.Input<boolean>, value: pulumi.Input<string> }>[]>;
    /**
     * A human-readable name for the application.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Rails environment for application of type `rails`.
     */
    readonly railsEnv?: pulumi.Input<string>;
    /**
     * A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
     */
    readonly shortName?: pulumi.Input<string>;
    /**
     * The SSL configuration of the app. Object is described below.
     */
    readonly sslConfigurations?: pulumi.Input<pulumi.Input<{ certificate: pulumi.Input<string>, chain?: pulumi.Input<string>, privateKey: pulumi.Input<string> }>[]>;
    /**
     * The id of the stack the application will belong to.
     */
    readonly stackId: pulumi.Input<string>;
    /**
     * The type of source to use. For example, "archive".
     */
    readonly type: pulumi.Input<string>;
}
