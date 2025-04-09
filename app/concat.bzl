def _concat(ctx):
    output = ctx.outputs.output
    concat_binary = ctx.executable._concat_binary

    args = [output.path] + [input.path for input in ctx.files.srcs]

    ctx.actions.run(
        executable = concat_binary,
        arguments = args,
        inputs = depset(ctx.files.srcs),
        outputs = [output],
        progress_message = "Concatenating files into %s" % output.path,
    )

    return [DefaultInfo(files = depset([output]))]

concat = rule(
    implementation = _concat,
    attrs = {
        "srcs": attr.label_list(
            mandatory = True,
            allow_files = True,
            doc = "Input files to concatenate",
        ),
        "output": attr.output(
            mandatory = True,
            doc = "Output file",
        ),
        "_concat_binary": attr.label(
            default = Label("//:the_concatenator"),
            executable = True,
            cfg = "exec",
            doc = "Concatenation binary",
        ),
    },
    doc = "Concatenates multiple files into a single output file",
)
