


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
![](https://github.com/myJPQ/gqlgen-todos/raw/master/testimage/create.jpg)
![](https://github.com/myJPQ/gqlgen-todos/raw/master/testimage/update.jpg)

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
![](https://github.com/myJPQ/gqlgen-todos/raw/master/testimage/query.jpg)

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
![](https://github.com/myJPQ/gqlgen-todos/raw/master/testimage/delete.jpg)
