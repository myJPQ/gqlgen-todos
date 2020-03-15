playground:(http://45.77.6.185:8080/)


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


delete Meeting

``` 
mutation deleteMeeting {
  deleteMeeting(input:[3]) {
    id
    meeting{
      title
      description
      startTime
    }
}
``` 

