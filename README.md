# Message_Queuing_System

### Requirement :- 
Design and implement (with tests) a _message queuing system_ using Go programming
language and RabbitMQ or Kafka.

### Tools or frameworks used:-

    Mysql
    RabbitMQ
    Go
    Gin

### Details :- 

#### Schema :- 
```mysql
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| zocket_db          |
+--------------------+


+---------------------+
| Tables_in_zocket_db |
+---------------------+
| Image_Store         |
| Products            |
| Users               |
+---------------------+
```

**Users**
```mysql
+------------+--------------+------+-----+-------------------+-----------------------------------------------+
| Field      | Type         | Null | Key | Default           | Extra                                         |
+------------+--------------+------+-----+-------------------+-----------------------------------------------+
| id         | int          | NO   | PRI | NULL              | auto_increment                                |
| name       | varchar(255) | YES  |     | NULL              |                                               |
| mobile     | varchar(255) | YES  |     | NULL              |                                               |
| latitude   | varchar(255) | YES  |     | NULL              |                                               |
| longitude  | varchar(255) | YES  |     | NULL              |                                               |
| created_at | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| updated_at | timestamp    | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+------------+--------------+------+-----+-------------------+-----------------------------------------------+
```

**Products**
```mysql
+---------------------+---------------+------+-----+-------------------+-----------------------------------------------+
| Field               | Type          | Null | Key | Default           | Extra                                         |
+---------------------+---------------+------+-----+-------------------+-----------------------------------------------+
| product_id          | int           | NO   | PRI | NULL              | auto_increment                                |
| product_name        | varchar(255)  | YES  |     | NULL              |                                               |
| product_description | text          | YES  |     | NULL              |                                               |
| product_price       | decimal(10,3) | YES  |     | NULL              |                                               |
| created_at          | timestamp     | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| updated_at          | timestamp     | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+---------------------+---------------+------+-----+-------------------+-----------------------------------------------+
```

**Image_Store**
```Mysql
+---------------------------+------+------+-----+---------+-------+
| Field                     | Type | Null | Key | Default | Extra |
+---------------------------+------+------+-----+---------+-------+
| product_id                | int  | YES  |     | NULL    |       |
| product_images            | text | YES  |     | NULL    |       |
| compressed_product_images | text | YES  |     | NULL    |       |
+---------------------------+------+------+-----+---------+-------+
```

#### How to Run?

>To run Api with Producer use  **go run main.go** and to terminate use **control + c**

>To run Consumer use **go run ./consumer/main.go ./consumer/consumer.go** and to terminate use **control + c**


**Note:-** 
1. Here Image_store used to store array of product_images comes from Api and compressed_product_images path.
2. I've tested in Mac OS. To resolve dependency issue please follow below steps to resolve.

  To install the dependencies using brew do:
  
    brew install pngquant
    brew install mozjpeg