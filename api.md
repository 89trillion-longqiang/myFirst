#1.输入士兵id获取稀有度
##接口地址
```
/idGetRarity
```
## 请求方式
GET
## 请求示例
```
http://127.0.0.1:8000/idGetRarity?id=10101
```
## 参数  说明

``` 
id  string类型 输入需要查询的士兵id
```

```
成功示例 
{
    "Rarity": "1",
    "condition": "success",
    "id": "10101"
}
#错误示列 
{
    "Rarity": "error",
    "condition": "error",
    "id": ""
}
```

#2.输入士兵id获取战力
```
接口地址 
/idGetCombatPoints
```
## 请求方式
GET
## 请求示例
```
http://127.0.0.1:8000/idGetCombatPoints?id=10101
```
## 参数  说明
``` 
id  string类型 输入需要查询的士兵id
```
```
成功示例 
{
    "CombatPoints": "167",
    "condition": "success",
    "id": "10101"
}
错误示列 
{
    "CombatPoints": "error",
    "condition": "not_exist",
    "id": "101"
}
```
#3.获取每个阶段解锁相应士兵的json数据

##接口地址
```
/getSoldiersJson
```
## 请求方式
GET
## 请求示例
```
http://127.0.0.1:8000/getSoldiersJson?stage=army_desc_10101
```
## 参数  说明
``` 
id  string类型 输入需要查询的阶段
```
```
成功示例
{
    "condition": "success",
    "data": [
        {
            "id": "10102",
            "name": "Swordsman",
            "UnlockArena": "0",
            "rarity": "1",
            "combatPoints": "342",
            "desc": "army_desc_10101"
        },
        {
            "id": "10103",
            "name": "Swordsman",
            "UnlockArena": "0",
            "rarity": "1",
            "combatPoints": "691",
            "desc": "army_desc_10101"
        }
    ],
    "stage": "army_desc_10101"
}
错误示例
{
    "condition": "not_exist",
    "data": "",
    "stage": "army_desc_"
}
```