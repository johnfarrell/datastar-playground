# datastar-playground

Sample project for playing around and learning
[datastar](https://github.com/starfederation/datastar) and
[gostar](https://github.com/delaneyj/gostar).

The goal of this project is to primarily get more experienced with these two libraries, but the application itself will
be used for planning and tracking personal goals in the game
[Old School RuneScape](https://oldschool.runescape.com/).

I think this will be a good test of the capabilities of the frameworks since I intend to implement a few features that
may be common in regular web applications:
 - [ ] Integration with third party APIs
 - [ ] User profiles and sessions
 - [ ] Graph-based editor for structuring goal trees. (This one might not be common but hopefully will be a good test for extending the frameworks.)

Other general goals of this project is to serve as a good example of general Golang best-practices,
usage of [Zap](https://github.com/uber-go/zap) for logging, [Cobra](https://github.com/spf13/cobra) for CLI control,
[Viper](https://github.com/spf13/viper) for configuration, and the built-in `net/http` for that stuff.

# Usage

`datastar-playground` uses a [Taskfile](https://github.com/go-task/task?tab=readme-ov-file) for the build system. To install
Task on your system, follow the instructions [here](https://taskfile.dev/installation/).

Building and running DSP can be done by:
```shell
task -- server
```

This will download and verify all dependencies, build the DSP binary, and run it.

To see all available tasks:
```shell
task
```

Useful tasks include `build`, and `run`.