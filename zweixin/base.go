/*
微信包初始化
*/
package zweixin

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zouhuigang/package/zhttp"
	"github.com/zouhuigang/package/ztime"
)

const API_URL = "https://api.weixin.qq.com"
const API_URL_SH = "https://api.weixin.qq.com"
const API_URL_SZ = "https://api.weixin.qq.com"
const API_URL_HK = "https://api.weixin.qq.com"

type Zweixin struct {
	appid       string
	appsecret   string
	mysqlHost   string
	mysqlPort   int
	mysqlDB     string
	mysqlPasswd string
	db          *sql.DB
}

func NewZweixin(appid, appsecret string, mysqlHost string, mysqlPort int, mysqlDB string, mysqlPasswd string) (*Zweixin, error) {
	wx := &Zweixin{
		appid:       appid,
		appsecret:   appsecret,
		mysqlHost:   mysqlHost,
		mysqlPort:   mysqlPort,
		mysqlDB:     mysqlDB,
		mysqlPasswd: mysqlPasswd,
	}

	var err error
	wx.db, err = wx.GetConn() //初始化数据库
	if err != nil {
		return nil, err
	}
	return wx, nil
}

func (wx *Zweixin) GetConn() (*sql.DB, error) {
	//配置数据库连接地址.统一配置.
	conn := fmt.Sprintf("root:%s@tcp(%s:%d)/%s?charset=utf8", wx.mysqlPasswd, wx.mysqlHost, wx.mysqlPort, wx.mysqlDB)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Printf("failed to open database:%s\n", err.Error())
		return nil, err
	}

	db.SetMaxIdleConns(100)

	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		fmt.Printf("Error on opening database connection: %s\n", err.Error())
		return nil, err
	}
	return db, nil
}

func (wx *Zweixin) GetAccessToken() string {
	db := wx.db

	//查询数据
	rows, err := db.Query("SELECT access_token,expires_in,mtime FROM wx_js_token where id=1")
	checkErr(err)
	var expires_in int
	var atoken string
	var mtime int64
	for rows.Next() {
		err := rows.Scan(&atoken, &expires_in, &mtime)
		checkErr(err)
	}

	nowtimestamp := ztime.NowTimeStamp()
	if nowtimestamp < (int64(expires_in) + mtime) {
		return atoken
	}

	callback := &access_token{}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", wx.appid, wx.appsecret)
	zhttp.GETWithUnmarshal(url, callback)

	//修改数据
	stmt, err := db.Prepare("update wx_js_token SET access_token=?,expires_in=?,mtime=? where id=1")
	checkErr(err)

	mt := ztime.NowTimeStamp()
	res, err := stmt.Exec(callback.Access_token, callback.Expires_in, mt)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	return callback.Access_token
}

func (wx *Zweixin) GetJsApiTicket() string {
	db := wx.db
	//查询数据
	rows, err := db.Query("SELECT access_token,expires_in,mtime FROM wx_js_token where id=2")
	checkErr(err)
	var expires_in int
	var atoken string
	var mtime int64
	for rows.Next() {
		err := rows.Scan(&atoken, &expires_in, &mtime)
		checkErr(err)
	}

	nowtimestamp := ztime.NowTimeStamp()
	if nowtimestamp < (int64(expires_in) + mtime) {
		return atoken
	}

	accessToken := wx.GetAccessToken()
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token=%s", accessToken)
	callback := &ticket_callback{}
	zhttp.GETWithUnmarshal(url, callback)

	//修改数据
	stmt, err := db.Prepare("update wx_js_token SET access_token=?,expires_in=?,mtime=? where id=2")
	checkErr(err)

	mt := ztime.NowTimeStamp()
	res, err := stmt.Exec(callback.Ticket, callback.Expires_in, mt)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	return callback.Ticket
}
