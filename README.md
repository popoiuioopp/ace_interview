# ace_interview

# Run
+ To run this code, you can simply use the command below in this directory
```bash
go run main.go
```

+ you can change your test cases for each question in the file main.go 

+ or you can edit [this file](./questions/question_test.go) and use your use cases and expected results there. You can run the testing result by the command below
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

# Unit Test
+ You can see the testing coverage by using the command 
```bash
go tool cover -html=coverage.out
```
+ If you need to re-test the package again yourself.
```bash
go test ./... -coverprofile=coverage.out
```