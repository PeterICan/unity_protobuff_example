using Google.Protobuf;

namespace ProtoBufferExample.Client
{
    public interface ISerializer
    {
        byte[] Serialize<T>(T obj) where T : IMessage<T>;
        T Deserialize<T>(byte[] data) where T : IMessage<T>, new();
    }
}
