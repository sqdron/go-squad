package middleware

import (
	"github.com/sqdron/squad/endpoint"
)

func EncryptionMiddleware() endpoint.EndpointFunc {
	return func(c *endpoint.EndpointContext, m *endpoint.Message) {

		//token := c.Request.FormValue("api_token")
		//
		//if token == "" {
		//	respondWithError(401, "API token required", c)
		//	return
		//}
		//
		//if token != os.Getenv("API_TOKEN") {
		//	respondWithError(401, "Invalid API token", c)
		//	return
		//}
		//
		c.Next(m)
	}
}