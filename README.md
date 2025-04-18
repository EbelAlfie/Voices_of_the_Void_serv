# Voices Of the Void Service
This is the backend service that acts as a middleware between n8n and the android client app in another repository.

The AI LLM inside n8n will do a http request to this in order to pass it's messages along with the user id that it replied to. Then, this backend service will forward either the messages, or a command to connect the client between the customer and customer service via websocket integration. (Since n8n supports the use of code and external node js module, maybe next time, I will try to use that instead)

To run the backend service, just use
```go run votv.co```

To run the n8n, use this command
```start_8n.sh```