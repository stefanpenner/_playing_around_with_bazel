load("@bazel_skylib//lib:paths.bzl", "paths")

def _compiler(ctx):
  compiler = ctx.files.compiler[0];

  outputs = []
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
  attrs={
    'srcs': attr.label_list(allow_files=True),
    'compiler': attr.label(allow_single_file=True, default = "//compiler:compiler")
  },
)
