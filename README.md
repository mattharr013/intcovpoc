# Integration test coverage example

An example for getting more accurate code coverage by including code covered by integration tests.

This is based off the idea presented in [this](https://www.ory.sh/golang-go-code-coverage-accurate/) article.

## Example usage

```bash
# install the required tools (need to do this before the other commands will work).
make deps

# run all coverage
make coverage

# run coverage for a specific package
make code/coverage.txt
make strcheck/coverage.txt
```


## Open questions

### 1. Should exercising code in other unit tests count towards coverage?

See the `@todo` item. Utilizing the application the default way would mean that this is the case. If that is
undesirable there might be a way around it.
