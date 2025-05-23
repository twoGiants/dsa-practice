# To-Dos

## Boilerplate Generator

- [x] first boilerplate generator implementation
  - [x] create e2e boilerplate.md Go template
  - [x] parametrize: title, solution subtitle and link
  - [x] load docs.tmpl file
  - [x] store file to path
  - [x] store generated boilerplate docs into correct folders
    - [x] create directories (if they don't exist)
      - [x] take three inputs: pattern, difficulty, title
    - [x] generate, then store
  - [x] execute from cmd i.e. get inputs from cli
- [x] refactoring
  - WHAT YOU DID:
    - [x] command line args parsing and validation
    - [x] e2e creation of a sample in a temp dir: `go run . create arrays easy "Missing Number"`
    - [x] deletion of sample: `go run . delete arrays easy "Missing Number"`
    - [x] create Generator
      - [x] setup config data: target dir, template location, exercise metadata from args processing
      - [x] init generator with config data
      - [x] execute e2e generation
    - [x] move args parsing, exercise struct and validation to generator
    - [x] rename generator package to boilerplate
    - [x] refactor e2e test and design => make testable
      - [x] extracted metadata
      - [x] added validating "decorator" metadata
      - [x] moved args parsing to main, removed args.go
- [ ] Elegant Objects version
  - WHAT YOU DID
    - [x] implementing new idea from Elegant Objects
      - [x] add `package generator` with `Boilerplate` struct and `StrValue` method so you can use it as: `docs := NewBoilerplate(tmpl, "Missing Number"); docs.StrValue()`
        - it represents a boilerplate file which knows how to generate itself => doesn't feel right
      - [x] add second implementation in `eoBoilerplate.go`
        - [x] implemented the structure from below using an interface for the Filesystem => we're getting there
        - [x] 22.03.25
          - [x] implement the FS so that you can test the EO version => works
          - [x] add DocsFile interface to use Content
          - [x] add docFilePath and smallKebab methods to DocsBoilerplate => doesn't feel right
          - [x] create `eoDocs.go` entrypoint for manual testing => works
        - [x] 27.04.25
          - [x] understand eo implementation => not easy after a month
            - next time read through the example code below
            - go from EoDocsCreate object by object
            - the boilerplate can save itself to provided Store and delete itself
            - the store is implemented by the filesystem
            - a common DocsFile interface is implemented by Boilerplate and Template so they can interact internally
            - actual data is provided by DocsTemplateData and put into the Template by the Boilerplate
          - [x] add `Delete(fs filesystem)` method to DocsBoilerplate
          - [x] implement integration tests for Create and Delete
        - [x] 28.04.15
          - [x] fix ci test
          - [x] refactor docsFilePath and smallKebab
          - [x] CLI: reads command, pattern, difficulty and title from command line user input => separate package, no connection to DocsBoilerplate
          - [x] filesystem creates directory chain if needed
          - [x] create `Request` object which contains everything the Boilerplate needs
    - [x] chore: switch git setting repo to ssh
  - NEXT:
    - [ ] CONTINUE HERE: read what you did on 27.04.25, THEN: have something e2e and then refactor and try other implementations like from James Shore, Hexagonal Architecture, procedural etc.
    - Whats left for E2E execution?
      - [ ] implement `*-solution.md` boilerplate generation
      - [ ] generate both concurrently, use go go routines
      - [ ] `Pattern` object (like App/Entrypoint) which centralize the creation of all the boilerplate for one pattern
        - uses DocsBoilerplate
        - uses CLI
      - [ ] move filesystem to infrastructure
      - [ ] create Nullable StubbedCli
      - [ ] use polymorphism for Save/Delete
      - [ ] tests for Save and Delete using Nullable filesystem
    - [ ] can I use accessor functions w/o get? not sure tmpl.Execute accepts
    - [ ] clean up main
    - [ ] clarify what was this task: "exercise directory deletion => no test"
    - [ ] idea: create Generator struct with ?Process(request) and ?Execute() methods

This is procedural thinking:

- receive config info with:
  - where is template, where to save boilerplate
  - data for template parameters
- load template from disk
- replace template parameters with data which results in the boilerplate
- store boilerplate to disk

Is this living objects thinking?:

- The boilerplate docs file stores itself
  - it wraps a template file
    - which provides its content
    - knows how to load itself
  - it wraps a config to know where to store itself
  - it wraps data which is used to populate the template

*Elegant Objects* inspired prototype:

```go
type Filesystem interface {
  Load(from string) (string, error)
  Save(data, to string) error
}

filesystem := Filesystem()
config := Config()
OnDiskFile( // StoredFile
  config, // config.To
  DocsBoilerplate(
      InMemoryFile(config, filesystem), // config.From, LoadedFile
      DocsTemplateData( /* ??? */),
  ),
).Save(filesystem)
```

## Testing

WHAT YOU DID

- [x] 24.03.25
  - [x] add tests to assert package
    - [x] make FirstDimensionLengthEqual testable
    - [x] add first positive test

CONTINUE HERE

- [ ] add tests to assert package
  - [ ] add negative FirstDimensionLengthEqual test
  - [ ] make other assert functions testable
  - [ ] refactor FirstDimensionLengthEqual to assert all dimensions; update tests
- [ ] assert package: provide 'name' instead of input

## CI

WHAT YOU DID

- [x] add Dockerfile => the image is large => you can run tests there
- [x] run all tests in container using the Dockerfile
- [x] you added a dsa-pod.yaml and you can deploy it in kind using `ko` => just executes main and prints the temp message
- [x] added a test-pod.yaml with command to run tests but it crashes => looks like `ko` images are not supposed to run tests

CONTINUE HERE

- [ ] add pre-commit or pre-push hook to run CI
- [ ] remove tools installation in Pipeline use existing GitHub action
- [ ] use Copilot to build strong CI/CD
  - [x] setup.sh script
  - [x] add Makefile
  - [x] add golangci-lint with configuration
  - [x] add Github Actions CI/verify pipeline
- [ ] maybe testcontainers/devcontainers
- [ ] do the docker.com tutorial on Go images, multistage builds and kubernetes deployments [here](https://docs.docker.com/guides/golang/build-images/)
- [ ] decide what to do: create an ephemeral pod to run tests in or stay with pure docker
- [ ] configure yaml extension => its not recognizing kubernetes manifests
- [ ] run CI on every push to master - see fabric samples

## Refactoring

- [x] add refactoring section with first sample
- [x] practice theatrical sample
  - [x] page 14: Extracting Volume Credits
  - [x] add non refactored version
  - [x] resolve ci issues
  - [x] page 25: Splitting the Phases of Calculation and Formatting
  - [x] page 28: move amountFor and replace
  - [x] page 36: Creating a Performance Calculator
  - [x] page 39: at the top, polymorphic calculator
  - [ ] page 34 at the end: add tests that probe the intermediate data structure
- [x] practice retention => retained

## Resources

- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [ko](https://ko.build/features/k8s/)
