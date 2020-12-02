// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getExport";
export * from "./getStack";
export * from "./stack";
export * from "./stackSet";
export * from "./stackSetInstance";

// Import resources to register:
import { Stack } from "./stack";
import { StackSet } from "./stackSet";
import { StackSetInstance } from "./stackSetInstance";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:cloudformation/stack:Stack":
                return new Stack(name, <any>undefined, { urn })
            case "aws:cloudformation/stackSet:StackSet":
                return new StackSet(name, <any>undefined, { urn })
            case "aws:cloudformation/stackSetInstance:StackSetInstance":
                return new StackSetInstance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "cloudformation/stack", _module)
pulumi.runtime.registerResourceModule("aws", "cloudformation/stackSet", _module)
pulumi.runtime.registerResourceModule("aws", "cloudformation/stackSetInstance", _module)
