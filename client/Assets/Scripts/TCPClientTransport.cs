using System;
using System.Net.Sockets;
using System.Threading;

namespace ProtoBufferExample.Client
{
    public class TCPClientTransport : IConnection
    {
        private TcpClient _tcpClient;
        private NetworkStream _stream;
        private Thread _receiveThread;
        private bool _isConnected;

        public event Action<byte[]> OnMessageReceived;

        public bool IsConnected => _isConnected;

        public void Connect(string address, int port)
        {
            if (_isConnected)
            {
                return;
            }

            try
            {
                _tcpClient = new TcpClient(address, port);
                _stream = _tcpClient.GetStream();
                _isConnected = true;

                _receiveThread = new Thread(ReceiveLoop);
                _receiveThread.IsBackground = true;
                _receiveThread.Start();

                UnityEngine.Debug.Log($"Connected to server at {address}:{port}");
            }
            catch (Exception e)
            {
                UnityEngine.Debug.LogError($"Failed to connect to server: {e.Message}");
                _isConnected = false;
            }
        }

        public void Disconnect()
        {
            if (!_isConnected)
            {
                return;
            }

            _isConnected = false;
            _receiveThread?.Join();
            _stream?.Close();
            _tcpClient?.Close();

            UnityEngine.Debug.Log("Disconnected from server.");
        }

        public void Send(byte[] data)
        {
            if (!_isConnected || _stream == null || !_stream.CanWrite)
            {
                UnityEngine.Debug.LogWarning("Not connected. Cannot send data.");
                return;
            }

            try
            {
                _stream.Write(data, 0, data.Length);
            }
            catch (Exception e)
            {
                UnityEngine.Debug.LogError($"Failed to send data: {e.Message}");
                Disconnect();
            }
        }

        private void ReceiveLoop()
        {
            var headerBuffer = new byte[12]; // antnet MessageHead is 12 bytes
            while (_isConnected)
            {
                try
                {
                    if (!_stream.CanRead)
                    {
                        Thread.Sleep(100);
                        continue;
                    }

                    // 1. Read the message header
                    int bytesRead = 0;
                    while (bytesRead < headerBuffer.Length)
                    {
                        int read = _stream.Read(headerBuffer, bytesRead, headerBuffer.Length - bytesRead);
                        if (read == 0)
                        {
                            // Server disconnected
                            _isConnected = false;
                            break;
                        }
                        bytesRead += read;
                    }

                    if (!_isConnected) break;

                    // 2. Get the message body length from the header
                    // The 'Len' field is the first 4 bytes (uint32) in little-endian format.
                    uint bodyLength = BitConverter.ToUInt32(headerBuffer, 0);

                    // 3. Read the message body
                    byte[] bodyBuffer = new byte[bodyLength];
                    bytesRead = 0;
                    while (bytesRead < bodyLength)
                    {
                        int read = _stream.Read(bodyBuffer, bytesRead, (int)bodyLength - bytesRead);
                        if (read == 0)
                        {
                            // Server disconnected
                            _isConnected = false;
                            break;
                        }
                        bytesRead += read;
                    }

                    if (!_isConnected) break;

                    // 4. Combine header and body and invoke event
                    byte[] fullMessage = new byte[headerBuffer.Length + bodyBuffer.Length];
                    Buffer.BlockCopy(headerBuffer, 0, fullMessage, 0, headerBuffer.Length);
                    Buffer.BlockCopy(bodyBuffer, 0, fullMessage, headerBuffer.Length, bodyBuffer.Length);

                    OnMessageReceived?.Invoke(fullMessage);
                }
                catch (Exception e)
                {
                    if (_isConnected)
                    {
                        UnityEngine.Debug.LogError($"Error receiving data: {e.Message}");
                        _isConnected = false;
                    }
                }
            }
            Disconnect();
        }
    }
}
