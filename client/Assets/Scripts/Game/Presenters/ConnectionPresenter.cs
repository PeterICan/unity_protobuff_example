using ProtoBufferExample.Client.Game.Models;
using ProtoBufferExample.Client.Game.Views;
using UnityEngine; // For MonoBehaviour

namespace ProtoBufferExample.Client.Game.Presenters
{
    public class ConnectionPresenter
    {
        private ConnectionModel _model;
        private IConnectionView _view;

        public ConnectionPresenter(ConnectionModel model, IConnectionView view)
        {
            _model = model ?? throw new System.ArgumentNullException(nameof(model));
            _view = view ?? throw new System.ArgumentNullException(nameof(view));

            _model.OnConnectionStatusChanged += OnConnectionStatusChanged;
            _model.OnMessageLogged += OnMessageLogged;

            // Initial UI update
            _view.UpdateConnectionStatus(_model.IsConnected);
        }

        private void OnConnectionStatusChanged(bool isConnected)
        {
            _view.UpdateConnectionStatus(isConnected);
        }

        private void OnMessageLogged(string message)
        {
            _view.LogMessage(message);
        }

        public void Connect(string address, int port)
        {
            _model.Connect(address, port);
        }

        public void Disconnect()
        {
            _model.Disconnect();
        }
    }
}
