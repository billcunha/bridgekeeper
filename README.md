# Bridgekeeper

Bridgekeeper is a API to protect your transactions from suspects payment orders. It's based in some little infos, like the age of account, phone area code and some other stuffs.

![Bridgekeeper](https://i.imgur.com/wHr5znt.png?1)

## Principles

The idea of the API is detect some attention points and evaluated one a one. All these values are summed and the value reached cannot pass 100 points. Any request with values above this will be denied.

The attention points are:

- Account age, in days
- Profile status (verified or not)
- Total orders realized with this account
- Cardholder name empty or too short
- Profile name different with cardholder name
- Area code phone different in profile and billing address, with low orders in account

## Running

Make a clone of this repo in your GOPATH and run:

```bash
go run ./cmd/server/...
```

## Using

To call the Api, you need to use GRPC. For local use, I suggest the [evans](https://github.com/ktr0731/evans) client. I love it!

The messages are all detailed in the `api/proto` files.

## Testing

The Api was developed using TDD approach. To run all tests, use:

```bash
go test ./...
```

## How did I feel using Go?

The Golang it's a much pleasant language. It's have a powerful tools, a lot of libs, easy/short declarations and zero-valued variables. I love when I have a single one way to solve a problem, and it's enough!