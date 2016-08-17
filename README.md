# Template Generating and Testing Flamework using go-jstmpl

tempcube is a go-jstmpl support framework to building and testing.
that is Golang code generater from a JSON Schema setting file using the Golang text/template.  

# Create Template Project

If you create initial template project, init command create a set of files.

```sh:
tmpcube init sample

create starting ...

---> create  ./sample
---> create  ./sample/sample.go              // template
---> create  ./sample/sample_test.go         // testing for generated go code from template and Schema file
---> create  ./sample/sample_helper.go       // helper function used from main file 
---> create  ./sample/sample_helper_test.go  // helper function helper

... Finished creating initial files.

```


# Build

It is simple. build command specify source project folder, destination folder what you want to create,
and JSON Schema file path.
Building from template file 

```sh:

tempcube build ./sample ./sample_builded -s /path/to/schema

```

# Testing

tempcube use 2 type of testing. usual Golang testing and validation test to generated Golang code.

```sh:

tempcube test ./sample -s /path/to/schema

validation test ...

---> ./sample/sample.go  PASS
---> ./sample/sample_test.go PASS
---> ./sample/sample_helper.go PASS
---> ./sample/sample_helper_test.go PASS

... validation test OK

testing ...

... testiong pass

```