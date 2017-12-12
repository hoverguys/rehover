# TEV shader file format

TEV files are files containing GX calls that affect the TEV pipeline.

They use a subset of C with only function calls and single line comments.

No constructs (for/while/if/..) are allowed.

Functions can't span over multiple lines. The files are meant to be machine-generated, although comments are kept so they can retain some human readability.

# TDL compiled shader format

TDL files are GCN display lists generated from TEV files.