using System;
using UnityEngine; // For Debug.Log, etc.
using Newtonsoft.Json.Linq;
using JsonApi;
using Google.Protobuf;

namespace ProtoBufferExample.Client.Game.Models
{
    public class ConnectionModel
    {
        private IConnection _connection;
        private ISerializer _serializer;

        public event Action<bool> OnConnectionStatusChanged;
        public event Action<string> OnMessageLogged;
        public event Action<IMessage> OnRawMessageReceived;

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
            try
            {
                //解析傳來的 rawData 成 JSON
                JObject jsonObject = _serializer.DeserializeToJObject(rawData);
                Debug.Log($"Received raw data as JSON: {jsonObject}");
                string route = jsonObject["route"]?.ToString();
                
                if (string.IsNullOrEmpty(route))
                {
                    Debug.LogWarning("Received message without route");
                    return;
                }
                
                IMessage message = null;
                switch (route)
                {
                    case "position/update":
                        message = _serializer.Deserialize<S2CPositionUpdate>(rawData);
                        Debug.Log($"Parsed Position Update Response: {message as S2CPositionUpdate}");
                        break;
                    case "position/notify":
                        message = _serializer.Deserialize<S2CNotifyWorldPositionChange>(rawData);
                        Debug.Log($"Parsed Position Notify: {message as S2CNotifyWorldPositionChange}");
                        break;
                    default:
                        Debug.LogWarning($"Unknown route: {route}");
                        return;
                }
                
                if (message == null)
                {
                    Debug.LogWarning($"Failed to deserialize message for route: {route}");
                    return;
                }
                
                if (message is not S2CPositionUpdate
                    && message is not S2CNotifyWorldPositionChange)
                {
                    Debug.Log($"Unexpected Log Message From ConnectionPresenter: {message.Descriptor.FullName}");
                }
                
                string logMessage = message switch
                {
                    S2CPositionUpdate posUpdate when posUpdate.Status != "" => 
                        $"[位置更新回應失敗] 狀態: {posUpdate.Status}",
                    S2CPositionUpdate => 
                        $"[位置更新回應成功]",
                    S2CNotifyWorldPositionChange => 
                        $"[位置變更通知]",
                    _ => 
                        $"[未知消息類型] {message.Descriptor.FullName}"
                };
                
                Debug.Log(logMessage);
                LogMessage(logMessage);
                
                // Trigger the raw message received event
                OnRawMessageReceived?.Invoke(message);
            }
            catch (Exception ex)
            {
                Debug.LogError($"Error handling received message: {ex.Message}");
                LogMessage($"Message handling error: {ex.Message}");
                // Don't rethrow the exception to prevent connection disruption
            }
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
            OnMessageLogged?.Invoke(message);
        }
    }
}