load("@bazel_skylib//lib:paths.bzl", "paths")

CompilerInfo = provider(
    doc = "The Amazing Compiler",
    # In the real world, compiler_path and system_lib might hold File objects,
    # but for simplicity they are strings for this example. arch_flags is a list
    # of strings.
    fields = {
      "compiler_path": "the path to the compiler",
      "arch_flags": "architecture flags",
      "system_lib": "system library"
    }
)

def _compiler(ctx):
  outputs = []
  info = ctx.toolchains['//compiler:toolchain_type'].info
  compiler = info.compiler_path.files_to_run.executable

  for input in ctx.files.srcs:
    output = ctx.actions.declare_file(paths.replace_extension(input.path, ".js"))
    outputs.append(output)

    command ="./{compiler} -o {output} {input}".format(
      compiler=compiler.path,
      output=output.path,
      input=input.path
    )

    ctx.actions.run_shell(inputs = [input, compiler],
                          outputs = [output],
                          progress_message = "Compiling {input}".format(input=input),
                          command = command
    )


  return [DefaultInfo(files=depset(outputs))]

compiler = rule(
  implementation=_compiler,
  attrs= {
    'srcs': attr.label_list(allow_files=True),
  },
   toolchains = [
        "//compiler:toolchain_type"
    ],
)

def _compiler_toolchain_impl(ctx):
    toolchain_info = platform_common.ToolchainInfo(
        info = CompilerInfo(
            compiler_path = ctx.attr.compiler_path,
            system_lib = ctx.attr.system_lib,
            arch_flags = ctx.attr.arch_flags,
        ),
    )
    return [toolchain_info]

compiler_toolchain = rule(
    implementation = _compiler_toolchain_impl,
    attrs = {
        "compiler_path": attr.label(allow_single_file=True),
        "system_lib": attr.string(mandatory=False),
        "arch_flags": attr.string_list(mandatory=False),
    },
)
