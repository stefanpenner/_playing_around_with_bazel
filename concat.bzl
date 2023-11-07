def _concat(ctx):
  output = ctx.outputs.output
  command = "cat {inputs} > {output}".format(
    inputs = " ".join([input.path for input in ctx.files.srcs]),
    output = output.path,
  )
  ctx.actions.run_shell(inputs=ctx.files.srcs,
                        outputs=[output],
                        progress_message="Concatenating {output}".format(output=output.path),
                        command=command
  )
  return [DefaultInfo(files=depset([output]))]

concat = rule(
  implementation=_concat,
  attrs={
    "srcs": attr.label_list(mandatory=True, allow_files=True),
    "output": attr.output(mandatory=True),
  },
  doc="Concatenate files",
)
