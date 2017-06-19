// Package shape allows struct fields to be populated directly from environment
// variables.
//
// To use, create a struct tag called "env" and call shape.Env, passing a pointer
// to the struct you wish to populate.  You can optionally, provide a "required"
// tag to determine whether an error should be returned in the event of missing
// environment configuration.
//
// Like the encoding/* packages, shape.Env will return an error if a non-pointer
// type is provided.
package shape
