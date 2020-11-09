// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as alb from "./alb";
import * as applicationloadbalancing from "./applicationloadbalancing";
import * as autoscaling from "./autoscaling";
import * as ec2 from "./ec2";
import * as iam from "./iam";
import * as lambda from "./lambda";
import * as rds from "./rds";
import * as route53 from "./route53";
import * as s3 from "./s3";
import * as ssm from "./ssm";

export {
    alb,
    applicationloadbalancing,
    autoscaling,
    ec2,
    iam,
    lambda,
    rds,
    route53,
    s3,
    ssm,
};

export const Region = {
    AFSouth1: "af-south-1",
    APEast1: "ap-east-1",
    APNortheast1: "ap-northeast-1",
    APNortheast2: "ap-northeast-2",
    APSouth1: "ap-south-1",
    APSoutheast2: "ap-southeast-2",
    APSoutheast1: "ap-southeast-1",
    CACentral: "ca-central-1",
    CNNorth1: "cn-north-1",
    CNNorthwest1: "cn-northwest-1",
    EUCentral1: "eu-central-1",
    EUNorth1: "eu-north-1",
    EUWest1: "eu-west-1",
    EUWest2: "eu-west-2",
    EUWest3: "eu-west-3",
    EUSouth1: "eu-south-1",
    MESouth1: "me-south-1",
    SAEast1: "sa-east-1",
    USGovEast1: "us-gov-east-1",
    USGovWest1: "us-gov-west-1",
    USEast1: "us-east-1",
    USEast2: "us-east-2",
    USWest1: "us-west-1",
    USWest2: "us-west-2",
} as const;

/**
 * A Region represents any valid Amazon region that may be targeted with deployments.
 */
export type Region = (typeof Region)[keyof typeof Region];
