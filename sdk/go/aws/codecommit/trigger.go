// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codecommit

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CodeCommit Trigger Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codecommit"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testRepository, err := codecommit.NewRepository(ctx, "testRepository", &codecommit.RepositoryArgs{
// 			RepositoryName: pulumi.String("test"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codecommit.NewTrigger(ctx, "testTrigger", &codecommit.TriggerArgs{
// 			RepositoryName: testRepository.RepositoryName,
// 			Triggers: codecommit.TriggerTriggerArray{
// 				&codecommit.TriggerTriggerArgs{
// 					Name: pulumi.String("all"),
// 					Events: pulumi.StringArray{
// 						pulumi.String("all"),
// 					},
// 					DestinationArn: pulumi.Any(aws_sns_topic.Test.Arn),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Trigger struct {
	pulumi.CustomResourceState

	ConfigurationId pulumi.StringOutput `pulumi:"configurationId"`
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName pulumi.StringOutput       `pulumi:"repositoryName"`
	Triggers       TriggerTriggerArrayOutput `pulumi:"triggers"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.RepositoryName == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryName'")
	}
	if args.Triggers == nil {
		return nil, errors.New("invalid value for required argument 'Triggers'")
	}
	var resource Trigger
	err := ctx.RegisterResource("aws:codecommit/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("aws:codecommit/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	ConfigurationId *string `pulumi:"configurationId"`
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName *string          `pulumi:"repositoryName"`
	Triggers       []TriggerTrigger `pulumi:"triggers"`
}

type TriggerState struct {
	ConfigurationId pulumi.StringPtrInput
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName pulumi.StringPtrInput
	Triggers       TriggerTriggerArrayInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName string           `pulumi:"repositoryName"`
	Triggers       []TriggerTrigger `pulumi:"triggers"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName pulumi.StringInput
	Triggers       TriggerTriggerArrayInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}
