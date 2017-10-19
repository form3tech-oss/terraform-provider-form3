[![Build Status](https://travis-ci.org/ewilde/terraform-provider-form3.svg?branch=master)](https://travis-ci.org/ewilde/terraform-provider-form3)
# Terraform Form3 Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

The Form3 provider is used to create and manage Form3 resources using
the official [Form3 API](http://api-docs.form3.tech)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/ewilde/terraform-provider-form3`

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/ewilde/terraform-provider-form3
$ make build
```

## Using the provider

See [examples](examples/)

See [form3 providers documentation](website/docs)

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.9+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-form3
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ env FORM3_CLIENT_ID={your_client_id} FORM3_CLIENT_SECRET={your_client_secret} FORM3_HOST={some_env_host} FORM3_ORGANISATION_ID={your_organisation_id} make testacc
```

### Updating packages
* govendor fetch github.com/ewilde/go-form3/...