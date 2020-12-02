// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./alias";
export * from "./build";
export * from "./fleet";
export * from "./gameSessionQueue";

// Import resources to register:
import { Alias } from "./alias";
import { Build } from "./build";
import { Fleet } from "./fleet";
import { GameSessionQueue } from "./gameSessionQueue";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:gamelift/alias:Alias":
                return new Alias(name, <any>undefined, { urn })
            case "aws:gamelift/build:Build":
                return new Build(name, <any>undefined, { urn })
            case "aws:gamelift/fleet:Fleet":
                return new Fleet(name, <any>undefined, { urn })
            case "aws:gamelift/gameSessionQueue:GameSessionQueue":
                return new GameSessionQueue(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "gamelift/alias", _module)
pulumi.runtime.registerResourceModule("aws", "gamelift/build", _module)
pulumi.runtime.registerResourceModule("aws", "gamelift/fleet", _module)
pulumi.runtime.registerResourceModule("aws", "gamelift/gameSessionQueue", _module)
