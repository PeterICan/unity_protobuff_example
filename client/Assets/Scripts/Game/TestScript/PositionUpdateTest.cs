using UnityEngine;
using System.IO;
using ProtoBufferExample.Client.Generated;
using Newtonsoft.Json.Linq;
using JsonApi;
using ProtoBufferExample.Client;

namespace ProtoBufferExample.Client.Game.TestScript
{
    public class PositionUpdateTest : MonoBehaviour
    {
        public enum ConnectType
        {
            TCP,
            WebSocket
        }

        public enum SerializeType
        {
            Protobuf,
            MessagePack
        }

        public string serverAddress = "127.0.0.1";
        public int serverPort = 6666;

        public ConnectType connectionType = ConnectType.TCP;

        private IConnection _connection;
        private ISerializer _protoSerializer; // For TCP/Protobuf
        private JSONSerializer _jsonSerializer; // For WebSocket/JSON

        void Start()
        {
            _connection = connectionType == ConnectType.TCP ? new TCPClientTransport() : new WebSocketClientTransport();

            if (connectionType == ConnectType.TCP)
            {
                _protoSerializer = new ProtobufSerializer();
            }
            else // WebSocket
            {
                _jsonSerializer = new JSONSerializer();
            }

            _connection.OnMessageReceived += OnMessageReceived;

            Debug.Log("Connecting to server...");
            _connection.Connect(serverAddress, serverPort);
        }

        void Update()
        {
            if (Input.GetKeyDown(KeyCode.Space))
            {
                if (_connection.IsConnected)
                {
                    SendPositionUpdate();
                }
                else
                {
                    Debug.LogWarning("Not connected to server.");
                }
            }
        }

        void OnDestroy()
        {
            if (_connection != null && _connection.IsConnected)
            {
                _connection.Disconnect();
            }
        }

        private void SendPositionUpdate()
        {
            byte[] message = null;

            if (connectionType == ConnectType.TCP)
            {
                // 1. Create the Protobuf message payload
                var position = new PlayerPosition
                {
                    X = Random.Range(-10f, 10f),
                    Y = Random.Range(-10f, 10f),
                    Z = Random.Range(-10f, 10f)
                };

                // 2. Serialize the payload
                byte[] payloadBytes = _protoSerializer.Serialize(position);

                // 3. Create the antnet MessageHead using the generated enums
                byte[] headerBytes = CreateAntnetHeader((byte)Cmd.Position, (byte)ActPosition.Update, (uint)payloadBytes.Length);

                // 4. Combine header and payload
                message = new byte[headerBytes.Length + payloadBytes.Length];
                System.Buffer.BlockCopy(headerBytes, 0, message, 0, headerBytes.Length);
                System.Buffer.BlockCopy(payloadBytes, 0, message, headerBytes.Length, payloadBytes.Length);
                Debug.Log($"Sent position update: (X: {position.X}, Y: {position.Y}, Z: {position.Z})");
            }
            else if (connectionType == ConnectType.WebSocket)
            {
                // For WebSocket, create our new C2SPositionUpdate message
                var positionUpdate = new C2SPositionUpdate
                {
                    Route = "position/update",
                    // RequestId = System.Guid.NewGuid().ToString(), // request_id 暫時禁用
                    X = Random.Range(-10f, 10f),
                    Y = Random.Range(-10f, 10f),
                    Z = Random.Range(-10f, 10f)
                };

                // Serialize the C2SPositionUpdate message to JSON bytes
                message = _jsonSerializer.Serialize(positionUpdate);
                Debug.Log($"Sent position update: (X: {positionUpdate.X}, Y: {positionUpdate.Y}, Z: {positionUpdate.Z})");
            }

            // 5. Send the message
            _connection.Send(message);            
        }

        private void OnMessageReceived(byte[] messageBytes)
        {
            if (connectionType == ConnectType.WebSocket)
            {
                if (_jsonSerializer == null)
                {
                    Debug.LogError("JSONSerializer not initialized for WebSocket connection.");
                    return;
                }

                // 1. Deserialize to generic JObject to get the route
                JObject jsonObject = _jsonSerializer.DeserializeToJObject(messageBytes);
                if (jsonObject == null)
                {
                    Debug.LogError("Failed to parse WebSocket message as JSON.");
                    return;
                }

                string route = jsonObject["route"]?.ToString();
                // string requestId = jsonObject["request_id"]?.ToString(); // request_id 暫時禁用

                if (string.IsNullOrEmpty(route))
                {
                    Debug.LogWarning($"Received WebSocket message with no route: {jsonObject.ToString(Newtonsoft.Json.Formatting.None)}");
                    return;
                }

                // 2. Based on route, deserialize to specific S2C message type
                switch (route)
                {
                    case "position/update":
                        var positionResponse = _jsonSerializer.Deserialize<S2CPositionUpdate>(messageBytes);
                        Debug.Log($"WebSocket received position update response: Status: {positionResponse.Status}");
                        // Further handle positionResponse if needed
                        break;
                    case "gamer_info/retrieve":
                        var gamerInfoResponse = _jsonSerializer.Deserialize<S2CGamerInfoRetrieve>(messageBytes);
                        Debug.Log($"WebSocket received gamer info response: Nickname: {gamerInfoResponse.NickName}, Level: {gamerInfoResponse.Level}");
                        // Further handle gamerInfoResponse if needed
                        break;
                    default:
                        Debug.LogWarning($"Received WebSocket message with unknown route: {route}. Full message: {jsonObject.ToString(Newtonsoft.Json.Formatting.None)}");
                        break;
                }
                return;
            }

            // --- TCP Handling (remains largely the same) ---
            if (messageBytes.Length < 12)
            {
                Debug.LogError("Received message is too short to be a valid antnet message.");
                return;
            }

            uint len = System.BitConverter.ToUInt32(messageBytes, 0);
            ushort error = System.BitConverter.ToUInt16(messageBytes, 4);
            byte cmd = messageBytes[6];
            byte act = messageBytes[7];

            if (cmd == (byte)Cmd.Position && act == (byte)ActPosition.Update)
            {
                int payloadSize = messageBytes.Length - 12;
                byte[] payloadBytes = new byte[payloadSize];
                System.Buffer.BlockCopy(messageBytes, 12, payloadBytes, 0, payloadSize);

                var position = _protoSerializer.Deserialize<PlayerPosition>(payloadBytes);

                Debug.Log($"Received echo: (X: {position.X}, Y: {position.Y}, Z: {position.Z})");
                return;
            }
            
            Debug.LogWarning($"Received unknown message. Cmd: {cmd}, Act: {act}");
        }

        private byte[] CreateAntnetHeader(byte cmd, byte act, uint payloadLength)
        {
            using (var ms = new MemoryStream(12))
            {
                using (var writer = new BinaryWriter(ms))
                {
                    writer.Write(payloadLength); // Len (uint32)
                    writer.Write((ushort)0);    // Error (uint16)
                    writer.Write(cmd);           // Cmd (uint8)
                    writer.Write(act);           // Act (uint8)
                    writer.Write((ushort)0);    // Index (uint16)
                    writer.Write((ushort)0);    // Flags (uint16)
                    return ms.ToArray();
                }
            }
        }
    }
}
