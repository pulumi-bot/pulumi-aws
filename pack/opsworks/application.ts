// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Application extends fabric.Resource {
    public readonly appSource: fabric.Computed<{ password?: string, revision?: string, sshKey?: string, type: string, url?: string, username?: string }[]>;
    public readonly autoBundleOnDeploy?: fabric.Computed<string>;
    public readonly awsFlowRubySettings?: fabric.Computed<string>;
    public readonly dataSourceArn?: fabric.Computed<string>;
    public readonly dataSourceDatabaseName?: fabric.Computed<string>;
    public readonly dataSourceType?: fabric.Computed<string>;
    public readonly description?: fabric.Computed<string>;
    public readonly documentRoot?: fabric.Computed<string>;
    public readonly domains?: fabric.Computed<string[]>;
    public readonly enableSsl?: fabric.Computed<boolean>;
    public readonly environment?: fabric.Computed<{ key: string, secure?: boolean, value: string }[]>;
    public /*out*/ readonly applicationId: fabric.Computed<string>;
    public readonly name: fabric.Computed<string>;
    public readonly railsEnv?: fabric.Computed<string>;
    public readonly shortName: fabric.Computed<string>;
    public readonly sslConfiguration?: fabric.Computed<{ certificate: string, chain?: string, privateKey: string }[]>;
    public readonly stackId: fabric.Computed<string>;
    public readonly type: fabric.Computed<string>;

    constructor(urnName: string, args: ApplicationArgs, dependsOn?: fabric.Resource[]) {
        if (args.stackId === undefined) {
            throw new Error("Missing required property 'stackId'");
        }
        if (args.type === undefined) {
            throw new Error("Missing required property 'type'");
        }
        super("aws:opsworks/application:Application", urnName, {
            "appSource": args.appSource,
            "autoBundleOnDeploy": args.autoBundleOnDeploy,
            "awsFlowRubySettings": args.awsFlowRubySettings,
            "dataSourceArn": args.dataSourceArn,
            "dataSourceDatabaseName": args.dataSourceDatabaseName,
            "dataSourceType": args.dataSourceType,
            "description": args.description,
            "documentRoot": args.documentRoot,
            "domains": args.domains,
            "enableSsl": args.enableSsl,
            "environment": args.environment,
            "name": args.name,
            "railsEnv": args.railsEnv,
            "shortName": args.shortName,
            "sslConfiguration": args.sslConfiguration,
            "stackId": args.stackId,
            "type": args.type,
            "applicationId": undefined,
        }, dependsOn);
    }
}

export interface ApplicationArgs {
    readonly appSource?: fabric.ComputedValue<{ password?: fabric.ComputedValue<string>, revision?: fabric.ComputedValue<string>, sshKey?: fabric.ComputedValue<string>, type: fabric.ComputedValue<string>, url?: fabric.ComputedValue<string>, username?: fabric.ComputedValue<string> }>[];
    readonly autoBundleOnDeploy?: fabric.ComputedValue<string>;
    readonly awsFlowRubySettings?: fabric.ComputedValue<string>;
    readonly dataSourceArn?: fabric.ComputedValue<string>;
    readonly dataSourceDatabaseName?: fabric.ComputedValue<string>;
    readonly dataSourceType?: fabric.ComputedValue<string>;
    readonly description?: fabric.ComputedValue<string>;
    readonly documentRoot?: fabric.ComputedValue<string>;
    readonly domains?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly enableSsl?: fabric.ComputedValue<boolean>;
    readonly environment?: fabric.ComputedValue<{ key: fabric.ComputedValue<string>, secure?: fabric.ComputedValue<boolean>, value: fabric.ComputedValue<string> }>[];
    readonly name?: fabric.ComputedValue<string>;
    readonly railsEnv?: fabric.ComputedValue<string>;
    readonly shortName?: fabric.ComputedValue<string>;
    readonly sslConfiguration?: fabric.ComputedValue<{ certificate: fabric.ComputedValue<string>, chain?: fabric.ComputedValue<string>, privateKey: fabric.ComputedValue<string> }>[];
    readonly stackId: fabric.ComputedValue<string>;
    readonly type: fabric.ComputedValue<string>;
}

