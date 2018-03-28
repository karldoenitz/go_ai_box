# go_ai_box
AI音箱的后台服务，golang开发

# 接口简介

| uri                 | 功能                   |  测试  |
| --------            | :-------------------  | :---: |
| /lean/home          | 获取home页面布局        |  通过  |
| /lean/detail        | 获取detail页面布局      |  通过  |
| /service/search     | 查询数据               |  通过  |
| /source/music       | 获取music资源数据       |  通过  |
| /source/singer      | 获取singer资源数据      |  通过  |
| /source/playbill    | 获取playbill资源数据    |  通过  |


# 依赖库
```
go get github.com/rtt/Go-Solr
go get github.com/go-redis/redis
```
