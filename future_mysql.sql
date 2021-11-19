-- drop DATABASE evolve
CREATE DATABASE /*!32312 IF NOT EXISTS*/`evolve` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE evolve;
SET time_zone='+8:00';
FLUSH PRIVILEGES;
#SHOW VARIABLES LIKE '%zone%';

DROP TABLE IF EXISTS role;
CREATE TABLE role (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL COMMENT '角色名称',
  description VARCHAR(255) NULL COMMENT '角色描述',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='角色表';
 INSERT INTO role(id,name,description) VALUES (1,'管理员','hehe'),(31,'11','1122'),(35,'33','44'),(42,'222','333'),(45,'44','22'),(46,'44','44'),(47,'123','123'),(53,'123','123'),(54,'123','123'),(55,'222','222');

DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(255) NOT NULL COMMENT '(邮箱)账号',
  salt VARCHAR(255) NOT NULL COMMENT '盐',
  pwd VARCHAR(255) DEFAULT NULL COMMENT '密码',
  name VARCHAR(16) DEFAULT NULL COMMENT '名称',
  mobile VARCHAR(16) DEFAULT NULL COMMENT '电话',
  lastOn DATETIME DEFAULT NULL COMMENT '最后登录时间',
  access VARCHAR(2000) DEFAULT NULL COMMENT '权限',
  avatar VARCHAR(255) DEFAULT NULL COMMENT '头像',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='用户表';

 INSERT INTO user(id,email,salt,pwd,name,mobile,lastOn) VALUES (3,'453017973@qq.com','e18f5b2c-2df9-4d4d-aa96-b0edd7b2dd8b','86950062ba25c65a2695c3193f9c8bfa83587bc906e0bbe3e4b4e4f1ccec334b','super_admin','13100000000','2018-12-17 09:36:32'),(5,'yu-liu@qulv.com','e18f5b2c-2df9-4d4d-aa96-b0edd7b2dd8b','86950062ba25c65a2695c3193f9c8bfa83587bc906e0bbe3e4b4e4f1ccec334b','yu-liu@qulv.com','','2018-10-16 10:35:02'),(7,'59862@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','59862@qulv.com',NULL,NULL),(8,'89004@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','89004@qulv.com',NULL,NULL),(9,'65435@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','65435@qulv.com',NULL,NULL),(10,'60165@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','60165@qulv.com',NULL,NULL),(11,'54519@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','54519@qulv.com',NULL,NULL),(12,'92099@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','92099@qulv.com',NULL,NULL),(13,'96938@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','96938@qulv.com',NULL,NULL),(14,'58393@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','58393@qulv.com',NULL,NULL),(15,'51155@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','51155@qulv.com',NULL,NULL),(16,'80593@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','80593@qulv.com',NULL,NULL),(17,'99501@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','99501@qulv.com',NULL,NULL),(18,'55725@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','55725@qulv.com',NULL,NULL),(19,'80125@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','80125@qulv.com',NULL,NULL),(20,'83451@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','83451@qulv.com',NULL,NULL),(21,'76878@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','76878@qulv.com',NULL,NULL),(22,'84038@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','84038@qulv.com',NULL,NULL),(23,'89558@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','89558@qulv.com',NULL,NULL),(24,'95678@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','95678@qulv.com',NULL,NULL),(25,'59714@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','59714@qulv.com',NULL,NULL),(26,'61537@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','61537@qulv.com',NULL,NULL),(27,'78545@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','78545@qulv.com',NULL,NULL),(28,'58113@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','58113@qulv.com',NULL,NULL),(29,'54929@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','54929@qulv.com',NULL,NULL),(30,'50310@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','50310@qulv.com',NULL,NULL),(31,'86762@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','86762@qulv.com',NULL,NULL),(32,'82879@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','82879@qulv.com',NULL,NULL),(33,'54112@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','54112@qulv.com',NULL,NULL),(34,'71923@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','71923@qulv.com',NULL,NULL),(35,'97277@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','97277@qulv.com',NULL,NULL),(36,'70619@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','70619@qulv.com',NULL,NULL),(37,'61265@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','61265@qulv.com',NULL,NULL),(38,'94468@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','94468@qulv.com',NULL,NULL),(39,'88546@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','88546@qulv.com',NULL,NULL),(40,'59325@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','59325@qulv.com',NULL,NULL),(41,'80986@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','80986@qulv.com',NULL,NULL),(42,'76957@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','76957@qulv.com',NULL,NULL),(43,'91828@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','91828@qulv.com',NULL,NULL),(44,'78271@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','78271@qulv.com',NULL,NULL),(45,'65870@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','65870@qulv.com',NULL,NULL),(46,'94538@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','94538@qulv.com',NULL,NULL),(47,'75081@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','75081@qulv.com',NULL,NULL),(48,'91791@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','91791@qulv.com',NULL,NULL),(49,'83710@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','83710@qulv.com',NULL,NULL),(50,'93181@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','93181@qulv.com',NULL,NULL),(51,'64775@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','64775@qulv.com',NULL,NULL),(52,'94334@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','94334@qulv.com',NULL,NULL),(53,'77342@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','77342@qulv.com',NULL,NULL),(54,'53710@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','53710@qulv.com',NULL,NULL),(55,'86526@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','86526@qulv.com',NULL,NULL),(56,'71500@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','71500@qulv.com',NULL,NULL),(57,'97921@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','97921@qulv.com',NULL,NULL),(58,'75108@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','75108@qulv.com',NULL,NULL),(59,'81775@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','81775@qulv.com',NULL,NULL),(60,'83551@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','83551@qulv.com',NULL,NULL),(61,'72430@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','72430@qulv.com',NULL,NULL),(62,'61497@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','61497@qulv.com',NULL,NULL),(63,'90197@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','90197@qulv.com',NULL,NULL),(64,'66492@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','66492@qulv.com',NULL,NULL),(65,'61871@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','61871@qulv.com',NULL,NULL),(66,'59881@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','59881@qulv.com',NULL,NULL),(67,'63793@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','63793@qulv.com',NULL,NULL),(68,'89320@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','89320@qulv.com',NULL,NULL),(69,'55223@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','55223@qulv.com',NULL,NULL),(70,'58155@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','58155@qulv.com',NULL,NULL),(71,'75105@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','75105@qulv.com',NULL,NULL),(72,'51062@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','51062@qulv.com',NULL,NULL),(73,'79995@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','79995@qulv.com',NULL,NULL),(74,'96789@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','96789@qulv.com',NULL,NULL),(75,'93961@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','93961@qulv.com',NULL,NULL),(76,'79436@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','79436@qulv.com',NULL,NULL),(77,'65301@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','65301@qulv.com',NULL,NULL),(78,'88198@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','88198@qulv.com',NULL,NULL),(79,'95086@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','95086@qulv.com',NULL,NULL),(80,'60835@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','60835@qulv.com',NULL,NULL),(81,'68917@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','68917@qulv.com',NULL,NULL),(82,'62083@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','62083@qulv.com',NULL,NULL),(83,'53665@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','53665@qulv.com',NULL,NULL),(84,'82074@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','82074@qulv.com',NULL,NULL),(85,'99378@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','99378@qulv.com',NULL,NULL),(86,'50668@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','50668@qulv.com',NULL,NULL),(87,'55208@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','55208@qulv.com',NULL,NULL),(88,'74034@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','74034@qulv.com',NULL,NULL),(89,'54547@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','54547@qulv.com',NULL,NULL),(90,'50633@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','50633@qulv.com',NULL,NULL),(91,'89524@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','89524@qulv.com',NULL,NULL),(92,'95724@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','95724@qulv.com',NULL,NULL),(93,'60048@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','60048@qulv.com',NULL,NULL),(94,'63070@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','63070@qulv.com',NULL,NULL),(95,'85204@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','85204@qulv.com',NULL,NULL),(96,'86810@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','86810@qulv.com',NULL,NULL),(97,'78437@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','78437@qulv.com',NULL,NULL),(98,'81759@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','81759@qulv.com',NULL,NULL),(99,'73482@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','73482@qulv.com',NULL,NULL),(100,'72134@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','72134@qulv.com',NULL,NULL),(101,'90223@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','90223@qulv.com',NULL,NULL),(102,'84716@qulv.com','d6255e12-54c3-48b9-9686-684ad52c0613','F89E98869DC27441E13F078EAC20A31B','84716@qulv.com',NULL,NULL);

DROP TABLE IF EXISTS user_role_map;
CREATE TABLE user_role_map (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL COMMENT '用户ID',
  role_id BIGINT NOT NULL COMMENT '角色ID'
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='角色用户映射表';

 INSERT INTO user_role_map(id,user_id,role_id) VALUES(1,3,1);


DROP TABLE IF EXISTS message;
CREATE TABLE message (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `from` VARCHAR(255) NOT NULL COMMENT '发送人',
  title  VARCHAR(255) NULL COMMENT '标题',
  content TEXT DEFAULT NULL COMMENT '内容',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='消息表';


DROP TABLE IF EXISTS message_to;
CREATE TABLE message_to (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  message_id BIGINT NOT NULL,
  `to` VARCHAR(255) NOT NULL COMMENT '接收人',
  status INT DEFAULT NULL COMMENT '状态',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='消息映射表';


