## Part 2 exercise 10

Apps and manifests located located in [/apps](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps) at commit https://github.com/mtuomiko/kubernetes-devops/commit/0919e925f1a2badcc5db878bff13f1c11018fb4c. Published through Docker Hub at [mtuomiko/todo-app-frontend](https://hub.docker.com/r/mtuomiko/todo-app-frontend) with tag `0.0.3` and [mtuomiko/todo-app-backend](https://hub.docker.com/r/mtuomiko/todo-app-backend) with tag `0.0.5`.

1. I added the 140 char limit for todo titles to both frontend and backend. Backend has the default Echo framework Logger middleware enabled but the logger doesn't have direct access to the request body (if we want to actually see the todo titles sent to the backend). I could've cobbled something together or just used the BodyDump middleware but instead I opted to make things easier for myself and added custom log messages in the route handler with tokens `todo_added` and `todo_validation_error` that include the actual title.

2. We can see our logs in Grafana by using the log stream selector `{app="todo-app-backend"}` and log pipeline `|~ "todo_added|todo_validation_error"`.

```
2021-04-02 00:01:47	2021-04-01T21:01:47.6840696Z stderr F 2021/04/01 21:01:47 todo_validation_error: title: "141char89012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901" Key: 'Todo.Title' Error:Field validation for 'Title' failed on the 'max' tag
2021-04-02 00:01:29	2021-04-01T21:01:29.0006364Z stderr F 2021/04/01 21:01:29 todo_added: title: "LONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCATLONGCAT"
2021-04-01 23:52:33	2021-04-01T20:52:33.7064859Z stderr F 2021/04/01 20:52:33 todo_validation_error: title: "LONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRINGLONGSTRING" Key: 'Todo.Title' Error:Field validation for 'Title' failed on the 'max' tag
2021-04-01 23:52:21	2021-04-01T20:52:21.5927155Z stderr F 2021/04/01 20:52:21 todo_added: title: "Hello Grafana!"
2021-04-01 23:34:45	2021-04-01T20:34:45.1667415Z stderr F 2021/04/01 20:34:45 todo_added: title: "140char8901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"
2021-04-01 23:34:32	2021-04-01T20:34:32.6484046Z stderr F 2021/04/01 20:34:32 todo_validation_error: title: "141456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901" Key: 'Todo.Title' Error:Field validation for 'Title' failed on the 'max' tag
2021-04-01 23:33:19	2021-04-01T20:33:19.2978862Z stderr F 2021/04/01 20:33:19 todo_added: title: "Test4"
```

3. Now we wait to see if the `todo-generator` can find a URL longer than 135 characters 😅
