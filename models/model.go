package models

import (
	"bytes"
	"fclink.cn/ethcoldwallet/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type Model struct {
	CreatedAt time.Time `gorm:"column:created" json:"created"`
	UpdatedAt time.Time `gorm:"column:updated" json:"updated"`
}

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

// db init
func Setup() {
	var err error
	c := conf.Context.Db

	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.Username, c.Password, c.Address, c.Port, c.Dbname)

	db, err = gorm.Open("mysql", args)

	if err != nil {
		log.Fatalln(err)
	}
	db.DB().SetMaxIdleConns(c.Maxidleconns)
	db.DB().SetMaxOpenConns(c.Maxopenconns)
	db.DB().SetConnMaxLifetime(c.Maxlifetime * time.Minute)
	db.DB().Ping()
	db.SingularTable(true)
	db.LogMode(c.LogMode)
}

type DeBcBlockTx struct {
	TxHash      string `gorm:"column:TX_HASH;type:varchar(128);primary_key;" json:"tx_hash"`
	BlockHash   string `gorm:"column:BLOCK_HASH;type:varchar(128);" json:"block_hash"`
	BlockNumber string `gorm:"column:BLOCK_NUMBER;type:varchar(64);" json:"block_number"`
	TxValue     string `gorm:"column:TX_VALUE;type:varchar(64);" json:"tx_value"`
	TokenCount  string `gorm:"column:TOKEN_COUNT;type:varchar(64);" json:"omitempty"`

	TxFrom        string    `gorm:"column:TX_FROM;type:varchar(64);" json:"tx_from"`
	TxTo          string    `gorm:"column:TX_TO;type:varchar(64);" json:"tx_to"`
	ContractAddr  string    `gorm:"column:CONTRACT_ADDR;type:varchar(200);" json:"contract_addr"`
	CreatedTime   time.Time `gorm:"column:CREATED_TIME;type:datetime;" json:"omitempty"`
	ModifiedTime  time.Time `gorm:"column:MODIFIED_TIME;type:datetime;" json:"omitempty"`
	Created       int64     `json:"create_time"`
	TxStatus      string    `gorm:"column:TX_STATUS;type:varchar(10);" json:"tx_status"`
	OmnilayerType string    `gorm:"column:OMNILAYER_TYPE;type:varchar(10);" json:"omnilayer_type"`
}
type DeAct2cFund struct {
	FundNo  string `gorm:"column:fund_no;type:varchar(128);primary_key`
	Address string `gorm:"column:address;type:varchar(128);`
}

func FindByAddress(op, address, contractAddr string, page, size int) []DeBcBlockTx {
	data := make([]DeBcBlockTx, 0)
	if op == "1" {
		sql := db.Model(&DeBcBlockTx{}).Where("tx_from=?", address)
		if len(contractAddr) > 0 {
			sql = sql.Where("contract_addr=?", contractAddr)
		} else {
			sql = sql.Where("omnilayer_type='ETH'")
		}
		sql.Order("created_time desc").Select("*,UNIX_TIMESTAMP(created_time)*1000 created").Offset(page * size).Limit(size).Scan(&data)
		//sql.Find(&data).Limit(10)
	} else if op == "2" {
		sql := db.Model(&DeBcBlockTx{}).Where("tx_to=?", address)
		if len(contractAddr) > 0 {
			sql = sql.Where("contract_addr=?", contractAddr)
		} else {
			sql = sql.Where("omnilayer_type='ETH'")
		}
		sql.Order("created_time desc").Select("*,UNIX_TIMESTAMP(created_time)*1000 created").Offset(page * size).Limit(size).Scan(&data)
	} else if op == "0" {

		var buffer bytes.Buffer
		buffer.WriteString("select *,UNIX_TIMESTAMP(created_time)*1000 created from (")
		buffer.WriteString(fmt.Sprintf("select * from de_bc_block_tx where tx_from='%s'", address))
		if len(contractAddr) > 0 {
			buffer.WriteString(fmt.Sprintf(" and contract_addr='%s'", contractAddr))
		} else {
			buffer.WriteString("and omnilayer_type='ETH'")
		}
		buffer.WriteString(" union ")
		buffer.WriteString(fmt.Sprintf("select * from de_bc_block_tx where tx_to='%s'", address))
		if len(contractAddr) > 0 {
			buffer.WriteString(fmt.Sprintf(" and contract_addr='%s'", contractAddr))
		} else {
			buffer.WriteString("and omnilayer_type='ETH'")
		}
		buffer.WriteString(fmt.Sprintf(") a order by a.created_time desc limit %d,%d", page*size, size))
		db.Model(&DeBcBlockTx{}).Raw(buffer.String()).Scan(&data)
	}
	return data
}

func CreateFund(address string) error {
	// User not found, create a new record with give conditions
	data := DeAct2cFund{
		Address: address, FundNo: address,
	}
	return db.FirstOrCreate(&data, DeAct2cFund{FundNo: address}).Error
}
