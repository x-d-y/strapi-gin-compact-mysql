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
    name:"xdy"
  }

  let res = await Request("http://localhost:8080/template/post-test",'post',data)
  data.
  res = await Request("http://localhost:8080/template/update-Test/111",'put',data)
  res = await Request("http://localhost:8080/template/gets-test?name=xdy",'get',data)
  // res = await Request("http://localhost:8080/test/delete-test/111",'delete',data)
  // res = await Request("http://localhost:8080/test/get-test/5d11a7aab07436cdcb64eec0",'get',data)

} 

post()

