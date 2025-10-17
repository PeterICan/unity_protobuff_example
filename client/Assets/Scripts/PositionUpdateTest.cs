using UnityEngine;
using System.IO;
using System.Buffers;
using ProtoBufferExample.Client.Generated;

namespace ProtoBufferExample.Client
{
    public class PositionUpdateTest : MonoBehaviour
    {
        public enum ConnectType
        {
            TCP,
            WebSocket
        }
        public string serverAddress = "127.0.0.1";
        public int serverPort = 6666;

        public ConnectType connectionType = ConnectType.TCP;

        private IConnection _connection;
        private ISerializer _serializer;

        void Start()
        {
            _connection = connectionType == ConnectType.TCP ? new TCPClientTransport() : new WebSocketClientTransport();
            _serializer = new ProtobufSerializer();

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
            // 1. Create the Protobuf message payload
            var position = new PlayerPosition
            {
                X = Random.Range(-10f, 10f),
                Y = Random.Range(-10f, 10f),
                Z = Random.Range(-10f, 10f)
            };

            // 2. Serialize the payload
            byte[] payloadBytes = _serializer.Serialize(position);

            // 3. Create the antnet MessageHead using the generated enums
            byte[] headerBytes = CreateAntnetHeader((byte)Cmd.Position, (byte)ActPosition.Update, (uint)payloadBytes.Length);

            // 4. Combine header and payload
            byte[] message = null;
            if (connectionType == ConnectType.TCP)
            {
                message = new byte[headerBytes.Length + payloadBytes.Length];
                System.Buffer.BlockCopy(headerBytes, 0, message, 0, headerBytes.Length);
                System.Buffer.BlockCopy(payloadBytes, 0, message, headerBytes.Length, payloadBytes.Length);
            }
            else if (connectionType == ConnectType.WebSocket)
            {

                message = new byte[headerBytes.Length + payloadBytes.Length];
                System.Buffer.BlockCopy(headerBytes, 0, message, 0, headerBytes.Length);
                System.Buffer.BlockCopy(payloadBytes, 0, message, headerBytes.Length, payloadBytes.Length);
                // //WebSocket 直接略過 header
                // message = new byte[payloadBytes.Length];
                // System.Buffer.BlockCopy(payloadBytes, 0, message, 0, payloadBytes.Length);
            }

            // 5. Send the message
            _connection.Send(message);

            Debug.Log($"Sent position update: (X: {position.X}, Y: {position.Y}, Z: {position.Z})");
        }

        private void OnMessageReceived(byte[] message)
        {
            // The server echoes the message back. Let's parse it.
            if (message.Length < 12)
            {
                Debug.LogError("Received message is too short to be a valid antnet message.");
                return;
            }

            // We can read the header to confirm it's the echo
            uint len = System.BitConverter.ToUInt32(message, 0);
            ushort error = System.BitConverter.ToUInt16(message, 4);
            byte cmd = message[6];
            byte act = message[7];

            if (cmd == (byte)Cmd.Position && act == (byte)ActPosition.Update)
            {
                // Extract the payload
                int payloadSize = message.Length - 12;
                byte[] payloadBytes = new byte[payloadSize];
                System.Buffer.BlockCopy(message, 12, payloadBytes, 0, payloadSize);

                // Deserialize the payload
                var position = _serializer.Deserialize<PlayerPosition>(payloadBytes);

                Debug.Log($"Received echo: (X: {position.X}, Y: {position.Y}, Z: {position.Z})");
            }
            else
            {
                Debug.LogWarning($"Received unknown message. Cmd: {cmd}, Act: {act}");
            }
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
