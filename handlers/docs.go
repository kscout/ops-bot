/*
HTTP handlers.

Unless otherwise specified the following is true of all endpoints:

- No authentication required
- URLs can have embedded parameters (/api/v0/user/<username>) and query parameters
- Request bodies must be JSON formatted
- Response bodies will always be JSON
*/
package handlers
