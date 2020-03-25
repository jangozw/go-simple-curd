# go curd 大法入门

go curd 基础示例. 需要用到的外部包只有gorm， 其余都是go自带的。


# 准备工作

1, 拉依赖包

根目录执行:

```
go mod download
```


2,  去database/database.go中修改数据库连接信息， 然后在你的数据库中导入：
```
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '手机号',
  `password` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '密码',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';
```

3, 启动


在项目根目录:
```
go run main.go 
```

3， 请求

* 添加用户接口
```
jangozw@jianggedeMacBook-Pro:~|⇒  curl http://127.0.0.1:8008/user/add\?hello
{"Code":0,"Msg":"success","Data":"add user succ"}%
```

* 用户列表接口
```
jangozw@jianggedeMacBook-Pro:~|⇒  curl http://127.0.0.1:8008/user/list
{"Code":0,"Msg":"success","Data":[{"Id":1,"Name":"test"},{"Id":2,"Name":"test"},{"Id":3,"Name":"hello"},{"Id":4,"Name":"hello"},{"Id":5,"Name":"hello"},{"Id":6,"Name":"hello"},{"Id":7,"Name":"hello"},{"Id":8,"Name":"hello"},{"Id":9,"Name":"hello"},{"Id":10,"Name":"hello"},{"Id":11,"Name":"hello"},{"Id":12,"Name":"hello"},{"Id":13,"Name":"hello"},{"Id":14,"Name":"hello"},{"Id":15,"Name":"hello"},{"Id":16,"Name":"hello"},{"Id":17,"Name":"hhh_test"},{"Id":18,"Name":"_test"}]}
```


接下来看代码怎么实现的，保证秒懂!


# 文档

更多curd花式操作请查询：http://gorm.book.jasperxu.com/







