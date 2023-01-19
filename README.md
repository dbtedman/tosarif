# [ToSARIF](https://github.com/dbtedman/tosarif)

[![CI GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/tosarif/ci.yml?branch=main&style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/tosarif/actions/workflows/ci.yml?query=branch%3Amain)
[![SAST GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/tosarif/sast.yml?branch=main&style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/tosarif/actions/workflows/sast.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dbtedman/tosarif?style=for-the-badge)](https://goreportcard.com/report/github.com/dbtedman/tosarif)

> ⚠️ WARNING! This project is in early development, and is not ready for production use.

-   [Usage Instructions](#usage-instructions)
-   [Contributing](#contributing)
-   [Security](#security)
-   [License](#license)

## Usage Instructions

### OSV Scanner

> ⚠️ Proposal Only

```shell
# Support pipes, and auto-detection for input format
osv-scanner -json ./ | tosarif > out.sarif

# Manually select input format
osv-scanner -json ./ | tosarif -format=osv-scanner > out.sarif

# Write to file directly
osv-scanner -json ./ | tosarif --out=out.sarif

# Read from file directly
tosarif in=in.json > out.sarif
```

## Contributing

See our [Contributing Guide](./CONTRIBUTING.md) to learn how to contribute to this tool.

## Security

Read our [Security Guide](SECURITY.md) to learn how security is considered during the development and operation of this
tool.

## License

See [LICENSE.md](./LICENSE.md) for details.
