module talos-go-example

go 1.18

require (
        github.com/siderolabs/pulumi-provider-talos/sdk v0.0.0-20220714134323-6b2a85548bdd
        github.com/pulumi/pulumi/sdk/v3 v3.33.2
)

replace github.com/siderolabs/pulumi-provider-talos/sdk => ../../sdk