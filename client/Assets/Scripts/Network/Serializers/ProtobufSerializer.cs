using Google.Protobuf;
using Newtonsoft.Json.Linq;

namespace ProtoBufferExample.Client
{
    public class ProtobufSerializer : ISerializer
    {
        public byte[] Serialize<T>(T obj) where T : IMessage<T>
        {
            if (obj is IMessage message)
            {
                return message.ToByteArray();
            }
            throw new System.ArgumentException("Object must be of type Google.Protobuf.IMessage", nameof(obj));
        }

        public T Deserialize<T>(byte[] data) where T : IMessage<T>, new()
        {
            // This requires a way to create an instance of T and then parse.
            // We assume T has a parameterless constructor and a Parser property.
            var instance = System.Activator.CreateInstance<T>();
            if (instance is IMessage message)
            {
                message.MergeFrom(data);
                return (T)message;
            }
            throw new System.ArgumentException("Type T must be a Google.Protobuf.IMessage", nameof(T));
        }

        public JObject DeserializeToJObject(byte[] data)
        {
            var json = System.Text.Encoding.UTF8.GetString(data);
            return JObject.Parse(json);
        }
    }
}
