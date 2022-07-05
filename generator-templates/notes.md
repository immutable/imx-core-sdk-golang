# Client generator templates

## Source

These template files https://github.com/go-swagger/go-swagger/tree/master/generator/templates/client were sourced from latest stable release of `v0.29.0`

## Code generation using custom templates

See Makefile at root path to generate the code using these template files.

## Developer Notes 

The parameters template file will be updated to include the custom header parameters.

### Generating Template File (for developer understanding only)

The `template/client/parameters.gotmpl` file will be generated from `parameters_template.gotmpl` file. This generating step will stamp the version number during the generation time.

For generaing template file, using this tool https://github.com/gliderlabs/sigil

See the follwoing sample command to generate

```
SIGIL_DELIMS={{{,}}} gliderlabs-sigil -f parameter_template.gotmpl GOLANG_SDK_VERSION="0.1.0" > parameter.gotmpl
```

*Note: These steps are all integrated with Makefile.*