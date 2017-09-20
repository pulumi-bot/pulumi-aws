// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class IdentityPool extends fabric.Resource {
    public readonly allowUnauthenticatedIdentities?: fabric.Computed<boolean>;
    public readonly cognitoIdentityProviders?: fabric.Computed<{ clientId?: string, providerName?: string, serverSideTokenCheck?: boolean }[]>;
    public readonly developerProviderName?: fabric.Computed<string>;
    public readonly identityPoolName: fabric.Computed<string>;
    public readonly openidConnectProviderArns?: fabric.Computed<string[]>;
    public readonly samlProviderArns?: fabric.Computed<string[]>;
    public readonly supportedLoginProviders?: fabric.Computed<{[key: string]: string}>;

    constructor(urnName: string, args: IdentityPoolArgs, dependsOn?: fabric.Resource[]) {
        if (args.identityPoolName === undefined) {
            throw new Error("Missing required property 'identityPoolName'");
        }
        super("aws:cognito/identityPool:IdentityPool", urnName, {
            "allowUnauthenticatedIdentities": args.allowUnauthenticatedIdentities,
            "cognitoIdentityProviders": args.cognitoIdentityProviders,
            "developerProviderName": args.developerProviderName,
            "identityPoolName": args.identityPoolName,
            "openidConnectProviderArns": args.openidConnectProviderArns,
            "samlProviderArns": args.samlProviderArns,
            "supportedLoginProviders": args.supportedLoginProviders,
        }, dependsOn);
    }
}

export interface IdentityPoolArgs {
    readonly allowUnauthenticatedIdentities?: fabric.ComputedValue<boolean>;
    readonly cognitoIdentityProviders?: fabric.ComputedValue<{ clientId?: fabric.ComputedValue<string>, providerName?: fabric.ComputedValue<string>, serverSideTokenCheck?: fabric.ComputedValue<boolean> }>[];
    readonly developerProviderName?: fabric.ComputedValue<string>;
    readonly identityPoolName: fabric.ComputedValue<string>;
    readonly openidConnectProviderArns?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly samlProviderArns?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly supportedLoginProviders?: fabric.ComputedValue<{[key: string]: fabric.ComputedValue<string>}>;
}

