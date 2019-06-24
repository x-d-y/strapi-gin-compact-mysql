const request = require('request')
async function sleep(ms){
  return new Promise(resolve=>{
    setTimeout(resolve,ms)
  })
}
async function Request(Url,method,data){
    return new Promise(function (resolve, reject){
      request({
        url:`${Url}`,
        method,
        json: true,
        headers: {
            "content-type": "application/json",
        },
        body:data
      },
      function(error, respond,body) {
        if(respond){
          if(respond.statusCode==200){
            return resolve(body);
          }else{
            return reject(respond.statusMessage);
          }
        }
        else{
          return reject(error);
        }
      });
    })
  }

async function post(){
  data ={
    age:18
  }

  let res = await Request("http://localhost:8080/test/post-test",'post',data)
  data.age = 17
  res = await Request("http://localhost:8080/test/put-test/111",'put',data)
  res = await Request("http://localhost:8080/test/get-test?name=xdy&age=13",'get',data)

} 

post()

