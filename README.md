# Experimenting with Bazel

This ia a playground to better understand some specific bazel features

Directories:
* **/app/** a silly TS app that needs its TS files compiled, and those files concattenated
* **/compiler/** a silly "compiler" that takes TS files and produced js files


Stuff to explore:
* [x] convert from WORKSPACE to Bazel Modules
  * [x] figure out why compiler/WORKSPACE is needed
  * [x] rename the-compiler back to compiler
* [x] using toolchains
  * [ ] polish toolchain usage
* [ ] using go workspaces
* [ ] using go modules
* [ ] using node modules
