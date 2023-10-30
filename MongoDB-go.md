# 1.MongoDB

## MongoDB简介

MongoDB是一个高性能，开源，无模式的，基于分布式文件存储的**文档型**数据库，由C++语言编写，其名称来源取自“hu*mongo*us”，是一种开源的文档数据库──NoSql数据库的一种。NoSql，全称是 Not Only Sql，指的是非关系型的数据库。

## MongoDB特点

MongoDB数据库的特点是高性能、易部署、易使用，存储数据非常方便。主要功能特性有：

- 面向集合存储，易存储对象类型的数据。
- 模式自由。
- 支持动态查询。
- 支持完全索引，包含内部对象。
- 支持查询。
- 支持复制和故障恢复。
- 使用高效的二进制数据存储，包括大型对象（如视频等）。
- 自动处理碎片，以支持云计算层次的扩展性
- 支持RUBY，PYTHON，JAVA，C++，PHP等多种语言。
- 文件存储格式为BSON（一种JSON的扩展）
- 可通过网络访问

## MongoDB官网

```
https://www.mongodb.com/
```



# 2.MongoDB数据库

## MongoDB基本概念

| SQL术语/概念 | MongoDB术语/概念 | 解释/说明                           |
| :----------- | :--------------- | :---------------------------------- |
| database     | database         | 数据库                              |
| table        | collection       | 数据库表/集合                       |
| row          | document         | 数据记录行/文档                     |
| column       | field            | 数据字段/域                         |
| index        | index            | 索引                                |
| table joins  |                  | 表连接,MongoDB不支持                |
| primary key  | primary key      | 主键,MongoDB自动将_id字段设置为主键 |

## MongoDB数据库

一个MongoDB中可以建立多个数据库。

MongoDB的默认数据库为"db"，该数据库存储在data目录中。

MongoDB的单个实例可以容纳多个独立的数据库，每一个都有自己的集合和权限，不同的数据库也放置在不同的文件中。

### 显示所有数据库

```
show dbs
show databases
```

### 创建或打开数据库

```
use local
use test
use golang_tech_stack
db.golang_tech_stack.insert({"name":"golang技术栈"}) # 插入数据才会显示
```

> use 命令数据库存在打开，不存在创建
>
> 插入数据才会显示

### 显示当前数据库

```
db
```

### 删除数据库

```
use golang_tech_stack
db.dropDatabase()
show dbs
```

## 系统默认数据库

- **admin**： 从权限的角度来看，这是"root"数据库。要是将一个用户添加到这个数据库，这个用户自动继承所有数据库的权限。一些特定的服务器端命令也只能从这个数据库运行，比如列出所有的数据库或者关闭服务器。
- **local:** 这个数据永远不会被复制，可以用来存储限于本地单台服务器的任意集合
- **config**: 当Mongo用于分片设置时，config数据库在内部使用，用于保存分片的相关信息。

# 3.MongoDB集合Collection

MongoDB 中使用 **createCollection()** 方法来创建集合。

语法格式：

```
db.createCollection(name, options)
```

参数说明：

- name: 要创建的集合名称
- options: 可选参数, 指定有关内存大小及索引的选项

options 可以是如下参数：

| 字段        | 类型 | 描述                                                         |
| :---------- | :--- | :----------------------------------------------------------- |
| capped      | 布尔 | （可选）如果为 true，则创建固定集合。固定集合是指有着固定大小的集合，当达到最大值时，它会自动覆盖最早的文档。 **当该值为 true 时，必须指定 size 参数。** |
| autoIndexId | 布尔 | 3.2 之后不再支持该参数。（可选）如为 true，自动在 _id 字段创建索引。默认为 false。 |
| size        | 数值 | （可选）为固定集合指定一个最大值，即字节数。 **如果 capped 为 true，也需要指定该字段。** |
| max         | 数值 | （可选）指定固定集合中包含文档的最大数量。                   |

在插入文档时，MongoDB 首先检查固定集合的 `size` 字段，然后检查 `max` 字段。

## 3.1创建集合实例

```
> use test   # 打开test数据库
switched to db test
> db.createCollection("golang")  # 创建集合golang
{ "ok" : 1 }
> show collections   # 显示集合，也可以使用 show tables
golang
```

## 3.2显示集合

```
show collections
show tables
```

## 3.3删除集合

**语法格式：**

```
db.collection.drop()
```

**实例**

```
db.golang.drop()  # 删除集合
true              # 删除成功
show collections  # 显示集合
```



# 4.MongoDB文档

MongoDB 将数据记录存储为 BSON 文档。BSON 是JSON文档的**二进制**表示，尽管它包含比 JSON 更多的数据类型。

![doc](https://myresou.oss-cn-shanghai.aliyuncs.com/img/11J8B8KJ0CAR88GMEKH3TO9JKQ.png)

## BSON数据类型

| 类型                       | 序号 | 别名                  | 备注  |
| :------------------------- | :--- | :-------------------- | :---- |
| Double                     | 1    | "double"              |       |
| String                     | 2    | "string"              |       |
| Object                     | 3    | "object"              |       |
| Array                      | 4    | "array"               |       |
| Binary data                | 5    | "binData"             |       |
| Undefined                  | 6    | "undefined"           | 废弃. |
| ObjectId                   | 7    | "objectId"            |       |
| Boolean                    | 8    | "bool"                |       |
| Date                       | 9    | "date"                |       |
| Null                       | 10   | "null"                |       |
| Regular Expression         | 11   | "regex"               |       |
| DBPointer                  | 12   | "dbPointer"           | 废弃. |
| JavaScript                 | 13   | "javascript"          |       |
| Symbol                     | 14   | "symbol"              | 废弃. |
| JavaScript code with scope | 15   | "javascriptWithScope" | 废弃. |
| 32-bit integer             | 16   | "int"                 |       |
| Timestamp                  | 17   | "timestamp"           |       |
| 64-bit integer             | 18   | "long"                |       |
| Decimal128                 | 19   | "decimal"             |       |
| Min key                    | -1   | "minKey"              |       |
| Max key                    | 127  | "maxKey"              |       |

## 文档结构

MongoDB 文档由字段-值对组成，具有以下结构：

```json
{
   field1: value1,
   field2: value2,
   field3: value3,
   ...
   fieldN: valueN
}
```

字段的值可以是任何 BSON数据类型，包括其他文档、数组和文档数组。例如，以下文档包含不同类型的值：

```
var mydoc = {
               _id: ObjectId("5099803df3f4948bd2f98391"),
               name: { first: "Alan", last: "Turing" },
               birth: new Date('Jun 23, 1912'),
               death: new Date('Jun 07, 1954'),
               contribs: [ "Turing machine", "Turing test", "Turingery" ],
               views : NumberLong(1250000)
            }
```

上述字段具有以下数据类型：

- `_id`持有一个ObjectId。
- `name`包含一个*嵌入文档*，其中包含字段 `first`和`last`.
- `birth`并`death`保存*Date*类型的值。
- `contribs`包含一个*字符串数组*。
- `views`保存*NumberLong*类型的值。

## 数组

要通过从零开始的索引位置指定或访问数组的元素，请将数组名称与点 ( `.`) 和从零开始的索引位置连接起来，并用引号引起来：

```
"<array>.<index>"
```

例如，给定文档中的以下字段：

```
{
   ...
   langs: [ "java", "python", "golang" ],
   ...
}
```

要指定数组中的第三个元素`langs`，请使用点表示法`"langs.2"`。

## 嵌入式文档

要使用点表示法指定或访问嵌入文档的字段，请将嵌入文档名称与点 ( `.`) 和字段名称连接起来，并用引号引起来：

```
"<embedded document>.<field>"
```

例如，给定文档中的以下字段：

```
{
   ...
   name: { first: "guo", last: "hongzhi" },
   contact: { phone: { type: "cell", number: "15011039899" } },
   ...
}
```

- 要指定字段中命名`last`的`name`字段，请使用点表示法`"name.last"`。
- 要`number`在`phone`文档中的 `contact`字段中指定 ，请使用点符号`"contact.phone.number"`。

# 5.MongoDB插入文档

## 插入单个文档

**语法**

```
db.collection.insertOne()
```

**实例**

```json
db.articles.insertOne({title:"golang技术栈",views:1000,tags:["golang","专栏","教程"],author:{id:1,name:"yj",address:"江苏·苏州"}})

// 执行结果
{
        "acknowledged" : true,
        "insertedId" : ObjectId("653f193e2046750e0529db93")
}
// 查询
db.articles.find()
// 查询结果
{ "_id" : ObjectId("62d20c8777f0333dd8628a7c"), "title" : "golang技术", "views" : 1000, "tags" : [ "golang", "专栏", "教程" ], "author" : { "id" : 1, "name" : "老郭", "address" : "北京" } }
>
```

> 集合不存在创建集合，`articles`是集合名称
>
> 对象id自动生成

## 插入多个文档

**语法**

```
db.collection.insertMany()
```

**实例**

```json
db.articles.insertMany([
... {title: "golang技术", views: 1000, tags: ["golang","专栏", "教程"], author: {id:1, name: "老郭", address: "北京"}},
	{title: "java技术", views: 800, tags: ["java","spring", "教程"], author: {id:2, name: "张三", address: "上海"}},
	{title: "python技术", views: 600, tags: ["python","ai", "教程"], author: {id:3, name: "李四", address: "深圳"}},
]
... )

// 查询
db.articles.find()
```

## 其他插入方法

- `db.collection.updateOne()` 当 `upsert: true` 时.
- `db.collection.updateMany()` 当 `upsert: true` 时.
- `db.collection.findAndModify()` 当 `upsert: true` 时.
- `db.collection.findOneAndUpdate()` 当 `upsert: true` 时.
- `db.collection.findOneAndReplace()` 当 `upsert: true` 时.
- `db.collection.bulkWrite()`.

# 6.MongoDB查询文档

## 插入一些文档

```
db.inventory.insertMany([
   { item: "journal", qty: 25, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
   { item: "notebook", qty: 50, size: { h: 8.5, w: 11, uom: "in" }, status: "A" },
   { item: "paper", qty: 100, size: { h: 8.5, w: 11, uom: "in" }, status: "D" },
   { item: "planner", qty: 75, size: { h: 22.85, w: 30, uom: "cm" }, status: "D" },
   { item: "postcard", qty: 45, size: { h: 10, w: 15.25, uom: "cm" }, status: "A" }
]);
```

> 一些商品库存(inventory)信息：名称、数量、大小、状态。

## 查询所有文档

```
db.inventory.find( {} )
```

类似sql语句

```sql
SELECT * FROM inventory
```

## 指定条件

```
db.inventory.find( { status: "D" } )
```

类似sql语句

```sql
SELECT * FROM inventory WHERE status = "D"
```

## 使用`in` `and` `or`等

### `in`

```
db.inventory.find( { status: { $in: [ "A", "D" ] } } )
```

类似sql

```sql
SELECT * FROM inventory WHERE status in ("A", "D")
```

### `and`

```
db.inventory.find( { status: "A", qty: { $lt: 30 } } )
```

类似sql

```sql
SELECT * FROM inventory WHERE status = "A" AND qty < 30
```

### `or`

```
db.inventory.find( { $or: [ { status: "A" }, { qty: { $lt: 30 } } ] } )
```

类似sql

```sql
SELECT * FROM inventory WHERE status = "A" OR qty < 30
```

### `and`和`or`

```
db.inventory.find( {
     status: "A",
     $or: [ { qty: { $lt: 30 } }, { item: /^p/ } ]
} )
```

类似sql

```sql
SELECT * FROM inventory WHERE status = "A" AND ( qty < 30 OR item LIKE "p%")
```



# 7.MongoDB更新文档

## 准备数据

```
db.inventory.insertMany( [
   { item: "canvas", qty: 100, size: { h: 28, w: 35.5, uom: "cm" }, status: "A" },
   { item: "journal", qty: 25, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
   { item: "mat", qty: 85, size: { h: 27.9, w: 35.5, uom: "cm" }, status: "A" },
   { item: "mousepad", qty: 25, size: { h: 19, w: 22.85, uom: "cm" }, status: "P" },
   { item: "notebook", qty: 50, size: { h: 8.5, w: 11, uom: "in" }, status: "P" },
   { item: "paper", qty: 100, size: { h: 8.5, w: 11, uom: "in" }, status: "D" },
   { item: "planner", qty: 75, size: { h: 22.85, w: 30, uom: "cm" }, status: "D" },
   { item: "postcard", qty: 45, size: { h: 10, w: 15.25, uom: "cm" }, status: "A" },
   { item: "sketchbook", qty: 80, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
   { item: "sketch pad", qty: 95, size: { h: 22.85, w: 30.5, uom: "cm" }, status: "A" }
] );
```

> 一些商品库存(inventory)信息：名称、数量、大小、状态。

## 更新单个文档

```
db.inventory.updateOne(
   { item: "paper" },
   {
     $set: { "size.uom": "cm", status: "P" },
     $currentDate: { lastModified: true }
   }
)
```

> `$set` 和 `$currentDate`是字段更新操作符
>
> `$currentDate` 用于更新匹配条件为{ item: "paper" }的文档。它会将 "size.uom" 字段的值更新为 "cm"，并将 "status" 字段的值更新为 "P"。同时，它还会创建一个新的字段 "lastModified"，并将其值设置为当前的日期和时间。

## 更新多个文档

```
db.inventory.updateMany(
   { "qty": { $lt: 50 } },
   {
     $set: { "size.uom": "in", status: "P" },
     $currentDate: { lastModified: true }
   }
)
```

## 替换更新

```
db.inventory.replaceOne(
   { item: "paper" },
   { item: "paper", instock: [ { warehouse: "A", qty: 60 }, { warehouse: "B", qty: 40 } ] }
)
```



# 8.MongoDB删除文档

## 准备数据

```
db.inventory.insertMany( [
   { item: "journal", qty: 25, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
   { item: "notebook", qty: 50, size: { h: 8.5, w: 11, uom: "in" }, status: "P" },
   { item: "paper", qty: 100, size: { h: 8.5, w: 11, uom: "in" }, status: "D" },
   { item: "planner", qty: 75, size: { h: 22.85, w: 30, uom: "cm" }, status: "D" },
   { item: "postcard", qty: 45, size: { h: 10, w: 15.25, uom: "cm" }, status: "A" },
] );
```

> 一些商品库存(inventory)信息：名称、数量、大小、状态。

## 删除所有文档

```
db.inventory.deleteMany({})
```

## 条件删除

```
db.inventory.deleteMany({ status : "A" })
```

## 删除单个文档

```
db.inventory.deleteOne( { status: "D" } )
```



# 9.golang操作MongoDB

下载地址：

```
https://www.mongodb.com/download-center/community
```

打开客户端

```
mongo.exe
```

创建数据库

```
use go_db;
```

创建集合

```
 db.createCollection("student");
```

下载驱动

```
go get github.com/mongodb/mongo-go-driver
```

连接mongoDB

```go
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)


var client *mongo.Client

func initDB()  {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func main() {
	initDB()
}
```

运行结果

```
Connected to MongoDB!
```



## 9.1 go-添加文档

**创建一个结构体**

```go
type Student struct {
	Name string
	Age int
}
```

**添加单个文档**

使用`collection.InsertOne()`方法插入一条文档记录：

```go
func insertOne(s Student) {
	initDB()
	collection := client.Database("go_db").Collection("student")
	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
```

**测试**

```go
func main() {
	s := Student{Name: "tom", Age: 20}
	insertOne(s)
}
```

运行结果

```
Connected to MongoDB!
Inserted a single document:  ObjectID("61124558682f5c9583330222")
```

客户端查看

```
mongodb 打开客户端
use go_db
db.student.find()
db.student.remove({}) // 删除所有
```

**插入多个文档**

使用`collection.InsertMany()`方法插入多条文档记录：

```go
func insertMore(students []interface{}) {
	//students := []interface{}{s2, s3}
	initDB()
	collection := client.Database("go_db").Collection("student")
	insertManyResult, err := collection.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
```

**测试**

```go
func main() {
	s := Student{Name: "tom", Age: 20}
	s1 := Student{Name: "kite", Age: 21}
	s2 := Student{Name: "rose", Age: 22}
	students := []interface{}{s, s1, s2}
	insertMore(students)
}
```

运行结果

```
Connected to MongoDB!
Inserted multiple documents:  [ObjectID("611246c56637c3554426bc92") ObjectID("611246c56637c3554426bc93") ObjectID("
```

## 9.2 go-查找文档

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
func initDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	co := options.Client().ApplyURI("mongodb://localhost:27017")
	mongo.Connect(context.TODO(), co)
	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	client.Ping(context.TODO(), nil)
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func find() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("go_db").Collection("student")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)
		marshalByte, err := bson.Marshal(result)
		err = bson.Unmarshal(marshalByte, &s)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}



func main() {
	initDB()
	find()
}
```

根据条件查询文档

and

```go
func findWithAnd() {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    collection := global.Client.Database("go_db").Collection("student")
    
    filter := bson.D{
        {"$and", bson.A{
            bson.D{{"name", "tim"}},
            bson.D{{"age", 20}},
        }},
    }
    
    var students []Student
    cur, err := collection.Find(ctx, filter)
    if err != nil {
        log.Fatal(err)
    }
    defer cur.Close(ctx)
    
    for cur.Next(ctx) {
        var student Student
        err := cur.Decode(&student)
        if err != nil {
            log.Fatal(err)
        }
        
        students = append(students, student)
    }
    
    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("results: %v\n", students)
}
```

or

```go
func findWithOr() {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    collection := global.Client.Database("go_db").Collection("student")
    
    filter := bson.M{
        "$or": []bson.M{
            bson.M{"name": "tim"},
            bson.M{"age": 20},
        },
    }
    
    var students []Student
    cur, err := collection.Find(ctx, filter)
    if err != nil {
        log.Fatal(err)
    }
    defer cur.Close(ctx)
    
    for cur.Next(ctx) {
        var student Student
        err := cur.Decode(&student)
        if err != nil {
            log.Fatal(err)
        }
        
        students = append(students, student)
    }
    
    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("results: %v\n", students)
}
```



## 9.3 go-更新文档

```go
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string
	Age  int
}

var client *mongo.Client

func initDb() {
	co := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.Connect(context.TODO(), co)
	if err != nil {
		log.Fatal(err)
	}

	err2 := c.Ping(context.TODO(), nil)
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println("连接成功！")
	}
	client = c
}

func update() {
	ctx := context.TODO()
	defer client.Disconnect(ctx)
	c := client.Database("go_db").Collection("Student")

	update := bson.D{{"$set", bson.D{{"Name", "big tom"}, {"Age", 22}}}}

	ur, err := c.UpdateMany(ctx, bson.D{{"name", "tom"}}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", ur.ModifiedCount)
}

func main() {
	initDb()
	update()
}
```

## 9.4 go-删除文档

```go
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string
	Age  int
}

var client *mongo.Client

func initDb() {
	co := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.Connect(context.TODO(), co)
	if err != nil {
		log.Fatal(err)
	}

	err2 := c.Ping(context.TODO(), nil)
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println("连接成功！")
	}
	client = c
}

func del() {

	initDb()
	c := client.Database("go_db").Collection("Student")
	ctx := context.TODO()

	dr, err := c.DeleteMany(ctx, bson.D{{"Name", "big kite"}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", dr.DeletedCount)

}

func main() {
	del()
}
```

