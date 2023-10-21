# go-gqlgen

A graphQL based server using gqlgen lib

## How to run

1. Clone the repo
2. Run db.sh script to setup mysql db
3. Run `go run server.go` in the root directory
4. Open `http://localhost:8080` in the browser
5. Run the following query in the graphql playground to get all the links

```graphql
query {
  links {
    id
    title
    address
    users {
      id
      username
    }
  }
}
```

6. Run the following mutation to create a user

```graphql
mutation {
  createUser(input: { username: "aba", password: "123456" })
}
```

and then run the following mutation to login

```graphql
mutation {
  login(input: { username: "aba", password: "123456" })
}
```

7. Run the following mutation to create a new link (make sure to add the jwt token in the authorization header)

```graphql
mutation {
  createLink(input: { title: "Google", address: "https://www.google.com" }) {
    id
    title
    address
  }
}
```
