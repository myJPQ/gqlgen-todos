


Create and Update Meeting

``` 
mutation updateMeeting {
  updateMeeting(input:[{id: 2, title:"1",description: "23" ,startTime: 1},{id: 3, title:"1",description: "23" ,startTime: 1}]) {
    id
    meeting{
      title
      description
      startTime
    }
  }
}
```
![](https://raw.githubusercontent.com/myJPQ/gqlgen-todos/master/testimage/create.jpg)

![](https://raw.githubusercontent.com/myJPQ/gqlgen-todos/master/testimage/update.jpg)

query Meeting

``` 
query  getMeeting {  
   getMeeting{
    id
    meeting{
      title
      description
      startTime
    }
  }
}
``` 
![](https://raw.githubusercontent.com/myJPQ/gqlgen-todos/master/testimage/query.jpg)

delete Meeting

``` 
mutation deleteMeeting {
  deleteMeeting(input:[1]) {
    id
    meeting{
      title
      description
      startTime
    }
}
}
``` 
![](https://raw.githubusercontent.com/myJPQ/gqlgen-todos/master/testimage/delete.jpg)
