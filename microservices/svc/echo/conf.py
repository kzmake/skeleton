import os

APP_ADDRESS = os.environ.get("ECHO_ADDRESS", "0.0.0.0:5050")
APP_IP_ADDRESS = APP_ADDRESS.split(":")[0]
APP_PORT = int(APP_ADDRESS.split(":")[1])
