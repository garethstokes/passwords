passwords
=========

standard functions needed for user auth

```go
	// pass in the plaintext password with the number
	// of iterations you want the hashing algo to pass through
	password := passwords.Compute("a secret password", 3)

	fmt.Println("hash: ", password.Hash)
	fmt.Println("salt: ", password.Salt)

	result := passwords.ComputeWithSalt("a secret password", 3, password.Salt)
	if result.Hash != password.Hash {
		fmt.Println("ERROR: incorrect password")
	}
```
