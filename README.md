# goEncryptionExample
Basic SHA 265 example in golang.

Could be used to store passwords, emails, ect.

# How to use

```golang
var username string = "username"
var password string = "password"
hashedLogin := hashing.HashLoginInfo(username, password)
fmt.Println(hashedLogin)
```
# main.go output

```
Mathiss-iMac:goEncryptionExample mathis$ go run main.go
Type username: mathis
Type password: test123
Hashed username: c413eb8a7c9415f5f8ab446ee87b55f6d636d63c1af34c861949d6a300b8f6c7
Hashed password: 28742e3f48057028d5bfe3cf825e6a94cc6cbfbaa27ca7e7269a60563b76e90c
Hashing salt: k47DbTjjPeqVuDq9xrh2KqQQGXJqfmp99arntrMvfkTPrm7Pzesn9mByiSXYU65c
```
If we run that a second time, you'll see the username hash is the same, but the password is diffirent (due to a diffirent salt)!
```
Mathiss-iMac:goEncryptionExample mathis$ go run main.go
Type username: mathis
Type password: test123
Hashed username: c413eb8a7c9415f5f8ab446ee87b55f6d636d63c1af34c861949d6a300b8f6c7
Hashed password: 4982f30b1e5ea11a977dde2f51e21a686a8058572856c47c9347f5a8021174f3
Hashing salt: 1kjBCLvctOMDE8QvWZ4PYsMfP822IVnWqAqHLcH2DUhYJbw3zJWnUEL35WwzFClh
```