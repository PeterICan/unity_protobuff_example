using System;
using System.Diagnostics;
using System.Net.WebSockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using UnityEngine; // Assuming Unity context for UnityEngine.Debug.Log

namespace ProtoBufferExample.Client
{
    public class WebSocketClientTransport : IConnection
    {
        private ClientWebSocket _webSocket;
        private CancellationTokenSource _cancellationTokenSource;
        private bool _isConnected;

        public event Action<byte[]> OnMessageReceived;
        public event Action<bool> OnConnectionStatusChanged; // Implements IConnection event

        public bool IsConnected => _isConnected;

        public async void Connect(string address, int port)
        {
            if (_isConnected)
            {
                UnityEngine.Debug.LogWarning("WebSocket is already connected.");
                return;
            }

            _webSocket = new ClientWebSocket();
            _cancellationTokenSource = new CancellationTokenSource();
            Uri uri = new Uri($"ws://{address}:{port}/ws"); // Assuming /ws endpoint for WebSocket

            try
            {
                UnityEngine.Debug.Log($"Connecting to WebSocket: {uri}");
                await _webSocket.ConnectAsync(uri, _cancellationTokenSource.Token);
                _isConnected = true;
                OnConnectionStatusChanged?.Invoke(true); // Notify successful connection
                UnityEngine.Debug.Log("WebSocket connected.");

                // Start listening for messages
                _ = ReceiveLoop();
            }
            catch (WebSocketException e)
            {
                UnityEngine.Debug.LogError($"WebSocket connection error: {e.Message}");
                _isConnected = false;
                OnConnectionStatusChanged?.Invoke(false); // Notify connection failure
                Disconnect(); // Ensure cleanup
            }
            catch (Exception e)
            {
                UnityEngine.Debug.LogError($"General connection error: {e.Message}");
                _isConnected = false;
                OnConnectionStatusChanged?.Invoke(false); // Notify connection failure
                Disconnect(); // Ensure cleanup
            }
        }

        public async void Disconnect()
        {
            if (!_isConnected && _webSocket == null)
            {
                return;
            }

            _cancellationTokenSource?.Cancel(); // Signal cancellation to receive loop

            if (_webSocket.State == WebSocketState.Open || _webSocket.State == WebSocketState.Connecting)
            {
                try
                {
                    UnityEngine.Debug.Log("Disconnecting WebSocket...");
                    await _webSocket.CloseAsync(WebSocketCloseStatus.NormalClosure, "Client initiated disconnect", CancellationToken.None);
                }
                catch (Exception e)
                {
                    UnityEngine.Debug.LogError($"Error while closing WebSocket: {e.Message}");
                }
            }

            _webSocket?.Dispose();
            _webSocket = null;
            _isConnected = false;
            OnConnectionStatusChanged?.Invoke(false); // Notify disconnection
            UnityEngine.Debug.Log("WebSocket disconnected.");
        }

        public async void Send(byte[] data)
        {
            if (!_isConnected || _webSocket.State != WebSocketState.Open)
            {
                UnityEngine.Debug.LogWarning("Cannot send data: WebSocket not connected or not open.");
                return;
            }

            try
            {
                await _webSocket.SendAsync(new ArraySegment<byte>(data), WebSocketMessageType.Binary, true, _cancellationTokenSource.Token);
            }
            catch (Exception e)
            {
                UnityEngine.Debug.LogError($"Error sending data over WebSocket: {e.Message}");
                Disconnect(); // Disconnect will also invoke OnConnectionStatusChanged(false)
            }
        }

        private async Task ReceiveLoop()
        {
            byte[] buffer = new byte[4096]; // Buffer for incoming messages
            try
            {
                while (_webSocket.State == WebSocketState.Open && !_cancellationTokenSource.IsCancellationRequested)
                {
                    WebSocketReceiveResult result = await _webSocket.ReceiveAsync(new ArraySegment<byte>(buffer), _cancellationTokenSource.Token);

                    if (result.MessageType == WebSocketMessageType.Close)
                    {
                        UnityEngine.Debug.Log("WebSocket received close message from server.");
                        await _webSocket.CloseOutputAsync(WebSocketCloseStatus.NormalClosure, "Server initiated close", _cancellationTokenSource.Token);
                        // Disconnect() will be called, which invokes OnConnectionStatusChanged(false)
                        Disconnect();
                        break;
                    }

                    if (result.MessageType == WebSocketMessageType.Binary)
                    {
                        // Handle fragmented messages if necessary, but for now assume full message in one go
                        byte[] receivedData = new byte[result.Count];
                        Array.Copy(buffer, receivedData, result.Count);
                        OnMessageReceived?.Invoke(receivedData);
                    }
                    else if (result.MessageType == WebSocketMessageType.Text)
                    {
                        // Although we expect binary, handle text as well for robustness or debugging
                        string textMessage = Encoding.UTF8.GetString(buffer, 0, result.Count);
                        UnityEngine.Debug.LogWarning($"Received unexpected text message: {textMessage}");
                        // Potentially invoke OnMessageReceived with text data if needed, or log an error
                    }
                }
            }
            catch (OperationCanceledException)
            {
                UnityEngine.Debug.Log("WebSocket receive loop cancelled.");
            }
            catch (WebSocketException e)
            {
                UnityEngine.Debug.LogError($"WebSocket receive error: {e.Message}");
                Disconnect();
            }
            catch (Exception e)
            {
                UnityEngine.Debug.LogError($"General receive loop error: {e.Message}");
                Disconnect();
            }
            finally
            {
                if (_isConnected) 
                {
                    Disconnect();
                }
            }
        }
    }
}