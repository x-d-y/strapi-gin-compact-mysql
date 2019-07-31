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
    name:"xdy",
    salary:111,
    deptId:23
  }

  let res = await Request("http://localhost:8080/template/post-test",'post',data)
  console.log(JSON.parse(res))

  // data.name = "sbtaao"
  // data.salary = 1
  // data.deptId = 0

  // res = await Request("http://localhost:8080/template/gets-test?name=xdy",'get',data)
  // res = await Request("http://localhost:8080/template/get-test/31",'get',data)
  // console.log(JSON.parse(res))
  // res = await Request("http://localhost:8080/template/delete-test/35",'delete',data)
  // console.log(JSON.parse(res))
  // res = await Request("http://localhost:8080/template/update-Test/5",'put',data)
  // console.log(JSON.parse(res))


} 

post()

