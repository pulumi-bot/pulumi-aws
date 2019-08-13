// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ChannelHlsIngest {
    /**
     * A list of the ingest endpoints
     */
    ingestEndpoints?: pulumi.Input<pulumi.Input<inputApi.mediapackage.ChannelHlsIngestIngestEndpoint>[]>;
}

export interface ChannelHlsIngestIngestEndpoint {
    /**
     * The password
     */
    password?: pulumi.Input<string>;
    /**
     * The URL
     */
    url?: pulumi.Input<string>;
    /**
     * The username
     */
    username?: pulumi.Input<string>;
}
