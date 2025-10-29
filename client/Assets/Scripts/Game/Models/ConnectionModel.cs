using System;
using UnityEngine; // For Debug.Log, etc.
using ProtoBufferExample.Client;

namespace ProtoBufferExample.Client.Game.Models
{
    public class ConnectionModel
    {
        private IConnection _connection;
        private ISerializer _serializer;

        public event Action<bool> OnConnectionStatusChanged;
        public event Action<string> OnMessageLogged;
        public event Action<byte[]> OnRawMessageReceived;

        public bool IsConnected => _connection?.IsConnected ?? false;

        public ConnectionModel(IConnection connection, ISerializer serializer)
        {
            _connection = connection ?? throw new ArgumentNullException(nameof(connection));
            _serializer = serializer ?? throw new ArgumentNullException(nameof(serializer));
            _connection.OnMessageReceived += HandleRawMessageReceived;
            _connection.OnConnectionStatusChanged += HandleConnectionStatusChanged; // Subscribe to underlying connection status changes
        }

        public void Connect(string address, int port)
        {
            if (IsConnected)
            {
                LogMessage("Already connected.");
                return;
            }

            LogMessage($"Attempting to connect to {address}:{port}...");
            try
            {
                _connection.Connect(address, port);
            }
            catch (Exception ex)
            {
                LogMessage($"Connection failed: {ex.Message}");
                Debug.LogError($"Connection failed: {ex.Message}");
            }
        }

        public void Disconnect()
        {
            if (!IsConnected)
            {
                LogMessage("Already disconnected.");
                return;
            }

            LogMessage("Attempting to disconnect...");
            try
            {
                _connection.Disconnect();
            }
            catch (Exception ex)
            {
                LogMessage($"Disconnection failed: {ex.Message}");
                Debug.LogError($"Disconnection failed: {ex.Message}");
            }
        }

        public void SendMessage<T>(T message) where T : Google.Protobuf.IMessage<T>
        {
            if (!IsConnected)
            {
                LogMessage("Cannot send message: Not connected.");
                return;
            }

            try
            {
                byte[] data = _serializer.Serialize(message);
                _connection.Send(data);
                LogMessage($"Sent message: {message.GetType().Name}");
            }
            catch (Exception ex)
            {
                LogMessage($"Failed to send message: {ex.Message}");
                Debug.LogError($"Failed to send message: {ex.Message}");
            }
        }

        private void HandleRawMessageReceived(byte[] rawData)
        {
            OnRawMessageReceived?.Invoke(rawData);
            LogMessage($"Received raw data, length: {rawData.Length}");
        }

        private void HandleConnectionStatusChanged(bool isConnected)
        {
            // Propagate the status change from the underlying connection to ConnectionModel's subscribers
            OnConnectionStatusChanged?.Invoke(isConnected);
            if (isConnected)
            {
                LogMessage("Connection status changed: Connected.");
            }
            else
            {
                LogMessage("Connection status changed: Disconnected.");
            }
        }

        public void LogMessage(string message)
        {
            OnMessageLogged?.Invoke($"[{DateTime.Now:HH:mm:ss}] {message}");
        }
    }
}