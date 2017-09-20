// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class ReceiptRule extends fabric.Resource {
    public readonly addHeaderAction?: fabric.Computed<{ headerName: string, headerValue: string, position: number }[]>;
    public readonly after?: fabric.Computed<string>;
    public readonly bounceAction?: fabric.Computed<{ message: string, position: number, sender: string, smtpReplyCode: string, statusCode?: string, topicArn?: string }[]>;
    public readonly enabled: fabric.Computed<boolean>;
    public readonly lambdaAction?: fabric.Computed<{ functionArn: string, invocationType: string, position: number, topicArn?: string }[]>;
    public readonly name: fabric.Computed<string>;
    public readonly recipients?: fabric.Computed<string[]>;
    public readonly ruleSetName: fabric.Computed<string>;
    public readonly s3Action?: fabric.Computed<{ bucketName: string, kmsKeyArn?: string, objectKeyPrefix?: string, position: number, topicArn?: string }[]>;
    public readonly scanEnabled: fabric.Computed<boolean>;
    public readonly snsAction?: fabric.Computed<{ position: number, topicArn: string }[]>;
    public readonly stopAction?: fabric.Computed<{ position: number, scope: string, topicArn?: string }[]>;
    public readonly tlsPolicy: fabric.Computed<string>;
    public readonly workmailAction?: fabric.Computed<{ organizationArn: string, position: number, topicArn?: string }[]>;

    constructor(urnName: string, args: ReceiptRuleArgs, dependsOn?: fabric.Resource[]) {
        if (args.ruleSetName === undefined) {
            throw new Error("Missing required property 'ruleSetName'");
        }
        super("aws:ses/receiptRule:ReceiptRule", urnName, {
            "addHeaderAction": args.addHeaderAction,
            "after": args.after,
            "bounceAction": args.bounceAction,
            "enabled": args.enabled,
            "lambdaAction": args.lambdaAction,
            "name": args.name,
            "recipients": args.recipients,
            "ruleSetName": args.ruleSetName,
            "s3Action": args.s3Action,
            "scanEnabled": args.scanEnabled,
            "snsAction": args.snsAction,
            "stopAction": args.stopAction,
            "tlsPolicy": args.tlsPolicy,
            "workmailAction": args.workmailAction,
        }, dependsOn);
    }
}

export interface ReceiptRuleArgs {
    readonly addHeaderAction?: fabric.ComputedValue<{ headerName: fabric.ComputedValue<string>, headerValue: fabric.ComputedValue<string>, position: fabric.ComputedValue<number> }>[];
    readonly after?: fabric.ComputedValue<string>;
    readonly bounceAction?: fabric.ComputedValue<{ message: fabric.ComputedValue<string>, position: fabric.ComputedValue<number>, sender: fabric.ComputedValue<string>, smtpReplyCode: fabric.ComputedValue<string>, statusCode?: fabric.ComputedValue<string>, topicArn?: fabric.ComputedValue<string> }>[];
    readonly enabled?: fabric.ComputedValue<boolean>;
    readonly lambdaAction?: fabric.ComputedValue<{ functionArn: fabric.ComputedValue<string>, invocationType?: fabric.ComputedValue<string>, position: fabric.ComputedValue<number>, topicArn?: fabric.ComputedValue<string> }>[];
    readonly name?: fabric.ComputedValue<string>;
    readonly recipients?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly ruleSetName: fabric.ComputedValue<string>;
    readonly s3Action?: fabric.ComputedValue<{ bucketName: fabric.ComputedValue<string>, kmsKeyArn?: fabric.ComputedValue<string>, objectKeyPrefix?: fabric.ComputedValue<string>, position: fabric.ComputedValue<number>, topicArn?: fabric.ComputedValue<string> }>[];
    readonly scanEnabled?: fabric.ComputedValue<boolean>;
    readonly snsAction?: fabric.ComputedValue<{ position: fabric.ComputedValue<number>, topicArn: fabric.ComputedValue<string> }>[];
    readonly stopAction?: fabric.ComputedValue<{ position: fabric.ComputedValue<number>, scope: fabric.ComputedValue<string>, topicArn?: fabric.ComputedValue<string> }>[];
    readonly tlsPolicy?: fabric.ComputedValue<string>;
    readonly workmailAction?: fabric.ComputedValue<{ organizationArn: fabric.ComputedValue<string>, position: fabric.ComputedValue<number>, topicArn?: fabric.ComputedValue<string> }>[];
}

