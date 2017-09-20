// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class DefaultSecurityGroup extends fabric.Resource {
    public readonly egress?: fabric.Computed<{ cidrBlocks?: string[], fromPort: number, ipv6CidrBlocks?: string[], prefixListIds?: string[], protocol: string, securityGroups?: string[], self?: boolean, toPort: number }[]>;
    public readonly ingress?: fabric.Computed<{ cidrBlocks?: string[], fromPort: number, ipv6CidrBlocks?: string[], protocol: string, securityGroups?: string[], self?: boolean, toPort: number }[]>;
    public /*out*/ readonly name: fabric.Computed<string>;
    public /*out*/ readonly ownerId: fabric.Computed<string>;
    public readonly tags?: fabric.Computed<{[key: string]: any}>;
    public readonly vpcId: fabric.Computed<string>;

    constructor(urnName: string, args?: DefaultSecurityGroupArgs, dependsOn?: fabric.Resource[]) {
        super("aws:ec2/defaultSecurityGroup:DefaultSecurityGroup", urnName, {
            "egress": args.egress,
            "ingress": args.ingress,
            "tags": args.tags,
            "vpcId": args.vpcId,
            "name": undefined,
            "ownerId": undefined,
        }, dependsOn);
    }
}

export interface DefaultSecurityGroupArgs {
    readonly egress?: fabric.ComputedValue<{ cidrBlocks?: fabric.ComputedValue<fabric.ComputedValue<string>>[], fromPort: fabric.ComputedValue<number>, ipv6CidrBlocks?: fabric.ComputedValue<fabric.ComputedValue<string>>[], prefixListIds?: fabric.ComputedValue<fabric.ComputedValue<string>>[], protocol: fabric.ComputedValue<string>, securityGroups?: fabric.ComputedValue<fabric.ComputedValue<string>>[], self?: fabric.ComputedValue<boolean>, toPort: fabric.ComputedValue<number> }>[];
    readonly ingress?: fabric.ComputedValue<{ cidrBlocks?: fabric.ComputedValue<fabric.ComputedValue<string>>[], fromPort: fabric.ComputedValue<number>, ipv6CidrBlocks?: fabric.ComputedValue<fabric.ComputedValue<string>>[], protocol: fabric.ComputedValue<string>, securityGroups?: fabric.ComputedValue<fabric.ComputedValue<string>>[], self?: fabric.ComputedValue<boolean>, toPort: fabric.ComputedValue<number> }>[];
    readonly tags?: fabric.ComputedValue<{[key: string]: any}>;
    readonly vpcId?: fabric.ComputedValue<string>;
}

