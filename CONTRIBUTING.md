# [ToSARIF](https://github.com/dbtedman/tosarif) Contributing

-   [Getting Started](#getting-started)
-   [Verification](#verification)
-   [Design](#design)
-   [References](#references)

## Getting Started

```shell
nvm use && make
```

## Verification

### Linting

```shell
make lint
```

These rules can then be automatically applied:

```shell
make format
```

### Unit Testing

```shell
make test
```

## Design

### Repository Structure

| Directory           | Purpose                                          |
| ------------------- | ------------------------------------------------ |
| `.github/`          | Dependabot configuration and pipeline workflows. |
| `.husky/`           | Husky git hook configuration.                    |
| `cmd/`              | Entrypoint to go application.                    |
| `internal/domain/`  | Core application logic.                          |
| `internal/gateway/` | Encapsulated access to external systems.         |

### Stack

-   [✅ Cobra (cobra.dev)](https://cobra.dev)
-   [✅ CodeQL (codeql.github.com)](https://codeql.github.com)
-   [✅ Dependabot (docs.github.com)](https://docs.github.com/en/code-security/dependabot)
-   [✅ GitHub Actions (github.com)](https://github.com/features/actions)
-   [✅ GitHub Security (github.com)](https://github.com/security)
-   [✅ Go (go.dev)](https://go.dev)
-   [✅ Go Embed (pkg.go.dev)](https://pkg.go.dev/embed)
-   [✅ Go Modules (go.dev)](https://go.dev/ref/mod)
-   [✅ Go Report Card (goreportcard.com)](https://goreportcard.com/)
-   [✅ GolangCI Lint (github.com)](https://github.com/golangci/golangci-lint)
-   [✅ Husky (typicode.github.io)](https://typicode.github.io/husky/#/)
-   [✅ Node Version Manager (github.com)](https://github.com/nvm-sh/nvm)
-   [✅ NodeJS (nodejs.org)](https://nodejs.org/en/)
-   [✅ OSV Scanning (osv.dev)](https://osv.dev/)
-   [✅ PNPM (pnpm.io)](https://pnpm.io/)
-   [✅ Prettier (prettier.io)](https://prettier.io/)
-   [✅ Security Scorecards (securityscorecards.dev)](https://securityscorecards.dev/)
-   [✅ Snyk (snyk.io)](https://snyk.io)
-   [✅ Testify (github.com)](https://github.com/stretchr/testify)
-   [✅ TypeScript (typescriptlang.org)](https://www.typescriptlang.org/)

## References

> 💡 Resources referenced during the development of this project.

-   [Accretion (github.com)](https://github.com/dbtedman/accretion)
