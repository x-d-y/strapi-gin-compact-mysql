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
  let res = await Request("http://localhost:8080/test/get-test?name=xdy&age=13",'get',data)
} 

post()
