DROP TABLE IF EXISTS `students`;

CREATE TABLE `students` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(150) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `class` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `students` */

insert  into `students`(`id`,`name`,`age`,`class`) values 
(1,'Jason Bourne',29,'10A'),
(2,'James Bond',27,'10B'),
(3,'Ethan Hunt',27,'10A'),
(4,'John Wick',28,'10C');
