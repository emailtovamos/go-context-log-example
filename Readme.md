This is an example to show how to use context and zerolog in Go. 

The context has to be initialized from background at first. Then this context can be passed through the loggers so that we can have the logs in other functions in that context. This makes things easier to trace later in the logs when the code is implemented. 