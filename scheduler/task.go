package scheduler

import (
	"CronIbsforMbs/database"
	"CronIbsforMbs/functions"
	_ "github.com/shopspring/decimal"
	"github.com/vjeantet/jodaTime"
	_ "math/big"
	"os"
	"strconv"
	"time"
)

var conn = database.ConnectDB()
var connWA = database.ConnectDBWA()


type SelectTabWa struct {
	TransId     string  `json:"trans_id"`
	NoRekening  string  `json:"no_rekening"`
	NamaNasabah string  `json:"nama_nasabah"`
	Keterangan  string  `json:"keterangan"`
	Nominal     float64 `json:"nominal"`
	MyKodeTrans string  `json:"my_kode_trans"`
	KodeTrans   string  `json:"kode_trans"`
	Hp          string  `json:"hp"`
	TglTrans    string  `json:"tgl_trans"`
	MyTrans		string	`json:"my_trans"`
}

type SelectTabEmail struct {
	TransId     string  `json:"trans_id"`
	NoRekening  string  `json:"no_rekening"`
	NamaNasabah string  `json:"nama_nasabah"`
	Keterangan  string  `json:"keterangan"`
	Nominal     float64 `json:"nominal"`
	MyKodeTrans string  `json:"my_kode_trans"`
	KodeTrans   string  `json:"kode_trans"`
	Email       string  `json:"email"`
	TglTrans    string  `json:"tgl_trans"`
	MyTrans		string	`json:"my_trans"`
}

type SelectDepWa struct {
	TransId     string  `json:"trans_id"`
	NoRekening  string  `json:"no_rekening"`
	NamaNasabah string  `json:"nama_nasabah"`
	Keterangan  string  `json:"keterangan"`
	Nominal     float64 `json:"nominal"`
	Pokok       float64 `json:"pokok"`
	Bunga       float64 `json:"bunga"`
	MyKodeTrans string  `json:"my_kode_trans"`
	KodeTrans   string  `json:"kode_trans"`
	Hp          string  `json:"hp"`
	TglTrans    string  `json:"tgl_trans"`
	MyTrans		string	`json:"my_trans"`
}

type SelectDepEmail struct {
	TransId     string  `json:"trans_id"`
	NoRekening  string  `json:"no_rekening"`
	NamaNasabah string  `json:"nama_nasabah"`
	Keterangan  string  `json:"keterangan"`
	Nominal     float64 `json:"nominal"`
	Pokok       float64 `json:"pokok"`
	Bunga       float64 `json:"bunga"`
	MyKodeTrans string  `json:"my_kode_trans"`
	KodeTrans   string  `json:"kode_trans"`
	Email       string  `json:"email"`
	TglTrans    string  `json:"tgl_trans"`
	MyTrans		string	`json:"my_trans"`
}

type SelectKreWa struct {
	TransId     string  `json:"trans_id"`
	NoRekening  string  `json:"no_rekening"`
	NamaNasabah string  `json:"nama_nasabah"`
	Keterangan  string  `json:"keterangan"`
	Nominal     float64 `json:"nominal"`
	MyKodeTrans string  `json:"my_kode_trans"`
	KodeTrans   string  `json:"kode_trans"`
	Hp          string  `json:"hp"`
	TglTrans    string  `json:"tgl_trans"`
	MyTrans		string	`json:"my_trans"`
}

type SelectKreEmail struct {
	TransId     string  `json:"trans_id"`
	NoRekening  string  `json:"no_rekening"`
	NamaNasabah string  `json:"nama_nasabah"`
	Keterangan  string  `json:"keterangan"`
	Nominal     float64 `json:"nominal"`
	MyKodeTrans string  `json:"my_kode_trans"`
	KodeTrans   string  `json:"kode_trans"`
	Email       string  `json:"email"`
	TglTrans    string  `json:"tgl_trans"`
	MyTrans		string	`json:"my_trans"`
}

type SelectUltahWa struct {
	NamaNasabah string  `json:"nama_nasabah"`
	Hp          string  `json:"hp"`
	TglTrans    string  `json:"tgl_trans"`
}

type SelectUltahEmail struct {
	NamaNasabah string  `json:"nama_nasabah"`
	Email       string  `json:"email"`
	TglTrans    string  `json:"tgl_trans"`
}

type SelectTagihanWa struct {
	TransId     string  `json:"trans_id"`
	NoRekening  string  `json:"no_rekening"`
	NamaNasabah string  `json:"nama_nasabah"`
	Keterangan  string  `json:"keterangan"`
	Nominal     float64 `json:"nominal"`
	MyKodeTrans string  `json:"my_kode_trans"`
	KodeTrans   string  `json:"kode_trans"`
	Hp          string  `json:"hp"`
	TglTrans    string  `json:"tgl_trans"`
	MyTrans		string	`json:"my_trans"`
}

type SelectTagihanEmail struct {
	TransId     string  `json:"trans_id"`
	NoRekening  string  `json:"no_rekening"`
	NamaNasabah string  `json:"nama_nasabah"`
	Keterangan  string  `json:"keterangan"`
	Nominal     float64 `json:"nominal"`
	MyKodeTrans string  `json:"my_kode_trans"`
	KodeTrans   string  `json:"kode_trans"`
	Email       string  `json:"email"`
	TglTrans    string  `json:"tgl_trans"`
	MyTrans		string	`json:"my_trans"`
}

type SelectNasabahWA struct {
	NasabahId   string  `json:"nasabah_id"`
	NamaNasabah string  `json:"nama_nasabah"`
	Hp          string  `json:"hp"`
	TglTrans    string  `json:"tgl_trans"`
}

type SelectNasabahEmail struct {
	NasabahId   string  `json:"nasabah_id"`
	NamaNasabah string  `json:"nama_nasabah"`
	Email       string  `json:"email"`
	TglTrans    string  `json:"tgl_trans"`
}

//TABTRANS
func CekTabtransWA() {

	cTgl := jodaTime.Format("YYYYMMdd", time.Now())
	sqlStatement := "SELECT tabtrans_id trans_id,t.no_rekening,nama_nasabah,ifnull(t.keterangan,'') keterangan, "+
		"coalesce(pokok,0) nominal,my_kode_trans,kode_trans,if (length(n.hp)=11, "+
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 10 ) ) )," +
		"IF (length( n.hp ) = 12, trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 11 ) ) ) ," +
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 12 ) ) ) )) AS hp," +
		"tgl_trans,floor(my_kode_trans/100) myTrans " +
		"FROM " +
		"tabtrans t," +
		"nasabah n," +
		"tabung m " +
		"WHERE " +
		"t.no_rekening = m.no_rekening " +
		"AND n.nasabah_id = m.nasabah_id " +
		"AND tgl_trans = '"+cTgl+"' " +
		"and hp<>'' and floor(my_kode_trans/100)<=2 " +
		"and tabtrans_id not in (select trans_id from `wa_api`.wa_notifikasi where hp<>'') having hp<>'62' or hp <>'0' or hp<>'620' or hp<>'' ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Tabtrans WA")
		messageList := SelectTabWa{}
		err = rows.Scan(&messageList.TransId, &messageList.NoRekening, &messageList.NamaNasabah,
			&messageList.Keterangan, &messageList.Nominal, &messageList.MyKodeTrans, &messageList.KodeTrans,
			&messageList.Hp, &messageList.TglTrans,&messageList.MyTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		MyTrans := messageList.MyTrans
		cModul := ""
		if MyTrans == "1" {
			cModul = "TABSETORAN"
		}else{
			cModul = "TABTARIK"
		}
		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(messageList.TransId,messageList.NoRekening,
				messageList.NamaNasabah,messageList.Keterangan,messageList.Nominal,0,messageList.MyKodeTrans,
				messageList.TglTrans,cModul,messageList.Hp,0,"",0,"")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Tabtrans WA")
	}
	rows.Close()
}

func CekTabtransEmail() {

	cTgl := jodaTime.Format("YYYYMMdd", time.Now())
	sqlStatement := "SELECT tabtrans_id trans_id,t.no_rekening,nama_nasabah,ifnull(t.keterangan,'') keterangan, "+
		"coalesce(pokok,0) nominal,my_kode_trans,kode_trans, "+
		"email," +
		"tgl_trans,floor(my_kode_trans/100) myTrans " +
		"FROM " +
		"tabtrans t," +
		"nasabah n," +
		"tabung m " +
		"WHERE " +
		"t.no_rekening = m.no_rekening " +
		"AND n.nasabah_id = m.nasabah_id " +
		"AND tgl_trans = '"+cTgl+"' " +
		"and email<>'' and floor(my_kode_trans/100)<=2 " +
		"and tabtrans_id not in (select trans_id from `wa_api`.wa_notifikasi where use_email=1) ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Tabtrans Email")
		messageList := SelectTabEmail{}
		err = rows.Scan(&messageList.TransId, &messageList.NoRekening, &messageList.NamaNasabah,
			&messageList.Keterangan, &messageList.Nominal, &messageList.MyKodeTrans, &messageList.KodeTrans,
			&messageList.Email, &messageList.TglTrans,&messageList.MyTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		MyTrans := messageList.MyTrans
		cModul := ""
		if MyTrans == "1" {
			cModul = "TABSETORAN"
		}else{
			cModul = "TABTARIK"
		}
		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(messageList.TransId,messageList.NoRekening,
				messageList.NamaNasabah,messageList.Keterangan,messageList.Nominal,0,messageList.MyKodeTrans,
				messageList.TglTrans,cModul,"",1,messageList.Email,0,"[nama_lembaga] : TRANSACTIONAL NOTIFICATION")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Tabtrans Email")
	}
	rows.Close()
}

//DEPTRANS
func CekDeptransWA() {

	cTgl := jodaTime.Format("YYYYMMdd", time.Now())
	sqlStatement := "SELECT deptrans_id trans_id,t.no_rekening,nama_nasabah,ifnull(t.keterangan,'') keterangan, "+
		"coalesce(pokok_trans,0)+coalesce(bunga_trans,0)-coalesce(pajak_trans,0) nominal," +
		"my_kode_trans,kode_trans,if (length(n.hp)=11, "+
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 10 ) ) )," +
		"IF (length( n.hp ) = 12, trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 11 ) ) ) ," +
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 12 ) ) ) )) AS hp," +
		"tgl_trans,floor(my_kode_trans/100) myTrans,coalesce(pokok_trans,0) pokok,coalesce(bunga_trans,0) bunga " +
		"FROM " +
		"deptrans t," +
		"nasabah n," +
		"deposito m " +
		"WHERE " +
		"t.no_rekening = m.no_rekening " +
		"AND n.nasabah_id = m.nasabah_id " +
		"AND tgl_trans = '"+cTgl+"' " +
		"and hp<>'' and floor(my_kode_trans/100)<=2 " +
		"and kode_trans<>'501' "+
		"and deptrans_id not in (select trans_id from `wa_api`.wa_notifikasi where hp<>'') " +
		"having hp<>'62' or hp <>'0' or hp<>'620' or hp<>'' ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Deptrans WA")
		messageList := SelectDepWa{}
		err = rows.Scan(&messageList.TransId, &messageList.NoRekening, &messageList.NamaNasabah,
			&messageList.Keterangan, &messageList.Nominal, &messageList.MyKodeTrans, &messageList.KodeTrans,
			&messageList.Hp, &messageList.TglTrans,&messageList.MyTrans,&messageList.Pokok,&messageList.Bunga)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		MyTrans := messageList.MyTrans
		cModul := ""
		if messageList.Pokok!=0 {
			if MyTrans == "1"{
				cModul = "DEPSETORPOKOK"
			}
			if MyTrans == "2"{
				cModul = "DEPTARIKPOKOK"
			}
		}
		if messageList.Bunga!=0 {
			if MyTrans == "2"{
				cModul = "DEPSETORBUNGA"
			}
		}
		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(messageList.TransId,messageList.NoRekening,
				messageList.NamaNasabah,messageList.Keterangan,messageList.Nominal,0,messageList.MyKodeTrans,
				messageList.TglTrans,cModul,messageList.Hp,0,"",0,"")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Deptrans WA")
	}
	rows.Close()
}

func CekDeptransEmail() {

	cTgl := jodaTime.Format("YYYYMMdd", time.Now())
	sqlStatement := "SELECT deptrans_id trans_id,t.no_rekening,nama_nasabah,ifnull(t.keterangan,'') keterangan, "+
		"coalesce(pokok_trans,0)+coalesce(bunga_trans,0)-coalesce(pajak_trans,0) nominal," +
		"my_kode_trans,kode_trans,email," +
		"tgl_trans,floor(my_kode_trans/100) myTrans,coalesce(pokok_trans,0) pokok,coalesce(bunga_trans,0) bunga " +
		"FROM " +
		"deptrans t," +
		"nasabah n," +
		"deposito m " +
		"WHERE " +
		"t.no_rekening = m.no_rekening " +
		"AND n.nasabah_id = m.nasabah_id " +
		"AND tgl_trans = '"+cTgl+"' " +
		"and email<>'' and floor(my_kode_trans/100)<=2 " +
		"and kode_trans<>'501' "+
		"and deptrans_id not in (select trans_id from `wa_api`.wa_notifikasi where use_email=1) ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Deptrans Email")
		messageList := SelectDepEmail{}
		err = rows.Scan(&messageList.TransId, &messageList.NoRekening, &messageList.NamaNasabah,
			&messageList.Keterangan, &messageList.Nominal, &messageList.MyKodeTrans, &messageList.KodeTrans,
			&messageList.Email, &messageList.TglTrans,&messageList.MyTrans,&messageList.Pokok,&messageList.Bunga)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		MyTrans := messageList.MyTrans
		cModul := ""
		if messageList.Pokok!=0 {
			if MyTrans == "1"{
				cModul = "DEPSETORPOKOK"
			}
			if MyTrans == "2"{
				cModul = "DEPTARIKPOKOK"
			}
		}
		if messageList.Bunga!=0 {
			if MyTrans == "2"{
				cModul = "DEPSETORBUNGA"
			}
		}
		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(messageList.TransId,messageList.NoRekening,
				messageList.NamaNasabah,messageList.Keterangan,messageList.Nominal,0,messageList.MyKodeTrans,
				messageList.TglTrans,cModul,"",1,messageList.Email,0,"[nama_lembaga] : TRANSACTIONAL NOTIFICATION")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Deptrans Email")
	}
	rows.Close()
}

//KRETRANS
func CekKretransWA() {

	cTgl := jodaTime.Format("YYYYMMdd", time.Now())
	sqlStatement := "SELECT kretrans_id trans_id,t.no_rekening,nama_nasabah,ifnull(t.keterangan,'') keterangan, "+
		"(coalesce(pokok,0)+coalesce(bunga,0)+coalesce(denda,0)) - (coalesce(disc_bunga,0)+coalesce(disc_denda,0)) nominal," +
		"my_kode_trans,kode_trans,if (length(n.hp)=11, "+
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 10 ) ) )," +
		"IF (length( n.hp ) = 12, trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 11 ) ) ) ," +
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 12 ) ) ) )) AS hp," +
		"tgl_trans,floor(my_kode_trans/100) myTrans " +
		"FROM " +
		"kretrans t," +
		"nasabah n," +
		"kredit m " +
		"WHERE " +
		"t.no_rekening = m.no_rekening " +
		"AND n.nasabah_id = m.nasabah_id " +
		"AND tgl_trans = '"+cTgl+"' " +
		"and hp<>'' and (floor(my_kode_trans/100)=1 or floor(my_kode_trans/100)=3) " +
		"and kode_trans<>'501' "+
		"and kretrans_id not in (select trans_id from `wa_api`.wa_notifikasi where hp<>'') " +
		"having hp<>'62' or hp <>'0' or hp<>'620' or hp<>'' ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Kretrans WA")
		messageList := SelectKreWa{}
		err = rows.Scan(&messageList.TransId, &messageList.NoRekening, &messageList.NamaNasabah,
			&messageList.Keterangan, &messageList.Nominal, &messageList.MyKodeTrans, &messageList.KodeTrans,
			&messageList.Hp, &messageList.TglTrans,&messageList.MyTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		MyTrans := messageList.MyTrans
		cModul := ""
		if MyTrans == "1"{
			cModul = "KREREALISASI"
		}
		if MyTrans == "3"{
			cModul = "KREANGSUR"
		}
		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(messageList.TransId,messageList.NoRekening,
				messageList.NamaNasabah,messageList.Keterangan,messageList.Nominal,0,messageList.MyKodeTrans,
				messageList.TglTrans,cModul,messageList.Hp,0,"",0,"")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Kretrans WA")
	}
	rows.Close()
}

func CekKretransEmail() {

	cTgl := jodaTime.Format("YYYYMMdd", time.Now())
	sqlStatement := "SELECT kretrans_id trans_id,t.no_rekening,nama_nasabah,ifnull(t.keterangan,'') keterangan, "+
		"(coalesce(pokok,0)+coalesce(bunga,0)+coalesce(denda,0)) - (coalesce(disc_bunga,0)+coalesce(disc_denda,0)) nominal," +
		"my_kode_trans,kode_trans, "+
		"email," +
		"tgl_trans,floor(my_kode_trans/100) myTrans " +
		"FROM " +
		"kretrans t," +
		"nasabah n," +
		"kredit m " +
		"WHERE " +
		"t.no_rekening = m.no_rekening " +
		"AND n.nasabah_id = m.nasabah_id " +
		"AND tgl_trans = '"+cTgl+"' " +
		"and email<>'' and (floor(my_kode_trans/100)=1 or floor(my_kode_trans/100)=3) " +
		"and kretrans_id not in (select trans_id from `wa_api`.wa_notifikasi where use_email=1) ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Kretrans Email")
		messageList := SelectKreEmail{}
		err = rows.Scan(&messageList.TransId, &messageList.NoRekening, &messageList.NamaNasabah,
			&messageList.Keterangan, &messageList.Nominal, &messageList.MyKodeTrans, &messageList.KodeTrans,
			&messageList.Email, &messageList.TglTrans,&messageList.MyTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		MyTrans := messageList.MyTrans
		cModul := ""
		if MyTrans == "1" {
			cModul = "KREREALISASI"
		}else{
			cModul = "KREANGSUR"
		}
		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(messageList.TransId,messageList.NoRekening,
				messageList.NamaNasabah,messageList.Keterangan,messageList.Nominal,0,messageList.MyKodeTrans,
				messageList.TglTrans,cModul,"",1,messageList.Email,0,"[nama_lembaga] : TRANSACTIONAL NOTIFICATION")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Kretrans Email")
	}
	rows.Close()
}

//ULTAH
func CekUltahWA() {
	cDay := jodaTime.Format("dd", time.Now())
	cMonth := jodaTime.Format("MM", time.Now())
	sqlStatement := "SELECT nama_nasabah,  "+
		"if (length(n.hp)=11, "+
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 10 ) ) )," +
		"IF (length( n.hp ) = 12, trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 11 ) ) ) ," +
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 12 ) ) ) )) AS hp," +
		"date(now()) tgl_trans " +
		"FROM " +
		"nasabah n " +
		"WHERE " +
		"(day(tgllahir)="+cDay+" and month(tgllahir)="+cMonth+") " +
		"and hp<>'' " +
		"having hp<>'62' or hp <>'0' or hp<>'620' or hp<>'' ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Ultah WA")
		messageList := SelectUltahWa{}
		err = rows.Scan(&messageList.NamaNasabah,
			&messageList.Hp, &messageList.TglTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		cModul := "ULANGTAHUN"
		keterangan :=GetFValueByFKeyValueWA("template","template_name","ULANGTAHUN","template_text")

		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(0,"",
				messageList.NamaNasabah,keterangan,0,0,100,
				messageList.TglTrans,cModul,messageList.Hp,0,"",0,"")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Ultah WA")
	}
	rows.Close()
}

func CekUltahEmail() {
	cDay := jodaTime.Format("dd", time.Now())
	cMonth := jodaTime.Format("MM", time.Now())
	sqlStatement := "SELECT nama_nasabah,  "+
		"email," +
		"date(now()) tgl_trans " +
		"FROM " +
		"nasabah n " +
		"WHERE " +
		"(day(tgllahir)="+cDay+" and month(tgllahir)="+cMonth+") " +
		"and email<>''  ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Ultah Email")
		messageList := SelectUltahEmail{}
		err = rows.Scan(&messageList.NamaNasabah,
			&messageList.Email, &messageList.TglTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		cModul := "ULANGTAHUN"
		keterangan :=GetFValueByFKeyValueWA("template_email","template_name","ULANGTAHUN","template_text")

		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?, " +
				"?,?,?,?,?,?,?,?) ")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(0,"",
				messageList.NamaNasabah,keterangan,0,0,100,
				messageList.TglTrans,cModul,"",1,messageList.Email,0,"[nama_lembaga] - Selamat Ulang Tahun untuk Anda")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Ultah Email")
	}
	rows.Close()
}


//TAGIHAN
func CekTagihanWA() {

	t := time.Now()
	t2 := time.Now()
	nSelisih, _ := strconv.Atoi(os.Getenv("DAY_SELISIH_BROADCAST_TAGIHAN"))
	t = t.AddDate(0,0,0)
	t2 = t2.AddDate(0,0,nSelisih)
	cTgl := jodaTime.Format("YYYYMMdd", t2)
	cTgl2 := jodaTime.Format("YYYYMMdd", t)
	sqlStatement := "SELECT kretrans_id trans_id,t.no_rekening,nama_nasabah,ifnull(t.keterangan,'') keterangan, "+
		"if(coalesce(pokok,0)<0,0,coalesce(pokok,0)) + if(coalesce(bunga,0)<0,0,coalesce(bunga,0)) as nominal," +
		"my_kode_trans,kode_trans,if (length(n.hp)=11, "+
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 10 ) ) )," +
		"IF (length( n.hp ) = 12, trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 11 ) ) ) ," +
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 12 ) ) ) )) AS hp," +
		"tgl_trans,floor(my_kode_trans/100) myTrans " +
		"FROM " +
		"kretrans t," +
		"nasabah n," +
		"kredit m " +
		"WHERE " +
		"t.no_rekening = m.no_rekening " +
		"AND n.nasabah_id = m.nasabah_id " +
		"AND tgl_trans <= '"+cTgl+"' " +
		"AND tgl_trans >= '"+cTgl2+"' " +
		"and hp<>'' and (floor(my_kode_trans/100)=2) " +
		"having (hp<>'62' or hp <>'0' or hp<>'620' or hp<>'') and nominal>0";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Tagihan WA")
		messageList := SelectTagihanWa{}
		err = rows.Scan(&messageList.TransId, &messageList.NoRekening, &messageList.NamaNasabah,
			&messageList.Keterangan, &messageList.Nominal, &messageList.MyKodeTrans, &messageList.KodeTrans,
			&messageList.Hp, &messageList.TglTrans,&messageList.MyTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		keterangan := GetFValueByFKeyValueWA("template","template_name","KRETAGIHAN","template_text")
		cModul := "KRETAGIHAN"
		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(messageList.TransId,messageList.NoRekening,
				messageList.NamaNasabah,keterangan,messageList.Nominal,0,messageList.MyKodeTrans,
				messageList.TglTrans,cModul,messageList.Hp,0,"",0,"")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Tagihan WA")
	}
	rows.Close()
}

func CekTagihanEmail() {

	t := time.Now()
	t2 := time.Now()
	nSelisih, _ := strconv.Atoi(os.Getenv("DAY_SELISIH_BROADCAST_TAGIHAN"))
	t = t.AddDate(0,0,0)
	t2 = t2.AddDate(0,0,nSelisih)
	cTgl := jodaTime.Format("YYYYMMdd", t2)
	cTgl2 := jodaTime.Format("YYYYMMdd", t)
	sqlStatement := "SELECT kretrans_id trans_id,t.no_rekening,nama_nasabah,ifnull(t.keterangan,'') keterangan, "+
		"if(coalesce(pokok,0)<0,0,coalesce(pokok,0)) + if(coalesce(bunga,0)<0,0,coalesce(bunga,0)) as nominal," +
		"my_kode_trans,kode_trans,email," +
		"tgl_trans,floor(my_kode_trans/100) myTrans " +
		"FROM " +
		"kretrans t," +
		"nasabah n," +
		"kredit m " +
		"WHERE " +
		"t.no_rekening = m.no_rekening " +
		"AND n.nasabah_id = m.nasabah_id " +
		"AND tgl_trans <= '"+cTgl+"' " +
		"AND tgl_trans >= '"+cTgl2+"' " +
		"and email<>'' and (floor(my_kode_trans/100)=2) " +
		"having nominal>0 ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Tagihan WA")
		messageList := SelectTagihanEmail{}
		err = rows.Scan(&messageList.TransId, &messageList.NoRekening, &messageList.NamaNasabah,
			&messageList.Keterangan, &messageList.Nominal, &messageList.MyKodeTrans, &messageList.KodeTrans,
			&messageList.Email, &messageList.TglTrans,&messageList.MyTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		keterangan := GetFValueByFKeyValueWA("template_email","template_name","KRETAGIHAN","template_text")
		cModul := "KRETAGIHAN"
		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(messageList.TransId,messageList.NoRekening,
				messageList.NamaNasabah,keterangan,messageList.Nominal,0,messageList.MyKodeTrans,
				messageList.TglTrans,cModul,"",1,messageList.Email,0,"[nama_lembaga] - Pengingat Tagihan Kredit")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Tagihan Email")
	}
	rows.Close()
}

//NASABAH
func CekNasabahWA() {
	cTgl := jodaTime.Format("YYYYMMdd", time.Now())
	sqlStatement := "SELECT nasabah_id,nama_nasabah,  "+
		"if (length(n.hp)=11, "+
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 10 ) ) )," +
		"IF (length( n.hp ) = 12, trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 11 ) ) ) ," +
		"trim( concat( REPLACE ( LEFT ( n.hp, 1 ), '0', '62' ), RIGHT ( n.hp, 12 ) ) ) )) AS hp," +
		"date(tgl_register) tgl_trans " +
		"FROM " +
		"nasabah n " +
		"WHERE " +
		"tgl_register='"+cTgl+"' " +
		"and hp<>''  " +
		"and nasabah_id not in (select no_rekening from `wa_api`.wa_notifikasi " +
		"where tgl_trans='"+cTgl+"' and modul='NASABAHBARU' and hp<>'') "+
		"having hp<>'62' or hp <>'0' or hp<>'620' or hp<>'' ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Nasabah Baru WA")
		messageList := SelectNasabahWA{}
		err = rows.Scan(&messageList.NasabahId,&messageList.NamaNasabah,
			&messageList.Hp, &messageList.TglTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		cModul := "NASABAHBARU"
		keterangan :=GetFValueByFKeyValueWA("template","template_name","NASABAHBARU","template_text")

		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?," +
				"?,?,?,?,?,?,?,?)")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(0,messageList.NasabahId,
				messageList.NamaNasabah,keterangan,0,0,100,
				messageList.TglTrans,cModul,messageList.Hp,0,"",0,"")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Nasabah Baru WA")
	}
	rows.Close()
}

func CekNasabahEmail() {
	cTgl := jodaTime.Format("YYYYMMdd", time.Now())
	sqlStatement := "SELECT nasabah_id,nama_nasabah,  "+
		"email," +
		"date(tgl_register) tgl_trans " +
		"FROM " +
		"nasabah n " +
		"WHERE " +
		"tgl_register='"+cTgl+"' "+
		"and nasabah_id not in (select no_rekening from `wa_api`.wa_notifikasi " +
		"where tgl_trans='"+cTgl+"' and modul='NASABAHBARU' and use_email=1) "+
		"and email<>''  ";
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return
	}
	defer database.ConnectDB().Close()

	for rows.Next() {
		functions.Logger().Info("Starting Scheduler Cek Nasabah Baru Email")
		messageList := SelectNasabahEmail{}
		err = rows.Scan(&messageList.NasabahId,&messageList.NamaNasabah,
			&messageList.Email, &messageList.TglTrans)
		if err != nil {
			functions.Logger().Error(err.Error())
			return
		}
		cModul := "NASABAHBARU"
		keterangan :=GetFValueByFKeyValueWA("template_email","template_name","NASABAHBARU","template_text")

		if (cModul!=""){
			stmt, err := connWA.Prepare("INSERT INTO `wa_api`.wa_notifikasi " +
				"(trans_id,no_rekening,nama_nasabah,keterangan,nominal,status_wa," +
				"my_kode_trans,tgl_trans,modul,hp,use_email,email,status_email,subject) " +
				"values " +
				"(?,?,?,?,?,?, " +
				"?,?,?,?,?,?,?,?) ")
			if err != nil {
				//panic(err.Error())
				functions.Logger().Error(err.Error())
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(0,messageList.NasabahId,
				messageList.NamaNasabah,keterangan,0,0,100,
				messageList.TglTrans,cModul,"",1,messageList.Email,0,"[nama_lembaga] - Selamat Datang")
			if err != nil {
				functions.Logger().Error(err.Error())
				return
			}
		}
		functions.Logger().Info("Successfully Cek Nasabah Baru Email")
	}
	rows.Close()
}


func GetFValueByFKeyValueIBS(Table string, FieldKey string, FieldKeyValue string, FieldTarget string) string {
	sqlStatement := "SELECT ifnull(" + FieldTarget + ",'') " + FieldTarget + " from " + Table + " where " + FieldKey + " = '" + FieldKeyValue + "'"
	rows, err := database.ConnectDB().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return ""
	}
	defer database.ConnectDB().Close()
	var Field string
	for rows.Next() {
		err = rows.Scan(&Field)
		if err != nil {
			return ""
		}
	}
	rows.Close()
	return Field

}

func GetFValueByFKeyValueWA(Table string, FieldKey string, FieldKeyValue string, FieldTarget string) string {
	sqlStatement := "SELECT ifnull(" + FieldTarget + ",'') " + FieldTarget + " from " + Table + " where " + FieldKey + " = '" + FieldKeyValue + "'"
	rows, err := database.ConnectDBWA().Query(sqlStatement)
	if err != nil {
		functions.Logger().Error(err.Error())
		return ""
	}
	defer database.ConnectDBWA().Close()
	var Field string
	for rows.Next() {
		err = rows.Scan(&Field)
		if err != nil {
			return ""
		}
	}
	rows.Close()
	return Field

}
