from echo.v1.echo_grpc import EchoBase
from echo.v1.echo_pb2 import EchoRequest, EchoResponse
from grpclib.server import Stream


class Echo(EchoBase):
    async def Echo(self, stream: Stream[EchoRequest, EchoResponse]) -> None:
        req = await stream.recv_message()
        assert req is not None

        res = EchoResponse(msg=req.msg)
        await stream.send_message(res)
