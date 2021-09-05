import asyncio

import conf
import server

if __name__ == "__main__":
    asyncio.run(server.run(host=conf.APP_IP_ADDRESS, port=conf.APP_PORT))
