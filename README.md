# ed-go-testing
Unit testing tutorial for Go

# Writing unit tests
To write unit tests in go:
1. Create a file in the same directory as the file you want to test. It should have the same name as the file you want to test, along with the _test.go suffix.
2. Add an `import "testing"` statement after the `package` statement.
3. Create a function that accepts a single argument: `t *testing.T`. For example, the following file is called <strong>main_test.go</strong> and resides in the <strong>./cmd/app</strong> directory:
``` go
package main

import "testing"

func TestItem(t *testing.T) {
  result := testSomething(5)
  if (result != 5) {
    t.Errorf("Result is %d; want 5", result)
  }
}
```

# Challenge Instructions
Clone this repository locally.

This program calculates factorials. It has a bug in it. Running the program gives the following output:

Factorial of 5: 120
Factorial of -3: -1

As you can see, it is not handling negative values correctly. While this logical error would be easy to spot in code, that is not always the case. Pretend you are unable to spot the cause of the logical error and write the following unit tests to help you identify it:

1. Write a test that passes in a positive value to the Factorial function.
2. Write a test that passes in a zero value to the Factorial function.
3. Write a test that passes in a negative value to the Factorial function.
4. Update the Factorial function so that all tests pass.


