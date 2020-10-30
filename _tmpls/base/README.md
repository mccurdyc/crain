<div align="center">
  <img src="./docs/imgs/logo.png"><br>
  <h3 align="center">{{{.Name}}}</h3>
  <p align="center">{{{.Description}}}</p>

  [![Build Status][build-badge]][build-url]
  [![GoDev][godev-badge]][godev-url]
  [![License][license-badge]][license-url]
  [![codecov][codecov-badge]][codecov-url]
  [![Release][release-badge]][release-url]
</div>

[build-badge]: https://circleci.com/gh/{{{.GitHubUser}}}/{{{.Name}}}/tree/master.svg?style=svg
[build-url]: https://circleci.com/gh/{{{.GitHubUser}}}/{{{.Name}}}/tree/master
[godev-badge]: https://pkg.go.dev/badge/github.com/{{{.GitHubUser}}}/{{{.Name}}}
[godev-url]: https://pkg.go.dev/github.com/{{{.GitHubUser}}}/{{{.Name}}}?tab=overview
[license-badge]: https://img.shields.io/github/license/{{{.GitHubUser}}}/{{{.Name}}}
[license-url]: LICENSE
[codecov-badge]: https://codecov.io/gh/{{{.GitHubUser}}}/{{{.Name}}}/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/{{{.GitHubUser}}}/{{{.Name}}}
[release-badge]: https://img.shields.io/github/release/{{{.GitHubUser}}}/{{{.Name}}}.svg
[release-url]: https://github.com/{{{.GitHubUser}}}/{{{.Name}}}/releases/latest

## Installing

{{{.Installing}}}

## Usage

{{{.Usage}}}

## Contributing

+ [Check out the CONTRIBUTING document.](./CONTRIBUTING.md)

## License

+ [{{{.License}}}](./LICENSE)
