package main

import "fmt"

func main() {
    //如果只是声明，没有初始化，默认是nil map,不能存放kv对;
    //必须用make再创建
    var countryCapitalMap map[string]string
    //var countryCapitalMap map[string]string = map[string]string{"France":"Paris"}
    //var countryCapitalMap = map[string]string{"France":"Paris"}

   /* 创建集合 */
   //可以直接用make声明
   /*var*/ countryCapitalMap = make(map[string]string)

   
   //map 插入 key-value 对，各个国家对应的首都 
   countryCapitalMap["France"] = "Paris"
   countryCapitalMap["Italy"] = "Rome"
   countryCapitalMap["Japan"] = "Tokyo"
   countryCapitalMap["India"] = "New Delhi"
   

   /* 使用 key 输出 map 值 */
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }

   /* 查看元素在集合中是否存在 */
   captial, ok := countryCapitalMap["United States"]
   /* 如果 ok 是 true, 则存在，否则不存在 */
   if(ok){
      fmt.Println("Capital of United States is", captial)  
   }else {
      fmt.Println("Capital of United States is not present") 
   }
}
