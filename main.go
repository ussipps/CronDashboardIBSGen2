package main

import (
	"CronIbsforMbs/functions"
	"CronIbsforMbs/scheduler"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	IntervalTimeNasabah, err := strconv.ParseUint(os.Getenv("INTERVAL_TIME_NASABAH"), 10, 64)
	IntervalTimeTabtrans, err := strconv.ParseUint(os.Getenv("INTERVAL_TIME_TABTRANS"), 10, 64)
	IntervalTimeDeptrans, err := strconv.ParseUint(os.Getenv("INTERVAL_TIME_DEPTRANS"), 10, 64)
	IntervalTimeKretrans, err := strconv.ParseUint(os.Getenv("INTERVAL_TIME_KRETRANS"), 10, 64)
	IntervalBroadcastTagihan := functions.ParseTimeScheduler(os.Getenv("INTERVAL_BROADCAST_TAGIHAN"))
	IntervalBroadcastUlangTahun := functions.ParseTimeScheduler(os.Getenv("INTERVAL_BROADCAST_ULTAH"))
	UseAutomaticBroadcast	:=	os.Getenv("AUTOMATIC_BROADCAST")
	UseEmail	:=	os.Getenv("USE_EMAIL")
	UseWA	:=	os.Getenv("USE_WA")

	if err != nil {

		log.Fatal("Error loading .env file")
	}
	functions.Logger().Info("Starting Scheduler Cron Ibs for MBS")


	if uint64(IntervalTimeNasabah) == 0 {
	}else{
		if (UseWA == "YA") && (UseEmail != "YA"){
			gocron.Every(uint64(IntervalTimeNasabah)).Seconds().Do(scheduler.CekNasabahWA)
		}
		if (UseWA != "YA") && (UseEmail == "YA"){
			gocron.Every(uint64(IntervalTimeNasabah)).Seconds().Do(scheduler.CekNasabahEmail)
		}

		if (UseWA == "YA") && (UseEmail == "YA"){
			gocron.Every(uint64(IntervalTimeNasabah)).Seconds().Do(scheduler.CekNasabahWA)
			gocron.Every(uint64(IntervalTimeNasabah)).Seconds().Do(scheduler.CekNasabahEmail)
		}

	}

	if uint64(IntervalTimeTabtrans) == 0 {
	}else{
		if (UseWA == "YA") && (UseEmail != "YA"){
			gocron.Every(uint64(IntervalTimeTabtrans)).Seconds().Do(scheduler.CekTabtransWA)
		}
		if (UseWA != "YA") && (UseEmail == "YA"){
			gocron.Every(uint64(IntervalTimeTabtrans)).Seconds().Do(scheduler.CekTabtransEmail)
		}

		if (UseWA == "YA") && (UseEmail == "YA"){
			gocron.Every(uint64(IntervalTimeTabtrans)).Seconds().Do(scheduler.CekTabtransWA)
			gocron.Every(uint64(IntervalTimeTabtrans)).Seconds().Do(scheduler.CekTabtransEmail)
		}

	}
	if uint64(IntervalTimeDeptrans) == 0 {
	}else{
		if (UseWA == "YA") && (UseEmail != "YA"){
			gocron.Every(uint64(IntervalTimeDeptrans)).Seconds().Do(scheduler.CekDeptransWA)
		}
		if (UseWA != "YA") && (UseEmail == "YA"){
			gocron.Every(uint64(IntervalTimeDeptrans)).Seconds().Do(scheduler.CekDeptransEmail)
		}

		if (UseWA == "YA") && (UseEmail == "YA"){
			gocron.Every(uint64(IntervalTimeDeptrans)).Seconds().Do(scheduler.CekDeptransWA)
			gocron.Every(uint64(IntervalTimeDeptrans)).Seconds().Do(scheduler.CekDeptransEmail)
		}
	}
	if uint64(IntervalTimeKretrans) == 0 {
	}else{
		if (UseWA == "YA") && (UseEmail != "YA"){
			gocron.Every(uint64(IntervalTimeKretrans)).Seconds().Do(scheduler.CekKretransWA)
		}
		if (UseWA != "YA") && (UseEmail == "YA"){
			gocron.Every(uint64(IntervalTimeKretrans)).Seconds().Do(scheduler.CekKretransEmail)
		}

		if (UseWA == "YA") && (UseEmail == "YA"){
			gocron.Every(uint64(IntervalTimeKretrans)).Seconds().Do(scheduler.CekKretransWA)
			gocron.Every(uint64(IntervalTimeKretrans)).Seconds().Do(scheduler.CekKretransEmail)
		}
	}
	if UseAutomaticBroadcast == "YA"{
		//TAGIHAN
		if (UseWA == "YA") && (UseEmail != "YA"){
			gocron.Every(1).Day().At(string(IntervalBroadcastTagihan)).Do(scheduler.CekTagihanWA)
		}
		if (UseWA != "YA") && (UseEmail == "YA"){
			gocron.Every(1).Day().At(string(IntervalBroadcastTagihan)).Do(scheduler.CekTagihanEmail)
		}

		if (UseWA == "YA") && (UseEmail == "YA"){
			gocron.Every(1).Day().At(string(IntervalBroadcastTagihan)).Do(scheduler.CekTagihanWA)
			gocron.Every(1).Day().At(string(IntervalBroadcastTagihan)).Do(scheduler.CekTagihanEmail)
		}

		//ULTAH
		if (UseWA == "YA") && (UseEmail != "YA"){
			gocron.Every(1).Day().At(string(IntervalBroadcastUlangTahun)).Do(scheduler.CekUltahWA)
		}
		if (UseWA != "YA") && (UseEmail == "YA"){
			gocron.Every(1).Day().At(string(IntervalBroadcastUlangTahun)).Do(scheduler.CekUltahEmail)
		}

		if (UseWA == "YA") && (UseEmail == "YA"){
			gocron.Every(1).Day().At(string(IntervalBroadcastUlangTahun)).Do(scheduler.CekUltahWA)
			gocron.Every(1).Day().At(string(IntervalBroadcastUlangTahun)).Do(scheduler.CekUltahEmail)
		}
	}
	<-gocron.Start()

}
