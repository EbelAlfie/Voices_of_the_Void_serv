# Welcome to the Abyss.. Where soul perist while AI persist
A simple integration with customer, AI customer service, and human customer service voice calling with N8N. The customer (you) will first talk to the AI about stuff they need to know, then the customer may request to call a real customer service and the AI will create a call between you and the customer service (Note, you must have two devices or optionaly, friends that acts as a customer service)

This is the backend service that acts as a middleware between n8n and the android client app in another repository.

The AI LLM inside n8n will do a http request to this in order to pass it's messages along with the user id that it replied to. Then, this backend service will forward either the messages, or a command to connect the client between the customer and customer service via websocket integration. (Since n8n supports the use of code and external node js module, maybe next time, I will try to use that instead)

To run the backend service, just use
```go run votv.co```

To run the n8n, use this command
```start_8n.sh```

Don't forget, you must have n8n first either localy or cloud. If you want to host it like I do, simply use this command 
```npm i -g n8n```

oh node js must be either 18, 20, or 22. Or so I heard...

The workflow is the ```n8n_workflow.json``` you can import it to your own workflow