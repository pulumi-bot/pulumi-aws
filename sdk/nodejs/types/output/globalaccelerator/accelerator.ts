// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface AcceleratorAttributes {
    /**
     * Indicates whether flow logs are enabled.
     */
    flowLogsEnabled?: boolean;
    /**
     * The name of the Amazon S3 bucket for the flow logs.
     */
    flowLogsS3Bucket?: string;
    /**
     * The prefix for the location in the Amazon S3 bucket for the flow logs.
     */
    flowLogsS3Prefix?: string;
}

export interface AcceleratorIpSet {
    /**
     * The array of IP addresses in the IP address set.
     */
    ipAddresses: string[];
    /**
     * The types of IP addresses included in this IP set.
     */
    ipFamily: string;
}
