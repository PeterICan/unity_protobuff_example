using Google.Protobuf;
using Newtonsoft.Json.Linq;

namespace ProtoBufferExample.Client
{
    public interface ISerializer
    {
        byte[] Serialize<T>(T obj) where T : IMessage<T>;
        T Deserialize<T>(byte[] data) where T : IMessage<T>, new();
        JObject DeserializeToJObject(byte[] data);
    }
}
