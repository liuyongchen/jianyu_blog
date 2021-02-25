package model

import (
	"fmt"

	"blog-service/global"
	"blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CREATE DATABASE
// IF
//     NOT EXISTS blog_service DEFAULT CHARACTER
//     SET utf8mb4 	DEFAULT COLLATE utf8mb4_general_ci;
//
// `created_on` INT(10) unsigned DEFAULT '0'  COMMENT '创建时间'，
// `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
// `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
// `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
// `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
// `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1为已删除',

// 2、创建标签表
// CREATE TABLE `blog_tag` (
// `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
// `name` varchar(100) DEFAULT '' COMMENT '标签名称',
// `created_on` INT(10) unsigned DEFAULT '0'  COMMENT '创建时间',
// `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
// `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
// `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
// `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
// `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1为已删除',
// `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用，1为启用',
//  PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

// 3、创建文章表
// CREATE TABLE `blog_article` (
// `id` int(10) unsigned  NOT NULL AUTO_INCREMENT,
// `title` varchar(100) DEFAULT '' COMMENT '文章标题',
// `desc` varchar(255) DEFAULT '' COMMENT '文章简述',
// `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
// `content` longtext COMMENT '文章内容',
// `created_on` INT(10) unsigned DEFAULT '0'  COMMENT '创建时间',
// `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
// `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
// `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
// `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
// `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1为已删除',
// `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
// PRIMARY KEY(`id`)
// ) ENGINE = InnoDB DEFAULT CHARSET =utf8mb4 COMMENT = '文章管理'
//
// 4、创建文章标签关联表
// CREATE TABLE `blog_article_tag` (
// `id` int(10) unsigned  NOT NULL AUTO_INCREMENT,
// `article_id` int(11) NOT NULL COMMENT '文章ID',
// `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
// `created_on` INT(10) unsigned DEFAULT '0'  COMMENT '创建时间',
// `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
// `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
// `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
// `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
// `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1为已删除',
// PRIMARY KEY (`id`)
// ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT ='文章标签关联'
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreateBy   string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)

	db, err := gorm.Open(databaseSetting.DBType, s)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}
