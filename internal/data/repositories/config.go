package repositories

// database
const hostDefault = "127.0.0.1"
const portDefault = "6001"
const userDefault = "postgres"
const passwordDefault = "mysecretpassword"
const databaseDefault = "postgres"

// migration
const packagePath = "/home/isjhar/Documents/Quests/Tms/api" // lokasi folder project. contoh : /home/isjhar/Documents/Quests/Tms/api

const traccarUrlDefault = "https://gps-silog.id-transport.net/api"
const traccarUsernameDefault = "gps@silog.co.id"
const traccarPasswordDefault = "gps@silog"
const traccarGroupID = "4"

const redisAddress = "localhost:6379"

// integrasi sig
const silogWebUrlDefault = "http://api-silog.id-transport.net/api"
const sigUrlDefault = "http://10.4.194.150"
const sipUrlDefault = "http://10.227.16.48"
const sipUrl2Default = "https://sip.solusibangunindonesia.com"

// sangu
const sanguUrlDefault = "http://202.180.26.147:8080/sangudriver/getTransfer.php"
const sanguTokenDefault = "silogjayaselamamnya"

// sms gateway
const smsGatewayUrlDefault = "https://otp.id-transport.net/services/send.php?key=20bff9f59961c2714a482dd7bf4456ecbda67246&devices=3%7C1&type=sms&prioritize=0"

const elasticsearchUrl = "http://wikalcts-logging.gpstracker.io"
const elasticsearchKey = "ZWxhc3RpYzpJRC1DbG91ZHMuTjNU"
const elasticsearchCategory = "silog-service-local-log"

const clientKey = "ac128e27350dae2c7fac9905ead5d9c5"

const qrCodeExpiredSeconds = "300"

// const telegramUrl = "https://api.telegram.org/bot7784709290:AAFKpkFWUorrSLCnMpv1nTS1Qp9PsOYkPps/sendMessage?chat_id=-1002466180663&text="
const telegramUrl = ""

const easyGoUrl = "http://api-silog.id-transport.net"
