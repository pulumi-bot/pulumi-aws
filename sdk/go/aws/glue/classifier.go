// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Glue Classifier resource.
//
// > **NOTE:** It is only valid to create one type of classifier (csv, grok, JSON, or XML). Changing classifier types will recreate the classifier.
//
// ## Example Usage
// ### Csv Classifier
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
// 			CsvClassifier: &glue.ClassifierCsvClassifierArgs{
// 				AllowSingleColumn:    pulumi.Bool(false),
// 				ContainsHeader:       pulumi.String("PRESENT"),
// 				Delimiter:            pulumi.String(","),
// 				DisableValueTrimming: pulumi.Bool(false),
// 				Headers: pulumi.StringArray{
// 					pulumi.String("example1"),
// 					pulumi.String("example2"),
// 				},
// 				QuoteSymbol: pulumi.String("'"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Grok Classifier
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
// 			GrokClassifier: &glue.ClassifierGrokClassifierArgs{
// 				Classification: pulumi.String("example"),
// 				GrokPattern:    pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### JSON Classifier
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
// 			JsonClassifier: &glue.ClassifierJsonClassifierArgs{
// 				JsonPath: pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### XML Classifier
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
// 			XmlClassifier: &glue.ClassifierXmlClassifierArgs{
// 				Classification: pulumi.String("example"),
// 				RowTag:         pulumi.String("example"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Glue Classifiers can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import aws:glue/classifier:Classifier MyClassifier MyClassifier
// ```
type Classifier struct {
	pulumi.CustomResourceState

	// A classifier for Csv content. Defined below.
	CsvClassifier ClassifierCsvClassifierPtrOutput `pulumi:"csvClassifier"`
	// A classifier that uses grok patterns. Defined below.
	GrokClassifier ClassifierGrokClassifierPtrOutput `pulumi:"grokClassifier"`
	// A classifier for JSON content. Defined below.
	JsonClassifier ClassifierJsonClassifierPtrOutput `pulumi:"jsonClassifier"`
	// The name of the classifier.
	Name pulumi.StringOutput `pulumi:"name"`
	// A classifier for XML content. Defined below.
	XmlClassifier ClassifierXmlClassifierPtrOutput `pulumi:"xmlClassifier"`
}

// NewClassifier registers a new resource with the given unique name, arguments, and options.
func NewClassifier(ctx *pulumi.Context,
	name string, args *ClassifierArgs, opts ...pulumi.ResourceOption) (*Classifier, error) {
	if args == nil {
		args = &ClassifierArgs{}
	}

	var resource Classifier
	err := ctx.RegisterResource("aws:glue/classifier:Classifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClassifier gets an existing Classifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClassifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClassifierState, opts ...pulumi.ResourceOption) (*Classifier, error) {
	var resource Classifier
	err := ctx.ReadResource("aws:glue/classifier:Classifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Classifier resources.
type classifierState struct {
	// A classifier for Csv content. Defined below.
	CsvClassifier *ClassifierCsvClassifier `pulumi:"csvClassifier"`
	// A classifier that uses grok patterns. Defined below.
	GrokClassifier *ClassifierGrokClassifier `pulumi:"grokClassifier"`
	// A classifier for JSON content. Defined below.
	JsonClassifier *ClassifierJsonClassifier `pulumi:"jsonClassifier"`
	// The name of the classifier.
	Name *string `pulumi:"name"`
	// A classifier for XML content. Defined below.
	XmlClassifier *ClassifierXmlClassifier `pulumi:"xmlClassifier"`
}

type ClassifierState struct {
	// A classifier for Csv content. Defined below.
	CsvClassifier ClassifierCsvClassifierPtrInput
	// A classifier that uses grok patterns. Defined below.
	GrokClassifier ClassifierGrokClassifierPtrInput
	// A classifier for JSON content. Defined below.
	JsonClassifier ClassifierJsonClassifierPtrInput
	// The name of the classifier.
	Name pulumi.StringPtrInput
	// A classifier for XML content. Defined below.
	XmlClassifier ClassifierXmlClassifierPtrInput
}

func (ClassifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*classifierState)(nil)).Elem()
}

type classifierArgs struct {
	// A classifier for Csv content. Defined below.
	CsvClassifier *ClassifierCsvClassifier `pulumi:"csvClassifier"`
	// A classifier that uses grok patterns. Defined below.
	GrokClassifier *ClassifierGrokClassifier `pulumi:"grokClassifier"`
	// A classifier for JSON content. Defined below.
	JsonClassifier *ClassifierJsonClassifier `pulumi:"jsonClassifier"`
	// The name of the classifier.
	Name *string `pulumi:"name"`
	// A classifier for XML content. Defined below.
	XmlClassifier *ClassifierXmlClassifier `pulumi:"xmlClassifier"`
}

// The set of arguments for constructing a Classifier resource.
type ClassifierArgs struct {
	// A classifier for Csv content. Defined below.
	CsvClassifier ClassifierCsvClassifierPtrInput
	// A classifier that uses grok patterns. Defined below.
	GrokClassifier ClassifierGrokClassifierPtrInput
	// A classifier for JSON content. Defined below.
	JsonClassifier ClassifierJsonClassifierPtrInput
	// The name of the classifier.
	Name pulumi.StringPtrInput
	// A classifier for XML content. Defined below.
	XmlClassifier ClassifierXmlClassifierPtrInput
}

func (ClassifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*classifierArgs)(nil)).Elem()
}

type ClassifierInput interface {
	pulumi.Input

	ToClassifierOutput() ClassifierOutput
	ToClassifierOutputWithContext(ctx context.Context) ClassifierOutput
}

type ClassifierPtrInput interface {
	pulumi.Input

	ToClassifierPtrOutput() ClassifierPtrOutput
	ToClassifierPtrOutputWithContext(ctx context.Context) ClassifierPtrOutput
}

func (Classifier) ElementType() reflect.Type {
	return reflect.TypeOf((*Classifier)(nil)).Elem()
}

func (i Classifier) ToClassifierOutput() ClassifierOutput {
	return i.ToClassifierOutputWithContext(context.Background())
}

func (i Classifier) ToClassifierOutputWithContext(ctx context.Context) ClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassifierOutput)
}

func (i Classifier) ToClassifierPtrOutput() ClassifierPtrOutput {
	return i.ToClassifierPtrOutputWithContext(context.Background())
}

func (i Classifier) ToClassifierPtrOutputWithContext(ctx context.Context) ClassifierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassifierPtrOutput)
}

type ClassifierOutput struct {
	*pulumi.OutputState
}

func (ClassifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassifierOutput)(nil)).Elem()
}

func (o ClassifierOutput) ToClassifierOutput() ClassifierOutput {
	return o
}

func (o ClassifierOutput) ToClassifierOutputWithContext(ctx context.Context) ClassifierOutput {
	return o
}

type ClassifierPtrOutput struct {
	*pulumi.OutputState
}

func (ClassifierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Classifier)(nil)).Elem()
}

func (o ClassifierPtrOutput) ToClassifierPtrOutput() ClassifierPtrOutput {
	return o
}

func (o ClassifierPtrOutput) ToClassifierPtrOutputWithContext(ctx context.Context) ClassifierPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClassifierOutput{})
	pulumi.RegisterOutputType(ClassifierPtrOutput{})
}
