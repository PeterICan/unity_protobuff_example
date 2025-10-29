using System.Text;
using Google.Protobuf;
using Newtonsoft.Json; // Add this
using Newtonsoft.Json.Linq; // Add this

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

            // Use Google.Protobuf's formatter for Protobuf messages
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
            // Use Google.Protobuf's parser for Protobuf messages
            return _parser.Parse<T>(jsonString);
        }

        // New method to deserialize raw byte[] into a generic JObject for route extraction
        public JObject DeserializeToJObject(byte[] data)
        {
            if (data == null || data.Length == 0)
            {
                return null;
            }
            string jsonString = Encoding.UTF8.GetString(data);
            return JObject.Parse(jsonString);
        }
    }
}
