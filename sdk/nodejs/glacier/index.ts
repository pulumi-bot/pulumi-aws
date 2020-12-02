// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./vault";
export * from "./vaultLock";

// Import resources to register:
import { Vault } from "./vault";
import { VaultLock } from "./vaultLock";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:glacier/vault:Vault":
                return new Vault(name, <any>undefined, { urn })
            case "aws:glacier/vaultLock:VaultLock":
                return new VaultLock(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "glacier/vault", _module)
pulumi.runtime.registerResourceModule("aws", "glacier/vaultLock", _module)
