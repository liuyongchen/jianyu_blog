package model

import "github.com/jinzhu/gorm"

//CREATE TABLE `blog_auth` (
//	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
// `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
//	`app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
//	#此处请写入公共字段
//	`created_on` INT(10) unsigned DEFAULT '0'  COMMENT '创建时间'，
//	`created_by` varchar(100) DEFAULT '' COMMENT '创建人',
//	`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
//	`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
//	`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
//	`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1为已删除',
// PRIMARY KEY (`id`) USING BTREE
//) ENGIN=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT="认证管理";
//
//INSERT INTO `blog_service`.`blog_auth`(
//	`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`,
//	`modified_by`, `deleted_on`, `is_del`)
//	VALUES (1, 'eddycjy', 'go-programming-tour-book', 0, 'eddycjy', 0,
// '', 0, 0)

// 1、Header
// alg: 签名算法(HMAC、SHA256、RSA)
// typ: 令牌类型(JWT)
// 2、Payload
// aud(audience):受众，即接受JWT的一方
// exp(ExpiresAt):所签发的JWT过期时间，过期时间必须大于签发时间
// jti(JWTId):JWT的唯一标识
// iat(IssuedAt):签发时间
// iss(Issuer):JWT的签发者
// nbf(NotBefore):JWT的生效时间，如果未到这个时间，则不可用
// sub(Subject):主题
// 3、Signature
// HMACSHA256(
// base64UrlEncode(header) + "." +
// base64UrlEncode(payload),
// secret)
//

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}
