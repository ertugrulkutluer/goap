package middleware

// func CORSMiddleware(h http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         log.Println("middleware", r.URL)
// 		w.WriteHeader().Set("Access-Control-Allow-Credentials", "true")
// 		w.WriteHeader().Set("Access-Control-Allow-Origin", "*")
// 		w.WriteHeader().Set("Content-Type", "application/json; charset=utf-8")
// 		w.WriteHeader().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
// 		w.WriteHeader().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
// 		w

//     })
// }
// w.WriteHeader(code)
// func JSONMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.WriterWriteHeader.Set("Content-Type", "application/json")
// 		c.Next()
// 	}
// }
