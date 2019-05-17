package item

/*	CREATE TABLE `tb_item` (
	`id` bigint(20) NOT NULL COMMENT '商品id，同时也是商品编号',
	`title` varchar(100) NOT NULL COMMENT '商品标题',
	`sell_point` varchar(500) DEFAULT NULL COMMENT '商品卖点',
	`price` bigint(20) NOT NULL COMMENT '商品价格，单位为：分',
	`num` int(10) NOT NULL COMMENT '库存数量',
	`barcode` varchar(30) DEFAULT NULL COMMENT '商品条形码',
	`image` varchar(500) DEFAULT NULL COMMENT '商品图片',
	`cid` bigint(10) NOT NULL COMMENT '所属类目，叶子类目',
	`status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '商品状态，1-正常，2-下架，3-删除',
	`created` datetime NOT NULL COMMENT '创建时间',
	`updated` datetime NOT NULL COMMENT '更新时间',
	PRIMARY KEY (`id`),
	KEY `cid` (`cid`),
	KEY `status` (`status`),
	KEY `updated` (`updated`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品表';*/

type TbItem struct {
	Id        int64
	Title     string
	SellPoint string
	Price     int
	Num       int
	Barcode   string
	Image     string
	Cid       int
	Status    int8
	Create    string
	Update    string
}

//给页面使用，实现商品类目
type TbItemChild struct {
	TbItem
	CategoryName string
}
