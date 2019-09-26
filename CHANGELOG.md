# CHANGELOG

<a name="unreleased"></a>
## [Unreleased]

- Use Cobra CLI instead of docopt ([#116](https://github.com/segmentio/terraform-docs/issues/116))
- Update Changelog.
- Escape pipe character when generating Markdown ([#114](https://github.com/segmentio/terraform-docs/issues/114))
- Add appropriate Changelog header.
- Complete development requirements documentation.
- Configure git-chglog to not show git merge commit messages.
- Add Changelog generation via git-chglog. ([#104](https://github.com/segmentio/terraform-docs/issues/104))
- Remove occurrence of gometalinter from CircleCI config.
- Replace dep with Go Modules ([#100](https://github.com/segmentio/terraform-docs/issues/100))
- Replace gometalinter with golangci-lint. ([#103](https://github.com/segmentio/terraform-docs/issues/103))
- Add 'enhancement' section to pull request template ([#101](https://github.com/segmentio/terraform-docs/issues/101))
- Fix typo in options documentation ([#98](https://github.com/segmentio/terraform-docs/issues/98))
- Bump Homebrew formula to 0.6.0.


<a name="v0.6.0"></a>
## [v0.6.0] - 2018-12-13

- Bump version to 0.6.0.
- Unify default values of inputs ([#97](https://github.com/segmentio/terraform-docs/issues/97))
- Unify description text of inputs and outputs ([#96](https://github.com/segmentio/terraform-docs/issues/96))
- Capitalize headings in documentation.
- Fix Markdown lint errors and enhancement in README ([#94](https://github.com/segmentio/terraform-docs/issues/94))
- Update project documentation.
- Move Terraform test configuration to folder 'examples'.
- Capitalize the word 'markdown' in documentation.
- Purge History.md file.
- Add support for rendering Markdown documents ([#81](https://github.com/segmentio/terraform-docs/issues/81))
- Migrate from github.com/tj/docopt to github.com/docopt/docopt-go ([#91](https://github.com/segmentio/terraform-docs/issues/91))
- Fix authors target in Makefile to get 'Author''s email not 'Committer' ([#90](https://github.com/segmentio/terraform-docs/issues/90))
- Add requirement to discuss suitability of a new feature in an issue before submitting a pull request.


<a name="v0.5.0"></a>
## [v0.5.0] - 2018-10-24

- Bump version to 0.5.0.
- Add support to print Markdown files with underscored variable names escaped ([#48](https://github.com/segmentio/terraform-docs/issues/48)) ([#63](https://github.com/segmentio/terraform-docs/issues/63))
- Add CircleCI badge.
- Fix homebrew formula. ([#75](https://github.com/segmentio/terraform-docs/issues/75))
- Add sort by "required" and then name ([#43](https://github.com/segmentio/terraform-docs/issues/43))
- Add Homebrew formula. ([#68](https://github.com/segmentio/terraform-docs/issues/68))


<a name="v0.4.5"></a>
## [v0.4.5] - 2018-10-07

- Bump version to 0.4.5.
- Allow unquoted item names. Fixes [#64](https://github.com/segmentio/terraform-docs/issues/64) ([#70](https://github.com/segmentio/terraform-docs/issues/70))
- Change build dir structure ([#74](https://github.com/segmentio/terraform-docs/issues/74))
- Update makefile to fix Windows build filename ([#72](https://github.com/segmentio/terraform-docs/issues/72))
- Remove extra newlines between comments and inputs/outputs to fix MarkDownLint warnings ([#66](https://github.com/segmentio/terraform-docs/issues/66))
- Fix loading of comments from main.tf on Windows ([#65](https://github.com/segmentio/terraform-docs/issues/65))


<a name="v0.4.0"></a>
## [v0.4.0] - 2018-09-23

- Bump version to 0.4.0.
- Add option --with-aggregate-type-defaults to enable printing of default values for types 'list' and 'map'. ([#53](https://github.com/segmentio/terraform-docs/issues/53))
- Add option --no-sort to omit sorted rendering of inputs and outputs. ([#61](https://github.com/segmentio/terraform-docs/issues/61))
- Refactor package 'doc' for better modularity. ([#60](https://github.com/segmentio/terraform-docs/issues/60))
- Refactor package 'print' for better modularity. ([#59](https://github.com/segmentio/terraform-docs/issues/59))
- Complete CircleCI config. Add vendor directory. ([#58](https://github.com/segmentio/terraform-docs/issues/58))
- Update AUTHORS.
- Add issue and pull request templates.
- Add contributing guidelines.
- Fix indentation.
- Update documentation and license to reflect the terraform-docs authors.
- Update documentation.
- Move packages 'doc' and 'print' to internal/pkg.
- Add automated tests for package 'print'.
- Add automated tests for package 'doc'.
- Refactor code in main and prepare for tests.
- Add documentation of --version option.
- Add make target to run Go tests.
- Add make target to create and push a Git tag.
- Add make target to check Go sources for errors and warnings. Remove unused code.
- Add make target to create AUTHORS file from git logs.
- Add make target to clean the workspace.
- Add dependency management using go deps.
- Add Makefile header and build target.
- Add base CI config ([#56](https://github.com/segmentio/terraform-docs/issues/56))
- Add Maintenance section to Readme ([#55](https://github.com/segmentio/terraform-docs/issues/55))
- Update Readme.md
- Merge pull request [#44](https://github.com/segmentio/terraform-docs/issues/44) from coveo/description-before-comments
- If there is a description on an output, it should be considered before the preceding comment


<a name="v0.3.0"></a>
## [v0.3.0] - 2017-10-22

- Release v0.3.0
- auto version
- Merge pull request [#39](https://github.com/segmentio/terraform-docs/issues/39) from BWITS/[#38](https://github.com/segmentio/terraform-docs/issues/38)
- bugfix/[#38](https://github.com/segmentio/terraform-docs/issues/38)
- Merge pull request [#36](https://github.com/segmentio/terraform-docs/issues/36) from nwalke/fix_version_string
- closes [#35](https://github.com/segmentio/terraform-docs/issues/35) Updated version string


<a name="v0.2.0"></a>
## [v0.2.0] - 2017-08-15

- Release v0.2.0


<a name="v0.1.1"></a>
## [v0.1.1] - 2017-08-15

- Release v0.1.1
- Merge pull request [#34](https://github.com/segmentio/terraform-docs/issues/34) from COzero/master
- Merge pull request [#1](https://github.com/segmentio/terraform-docs/issues/1) from COzero/unquoted_names
- fixed name handling to handle unquoted hcl variable names.
- Merge pull request [#31](https://github.com/segmentio/terraform-docs/issues/31) from BWITS/typo
- fix typo
- Merge pull request [#28](https://github.com/segmentio/terraform-docs/issues/28) from s-urbaniak/no-required
- Merge pull request [#25](https://github.com/segmentio/terraform-docs/issues/25) from fatmcgav/support_output_description
- Prefer leading comments over description for outputs to maintain compatability.
- *: add --no-required option
- doc: snakecase -> camelcase
- Merge pull request [#27](https://github.com/segmentio/terraform-docs/issues/27) from fatmcgav/support_printing_type
- Add support for printing the variable 'type' in Markdown. Currently only markdown supported, but trivial to add to other outputs.
- Add support for reading `description` tag from `output` resources. Fixes [#24](https://github.com/segmentio/terraform-docs/issues/24)
- Merge pull request [#23](https://github.com/segmentio/terraform-docs/issues/23) from jacobwgillespie/patch-1
- Add note about installing with Homebrew
- Merge pull request [#22](https://github.com/segmentio/terraform-docs/issues/22) from jacobwgillespie/patch-1
- Strip # prefix from comments
- add proper license


<a name="v0.1.0"></a>
## [v0.1.0] - 2017-03-21

- Release v0.1.0
- Merge pull request [#21](https://github.com/segmentio/terraform-docs/issues/21) from s-urbaniak/files
- main: add support for files
- Merge pull request [#20](https://github.com/segmentio/terraform-docs/issues/20) from nwalke/update_readme_example
- closes [#17](https://github.com/segmentio/terraform-docs/issues/17) Updated example in README
- Merge pull request [#19](https://github.com/segmentio/terraform-docs/issues/19) from nwalke/add_sorting
- Closes [#18](https://github.com/segmentio/terraform-docs/issues/18) Added a very basic sort to inputs and outputs
- Merge pull request [#16](https://github.com/segmentio/terraform-docs/issues/16) from paybyphone/master
- doc: Account for single whitespace after comment character in header
- print/markdown: Better markdown description normalizations
- print/markdown: Added line break conversion for outputs
- doc: placeholder for list types
- doc: Allow top-level comments for variables when description missing
- print/markdown: Replace table cell newlines with HTML line breaks
- Merge pull request [#13](https://github.com/segmentio/terraform-docs/issues/13) from jbussdieker/jbb-fix-heredoc-description
- Trim whitespace on markdown description too
- Fix HEREDOC descriptions


<a name="v0.0.2"></a>
## [v0.0.2] - 2016-06-29

- Release v0.0.2
- Merge pull request [#11](https://github.com/segmentio/terraform-docs/issues/11) from segmentio/fix-md
- print: wrap default values
- Merge pull request [#10](https://github.com/segmentio/terraform-docs/issues/10) from segmentio/fix-map-type
- doc: fix map type
- add more install notes
- add dist


<a name="v0.0.1"></a>
## v0.0.1 - 2016-06-15

- Merge pull request [#5](https://github.com/segmentio/terraform-docs/issues/5) from segmentio/fix-comment
- use /** comment for module commnet
- print: actually print head comment
- img
- Merge pull request [#4](https://github.com/segmentio/terraform-docs/issues/4) from segmentio/layout
- fix view
- docs
- doc: ignore comments with /** prefix
- add head comment
- cleanup
- cleanup
- ocd
- ocd
- update doc
- add installation
- cleanup
- add usage
- better md output
- add markdown output
- outputs: use comments as description
- ocd
- clean
- working
- Initial commit


[Unreleased]: https://github.com/segmentio/terraform-docs/compare/v0.6.0...HEAD
[v0.6.0]: https://github.com/segmentio/terraform-docs/compare/v0.5.0...v0.6.0
[v0.5.0]: https://github.com/segmentio/terraform-docs/compare/v0.4.5...v0.5.0
[v0.4.5]: https://github.com/segmentio/terraform-docs/compare/v0.4.0...v0.4.5
[v0.4.0]: https://github.com/segmentio/terraform-docs/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/segmentio/terraform-docs/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/segmentio/terraform-docs/compare/v0.1.1...v0.2.0
[v0.1.1]: https://github.com/segmentio/terraform-docs/compare/v0.1.0...v0.1.1
[v0.1.0]: https://github.com/segmentio/terraform-docs/compare/v0.0.2...v0.1.0
[v0.0.2]: https://github.com/segmentio/terraform-docs/compare/v0.0.1...v0.0.2
