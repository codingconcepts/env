// Package env makes satisfying factor III of the 12-factor methodology easy, by allowing struct fields to be populated directly from environment
// variables with the use of struct tags.
//
// To use, create a struct tag called "env" and call env.Set, passing a pointer
// to the struct you wish to populate.  You can optionally, provide a "required"
// tag to determine whether an error should be returned in the event of missing
// environment configuration.
//
// Like the encoding/* packages, env.Set will return an error if a non-pointer
// type is provided.
package env
