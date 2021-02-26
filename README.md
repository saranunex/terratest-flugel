# Create AWS S3 Bucket and object + Terratest

Set up your GitHub Actions workflow with a specific version of [Terraform](https://www.terraform.io/).

## Usage

For this example we are gonna to use Terraform 0.14.7and go 1.15 . Make sure you have installed them correctly.

You can fork this repo to your local computer, setup your AWS Access ID and Secret Key env using AWS CLI, or you can fork this repo and configure/create/modify your own pipeline using Github Actions and Secrets.

To run it in you personal computer, you can do this:
```
# aws configure

# git clone <repo>

# cd test

# go test -v -tags=integration
```


### Outputs

| Parameter | Description |
| --------- | ----------- |
| `bucket_id` | checks if the given S3 bucket exists in the given region and fail the test if it does not. |

### Supported platforms

This action has been tested on the following platforms:

* ubuntu-18.04.5


## Contributing

Contributions to this repository are very welcome! We follow a fairly standard [pull request process](
https://help.github.com/articles/about-pull-requests/) for contributions, subject to the following guidelines:

1. File a GitHub issue
1. Fork the repository
1. Update the documentation
1. Update the tests
1. Update the code
1. Create a pull request
1. (Merge and release)

The maintainers for this repo will review your code and provide feedback. If everything looks good, they will merge the
code and release a new version, which you'll be able to find in the [releases page](../../releases).
