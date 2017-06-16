// Package shape allows struct fields to be populated directly from environment
// variables.  Simply create a struct tag called "env", and optionally provided
// a "required" tag to determine whether an error should be returned in the even
// of missing environment configuration and call the shape.Env function.
package shape
