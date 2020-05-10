## Echotodo

A toy todo app to strengthen my understanding on backend development and Google Cloud services

## Things I've learned

- Creating CRUD endpoint with Go Echo framework
- How to deploy services to AppEngine
- How to deploy services to Cloud Run
- How to connect to Cloud SQL instances
- How to use Cloud SQL Proxy

# Sample endpoint
AppEngine: https://todo-dot-toyprojects.appspot.com/todos

# Endpoint

- Get all todos
```
curl -X GET https://todo-dot-toyprojects.appspot.com/todos
```

- Create todo
```
curl -X POST -d '{
    "title": "buy more coffee"
}' 'https://todo-dot-toyprojects.appspot.com/todos'
```

- Get a single todo
```
curl -X GET https://todo-dot-toyprojects.appspot.com/todos/1
```

- Update a todo
```
curl -X DELETE -d '{
    "title": "buy some coffee"
}' 'https://todo-dot-toyprojects.appspot.com/todos/1'
```

- Delete a todo
```
curl -X DELETE https://todo-dot-toyprojects.appspot.com/todos/1
```