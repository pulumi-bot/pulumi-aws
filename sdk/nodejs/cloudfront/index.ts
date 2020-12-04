// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./distribution";
export * from "./getDistribution";
export * from "./originAccessIdentity";
export * from "./publicKey";

// Import resources to register:
import { Distribution } from "./distribution";
import { OriginAccessIdentity } from "./originAccessIdentity";
import { PublicKey } from "./publicKey";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:cloudfront/distribution:Distribution":
                return new Distribution(name, <any>undefined, { urn })
            case "aws:cloudfront/originAccessIdentity:OriginAccessIdentity":
                return new OriginAccessIdentity(name, <any>undefined, { urn })
            case "aws:cloudfront/publicKey:PublicKey":
                return new PublicKey(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "cloudfront/distribution", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/originAccessIdentity", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/publicKey", _module)
