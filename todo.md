# To Do

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
- [ ] refactoring
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
    - [x] implementing new idea from Elegant Objects
      - [x] package generator with Boilerplate struct and StrValue so you can use it later as: `docs := NewBoilerplate(tmpl, "Missing Number")`
        - it represents a boilerplate file which knows how to generate itself => doesn't feel right
      - [x] eoBoilerplate in `generator` package 
        - implemented the structure from below using an interface for the Filesystem => we're getting there
    - [x] chore: switch git setting repo to ssh
  - NEXT:
    - [ ] EO version
      - [ ] CONTINUE HERE: implement the FS so that you can test the EO version
      - [ ] can I use accessor functions w/o get? not sure tmpl.Execute accepts
    - [ ] clean up main
    - [ ] clarify what was this task: "exercise directory deletion => no test"
    - [ ] idea: create Generator struct with ?Process(request) and ?Execute() methods

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

- [ ] assert package: provide 'name' instead of input

## CI

WHAT YOU DID
- [x] add Dockerfile => the image is large => you can run tests there
- [x] run all tests in container using the Dockerfile
- [x] you added a dsa-pod.yaml and you can deploy it in kind using `ko ` => just executes main and prints the temp message
- [x] added a test-pod.yaml with command to run tests but it crashes => looks like `ko` images are not supposed to run tests

CONTINUE HERE
- [ ] maybe testcontainers/devcontainers
- [ ] do the docker.com tutorial on Go images, multistage builds and kubernetes deployments [here](https://docs.docker.com/guides/golang/build-images/)
- [ ] decide what to do: create an ephemeral pod to run tests in or stay with pure docker
- [ ] configure yaml extension => its not recognizing kubernetes manifests
- [ ] run CI on every push to master - see fabric samples

## Resources

- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [ko](https://ko.build/features/k8s/)
