import configparser
import os


conf_file_path = ""
match os.getenv("APP_ENV"):
    case "development":
        conf_file_path = "settings/settings-development.ini"
    case "production":
        conf_file_path = "settings/settings-production.ini"
    case _:
        conf_file_path = "settings/settings-local.ini"

conf = configparser.ConfigParser()
conf.read(conf_file_path)

log_file = conf["api"]["log_file"]
port = int(conf["api"]["port"])

db_host = conf["db"]["db_host"]
db_port = conf["db"]["db_port"]
db_driver = conf["db"]["db_driver"]
db_name = conf["db"]["db_name"]
db_user = conf["db"]["db_user"]
db_password = conf["db"]["db_password"]
db_ssl_mode = conf["db"]["db_ssl_mode"]