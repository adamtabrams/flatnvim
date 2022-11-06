# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2022-11-05
### BREAKING CHANGE
- Support latest version of Neovim (#1).
- Neovim now uses env var `NVIM` instead of `NVIM_LISTEN_ADDRESS`.

### Added
- Support optional logfile.
- Improve path handling.

### Fixed
- Better buffer deleting.

### Changed
- Add to readme.
- Add unit tests.

## [0.2.1] - 2021-09-17
### Fixed
- Improve error message the for build.sh script.

### Changed
- Add basic testing script.

## [0.2.0] - 2021-09-16
### Added
- Add version info via logs and env var.

### Changed
- Add a changelog.

## [0.1.0] - 2021-09-15
### Added
- Create `flatnvim`.

[Unreleased]: https://github.com/adamtabrams/flatnvim/compare/v1.0.0...HEAD
[1.0.0]: https://adamtabrams@github.com/adamtabrams/flatnvim/compare/v0.2.1...v1.0.0
[0.2.1]: https://github.com/adamtabrams/flatnvim/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/adamtabrams/flatnvim/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/adamtabrams/flatnvim/releases/tag/v0.1.0
