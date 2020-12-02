// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./broker";
export * from "./configuration";
export * from "./getBroker";

// Import resources to register:
import { Broker } from "./broker";
import { Configuration } from "./configuration";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:mq/broker:Broker":
                return new Broker(name, <any>undefined, { urn })
            case "aws:mq/configuration:Configuration":
                return new Configuration(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "mq/broker", _module)
pulumi.runtime.registerResourceModule("aws", "mq/configuration", _module)
