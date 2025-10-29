using ProtoBufferExample.Client.Game.Models;
using ProtoBufferExample.Client.Game.Views;
using JsonApi;
using UnityEngine; // For MonoBehaviour

namespace ProtoBufferExample.Client.Game.Presenters
{
    public class ConnectionPresenter
    {
        private ConnectionModel _model;
        private IConnectionView _view;
        private bool _isConnecting;

        public ConnectionPresenter(ConnectionModel model, IConnectionView view)
        {
            _model = model ?? throw new System.ArgumentNullException(nameof(model));
            _view = view ?? throw new System.ArgumentNullException(nameof(view));

            _model.OnConnectionStatusChanged += OnConnectionStatusChanged;
            _model.OnMessageLogged += OnMessageLogged;

            UpdateUI();
        }

        private void UpdateUI()
        {
            bool isConnected = _model.IsConnected;
            
            // Update Connection Tab UI
            if (_isConnecting)
            {
                _view.SetConnectionStatusText("連線狀態: 連線中...");
                _view.SetConnectButtonEnabled(false);
            }
            else if (isConnected)
            {
                _view.SetConnectionStatusText("連線狀態: 已連線");
                _view.SetConnectButtonText("斷線");
                _view.SetConnectButtonEnabled(true);
            }
            else
            {
                _view.SetConnectionStatusText("連線狀態: 未連線");
                _view.SetConnectButtonText("連線");
                _view.SetConnectButtonEnabled(true);
                
            }

            // Update Test Function Tab UI
            _view.SetTestButtonsInteractable(isConnected);
        }

        private void OnConnectionStatusChanged(bool isConnected)
        {
            _isConnecting = false;
            UpdateUI();
        }

        private void OnMessageLogged(string message)
        {
            string formattedMessage = FormatLogMessage(message);
            _view.LogMessage(formattedMessage);
        }


        // Connection Tab Actions
        public void OnConnectButtonClicked()
        {
            if (_model.IsConnected)
            {
                Disconnect();
            }
            else
            {
                Connect();
            }
        }

        public void Connect()
        {
            _isConnecting = true;
            UpdateUI();

            var systemManager = Object.FindObjectOfType<Singleton.SystemManager>();
            if (systemManager != null)
            {
                _model.Connect(systemManager.ServerAddress, systemManager.ServerPort);
            }
        }

        public void Disconnect()
        {
            _model.Disconnect();
        }

        // Test Function Tab Actions
        public void SendPositionUpdate(float x, float y, float z)
        {
            if (!_model.IsConnected)
            {
                _view.ShowTestResult("❌ 無法發送：未連線");
                return;
            }

            var message = new C2SPositionUpdate {
                Route = "position/update",
                X = x,
                Y = y,
                Z = z
            };
            _model.SendMessage(message);

            _view.ShowTestResult($"✅ 已發送位置更新: ({x:F2}, {y:F2}, {z:F2})");
            Debug.Log($"Send Position Update: ({x}, {y}, {z})");
        }

        public void SendTestMessage(string messageType)
        {
            if (!_model.IsConnected)
            {
                _view.ShowTestResult($"❌ 無法發送 {messageType}：未連線");
                return;
            }

            _view.ShowTestResult($"✅ 已發送測試消息: {messageType}");
            Debug.Log($"Send Test Message: {messageType}");
        }

        private string FormatLogMessage(string message)
        {
            if (message.Contains("成功") || message.Contains("Connected"))
                return $"<color=green>✓</color> {message}";
            if (message.Contains("失敗") || message.Contains("failed"))
                return $"<color=red>✗</color> {message}";
            if (message.Contains("嘗試") || message.Contains("Attempting"))
                return $"<color=yellow>⟳</color> {message}";

            return message;
        }


    }
}
