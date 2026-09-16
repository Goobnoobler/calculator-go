# calculator-go

An HTTP API that evaluates arithmetic expressions. Expressions arrive in ordinary infix form (`2*(4+3)`), get converted to reverse Polish notation with the shunting-yard algorithm, and are then evaluated off a stack.

The frontend that consumes this API lives in [calculator-vue](https://github.com/Goobnoobler/calculator-vue).

This backend was almost completely hand written with the exception of some initial boilerplate code for the unit/integration tests and the writing of this Readme.

I did use AI to give me hints when I was really stuck with the shunting-yard algorithm logic or when I really didn't understand something to do with Go.

## Requirements

- Go 1.21 or newer

## Running

```sh
go run .
```

The server listens on `:8080`.

## API

### `POST /api/calculate`

Request body:

```json
{ "expression": "2*(4+3)" }
```

Response:

```json
{ "result": 14 }
```

Results are JSON numbers, so division returns a fractional value (`7/2` gives `3.5`) rather than truncating.

### Errors

Failures return `400 Bad Request` with the message as plain text:

| Message | Cause |
| --- | --- |
| `malformed json` | Request body could not be decoded |
| `unmatched closing parenthesis detected` | A `)` with no matching `(` |
| `unclosed parenthesis detected` | A `(` that is never closed |
| `calculator only accepts digits and +-*/` | Unsupported character in the expression |
| `cannot divide by zero` | Division with a zero divisor |
| `unknown operator in expression` | Operator token that is not `+ - * /` |
| `not enough operands for operater` | An operator with fewer than two values to consume |
| `expression left more than one value` | Tokens remained on the stack after evaluation |
| `no answer returned` | Empty expression |

## What the parser accepts

Digits, decimal points, the four operators `+ - * /`, and parentheses. Standard precedence applies, so `2+3*4` is `14`, not `20`.

Whitespace is not accepted — a space is an unsupported character. Unary minus is also unsupported, so `-5` fails with `not enough operands for operater` rather than returning `-5`.

## Layout

| Path | Purpose |
| --- | --- |
| `main.go` | Entry point; binds the router to `:8080` |
| `server/` | Router and the `/api/calculate` handler, request and response types |
| `calculate/` | `Shunt` (infix to RPN) and `Eval` (RPN to a result), plus the sentinel errors |

## Tests

```sh
go test ./...
```

Table-driven tests cover the parser and evaluator in `calculate/`, and the handler's success and error paths in `server/`.
