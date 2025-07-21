Находясь в корневой папке запустить проект командой go run main.go

End point'ы:
http://localhost:8080/createTask 
{
    "links":["https://upload.wikimedia.org/wikipedia/commons/f/ff/Wikipedia_logo_593.jpg?20060603094750","https://upload.wikimedia.org/wikipedia/commons/f/ff/Wikipedia_logo_593.jpg?20060603094750","https://upload.wikimedia.org/wikipedia/commons/f/ff/Wikipedia_logo_593.jpg?20060603094750"]
}

http://localhost:8080/addLink?taskid=36f1a9d1-87dc-4f54-bde2-c8435a8b6e73 
{
    "links":["https://upload.wikimedia.org/wikipedia/commons/f/ff/Wikipedia_logo_593.jpg?20060603094750","https://upload.wikimedia.org/wikipedia/commons/f/ff/Wikipedia_logo_593.jpg?20060603094750"]
}

http://localhost:8080/getTaskStatus?taskid=e7539251-ddaa-4dc6-a4c0-b69fc9fd820f

