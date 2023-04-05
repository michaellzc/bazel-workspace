This project uses the [sourcegraph/controller-cdktf/gen/tfe/workspace](https://github.com/sourcegraph/controller-cdktf/tree/main/gen/tfe/workspace) go module which contains `workspace` and causing weird behaviour with bazel.

This project can be setup with

```sh
bazel run //:gazelle
bazel run //:gazelle-update-repos
```

Then compile the binary

```sh
bazel build //cmd/test
```


You should see

```
INFO: Repository com_github_sourcegraph_controller_cdktf_gen_tfe instantiated at:
  /Users/michael/Code/playground/bazel-workspace/WORKSPACE:63:16: in <toplevel>
  /Users/michael/Code/playground/bazel-workspace/deps.bzl:67:18: in go_dependencies
Repository rule go_repository defined at:
  /private/var/tmp/_bazel_michael/24824c515b06b5dc144043a986bb2cb5/external/bazel_gazelle/internal/go_repository.bzl:295:32: in <toplevel>
ERROR: An error occurred during the fetch of repository 'com_github_sourcegraph_controller_cdktf_gen_tfe':
   Traceback (most recent call last):
        File "/private/var/tmp/_bazel_michael/24824c515b06b5dc144043a986bb2cb5/external/bazel_gazelle/internal/go_repository.bzl", line 285, column 17, in _go_repository_impl
                fail("failed to generate BUILD files for %s: %s" % (
Error in fail: failed to generate BUILD files for github.com/sourcegraph/controller-cdktf/gen/tfe: gazelle: read /private/var/tmp/_bazel_michael/24824c515b06b5dc144043a986bb2cb5/external/com_github_sourcegraph_controller_cdktf_gen_tfe/WORKSPACE: is a directory
ERROR: /Users/michael/Code/playground/bazel-workspace/WORKSPACE:63:16: fetching go_repository rule //external:com_github_sourcegraph_controller_cdktf_gen_tfe: Traceback (most recent call last):
        File "/private/var/tmp/_bazel_michael/24824c515b06b5dc144043a986bb2cb5/external/bazel_gazelle/internal/go_repository.bzl", line 285, column 17, in _go_repository_impl
                fail("failed to generate BUILD files for %s: %s" % (
Error in fail: failed to generate BUILD files for github.com/sourcegraph/controller-cdktf/gen/tfe: gazelle: read /private/var/tmp/_bazel_michael/24824c515b06b5dc144043a986bb2cb5/external/com_github_sourcegraph_controller_cdktf_gen_tfe/WORKSPACE: is a directory
ERROR: /Users/michael/Code/playground/bazel-workspace/cmd/test/BUILD.bazel:3:11: //cmd/test:test_lib depends on @com_github_sourcegraph_controller_cdktf_gen_tfe//provider:provider in repository @com_github_sourcegraph_controller_cdktf_gen_tfe which failed to fetch. no such package '@com_github_sourcegraph_controller_cdktf_gen_tfe//provider': failed to generate BUILD files for github.com/sourcegraph/controller-cdktf/gen/tfe: gazelle: read /private/var/tmp/_bazel_michael/24824c515b06b5dc144043a986bb2cb5/external/com_github_sourcegraph_controller_cdktf_gen_tfe/WORKSPACE: is a directory
ERROR: Analysis of target '//cmd/test:test' failed; build aborted:
INFO: Elapsed time: 2.693s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (29 packages loaded, 771 targets configured)
```
