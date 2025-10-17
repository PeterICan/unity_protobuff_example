using System.Text;
using Google.Protobuf;

namespace ProtoBufferExample.Client
{
    public class JSONSerializer : ISerializer
    {
        private readonly JsonFormatter _formatter = new JsonFormatter(JsonFormatter.Settings.Default);
        private readonly JsonParser _parser = new JsonParser(JsonParser.Settings.Default);

        public byte[] Serialize<T>(T obj) where T : IMessage<T>
        {
            if (obj == null)
            {
                return null;
            }

            string jsonString = _formatter.Format(obj);
            return Encoding.UTF8.GetBytes(jsonString);
        }

        public T Deserialize<T>(byte[] data) where T : IMessage<T>, new()
        {
            if (data == null || data.Length == 0)
            {
                return default(T);
            }

            string jsonString = Encoding.UTF8.GetString(data);
            return _parser.Parse<T>(jsonString);
        }
    }
}
