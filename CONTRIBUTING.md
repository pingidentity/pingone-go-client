# Contributing to the PingOne Go SDK

We appreciate your help!  We welcome contributions in the form of creating issues or pull requests.

Know that:

1. If you have any questions, please ask!  We'll help as best we can
2. While we appreciate perfect PRs, it's not essential. We'll fix up any housekeeping changes before merge.  Any PRs that need further work, we'll point you in the right direction or can take on ourselves.
3. We may not be able to respond quickly, our development cycles are on a priority basis.
4. We base our priorities on customer need and the number of votes on issues/PRs by the number of 👍 reactions.  If there is an existing issue or PR for something you'd like, please vote!
5. Some files are created and maintained by an internal generator project. These files are marked with comments at the top of the file. We may not be able to merge changes to these files as changes will be overwritten by the generator.

The following guides are there to help you get started with SDK development, logging issues and creating PRs.  Following the guides allows us to respond quicker and will result in faster PR merges.

## Code Generation and Contribution Scope

**Important:** The files in the `pingone` directory are automatically generated using [OpenAPI Generator](https://openapi-generator.tech/) from specifications found in the [PingOne OpenAPI Specifications](https://github.com/pingidentity/pingone-openapi-specifications) repository.

* ✅ **Welcome contributions:** All files outside the `pingone` directory
* ❌ **Will not merge:** Direct changes to files in the `pingone` directory (they will be overwritten)

You can still modify files in the `pingone` directory to test ideas or highlight issues, but these changes won't be merged to main. Instead, use them to demonstrate potential improvements that can be incorporated into the OpenAPI generation process.

## Beta Development

This SDK supports beta functionality for non-GA PingOne platform features. Beta releases are identified by a `-beta` suffix in the version tag (e.g., `v1.2.3-beta`).

**Key Beta Development Concepts:**

* **Go Build Tags**: Beta functionality uses the `//go:build beta` build tag to separate beta code from GA releases
* **Beta Files**: Include `//go:build beta` at the top of files containing beta-specific code
* **GA Stubs**: Use `//go:build !beta` for GA-only files that provide stubs for beta functionality
* **Automatic Building**: Goreleaser automatically includes beta code when building releases with `-beta` tags

For detailed beta development guidance, see the [Beta Development Guide](contributing/beta.md).

## Getting Started

- [Set Up Your Development Environment](contributing/development-environment.md)
- [SDK Design](contributing/sdk-design.md)
- [Beta Development Guide](contributing/beta.md)
- [Testing Strategy](contributing/acceptance-test-strategy.md)

## Process

- [Pull Request Checklist](contributing/pr-checklist.md)
