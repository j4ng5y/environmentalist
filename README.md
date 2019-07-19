# Environmentalist

`[tree_hugging_gopher.jpg]`

## Usage

```
Environmentalist is an application that provides a consistent API for using a number of secrets management tools including:
  * Hashicorp Vault
  * AWS SSM
  * Ansible Vault
        etc...

The Server runs as both a RESTful service as well as a gRPC service so it should be usable for almost any situation.

A RESTful request to access a secret looks something like this:
  curl -X GET https://environmentalist:5005/hashicorp-vault/get/mySharedSecret

A RESTful request to store a new secret looks something like this:
  curl -X POST -H "Content-Type: application/json" -d '{"mySharedSecret": "thisIsASuperSecretPassword"}' https://environmentalist:5005/hashicorp-vault/new/mySharedSecret

A RESTful request to delete a secret looks something like this:
  curl -X DELETE https://environmentalist:5005/hashicorp-vault/delete/mySharedSecret

A RESTful request to modify a secret looks something like this:
  curl -X PUT -H "Content-Type: application/json" -d '{"mySharedSecret": "thisIsANewSuperSecretPassword"}' https://environmentalist:5005/hashicorp-vault/update/mySharedSecret

Please see https://github.com/j4ng5y/envrionmentalist for a full API breakdown.

Usage:
  environmentalist [flags]
  environmentalist [command]

Available Commands:
  help        Help about any command
  run         run the environmentalist daemon
  stop        stop the envrionmentalist daemon

Flags:
  -s, --aws-ssm                            the aws-ssm flag tells envrionmentalist that we want to use aws ssm
      --aws-ssm-access-key-id string       the aws-ssm-access-key-id flag tells envrionmentalist what Access Key to use to connect to AWS with
      --aws-ssm-credential-type string     the aws-ssm-credential-type flag tells environmentalist what type of credentials to look for to access AWS. (Options:"profile", "manual", "role" (default "profile")
      --aws-ssm-profile-name string        the aws-ssm-profile-name flag tells envrionmentalist what AWS profile to connect to AWS with (default "Default")
      --aws-ssm-region string              the aws-ssm-region flag tells envrionmentalist what AWS region to connect to (default "us-east-1")
      --aws-ssm-secret-access-key string   the aws-ssm-secret-access-key flag tells envrionmentalist what Secret Key to use to connect to AWS with
  -v, --hashicorp-vault                    the hashicorp-vault flag tells environmentalist that we want to use the hashicorp vault
      --hashicorp-vault-auth-type string   the hashicorp-vault-auth-type flag tells envrionmentalist what authentication type to use to log into the hashi-corp vault (default "approle")
  -h, --help                               help for environmentalist
      --version                            version for environmentalist

Use "environmentalist [command] --help" for more information about a command.
```