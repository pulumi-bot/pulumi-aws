module github.com/pulumi/pulumi-aws/provider/v3

go 1.15

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/hashicorp/aws-sdk-go-base v0.7.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.18.0
	github.com/pulumi/pulumi/pkg/v2 v2.18.1-0.20210121021731-8f7dfbb5ae22
	github.com/pulumi/pulumi/sdk/v2 v2.18.1-0.20210121021731-8f7dfbb5ae22
	github.com/terraform-providers/terraform-provider-aws v0.0.0-20191010190908-1261a98537f2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-aws => github.com/pulumi/terraform-provider-aws v1.38.1-0.20210122131940-c9982cab7fa0
)

replace github.com/pulumi/pulumi-terraform-bridge/v2 => ../../pulumi-terraform-bridge
