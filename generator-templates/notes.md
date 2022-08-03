# Client generator templates

*Note: These steps are all integrated with Makefile no need to run these steps manually.*
## Source

These template files https://github.com/OpenAPITools/openapi-generator/tree/v6.0.1/modules/openapi-generator/src/main/resources/go were sourced from latest stable release of `v6.0.1`

## Code generation using custom templates

See Makefile at root path to generate the code using these template files.

## Developer Notes 

The `client.mustache` template file will be updated to include the custom header parameters.

### Generating Template File (for developer understanding only)

The `templates/client.mustache` file will be generated from ``client_template.mustache`` file. This generating step will stamp the version number during the generation time.

For generating template file, using this tool https://github.com/gliderlabs/sigil

See the following sample command to generate

```
SIGIL_DELIMS={{{{,}}}} gliderlabs-sigil -f api_template.mustache GOLANG_SDK_VERSION="0.1.0" > api.mustache
```
