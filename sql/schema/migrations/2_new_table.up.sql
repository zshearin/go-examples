CREATE TABLE new_table
(
    random_content TEXT NOT NULL,
    other_stuff TEXT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `new_table` VALUES ("this is random content", "this is other stuff"),
( "this is also random content", "this is also other stuff");